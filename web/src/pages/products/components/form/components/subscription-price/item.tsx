import { Box, IconButton, Stack } from "@mui/material"
import RHFTextField from "@components/rhf/text-field"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import { useCallback } from "react"
import { useFormContext } from "react-hook-form"

const SubscriptionPriceItem = ({ month }: { month: string }) => {
    const [monthNumber] = month.split("_")
    const { setValue, getValues } = useFormContext()

    const handleDelete = useCallback(() => {
        const temp = { ...getValues("subscriptionPrice") }
        delete temp[month]
        setValue("subscriptionPrice", temp)
    }, [month])

    return (
        <Stack direction='row' alignItems='center'>
            <RHFTextField
                name={`subscriptionPrice.${month}.value`}
                label={`Price ${monthNumber} month`}
            />
            <RHFTextField
                name={`subscriptionPrice.${month}.label`}
                label={`Label ${monthNumber} month`}
            />
            <Box>
                <IconButton
                    disabled={monthNumber === "1"}
                    color='error'
                    size='small'
                    onClick={() => handleDelete()}
                >
                    <RenderIcon icon={icons.delete} />
                </IconButton>
            </Box>
        </Stack>
    )
}

export default SubscriptionPriceItem
