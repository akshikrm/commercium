import server from "@utils/server"
import { decodeJWT, getToken } from "@utils/session"
import { AxiosResponse } from "axios"

export const profile = async (): Promise<Profile> => {
    try {
        const { data } = await server.get("/profile")
        const jwt = getToken()
        const decoded = decodeJWT(jwt)
        const profile: Profile = { ...data.data, role: decoded.role }
        return profile
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ message: "failed to get products" })
    }
}
