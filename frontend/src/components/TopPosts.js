import {useEffect, useState} from "react";
import {getTopPosts} from "../api/api";
import {Link} from "react-router-dom";

export default function TopPostsComponent() {
    const [data, setData] = useState([]);

    useEffect(() => {
        async function xs() {
            const [data, error] = await getTopPosts(0, 50);
            if (!error && data) setData(data);
        }

        xs();
    }, []);

    const posts = (xs) => {
        return (
            <div className="grid gap-2 grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:">
                {xs.map((element) => {
                    return (
                        <div className="bg-black w-32 h-32 drop-shadow-2xl rounded-lg md: w-40 h-40" key={element.id}>
                            <Link to={`/top/${element.id}`} key={element.id}>
                                <img className="object-fill w-32 h-32 rounded-lg md: w-40 h-40"
                                     src={`http://localhost:8085/${element.fileId}`} alt={element.fileId}/>
                            </Link>
                        </div>)
                })}
            </div>
        )

    }

    return (
        <div className="flex justify-center mt-2">
            {posts(data)}
        </div>

    );
}