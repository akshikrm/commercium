import { useMutation } from "@tanstack/react-query"
import { order } from "@api"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"

const ORDER_TOAST = "order_toast"

const usePlaceOrder = () => {
    const navigate = useNavigate()
    const mutation = useMutation({
        mutationFn: () => order.placeOrder(),
        onMutate: () => {
            toast.loading("placing order", {
                id: ORDER_TOAST
            })
        },
        onSuccess: data => {
            toast.success(data, {
                id: ORDER_TOAST
            })
            navigate("/purchase")
        },
        onError: err => {
            console.error(err)
            toast.error("failed to place order", {
                id: ORDER_TOAST
            })
        }
    })
    return mutation
}

export default usePlaceOrder
