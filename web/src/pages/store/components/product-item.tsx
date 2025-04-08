import {
    Button,
    Card,
    CardActions,
    CardContent,
    CardMedia,
    Grid2 as Grid,
    Typography
} from "@mui/material"
import { useEffect, useMemo, useState } from "react"
import icons from "@/icons"
import RenderIcon from "@components/render-icon"
import { Cloudinary } from "@cloudinary/url-gen"
import { scale } from "@cloudinary/url-gen/actions/resize"
import Render from "@components/render"
import NormalPrice from "./normal-price"
import SubscriptionTypePrice from "./subscription-type-price"

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
    const isSubscriptionType: boolean = useMemo(
        () => type === "subscription",
        [type]
    )

    const [price, setPrice] = useState<NewCartEntry>({
        price_id: 0,
        price: 0,
        quantity: 1
    })

    useEffect(() => {
        setPrice({
            price_id: prices[0].id,
            price: prices[0].price,
            quantity: 1
        })
    }, [prices])

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
                    show={
                        <SubscriptionTypePrice
                            data={prices}
                            value={price}
                            onChange={v => {
                                setPrice(v)
                            }}
                        />
                    }
                    otherwise={
                        <NormalPrice
                            quantity={price.quantity}
                            price={price.price}
                            onChange={v =>
                                setPrice(prev => {
                                    return {
                                        ...prev,
                                        quantity: v
                                    }
                                })
                            }
                        />
                    }
                />
            </CardContent>
            <CardActions>
                <Button
                    startIcon={<RenderIcon icon={icons.addToCart} />}
                    onClick={() => addToCart(price)}
                >
                    add to cart
                </Button>
                <Button
                    color='success'
                    onClick={() => buyNow(price)}
                    startIcon={<RenderIcon icon={icons.buyNow} />}
                >
                    buy now
                </Button>
            </CardActions>
        </Grid>
    )
}

export default ProductItem
