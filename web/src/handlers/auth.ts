import server from "@utils/server";
import { AxiosResponse } from "axios";

export const login = async (payload: LoginRequest): Promise<string> => {
  try {
    const { data } = await server.post("/login", payload);
    return data.data;
  } catch (err) {
    const { status, data } = err as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({
          message: "not found",
          data: {
            email: "email not found",
          },
        });
      }
      case 401: {
        return Promise.reject({
          message: "invalid credentials",
          data: {
            email: "wrong email or password",
            password: "wrong email or password",
          },
        });
      }
      default: {
        return Promise.reject({
          message: "something went wrong",
        });
      }
    }
  }
};
