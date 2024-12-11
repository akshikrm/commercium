import { Controller, useFormContext } from "react-hook-form";
import { FunctionComponent, HTMLInputTypeAttribute } from "react";
import { TextField } from "@mui/material";

type Props = {
  label: string;
  name: string;
  type?: HTMLInputTypeAttribute;
};

const RHFTextField: FunctionComponent<Props> = ({
  label,
  name,
  type = "text",
}) => {
  const { control } = useFormContext();
  return (
    <Controller
      name={name}
      control={control}
      render={({ field, fieldState }) => {
        const { onChange, name, value } = field;
        const { error, invalid } = fieldState;
        return (
          <TextField
            label={label}
            name={name}
            onChange={onChange}
            value={value}
            type={type}
            error={invalid}
            helperText={error?.message}
          />
        );
      }}
    />
  );
};

export default RHFTextField;
