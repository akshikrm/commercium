import HeaderBreadcrumbs from "@components/header"
import { Card } from "@mui/material"
import ProductCategoryForm from "../components/form"
import useCreateProductCategory from "@hooks/product-categories/use-create-product-category"

const AddProductCategory = () => {
    const { mutate } = useCreateProductCategory()
    return (
        <>
            <HeaderBreadcrumbs
                heading='Add Category'
                links={[
                    {
                        label: "home",
                        href: "/"
                    },
                    {
                        label: "products",
                        href: "/admin/products"
                    },
                    {
                        label: "categories",
                        href: "/admin/products/categories"
                    },
                    {
                        label: "add category"
                    }
                ]}
            />
            <Card>
                <ProductCategoryForm onSubmit={mutate} buttonLabel='create' />
            </Card>
        </>
    )
}

export default AddProductCategory
