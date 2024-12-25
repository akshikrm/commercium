import { user } from "@api"
import { useQuery } from "@tanstack/react-query"
import { useEffect } from "react"
import toast from "react-hot-toast"

const CUSTOMER_TOAST_ID = "CUSTOMER_TOAST_ID"

const useGetCustomerID = (): string => {
    const query = useQuery({
        queryKey: ["customer-id"],
        queryFn: () => user.getCustomerId(),
        initialData: ""
    })

    const { data, status } = query

    useEffect(() => {
        if (status === "error") {
            toast.error("failed to get customer id", {
                id: CUSTOMER_TOAST_ID
            })
        }
    }, [status])

    return data
}

export default useGetCustomerID
