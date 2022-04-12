import {Outlet} from "react-router-dom";
import './index.css';
import Header from "./components/Header";
import {useRecoilState} from "recoil";
import {authAtom} from "./state/states";
import {useEffect} from "react";
import {getLoggedInStatus} from "./api/api";


export default function App() {
    const [auth, setAuth] = useRecoilState(authAtom);

    useEffect(() => {
        async function xs() {
            const [data, error] = await getLoggedInStatus();
            if (!error) setAuth(true);
        }

        xs();
    }, [setAuth]);
    return (
        <div className="">
            <Header/>
            <Outlet/>
        </div>
    );
}