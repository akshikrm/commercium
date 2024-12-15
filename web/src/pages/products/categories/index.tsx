import { Button } from "@mui/material"
import { Link } from "react-router"
import CategoryFilter from "./components/filter"
import DeleteProductCategory from "./components/delete"
import { useState } from "react"
import HeaderBreadcrumbs from "@components/header"
import List from "./components/list"
import dayjs from "dayjs"
import useSelectOne from "@hooks/use-select-one"
import useGetProductCategories from "@hooks/product-categories/use-get-product-categories"
import { ADMIN_PATHS } from "@/paths"

const ProductCategories = () => {
    const [filter, setFilter] = useState<Filter>({
        start_date: dayjs().startOf("M").toISOString(),
        end_date: dayjs().endOf("M").toISOString()
    })

    const { data: categories, refetch } = useGetProductCategories(filter, {
        onlyErrorToast: true
    })

    const [selectedID, handleSetSelectedID, handleUnsetSelectedID] =
        useSelectOne()

    return (
        <>
            <HeaderBreadcrumbs
                heading='Categories'
                links={[
                    {
                        label: "home",
                        href: ADMIN_PATHS.root
                    },
                    {
                        label: "products",
                        href: ADMIN_PATHS.products.root
                    },
                    {
                        label: "categories"
                    }
                ]}
                action={
                    <Button
                        component={Link}
                        to={ADMIN_PATHS.products.categories.add}
                        color='primary'
                    >
                        add
                    </Button>
                }
            />

            <DeleteProductCategory
                selectedID={selectedID}
                onClose={handleUnsetSelectedID}
                reload={async () => {
                    await refetch()
                }}
            />

            <CategoryFilter
                defaultFilter={filter}
                filter={async (inputData: Filter) => {
                    setFilter(inputData)
                }}
            />
            <List categories={categories} onDelete={handleSetSelectedID} />
        </>
    )
}

export default ProductCategories
