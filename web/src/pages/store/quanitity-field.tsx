import { TextField } from "@mui/material";
import { ChangeEvent, useState } from "react";
import { z, ZodError } from "zod";

type Props = {
  value: number;
  onChange: (v: number) => void;
};

const qtySchema = z.coerce
  .number({
    required_error: "quantity is required",
    invalid_type_error: "quantity should be a number",
  })
  .int({ message: "quantity should be a number" })
  .gt(0, { message: "quantity should be a number greater than 0" });

const QuantityField = ({ value, onChange }: Props) => {
  const [errors, setErrors] = useState<string | null>(null);
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    try {
      const parsed = qtySchema.parse(e.target.value);
      onChange(parsed);
      setErrors(null);
    } catch (err) {
      const error = err as ZodError;
      setErrors(error.flatten().formErrors[0]);
    }
  };

  // console.log(qtySchema.parse("asldkfj"));

  return (
    <TextField
      label="quantity"
      type="number"
      value={value}
      onChange={handleChange}
      error={Boolean(errors)}
      helperText={errors}
      slotProps={{
        htmlInput: {
          min: 1,
        },
      }}
    />
  );
};

export default QuantityField;
