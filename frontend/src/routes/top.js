import { Outlet } from "react-router-dom";
import TopPostsComponent from "../components/TopPosts";

export default function TopPosts() {
    return (
        <div className="">
            <div className="flex justify-center">
                <TopPostsComponent />
            </div>
            <Outlet />
        </div>

    );
}