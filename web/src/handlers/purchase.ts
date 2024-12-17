import server from "@/utils/server"
import { AxiosResponse } from "axios"

export const getPurcahses = async (): Promise<Purchases[]> => {
    try {
        const { data } = await server.get("/purchase")
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ message: "failed to retreive purchase" })
    }
}

export const placeOrder = async () => {
    try {
        const { data } = await server.post("/purchase")
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to place order" })
    }
}
