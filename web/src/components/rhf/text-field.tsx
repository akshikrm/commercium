import { Controller, useFormContext } from "react-hook-form"
import { FunctionComponent } from "react"
import { TextField, TextFieldProps } from "@mui/material"

const RHFTextField: FunctionComponent<
    TextFieldProps & { name: string }
> = props => {
    const { control } = useFormContext()

    return (
        <Controller
            name={props.name}
            control={control}
            render={({ field, fieldState }) => {
                const { onChange, value } = field
                const { error, invalid } = fieldState
                return (
                    <TextField
                        onChange={onChange}
                        value={value}
                        error={invalid}
                        helperText={error?.message}
                        {...props}
                    />
                )
            }}
        />
    )
}

export default RHFTextField
