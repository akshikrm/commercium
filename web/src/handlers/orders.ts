import server from "@/utils/server"
import { AxiosResponse } from "axios"
import toast from "react-hot-toast"

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

export const getInvoiceURI = async (txnId: string) => {
    try {
        const { data } = await server.get(`/orders/invoice/${txnId}`)
        if (data.data) {
            return open(data.data, "_blank")
        }
    } catch (err) {
        toast.error("failed to download invoice")
        console.error(err)
    }
}

export const isOrderComplete = async (txnId: string): Promise<boolean> => {
    try {
        const { data } = await server.get(`/orders/status/${txnId}`)
        return data.data === "completed"
    } catch (err) {
        console.log(err)
        return false
    }
}
