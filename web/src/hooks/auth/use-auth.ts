import { authenticationContext } from "@/store/auth";
import { useContext } from "react";

const useAuth = () => useContext(authenticationContext);

export default useAuth;
