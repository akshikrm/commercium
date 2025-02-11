import { Button } from "@mui/material"
import { useFormContext } from "react-hook-form"

const useAddSubscriptionPrice = () => {
    const { getValues, setValue } = useFormContext()

    const addPrice = () => {
        const price = getValues("subscription_price")
        const power = Object.keys(price).length - 1
        const nextMonth = 2 ** power * 3
        const temp: SubscriptionPrice = {
            [`${nextMonth}_month`]: {
                label: "",
                price: ""
            }
        }
        setValue("subscription_price", {
            ...price,
            ...temp
        })
    }
    return { addPrice }
}

const AddPriceButton = () => {
    const { addPrice } = useAddSubscriptionPrice()
    return <Button onClick={() => addPrice()}>Add Price</Button>
}

export default AddPriceButton
