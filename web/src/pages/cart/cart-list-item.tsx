import { Currency } from "@components/prefix"
import RenderList from "@components/render-list"
import useDeleteCart from "@hooks/carts/use-delete"
import useUpdateCart from "@hooks/carts/use-update-cart"
import { Card, IconButton, Stack, TextField, Typography } from "@mui/material"
import icons from "@/icons"
import RenderIcon from "@components/render-icon"
import { Cloudinary } from "@cloudinary/url-gen"
import { scale } from "@cloudinary/url-gen/actions/resize"

const CartItemList = ({ data }: { data: Cart[] }) => {
    const { mutate: update } = useUpdateCart()
    const { mutate: deleteCart } = useDeleteCart()

    const handleUpdate = (payload: UpdateCart) => {
        update(payload)
    }

    const handleDelete = (payload: number) => {
        deleteCart(payload)
    }

    const cld = new Cloudinary({
        cloud: { cloudName: "commercium" }
    })

    return (
        <Stack>
            <RenderList
                list={data}
                render={cart => {
                    const { id, product, quantity } = cart
                    const totalAmount = parseFloat(product.price) * quantity
                    const img = cld
                        .image(product.image)
                        .resize(scale().width(100).height(100))

                    return (
                        <Card key={id}>
                            <Stack direction='row' alignItems='center'>
                                <img src={img.toURL()} width={100} />
                                <Stack spacing={0}>
                                    <Typography
                                        component='div'
                                        variant='subtitle2'
                                    >
                                        {product.name}
                                    </Typography>
                                    <Typography variant='caption'>
                                        {product.description.slice(0, 20)}
                                        ...
                                    </Typography>
                                </Stack>
                                <Typography variant='body2'>
                                    <TextField
                                        size='small'
                                        value={quantity}
                                        type='number'
                                        onChange={e => {
                                            handleUpdate({
                                                quantity: parseInt(
                                                    e.target.value
                                                ),
                                                cartID: id
                                            })
                                        }}
                                    />
                                </Typography>
                                <Typography variant='body2'>
                                    <Currency amount={totalAmount} />
                                </Typography>
                                <IconButton
                                    size='small'
                                    color='error'
                                    onClick={() => handleDelete(id)}
                                >
                                    <RenderIcon icon={icons.delete} />
                                </IconButton>
                            </Stack>
                        </Card>
                    )
                }}
            />
        </Stack>
    )
}

export default CartItemList
