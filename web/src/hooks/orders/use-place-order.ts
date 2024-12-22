import useGetCustomerID from "@hooks/users/use-get-customer-id"
import { useCallback } from "react"
import toast from "react-hot-toast"
import { paddle } from "@api"

const ORDER_TOAST = "order_toast"

const usePlaceOrder = (purchaseItems: PaddlePurchaseItem[]) => {
    const customerID = useGetCustomerID()

    const placeOrder = useCallback(async () => {
        try {
            const paddleInstance = await paddle.connect(customerID)
            if (paddleInstance) {
                paddleInstance.Checkout.open({
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
    }, [customerID, purchaseItems])

    return placeOrder
}

export default usePlaceOrder
