import { Button } from "@mui/material"
import { useFormContext } from "react-hook-form"

const AddPriceButton = () => {
    const { addPrice } = useAddSubscriptionPrice()
    return <Button onClick={() => addPrice()}>Add Price</Button>
}

const useAddSubscriptionPrice = () => {
    const { getValues, setValue } = useFormContext()

    const addPrice = () => {
        setValue("subscription_price", [
            ...getValues("subscription_price"),
            {
                price: 0,
                label: "",
                frequency: 1,
                interval: "month"
            }
        ])
    }
    return { addPrice }
}

export default AddPriceButton
