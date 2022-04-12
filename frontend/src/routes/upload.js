import {useForm} from "react-hook-form";
import {uploadPost} from "../api/protected_api";

export default function Upload() {
    const {register, handleSubmit} = useForm();
    const onSubmit = (data) => {
        console.log(data.file[0])
        const formData = new FormData();
        formData.append('file', data.file[0])
        uploadPost(formData)
    };
    return (
        <div className="max-w-screen-md mx-auto flex flex-wrap flex-col justify-center mt-16">
            <h1>Upload your file</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <input {...register("file", {required: true})}
                       className="form-control block px-3 py-1.5 text-base font-normal text-gray-700 bg-white
                       bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0
                       focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                       type="file"/>
                <button type="submit"
                        className="inline-block px-6 py-2.5 bg-gray-200 text-gray-900 font-medium text-xs mt-2
                        leading-tight uppercase rounded-md shadow-md hover:bg-gray-300 hover:shadow-lg
                        focus:bg-gray-300 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-gray-400
                        active:shadow-lg transition duration-150 ease-in-out">
                    Submit
                </button>
            </form>

        </div>

    );
}