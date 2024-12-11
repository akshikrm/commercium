import { Button, Stack } from "@mui/material";
import RHFProvider from "@components/rhf/provider";
import RHFTextField from "@components/rhf/text-field";
import useLoginForm from "@hooks/auth/use-login-form";

type Props = {
  onSubmit: (inputData: LoginRequest) => void;
  error: ValidationErrors;
};

const LoginForm = ({ onSubmit, error }: Props) => {
  const methods = useLoginForm(error);

  return (
    <RHFProvider onSubmit={onSubmit} methods={methods}>
      <Stack spacing={2}>
        <RHFTextField label="Email" name="email" />
        <RHFTextField label="Password" name="password" type="password" />
        <div className="mt-5">
          <Button fullWidth type="submit">
            login
          </Button>
        </div>
      </Stack>
    </RHFProvider>
  );
};

export default LoginForm;
