import RHFProvider from "@/components/rhf/provider"
import { Box, Button, Grid2 as Grid } from "@mui/material"
import useProductForm from "@hooks/products/use-product-form"
import RHFImageUpload from "@components/rhf/file-upload"
import ProductFormCard from "./components/product-form-card"
import GeneralInformation from "./components/general-information"
import RHFSelect from "@components/rhf/select"

type Add = (inputData: NewProduct) => void
type Edit = (inputData: EditProduct) => void

type Props = {
    defaultValues?: EditProduct
    buttonLabel: string
    onSubmit: Add | Edit
}

const ProductForm = ({ buttonLabel, defaultValues, onSubmit }: Props) => {
    const methods = useProductForm(defaultValues)
    return (
        <RHFProvider methods={methods} onSubmit={onSubmit}>
            <Box sx={{ textAlign: "right", marginBottom: 2 }}>
                <Button type='submit'>{buttonLabel}</Button>
            </Box>
            <Grid container spacing={2}>
                <Grid container size={4} spacing={2}>
                    <Grid size={12}>
                        <ProductFormCard title='Product Images'>
                            <RHFImageUpload name='image' label='Image' />
                        </ProductFormCard>
                    </Grid>
                    <Grid size={12}>
                        <ProductFormCard title='Status'>
                            <RHFSelect label='Status' name='status'>
                                <option value='enabled'>Enabled</option>
                                <option value='disabled'>Disabled</option>
                            </RHFSelect>
                        </ProductFormCard>
                    </Grid>
                </Grid>

                <Grid size={8}>
                    <GeneralInformation />
                </Grid>
            </Grid>
        </RHFProvider>
    )
}

export default ProductForm
