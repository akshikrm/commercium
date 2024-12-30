import useGetCart from "@hooks/carts/use-get-cart"
import { Grid2 as Grid } from "@mui/material"
import HeaderBreadcrumbs from "@components/header"
import CartItemList from "./cart-list-item"
import OrderInformation from "./order-information"

const Cart = () => {
    const { data: carts, total, paddlePurchaseItems } = useGetCart()

    return (
        <>
            <HeaderBreadcrumbs
                heading='Cart'
                links={[{ label: "home", href: "/" }]}
            />

            <Grid container spacing={4}>
                <Grid
                    size={{ md: 8 }}
                    sx={{
                        maxHeight: "60vh",
                        overflow: "scroll"
                    }}
                >
                    <CartItemList data={carts || []} />
                </Grid>

                <Grid size={{ md: 4 }}>
                    <OrderInformation
                        total={total || 0}
                        paddlePurchaseItems={paddlePurchaseItems}
                    />
                </Grid>
            </Grid>
        </>
    )
}

export default Cart
