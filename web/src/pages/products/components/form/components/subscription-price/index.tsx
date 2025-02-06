import { Stack } from "@mui/material"
import RenderList from "@components/render-list"
import { useFormContext } from "react-hook-form"
import ProductFormCard from "../product-form-card"
import SubscriptionPriceItem from "./item"
import AddPriceButton from "./add-button"

const SubscriptionPrice = () => {
    const { watch } = useFormContext()
    const price = watch("subscriptionPrice")

    return (
        <ProductFormCard title='Price'>
            <Stack>
                <RenderList
                    list={Object.entries(price)}
                    render={([k]) => (
                        <SubscriptionPriceItem month={k} key={k} />
                    )}
                />
                <AddPriceButton />
            </Stack>
        </ProductFormCard>
    )
}

export default SubscriptionPrice
