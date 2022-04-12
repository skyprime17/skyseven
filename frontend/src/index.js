import {
    BrowserRouter,
    Routes,
    Route, Navigate, Outlet
} from "react-router-dom";
import App from "./App";
import ReactDOMClient from "react-dom/client";
import './index.css';
import Home from "./routes/home";
import Meme from "./routes/meme";
import TopPosts from "./routes/top";
import NewPosts from "./routes/new";
import {RecoilRoot, useRecoilValue} from "recoil";
import Login from "./routes/login";
import Profile from "./routes/profile";
import {authAtom} from "./state/states";
import NotFound from "./routes/notfound";
import Upload from "./routes/upload";

const container = document.getElementById("root");

const root = ReactDOMClient.createRoot(container);

function PrivateOutlet() {
    const loggedIn = useRecoilValue(authAtom)
    return loggedIn ? <Outlet /> : <Navigate to="/login" />;
}


root.render(
    <RecoilRoot>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<App/>}>
                    <Route path="/" element={<Home/>}>
                    </Route>
                    <Route path="login" element={<Login/>}/>

                    <Route path="profile" element={<PrivateOutlet/>}>
                        <Route path="" element={<Profile/>}/>
                    </Route>

                    <Route path="upload" element={<PrivateOutlet/>}>
                        <Route path="" element={<Upload/>}/>
                    </Route>

                    <Route path="top" element={<TopPosts/>}/>
                    <Route path="/top/:memeId" element={<Meme/>}/>
                    <Route path="new" element={<NewPosts/>}/>
                    <Route path="*" element={ <NotFound/> } />
                </Route>
            </Routes>
        </BrowserRouter>
    </RecoilRoot>
);