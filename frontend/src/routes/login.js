import {useEffect} from "react";
import {postLogin} from "../api/api";
import {useRecoilState} from "recoil";
import {authAtom} from "../state/states";
import * as yup from "yup";
import {yupResolver} from '@hookform/resolvers/yup';
import {useNavigate} from "react-router-dom";
import {useForm} from "react-hook-form";


const SignupSchema = yup.object().shape({
    username: yup.string().min(5).max(36).required(),
    password: yup.string().min(3).max(200).required(),
});

export default function Login() {
    let navigate = useNavigate()

    const [auth, setAuth] = useRecoilState(authAtom);

    const {register, handleSubmit, formState: {errors}} = useForm({resolver: yupResolver(SignupSchema)});

    useEffect(() => {
        if (auth) navigate("/");
    }, [auth, navigate]);

    const onSubmit = async (x) => {
        const [data, error] = await postLogin(JSON.stringify(x))
        setAuth(!error);
        if (!error) navigate("/");
    };

    return (
        <div className="max-w-screen-md mx-auto flex justify-center">
            <div className="">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-2">
                        <input {...register("username")}
                               className="px-4 py-2 text-lg font-medium bg-gray-200 rounded"
                               placeholder="username"
                        />
                        {errors.username && <p className="text-red-600">{errors.username.message}</p>}
                    </div>
                    <div className="">
                        <input {...register("password")}
                               className="px-4 py-2 text-lg font-medium bg-gray-200 rounded"
                               placeholder="password"
                               type="password"
                        />
                        {errors.password && <p className="text-red-600">{errors.password.message}</p>}
                    </div>
                    <input className="mt-2 bg-gray-200 font-medium text-lg rounded" type="submit"/>
                </form>
            </div>
        </div>
    );
}