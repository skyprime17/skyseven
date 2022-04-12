import axios from "axios";
import {getError} from "./errors";
import {BASE_URL} from "./api";


export const refreshTokens = async () => {
    await axios.post(`${BASE_URL}/api/v1/user/refreshToken`, undefined, {withCredentials: true})
}

export const handleProtectedRequest = async (request) => {
    try {
        return await request()
    } catch (error) {
        if (error?.response?.status === 401) {
            try {
                await refreshTokens()
                return await request()
            } catch (innerError) {
                throw getError(innerError)
            }
        }

        throw getError(error)
    }
}


export const fetcherProtected = async (url) => {
    try {
        const request = () => axios.get(url, {withCredentials: true})
        const {data} = await handleProtectedRequest(request)
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}

export const posterProtected = async (url, payload) => {
    try {
        const request = () => axios.post(url, payload, {withCredentials: true})
        const {data} = await handleProtectedRequest(request)
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}


export const uploadPost = async (file) => {
    try {
        const request = () => axios.post(`${BASE_URL}/api/v1/upload`, file,
            {
                withCredentials: true,
                headers: {"content-type": "multipart/form-data"}
            })
        const {data} = await handleProtectedRequest(request)
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}
















