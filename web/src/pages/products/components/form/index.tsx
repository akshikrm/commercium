import RHFProvider from "@/components/rhf/provider"
import { Box, Button, Grid2 as Grid } from "@mui/material"
import useProductForm from "@hooks/products/use-product-form"
import ProductFormCard from "./components/product-form-card"
import GeneralInformation from "./components/general-information"
import RHFSelect from "@components/rhf/select"
import RHFTextField from "@components/rhf/text-field"
import { useMemo } from "react"
import Render from "@components/render"
import SubscriptionPrice from "./components/subscription-price"
import RHFImageUpload from "@components/rhf/file-upload"

type Add = (inputData: NewProduct) => void
type Edit = (inputData: EditProduct) => void

type Props = {
    defaultValues?: EditProduct
    buttonLabel: string
    onSubmit: Add | Edit
}

const ProductForm = ({ buttonLabel, defaultValues, onSubmit }: Props) => {
    const methods = useProductForm(defaultValues)

    const { watch } = methods
    const productType = watch("type")

    const isSubscriptionProduct = useMemo(() => {
        return productType === "subscription"
    }, [productType])

    return (
        <RHFProvider methods={methods} onSubmit={onSubmit}>
            <Box sx={{ textAlign: "right", marginBottom: 2 }}>
                <Button type='submit'>{buttonLabel}</Button>
            </Box>
            <Grid container spacing={2}>
                <Grid size={4} spacing={2} direction='column' container>
                    <Grid size={12}>
                        <ProductFormCard title='Product Type'>
                            <RHFSelect label='Product Type' name='type'>
                                <option value='one-time'>One Time</option>
                                <option value='subscription'>
                                    Subscription
                                </option>
                            </RHFSelect>
                        </ProductFormCard>
                    </Grid>
                    <Grid size={12}>
                        <Render
                            when={isSubscriptionProduct}
                            show={<SubscriptionPrice />}
                            otherwise={
                                <ProductFormCard title='Price'>
                                    <RHFTextField
                                        name='price'
                                        label='Price'
                                        fullWidth
                                    />
                                </ProductFormCard>
                            }
                        />
                    </Grid>
                </Grid>
                <Grid size={8} spacing={2} direction='column' container>
                    <Grid size={12}>
                        <GeneralInformation />
                    </Grid>
                    <Grid size={12}>
                        <ProductFormCard title='Product Images'>
                            <RHFImageUpload name='image' label='Image' />
                        </ProductFormCard>
                    </Grid>
                </Grid>
            </Grid>
        </RHFProvider>
    )
}
export default ProductForm
