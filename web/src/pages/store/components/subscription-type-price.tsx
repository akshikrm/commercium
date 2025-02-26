import { Stack, TextField, Typography } from "@mui/material"
import RenderList from "@components/render-list"
import { Currency } from "@components/prefix"

const SubscriptionTypePrice = ({
    data,
    value,
    onChange
}: {
    data: Prices[]
    value: NewCartEntry
    onChange: (v: NewCartEntry) => void
}) => {
    return (
        <Stack
            direction='row'
            justifyContent='space-between'
            alignItems='center'
        >
            <TextField
                select
                fullWidth
                slotProps={{
                    select: {
                        native: true
                    }
                }}
                onChange={e => {
                    const { id, price } = JSON.parse(e.target.value)
                    onChange({
                        price_id: id,
                        price,
                        quantity: 1
                    })
                }}
            >
                <RenderList
                    list={data}
                    render={(price: Prices) => {
                        return (
                            <option
                                key={price.id}
                                value={JSON.stringify({
                                    id: price.id,
                                    price: price.price
                                })}
                            >
                                {price.label}
                            </option>
                        )
                    }}
                />
            </TextField>
            <Typography>
                <Currency amount={value.price} />
            </Typography>
        </Stack>
    )
}

export default SubscriptionTypePrice
