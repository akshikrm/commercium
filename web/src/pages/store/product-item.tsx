import { Currency } from "@components/prefix"
import { BASE_URL_FILE } from "@config"
import {
    Button,
    Card,
    CardActions,
    CardContent,
    CardMedia,
    Grid2 as Grid,
    Stack,
    Typography
} from "@mui/material"
import { useState } from "react"
import QuantityField from "./quanitity-field"
import icons from "@/icons"
import RenderIcon from "@components/render-icon"

type Props = {
    product: Product
    addToCart: (payload: NewCart) => void
}
const ProductItem = ({ product, addToCart }: Props) => {
    const { id, name, description, image, price } = product
    const convertedURI = [BASE_URL_FILE, image].join("/")

    const [quantity, setQuanitity] = useState<number>(1)

    return (
        <Grid
            size={{ sm: 6, md: 4 }}
            component={Card}
            key={id}
            sx={{
                padding: 0,
                display: "flex",
                flexDirection: "column",
                justifyContent: "space-between"
            }}
        >
            <CardMedia
                sx={{ height: 200 }}
                component='img'
                image={convertedURI}
                title='green iguana'
                onError={e => {
                    e.target.onerror = null
                    e.target.rc = "https://placehold.co/400@2x.png"
                }}
            />
            <CardContent>
                <Typography
                    variant='h6'
                    component='div'
                    sx={{ whiteSpace: "nowrap" }}
                >
                    {name}
                </Typography>
                <Typography>{description}</Typography>
                <Stack
                    direction='row'
                    alignItems='center'
                    justifyContent='space-between'
                >
                    <QuantityField
                        value={quantity}
                        onChange={v => setQuanitity(v)}
                    />
                    <Typography variant='body1'>
                        <Currency>{quantity * price}</Currency>
                    </Typography>
                </Stack>
            </CardContent>
            <CardActions>
                <Button
                    startIcon={<RenderIcon icon={icons.addToCart} />}
                    onClick={() => {
                        addToCart({ product_id: id, quantity })
                    }}
                >
                    add to cart
                </Button>
                <Button
                    color='success'
                    startIcon={<RenderIcon icon={icons.buyNow} />}
                >
                    buy now
                </Button>
            </CardActions>
        </Grid>
    )
}

export default ProductItem
