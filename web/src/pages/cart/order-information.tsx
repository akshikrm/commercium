import { Currency } from "@components/prefix"
import { Button, Card, CardContent, Typography } from "@mui/material"
import RenderIcon from "@components/render-icon"
import useGetCustomerID from "@hooks/users/use-get-customer-id"
import { useEffect } from "react"

type Props = {
    total: number
}

const OrderInformation = ({ total }: Props) => {
    // const mutation = usePlaceOrder()
    const { customerID, getCustomerID } = useGetCustomerID()

    useEffect(() => {
        if (customerID) {
            console.log(customerID)
        }
    }, [customerID])

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
                    getCustomerID()
                }}
            >
                complete your order
            </Button>
        </>
    )
}

export default OrderInformation
