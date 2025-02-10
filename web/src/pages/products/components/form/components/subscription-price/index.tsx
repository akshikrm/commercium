import { Stack, Typography } from "@mui/material"
import RenderList from "@components/render-list"
import { useFormContext } from "react-hook-form"
import ProductFormCard from "../product-form-card"
import SubscriptionPriceItem from "./item"
import AddPriceButton from "./add-button"
import Render from "@components/render"
import { useMemo } from "react"

const SubscriptionPrice = () => {
    const {
        watch,
        formState: { errors }
    } = useFormContext()
    const subscriptionPrice = watch("subscription_price")

    const message = useMemo(
        () => (errors?.subscription_price?.root?.message as string) || "",
        [errors?.subscription_price]
    )

    return (
        <ProductFormCard title='Price'>
            <Stack>
                <RenderList
                    list={Object.entries(subscriptionPrice)}
                    render={([k]) => (
                        <SubscriptionPriceItem month={k} key={k} />
                    )}
                />
                <Render
                    when={Boolean(message)}
                    show={
                        <Typography variant='caption' color='error'>
                            {message}
                        </Typography>
                    }
                />
                <AddPriceButton />
            </Stack>
        </ProductFormCard>
    )
}

export default SubscriptionPrice
