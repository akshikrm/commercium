import useGetCustomerID from "@hooks/users/use-get-customer-id"
import { useCallback, useEffect } from "react"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"
import { USER_PATHS } from "@/paths"
import useConnectPaddle from "./use-connect-paddle"

const ORDER_TOAST = "order_toast"

const usePlaceOrder = (purchaseItems: PaddlePurchaseItem[]) => {
    const customerID = useGetCustomerID()
    const navigate = useNavigate()

    const { paddle, event } = useConnectPaddle(customerID)

    const placeOrder = useCallback(async () => {
        try {
            if (paddle) {
                paddle.Checkout.open({
                    settings: {
                        displayMode: "overlay",
                        variant: "one-page"
                    },
                    items: purchaseItems,
                    customer: {
                        id: customerID
                    }
                })
            }
        } catch (error) {
            const err = error as Error
            toast.error(err.message, { id: ORDER_TOAST })
            console.log(err)
        }
    }, [customerID, purchaseItems, paddle])

    useEffect(() => {
        if (event.name === "checkout.completed") {
            setTimeout(() => {
                navigate(USER_PATHS.orders.root)
                paddle?.Checkout.close()
            }, 1000)
        }
    }, [event, paddle])

    return placeOrder
}

export default usePlaceOrder
