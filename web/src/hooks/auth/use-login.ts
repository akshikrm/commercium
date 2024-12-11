import { auth } from "@api";
import { useMutation } from "@tanstack/react-query";
import { setSession } from "@utils/session";
import toast from "react-hot-toast";
import { useNavigate } from "react-router";
import useAuth from "./use-auth";
import { ADMIN_PATHS } from "@/paths";

const TOAST_ID = "login_toast";

const useLogin = () => {
  const { refresh } = useAuth();
  const navigate = useNavigate();
  const mutation = useMutation({
    mutationFn: (payload: LoginRequest) => {
      return auth.login(payload);
    },
    onMutate() {
      toast.loading("logging in...", { id: TOAST_ID });
    },
    onSuccess(token) {
      toast.success("welcome", { id: TOAST_ID });
      refresh();
      const role = setSession(token);
      if (role === "admin") {
        navigate(ADMIN_PATHS.products.root, { replace: true });
      } else {
        navigate("/", { replace: true });
      }
    },
    onError: (err: FailedResponse) => {
      toast.error(err.message, { id: TOAST_ID });
    },
  });

  return mutation;
};

export default useLogin;
