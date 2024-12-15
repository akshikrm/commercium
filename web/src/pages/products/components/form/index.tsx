import RHFProvider from "@/components/rhf/provider"
import RHFTextField from "@/components/rhf/text-field"
import { Button, Stack } from "@mui/material"
import ProductCategoryNames from "@/components/product-category-names"
import useProductForm from "@hooks/products/use-product-form"
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

    return (
        <RHFProvider methods={methods} onSubmit={onSubmit}>
            <Stack>
                <RHFTextField name='name' label='Product Name' />
                <RHFTextField name='slug' label='Product URL' />
                <RHFTextField name='description' label='Description' />
                <RHFImageUpload name='image' label='Image' />
                <ProductCategoryNames />
                <RHFTextField name='price' label='Price' />
                <Button type='submit'>{buttonLabel}</Button>
            </Stack>
        </RHFProvider>
    )
}

export default ProductForm
