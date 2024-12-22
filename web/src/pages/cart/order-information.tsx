import { Currency } from "@components/prefix"
import { Button, Card, CardContent, Typography } from "@mui/material"
import RenderIcon from "@components/render-icon"
import useGetCustomerID from "@hooks/users/use-get-customer-id"
import { useEffect } from "react"
import { initializePaddle, Paddle } from "@paddle/paddle-js"

type Props = {
    total: number
    paddlePurchaseItems: PaddlePurchaseItem[]
}

const handleConnectPaddle = async (
    customerID: string
): Promise<Paddle | undefined> => {
    try {
        return await initializePaddle({
            token: "test_f17f04fa410842584b9d336943d",
            environment: "sandbox",
            pwCustomer: {
                id: customerID
            }
        })
    } catch (err) {
        console.log(err)
        return undefined
    }
}

const OrderInformation = ({ total, paddlePurchaseItems }: Props) => {
    // const mutation = usePlaceOrder()
    const { customerID, getCustomerID } = useGetCustomerID()
    // const [purchase, setPurchase] = useState<Paddle | undefined>(undefined)

    useEffect(() => {
        if (customerID) {
            handleConnectPaddle(customerID)
                .then(data => {
                    data?.Checkout.open({
                        settings: {
                            displayMode: "overlay",
                            variant: "one-page"
                        },
                        items: paddlePurchaseItems,
                        customer: {
                            id: customerID
                        }
                    })
                })
                .catch(err => {
                    console.log(err)
                })
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
