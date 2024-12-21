import { Currency } from "@components/prefix"
import { Button, Card, CardContent, Typography } from "@mui/material"
import RenderIcon from "@components/render-icon"
import usePlaceOrder from "@hooks/orders/use-place-order"

type Props = {
    total: number
}
const OrderInformation = ({ total }: Props) => {
    const mutation = usePlaceOrder()
    return (
        <>
            <Card sx={{ mb: 3 }}>
                <CardContent>
                    <Typography variant='h6' fontWeight='bold'>
                        Total
                    </Typography>
                    <Typography variant='body1'>
                        <Currency>{total}</Currency>
                    </Typography>
                </CardContent>
            </Card>

            <Button
                color='secondary'
                fullWidth
                startIcon={<RenderIcon icon='mdi:cash-fast' />}
                onClick={() => {
                    mutation.mutate()
                }}
            >
                complete your order
            </Button>
        </>
    )
}

export default OrderInformation
