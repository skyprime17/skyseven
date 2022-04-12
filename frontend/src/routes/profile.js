import {useEffect, useState} from "react";
import {getProfile} from "../api/api";

export default function Profile() {
    const [data, setData] = useState();
    useEffect(() => {
        async function xs() {
            const [data, error] = await getProfile();
            if (!error && data) setData(data);
        }

        xs();
    }, []);
    return (
        <div className="max-w-screen-md mx-auto">
            <div className="flex items-center justify-center h-full text-center">
                <div className="space-y-4 text-center flex flex-col">
                    {
                        data &&
                        <div>
                            <p>Hi, {data.username} 👋</p>
                            <p>{data.id}</p>
                            <p>{data.createdAt}</p>
                            <p>{data.enabled.toString()}</p>
                        </div>
                    }
                </div>
            </div>
        </div>
    );
}