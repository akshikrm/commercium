import { ADMIN_PATHS, AUTH_PATHS } from "@/paths";
import { decodeJWT, getToken } from "@utils/session";
import { Navigate, Outlet } from "react-router";

export const AdminGuard = () => {
  const token = getToken();
  if (!token) {
    return <Navigate to={AUTH_PATHS.login.root} />;
  }

  const { role } = decodeJWT(token);
  if (role !== "admin") {
    return <Navigate to={"/"} />;
  }

  return <Outlet />;
};

export const UserGuard = () => {
  const token = getToken();
  if (!token) {
    return <Navigate to={AUTH_PATHS.login.root} />;
  }

  const { role } = decodeJWT(token);
  if (role !== "user") {
    return <Navigate to={ADMIN_PATHS.products.root} />;
  }

  return <Outlet />;
};

export const AuthGuard = () => {
  const token = getToken();

  if (token) {
    return <Navigate to={ADMIN_PATHS.products.root} />;
  }

  return <Outlet />;
};
