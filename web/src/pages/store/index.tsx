import HeaderBreadcrumbs from "@components/header"
import RenderList from "@components/render-list"
import useAddToCart from "@hooks/carts/use-add-to-cart"
import useGetProducts from "@hooks/products/use-get-products"
import { Grid2 as Grid } from "@mui/material"
import ProductItem from "./components/product-item"

const Store = () => {
    const { data: products } = useGetProducts()
    const { addToCart, buyNow } = useAddToCart()

    return (
        <>
            <HeaderBreadcrumbs
                heading='Store'
                links={[{ label: "home", href: "/" }, { label: "store" }]}
            />
            <Grid container spacing={2} alignItems='stretch'>
                <RenderList
                    list={products}
                    render={product => (
                        <ProductItem
                            product={product}
                            addToCart={addToCart}
                            buyNow={buyNow}
                        />
                    )}
                />
            </Grid>
        </>
    )
}

export default Store
