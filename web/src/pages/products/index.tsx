import { Link } from "react-router"
import { useState } from "react"
import ProductFilter from "./components/filter"
import DeleteProduct from "./components/delete"
import HeaderBreadcrumbs from "@components/header"
import useGetProducts from "@hooks/products/use-get-products"
import List from "./components/list"
import { Button } from "@mui/material"
import dayjs from "dayjs"
import useSelectOne from "@hooks/use-select-one"
import { ADMIN_PATHS } from "@/paths"

const Products = () => {
    const [filter, setFilter] = useState<Filter>({
        start_date: dayjs().startOf("M").toISOString(),
        end_date: dayjs().endOf("M").toISOString()
    })
    const { data: products, refetch } = useGetProducts(filter)
    const [productIdToDelete, handleOpenDelete, handleCloseDelete] =
        useSelectOne()

    return (
        <>
            <HeaderBreadcrumbs
                heading='Products'
                links={[
                    {
                        label: "home",
                        href: ADMIN_PATHS.root
                    },
                    {
                        label: "products"
                    }
                ]}
                action={
                    <Button
                        component={Link}
                        to={ADMIN_PATHS.products.add}
                        color='primary'
                    >
                        add
                    </Button>
                }
            />

            <ProductFilter
                defaultValues={filter}
                filter={async (inputData: Filter) => {
                    setFilter(inputData)
                }}
            />

            <DeleteProduct
                selectedID={productIdToDelete}
                onClose={handleCloseDelete}
                reload={async () => {
                    await refetch()
                }}
            />

            <List products={products} onDelete={handleOpenDelete} />
        </>
    )
}

export default Products
