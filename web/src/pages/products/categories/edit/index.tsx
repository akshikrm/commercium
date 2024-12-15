import HeaderBreadcrumbs from "@components/header"
import { Card } from "@mui/material"
import ProductCategoryForm from "../components/form"
import useGetProductCategoryID from "@hooks/product-categories/use-get-product-category-id"
import useGetProductCategoryByID from "@hooks/product-categories/use-get-product-category-by-id"
import useUpdateProductCategory from "@hooks/product-categories/use-update-product-category"

const EditProductCategory = () => {
    const categoryID = useGetProductCategoryID()
    const { data: category } = useGetProductCategoryByID(categoryID)
    const { mutate } = useUpdateProductCategory(categoryID)
    return (
        <>
            <HeaderBreadcrumbs
                heading='Edit Category'
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
                        label: "edit category"
                    }
                ]}
            />

            <Card>
                <ProductCategoryForm
                    defaultValues={category}
                    onSubmit={mutate}
                    buttonLabel='update'
                />
            </Card>
        </>
    )
}

export default EditProductCategory
