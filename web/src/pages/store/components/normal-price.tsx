import { Currency } from "@components/prefix"
import { Stack, Typography } from "@mui/material"
import QuantityField from "./quanitity-field"

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

export default NormalPrice
