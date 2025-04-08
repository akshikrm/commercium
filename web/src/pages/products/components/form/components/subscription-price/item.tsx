import { Stack, TextField } from "@mui/material"
import { ChangeEventHandler, ReactNode } from "react"
import { FieldError, FieldErrorsImpl, Merge } from "react-hook-form"

const PaymentIntervals = ({
    value,
    onChange,
    error = false,
    helperText = null
}: {
    value: BillingInterval
    onChange: ChangeEventHandler<HTMLInputElement>
    error: boolean
    helperText: ReactNode
}) => {
    return (
        <TextField
            label='Interval'
            select
            fullWidth
            name='interval'
            value={value}
            onChange={onChange}
            error={error}
            helperText={helperText}
            slotProps={{
                select: {
                    native: true
                }
            }}
        >
            <option value='day'>Day</option>
            <option value='month'>Month</option>
            <option value='week'>Week</option>
            <option value='year'>Year</option>
        </TextField>
    )
}

type Props = {
    subscriptionPrice: NewSubscriptionPrice
    onChange: (v: NewSubscriptionPrice) => void
    error?: Merge<
        FieldError,
        FieldErrorsImpl<{
            price: number
            label: string
            interval: NonNullable<BillingInterval>
            frequency: number
        }>
    >
}

const SubscrptionPriceForm = ({
    subscriptionPrice,
    onChange,
    error
}: Props) => {
    const { price, frequency, interval, label } = subscriptionPrice

    const handleChange: ChangeEventHandler<HTMLInputElement> = e => {
        const { name, value } = e.target
        onChange({ ...subscriptionPrice, [name]: value })
    }

    return (
        <Stack direction='row' alignItems='center'>
            <TextField
                name='price'
                value={price > 0 ? price : ""}
                label='Price'
                onChange={handleChange}
                error={Boolean(error?.price)}
                helperText={error?.price?.message}
            />
            <TextField
                name='label'
                value={label}
                label='Label'
                onChange={handleChange}
                error={Boolean(error?.label)}
                helperText={error?.label?.message}
            />
            <PaymentIntervals
                value={interval}
                onChange={handleChange}
                error={Boolean(error?.interval)}
                helperText={error?.interval?.message}
            />
            <TextField
                name='frequency'
                value={frequency}
                label='Frequency'
                onChange={handleChange}
                error={Boolean(error?.frequency)}
                helperText={error?.frequency?.message}
            />
        </Stack>
    )
}

export default SubscrptionPriceForm
