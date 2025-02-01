import useGetCustomerID from "@hooks/users/use-get-customer-id"
import { useCallback, useEffect } from "react"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"
import { USER_PATHS } from "@/paths"
import useConnectPaddle from "./use-connect-paddle"
import { order } from "@api"
import useAuth from "@hooks/auth/use-auth"

const ORDER_TOAST = "order_toast"

const usePlaceOrder = (purchaseItems: PaddlePurchaseItem[]) => {
    const customerID = useGetCustomerID()
    const { user } = useAuth()
    const { email } = user
    const navigate = useNavigate()

    const { paddle, event } = useConnectPaddle(customerID)

    const placeOrder = useCallback(async () => {
        const transactionID = await order.createTransaction()
        try {
            if (paddle) {
                paddle.Checkout.open({
                    settings: {
                        displayMode: "overlay",
                        variant: "multi-page"
                    },
                    transactionId: transactionID
                })
            }
        } catch (error) {
            const err = error as Error
            toast.error(err.message, { id: ORDER_TOAST })
            console.log(err)
        }
    }, [customerID, purchaseItems, paddle, email])

    const { name, data } = event || {}
    const { transaction_id } = data || {}

    useEffect(() => {
        if (name === "checkout.completed") {
            const checkStatus = async (txnId: string) => {
                const isComplete = await order.isOrderComplete(txnId)
                if (isComplete) {
                    navigate(USER_PATHS.orders.root)
                    paddle?.Checkout.close()
                }
            }

            if (!transaction_id) return

            const intervalID = setInterval(() => {
                checkStatus(transaction_id)
            }, 3000)

            return () => {
                clearInterval(intervalID)
            }
        }
    }, [name, transaction_id, paddle])

    return placeOrder
}

export default usePlaceOrder
