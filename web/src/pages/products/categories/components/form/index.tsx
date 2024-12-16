import RHFProvider from "@components/rhf/provider"
import RHFSelect from "@components/rhf/select"
import RHFTextField from "@components/rhf/text-field"
import useProductCategoryForm from "@hooks/product-categories/use-product-category-form"
import { Button, Stack } from "@mui/material"

type Add = (inputData: NewProductCategory) => void
type Edit = (inputData: EditProductCategory) => void

type Props = {
    defaultValues?: EditProductCategory
    buttonLabel: string
    onSubmit: Add | Edit
}

const ProductCategoryForm = ({
    onSubmit,
    buttonLabel,
    defaultValues
}: Props) => {
    const methods = useProductCategoryForm(defaultValues)

    return (
        <RHFProvider methods={methods} onSubmit={onSubmit}>
            <Stack>
                <RHFTextField name='name' label='Name' />
                <RHFTextField name='slug' label='Category URL' />
                <RHFTextField name='description' label='Description' />
                <RHFSelect name='enabled' label='Status'>
                    <option value='enabled'>yes</option>
                    <option value='disabled'>no</option>
                </RHFSelect>
                <Button type='submit'>{buttonLabel}</Button>
            </Stack>
        </RHFProvider>
    )
}

export default ProductCategoryForm
