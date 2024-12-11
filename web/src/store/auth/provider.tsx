import { user } from "@api";
import { useQuery } from "@tanstack/react-query";
import { ReactNode, useEffect } from "react";
import { authenticationContext } from "./index";
import { getToken } from "@utils/session";

type Props = {
  children: ReactNode;
};

const AuthenticationProvider = ({ children }: Props) => {
  const query = useQuery<Profile>({
    queryKey: ["profile-key"],
    queryFn: user.profile,
    enabled: false,
    initialData: { role: null, first_name: "", last_name: "" },
  });

  const { refetch, data } = query;

  useEffect(() => {
    const token = getToken();
    if (token) {
      refetch();
    }
  }, [refetch]);

  return (
    <authenticationContext.Provider
      value={{
        user: data,
        refresh: () => {
          refetch();
        },
      }}
    >
      {children}
    </authenticationContext.Provider>
  );
};

export default AuthenticationProvider;
