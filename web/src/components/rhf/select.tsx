import { Controller, useFormContext } from "react-hook-form"
import { FunctionComponent, ReactNode } from "react"
import { TextField } from "@mui/material"

type Props = {
    label: string
    name: string
    children: ReactNode
}

const RHFSelect: FunctionComponent<Props> = ({ label, name, children }) => {
    const { control } = useFormContext()
    return (
        <Controller
            name={name}
            control={control}
            render={({ field, fieldState }) => {
                const { onChange, name, value } = field
                const { error, invalid } = fieldState
                return (
                    <TextField
                        label={label}
                        select
                        slotProps={{
                            select: {
                                native: true
                            }
                        }}
                        name={name}
                        onChange={onChange}
                        value={value}
                        error={invalid}
                        helperText={error?.message}
                    >
                        {children}
                    </TextField>
                )
            }}
        />
    )
}

export default RHFSelect
