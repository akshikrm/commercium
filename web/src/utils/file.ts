import axios, { AxiosError } from "axios"
import toast from "react-hot-toast"

const cloudinary = axios.create({
    baseURL: "https://api.cloudinary.com/v1_1/commercium/",
    headers: {
        "Content-Type": "multipart/form-data"
    }
})

const TOAST_ID = "network_error"

cloudinary.interceptors.response.use(
    function (response) {
        return response
    },
    function (error) {
        const { response, code } = error as AxiosError<{
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

        return Promise.reject(response)
    }
)

export default cloudinary
