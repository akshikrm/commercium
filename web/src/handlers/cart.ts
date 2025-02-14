import server from "@/utils/server"
import { AxiosResponse } from "axios"

export const getCart = async (): Promise<Cart[]> => {
    try {
        const { data } = await server.get("/carts")
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ message: "failed to retreive cart" })
    }
}

export const addToCart = async (reqData: NewCartEntry) => {
    try {
        const { data } = await server.post("/carts", reqData)
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to create cart" })
    }
}

export const updateCart = async (payload: UpdateCart) => {
    const { cartID, quantity } = payload
    try {
        const { data } = await server.put(`/carts/${cartID}`, { quantity })
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to update cart" })
    }
}

export const deleteCart = async (cartID: number) => {
    try {
        const { data } = await server.delete(`/carts/${cartID}`)
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ messate: "failed to delete cart" })
    }
}
