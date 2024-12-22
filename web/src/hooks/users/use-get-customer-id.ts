import { user } from "@api"
import { useQuery } from "@tanstack/react-query"
import { useEffect } from "react"
import toast from "react-hot-toast"

const CUSTOMER_TOAST_ID = "CUSTOMER_TOAST_ID"

const useGetCustomerID = (): {
    customerID: string
    getCustomerID: () => void
} => {
    const query = useQuery({
        queryKey: ["customer-id"],
        queryFn: () => user.getCustomerId(),
        initialData: "",
        enabled: false
    })

    const { data, status, refetch } = query

    useEffect(() => {
        if (status === "error") {
            toast.error("failed to get customer id", {
                id: CUSTOMER_TOAST_ID
            })
        }
    }, [status])

    const getCustomerID = () => refetch()

    return { customerID: data, getCustomerID }
}

export default useGetCustomerID
