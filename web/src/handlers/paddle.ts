import { PADDLE_ENVIRONMENT, PADDLE_TOKEN } from "@config"
import { initializePaddle, Paddle } from "@paddle/paddle-js"

export const connectToPaddle = async (
    customerID: string
): Promise<Paddle | undefined> => {
    try {
        return await initializePaddle({
            token: PADDLE_TOKEN,
            environment: PADDLE_ENVIRONMENT,
            pwCustomer: {
                id: customerID
            }
        })
    } catch (err) {
        console.log(err)
        return undefined
    }
}
