import axios from 'axios'
import {fetcherProtected} from "./protected_api";

export const BASE_URL = "http://localhost:8080"

export const getTopPosts = async (offset, limit) => {
    try {
        const req = () => axios.get(`${BASE_URL}/api/v1/upload/top?offset=${offset}&limit=${limit}`, {withCredentials: true})
        let {data} = await req();
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}

export const getNewPosts = async (offset, limit) => {
    try {
        const req = () => axios.get(`${BASE_URL}/api/v1/upload/new?offset=${offset}&limit=${limit}`, {withCredentials: true})
        let {data} = await req();
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}


export const getPostById = async (itemId) => {
    try {
        const req = () => axios.get(`${BASE_URL}/api/v1/upload/${itemId}`, {withCredentials: true})
        let {data} = await req();
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}

export const getLoggedInStatus = async () => {
    return fetcherProtected(`${BASE_URL}/api/v1/user/isLoggedIn`)
}

export const getProfile = async () => {
    return fetcherProtected(`${BASE_URL}/api/v1/user/me`)
}

export const postLogin = async (xs) => {
    try {
        const req = () => axios.post(
            "http://localhost:8080/api/v1/user/login", xs, {withCredentials: true})
        let {data} = await req();
        return [data, null]
    } catch (error) {
        return [null, error]
    }
}

