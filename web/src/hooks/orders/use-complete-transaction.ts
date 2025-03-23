import { useCallback, useEffect } from "react"
import { order } from "@api"
import useConnectPaddle from "@hooks/orders/use-connect-paddle"
import useGetCustomerID from "@hooks/users/use-get-customer-id"

const useCompleteTransaction = (cb: () => void) => {
    const customerID = useGetCustomerID()
    const { paddle, event } = useConnectPaddle(customerID)

    const completeOrder = useCallback(
        async (txnID: string) => {
            try {
                if (paddle) {
                    paddle.Checkout.open({
                        settings: {
                            displayMode: "overlay",
                            variant: "multi-page"
                        },
                        transactionId: txnID
                    })
                }
            } catch (error) {
                const err = error as Error
                console.log(err)
            }
        },
        [customerID, paddle]
    )

    const { name, data } = event || {}
    const { transaction_id } = data || {}

    useEffect(() => {
        if (name === "checkout.completed") {
            const checkStatus = async (txnId: string) => {
                const isComplete = await order.isOrderComplete(txnId)
                if (isComplete) {
                    cb()
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
    return completeOrder
}

export default useCompleteTransaction
