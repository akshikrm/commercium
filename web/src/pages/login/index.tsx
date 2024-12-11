import { Card } from "@mui/material";
import LoginForm from "./components/form";
import useLogin from "@hooks/auth/use-login";

const Login = () => {
  const { mutate, error } = useLogin();
  return (
    <Card>
      <LoginForm onSubmit={mutate} error={error?.data || null} />
    </Card>
  );
};

export default Login;
