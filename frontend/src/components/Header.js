import {NavLink, useNavigate} from "react-router-dom";
import {useRecoilState, useRecoilValue} from "recoil";
import {authAtom} from "../state/states";
import {posterProtected} from "../api/protected_api";
import {BASE_URL} from "../api/api";


function Header() {
    const loggedIn = useRecoilValue(authAtom)
    let navigate = useNavigate();
    const [auth, setAuth] = useRecoilState(authAtom);

    const navActive = ({isActive}) => {
        return isActive ? "text-gray-900" : "text-gray-500 hover:text-gray-900"
    }

    const handleLogout = () => {
        posterProtected(`${BASE_URL}/api/v1/user/logout`, null)
        setAuth(false)
        navigate("/", {replace: true})
    }

    const GeneralNavLinks = () => {
        return (
            <div className="flex gap-4">
                <NavLink className={navActive} to="/top">
                    top
                </NavLink>
                <NavLink className={navActive} to="/new">
                    new
                </NavLink>
            </div>
        )
    }

    const SignedInNavLinks = () => {
        return (loggedIn &&
            <div className="flex gap-4">
                <NavLink className={navActive} to="/profile">
                    profile
                </NavLink>
                <NavLink className={navActive} to="/upload">
                    upload
                </NavLink>
                <button type="button" className="text-gray-500 hover:text-gray-900" onClick={handleLogout}>
                    logout
                </button>

            </div>
        )
    }

    const SignedOutNavLinks = () => {
        return (!loggedIn &&
            <>
                <NavLink className={navActive}
                         to="/login">login</NavLink>
            </>
        )
    }

    return (
        <div className="flex justify-center sticky top-0 z-50 bg-white">
            <div className="border-black flex justify-between
            text-xl mb-3 mt-3 w-[496px] md:w-[664px] lg:w-[832px]">

                <NavLink className={navActive} to="/">
                    skyseven
                </NavLink>

                <div className="flex">
                    <GeneralNavLinks/>
                </div>

                { /*
                                <div>
                    <input className="bg-gray-200 rounded-sm"/>
                </div>
                */}


                <div className="flex">
                    <SignedInNavLinks/>
                    <SignedOutNavLinks/>
                </div>
            </div>
        </div>
    )
}

export default Header;