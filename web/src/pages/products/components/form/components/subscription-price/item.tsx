import { Stack } from "@mui/material"
import RHFTextField from "@components/rhf/text-field"

const SubscriptionPriceItem = ({ month }: { month: string }) => {
    const [monthNumber] = month.split("_")
    return (
        <Stack direction='row'>
            <RHFTextField
                name={`price.${month}.value`}
                label={`Price ${monthNumber} month`}
            />
            <RHFTextField
                name={`price.${month}.label`}
                label={`Label ${monthNumber} month`}
            />
        </Stack>
    )
}

export default SubscriptionPriceItem
