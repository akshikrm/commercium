import { AUTH_PATHS } from "@/paths";
import { clearSession } from "@utils/session";
import { useNavigate } from "react-router";

const useLogout = () => {
  const navigate = useNavigate();
  const logout = () => {
    clearSession();
    navigate(AUTH_PATHS.login.root, {
      replace: true,
    });
  };

  return logout;
};

export default useLogout;
