import { order } from "@api"
import { useQuery } from "@tanstack/react-query"
import { useEffect } from "react"
import toast from "react-hot-toast"

const useGetPurchases = () => {
    const query = useQuery({
        queryKey: ["purchaseList"],
        queryFn: () => {
            return order.getMyOrders()
        }
    })
    const { status } = query

    useEffect(() => {
        if (status === "error") {
            toast.error("failed to load purchases")
        }
    }, [status])

    return query
}

export default useGetPurchases
