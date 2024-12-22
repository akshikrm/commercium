import { useCallback, useEffect, useState } from "react"
import { PADDLE_ENVIRONMENT, PADDLE_TOKEN } from "@config"
import { initializePaddle, Paddle, PaddleEventData } from "@paddle/paddle-js"

const useConnectPaddle = (
    customerID: string
): { paddle: Paddle | undefined; event: PaddleEventData } => {
    const [paddle, setPaddle] = useState<Paddle>()
    const [event, setEvent] = useState<PaddleEventData>({})
    const connectToPaddle = useCallback(async (customerID: string) => {
        try {
            const paddleInstance = await initializePaddle({
                token: PADDLE_TOKEN,
                environment: PADDLE_ENVIRONMENT,
                pwCustomer: {
                    id: customerID
                },
                eventCallback: event => {
                    setEvent(event)
                }
            })
            if (paddleInstance) {
                setPaddle(paddleInstance)
            }
        } catch (err) {
            console.log(err)
        }
    }, [])

    useEffect(() => {
        if (customerID) {
            connectToPaddle(customerID)
        }
    }, [customerID, connectToPaddle])

    return { paddle, event }
}

export default useConnectPaddle
