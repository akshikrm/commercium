import { Currency } from "@components/prefix"
import {
    Button,
    Card,
    CardActions,
    CardContent,
    CardMedia,
    Grid2 as Grid,
    Stack,
    TextField,
    Typography
} from "@mui/material"
import { useState } from "react"
import QuantityField from "./quanitity-field"
import icons from "@/icons"
import RenderIcon from "@components/render-icon"
import { Cloudinary } from "@cloudinary/url-gen"
import { scale } from "@cloudinary/url-gen/actions/resize"
import Render from "@components/render"
import RenderList from "@components/render-list"

type Props = {
    product: Product
    addToCart: (payload: NewCartEntry) => void
    buyNow: (payload: NewCartEntry) => void
}

const ProductItem = ({ product, addToCart, buyNow }: Props) => {
    const { id, name, description, image, prices, type } = product
    const cld = new Cloudinary({
        cloud: { cloudName: "commercium" }
    })

    const img = cld.image(image).resize(scale().width(100).height(100))
    const [quantity, setQuanitity] = useState<number>(1)

    const isSubscriptionType = type === "subscription"

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
                image={img.toURL()}
                title='green iguana'
                onError={e => {
                    e.target.onerror = null
                    e.target.src = "https://placehold.co/400@2x.png"
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
                <Render
                    when={isSubscriptionType}
                    show={<SubscriptionTypePrice data={prices} />}
                    otherwise={
                        <NormalPrice
                            quantity={quantity}
                            price={prices[0].price}
                            onChange={v => setQuanitity(v)}
                        />
                    }
                />
            </CardContent>
            <CardActions>
                <Button
                    startIcon={<RenderIcon icon={icons.addToCart} />}
                    onClick={() => {
                        if (isSubscriptionType) {
                            return
                        }
                        addToCart({ price_id: prices[0].id, quantity })
                    }}
                >
                    add to cart
                </Button>
                <Button
                    color='success'
                    onClick={() => {
                        if (isSubscriptionType) {
                            return
                        }
                        buyNow({ price_id: prices[0].id, quantity })
                    }}
                    startIcon={<RenderIcon icon={icons.buyNow} />}
                >
                    buy now
                </Button>
            </CardActions>
        </Grid>
    )
}

const SubscriptionTypePrice = ({ data }: { data: Prices[] }) => {
    return (
        <TextField
            select
            fullWidth
            slotProps={{
                select: {
                    native: true
                }
            }}
        >
            <RenderList
                list={data}
                render={(price: Prices) => {
                    return <option value={price.price}>{price.label}</option>
                }}
            />
        </TextField>
    )
}

const NormalPrice = ({
    quantity,
    price,
    onChange
}: {
    quantity: number
    price: number
    onChange: (v: number) => void
}) => {
    return (
        <Stack
            direction='row'
            alignItems='center'
            justifyContent='space-between'
        >
            <QuantityField value={quantity} onChange={onChange} />
            <Typography variant='body1'>
                <Currency amount={quantity * price} />
            </Typography>
        </Stack>
    )
}

export default ProductItem
