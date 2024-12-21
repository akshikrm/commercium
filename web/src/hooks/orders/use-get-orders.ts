import { order } from "@api"
import { useQuery } from "@tanstack/react-query"
import { useEffect } from "react"
import toast from "react-hot-toast"

const useGetOrders = () => {
    const query = useQuery({
        queryKey: ["ordersList"],
        queryFn: () => {
            return order.getMyOrders()
        }
    })
    const { status } = query

    useEffect(() => {
        if (status === "error") {
            toast.error("failed to load orders")
        }
    }, [status])

    return query
}

export default useGetOrders
