import { Stack } from "@mui/material"
import { useEffect } from "react"
import RenderList from "@components/render-list"
import { useFormContext } from "react-hook-form"
import ProductFormCard from "../product-form-card"
import SubscriptionPriceItem from "./item"
import AddPriceButton from "./add-button"

const SubscriptionPrice = () => {
    const { watch, setValue } = useFormContext()

    const [productType, price] = watch(["type", "price"])
    useEffect(() => {
        if (productType === "subscription") {
            setValue("price", {
                "1_month": {
                    label: "",
                    value: ""
                }
            })
        }
    }, [productType])

    return (
        <ProductFormCard title='Price'>
            <Stack>
                <RenderList
                    list={Object.entries(price)}
                    render={([k]) => <SubscriptionPriceItem month={k} />}
                />
                <AddPriceButton />
            </Stack>
        </ProductFormCard>
    )
}

export default SubscriptionPrice
