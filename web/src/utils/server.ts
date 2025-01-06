import axios, { AxiosError } from "axios"
import { BASE_URL } from "@config"
import { clearSession, getToken } from "./session"
import toast from "react-hot-toast"

const server = axios.create({
    baseURL: BASE_URL,
    headers: {
        "Content-Type": "application/json"
    }
})

const TOAST_ID = "network_error"
server.interceptors.response.use(
    function (response) {
        return response
    },
    function (error) {
        const { status, response, code } = error as AxiosError<{
            error: string
        }>
        if (code === "ERR_NETWORK") {
            toast.error(
                "there seems to be some problem with your network connection, trying again...",
                { id: TOAST_ID }
            )
            return Promise.reject(
                "there seems to be some problem with your network connection, trying again..."
            )
        }
        if (status === 401) {
            clearSession()
        }
        return Promise.reject(response)
    }
)

server.interceptors.request.use(
    function (config) {
        const jwt = getToken()
        if (jwt) {
            config.headers.Authorization = jwt
        }
        return config
    },
    function (error) {
        return Promise.reject(error)
    }
)

export default server
