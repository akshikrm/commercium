import RHFTextField from "@/components/rhf/text-field"
import { Grid2 as Grid } from "@mui/material"
import ProductCategoryNames from "@/components/product-category-names"
import ProductFormCard from "./product-form-card"
import RHFSelect from "@components/rhf/select"

const GeneralInformation = () => {
    return (
        <ProductFormCard title='General Information'>
            <Grid container spacing={2}>
                <Grid size={12}>
                    <RHFTextField name='name' label='Product Name' fullWidth />
                </Grid>
                <Grid size={12}>
                    <RHFTextField name='slug' label='Product URL' fullWidth />
                </Grid>
                <Grid size={6}>
                    <ProductCategoryNames />
                </Grid>
                <Grid size={6}>
                    <RHFSelect label='Status' name='status'>
                        <option value='enabled'>Enabled</option>
                        <option value='disabled'>Disabled</option>
                    </RHFSelect>
                </Grid>
                <Grid size={12}>
                    <RHFTextField
                        name='description'
                        label='Description'
                        fullWidth
                        multiline
                        minRows={3}
                    />
                </Grid>
            </Grid>
        </ProductFormCard>
    )
}

export default GeneralInformation
