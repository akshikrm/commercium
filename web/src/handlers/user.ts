import server from "@utils/server";
import { AxiosResponse } from "axios";

export const profile = async (): Promise<Profile> => {
  try {
    const { data } = await server.get("/profile");
    return data.data;
  } catch (err) {
    const { data } = err as AxiosResponse;
    console.error(data);
    return Promise.reject({ message: "failed to get products" });
  }
};
