import { Controller, useFormContext } from "react-hook-form"
import { FunctionComponent } from "react"
import { DatePicker } from "@mui/x-date-pickers"
import dayjs from "dayjs"

type Props = {
    label: string
    name: string
}

const RHFDatePicker: FunctionComponent<Props> = ({ label, name }) => {
    const { control, setValue } = useFormContext()
    return (
        <Controller
            name={name}
            control={control}
            render={({ field }) => {
                const { name, value } = field
                return (
                    <DatePicker
                        name={name}
                        label={label}
                        value={dayjs(value)}
                        onChange={newValue =>
                            setValue(name, newValue?.toISOString())
                        }
                    />
                )
            }}
        />
    )
}

export default RHFDatePicker
