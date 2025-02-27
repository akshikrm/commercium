import { Stack } from "@mui/material"
import RenderList from "@components/render-list"
import { useFormContext } from "react-hook-form"
import ProductFormCard from "../product-form-card"
import SubscriptionPriceItem from "./item"
import AddPriceButton from "./add-button"

const SubscriptionPrice = () => {
    const {
        watch,
        setValue,
        getValues,
        formState: { errors }
    } = useFormContext<NewProduct>()
    const subscriptionPrice: NewSubscriptionPrice[] =
        watch("subscription_price")

    const allErrors = errors.subscription_price
    return (
        <ProductFormCard title='Price'>
            <Stack>
                <RenderList
                    list={subscriptionPrice}
                    render={(price, i) => {
                        const error = allErrors ? allErrors[i] : undefined
                        return (
                            <SubscriptionPriceItem
                                key={i}
                                subscriptionPrice={price}
                                error={error}
                                onChange={v => {
                                    const temp: NewSubscriptionPrice[] = [
                                        ...getValues("subscription_price")
                                    ]
                                    temp.splice(i, 1, v)
                                    setValue("subscription_price", temp)
                                }}
                            />
                        )
                    }}
                />
                <AddPriceButton />
            </Stack>
        </ProductFormCard>
    )
}

export default SubscriptionPrice
