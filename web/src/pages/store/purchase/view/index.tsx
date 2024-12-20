import { order } from "@api"
import Card from "@mui/material/Card"
import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"
import { useQuery } from "@tanstack/react-query"
import { useParams } from "react-router"

const OrderView = () => {
    const params = useParams()
    const { data } = useQuery({
        queryKey: ["OrderView", params.id],
        queryFn: async ({ queryKey }) => {
            const paramID = queryKey[1]

            if (paramID) {
                return await order.gerOrderByID(paramID)
            }
        }
    })

    console.log(data)

    return (
        <Card>
            <Stack direction='row' justifyContent='space-between'>
                <Typography>{data?.product.name}</Typography>
                <Typography color='textDisabled'>#{data?.order_id}</Typography>
            </Stack>
        </Card>
    )
}

export default OrderView
