import {useParams} from "react-router-dom";
import {getPostById} from "../api/api";
import {useEffect, useState} from "react";


export default function Meme() {
    let params = useParams();
    const [data, setData] = useState();

    useEffect(() => {
        async function xs() {
            const [data, error] = await getPostById(params.memeId);
            if (!error && data) setData(data);
        }

        xs();
    }, [params]);
    console.log(data);

    return (
        <div className="flex justify-center py-5">
            {data &&
                <div className="">
                    <div className="bg-black drop-shadow-2xl max-w-6xl">
                        <img className="rounded-lg md: w-auto"
                             src={`http://localhost:8085/${data.fileId}`} alt={data.fileId}/>
                    </div>
                    <div className="bg-amber-50 drop-shadow-2xl rounded-lg ">
                        <div className="ml mr-2 flex justify-between">
                            <div>
                                <h1>{data.userId}</h1>
                            </div>
                            <div>
                                <h1>{data.createdAt}</h1>
                            </div>
                            <div className="flex gap-4">
                                <h1>{data.up}</h1><h1>{data.down}</h1>
                            </div>
                        </div>
                        <div className="ml mr-2 flex justify-between">
                            <p>lorem ipsum</p>
                        </div>


                    </div>
                </div>
            }
            {!data && <div> Loading.... </div>}
        </div>

    );
}