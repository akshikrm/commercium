import { files } from "@api";
import { TextField } from "@mui/material";
import { FunctionComponent } from "react";
import { Controller, useFormContext } from "react-hook-form";

type Props = {
  label: string;
  name: string;
};

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
  const { control, setValue } = useFormContext();

  const handleUpload = async (payload: FileList) => {
    if (payload.length === 1) {
      const uploaded = await files.single(payload[0]);
      setValue(name, uploaded.path);
    }
  };

  return (
    <Controller
      name={name}
      control={control}
      render={({ fieldState, field }) => {
        const { value } = field;
        const { error, invalid } = fieldState;
        console.log(value);
        return (
          <TextField
            type="file"
            onChange={async (e) => {
              const { files } = e.target as HTMLInputElement;
              if (files?.length) {
                handleUpload(files);
              }
            }}
            label={label}
            error={invalid}
            helperText={error?.message}
          />
        );
      }}
    />
  );
};

export default RHFImageUpload;
