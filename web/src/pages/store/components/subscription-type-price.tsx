import { TextField } from "@mui/material"
import RenderList from "@components/render-list"

const SubscriptionTypePrice = ({
    data,
    onChange
}: {
    data: Prices[]
    onChange: (v: NewCartEntry) => void
}) => {
    return (
        <TextField
            select
            fullWidth
            slotProps={{
                select: {
                    native: true
                }
            }}
            onChange={e => {
                onChange({
                    price_id: parseInt(e.target.value),
                    quantity: 1
                })
            }}
        >
            <RenderList
                list={data}
                render={(price: Prices) => {
                    return (
                        <option key={price.id} value={price.id}>
                            {price.price} - {price.label}
                        </option>
                    )
                }}
            />
        </TextField>
    )
}

export default SubscriptionTypePrice
