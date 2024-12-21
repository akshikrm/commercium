import server from "@/utils/server"
import { AxiosResponse } from "axios"

export const getOrders = async (): Promise<Order[]> => {
    try {
        const { data } = await server.get("/orders")
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ message: "failed to retreive orders" })
    }
}

export const placeOrder = async () => {
    try {
        const { data } = await server.post("/orders")
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to place order" })
    }
}

export const getByOrderID = async (orderID: string): Promise<OrderView> => {
    try {
        const { data } = await server.get(`orders/${orderID}`)
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to place order" })
    }
}
