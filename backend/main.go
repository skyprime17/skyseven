package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pr0clone/auth"
	"pr0clone/conf"
	"pr0clone/controllers"
	"pr0clone/middleware"
	repo "pr0clone/repository"
	service "pr0clone/services"
	"syscall"
	"time"
)

var client *sqlx.DB

func init() {
	dbClient, err := conf.DbClient()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	client = dbClient
}

func main() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowFiles = true
	router.Use(cors.New(corsConfig))

	authMiddleware := auth.NewJWTProvider("secret")
	authService := service.NewAuthService(repo.NewUserRepository(client), auth.NewJWTProvider("secret"))
	userController := controllers.UserController{IAuthService: authService}
	userUploadController := controllers.UserUploadController{
		IUserUploadService: service.NewUserUploadService(repo.NewUserUploadRepository(client), authService)}

	v1 := router.Group("api/v1")
	{
		v1.GET("/health", controllers.Health)

		userGroup := v1.Group("user")
		userGroup.POST("/login", userController.LogIn)
		userGroup.POST("/logout", userController.Logout)
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/refreshToken", userController.RefreshToken)
		userGroup.Use(middleware.AuthMiddleware(authMiddleware))
		userGroup.GET("/isLoggedIn", userController.IsLoggedIn)
		userGroup.GET("/me", userController.Me)

		uploadGroup := v1.Group("upload")
		uploadGroup.GET("/top", userUploadController.GetTopPosts)
		uploadGroup.GET("/new", userUploadController.GetNewestPosts)
		uploadGroup.GET("/:itemId", userUploadController.GetPostById)
		uploadGroup.Use(middleware.AuthMiddleware(authMiddleware)).POST("", userUploadController.UploadFile)

	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
