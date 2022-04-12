import { Outlet } from "react-router-dom";
import NewPostsComponent from "../components/NewPosts";

export default function NewPosts() {
    return (
        <div className="">
            <div className="flex justify-center">
                <NewPostsComponent />
            </div>
            <Outlet />
        </div>

    );
}