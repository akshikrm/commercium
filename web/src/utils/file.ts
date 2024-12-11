import axios, { AxiosError } from "axios";
import { BASE_URL } from "@config";
import toast from "react-hot-toast";

const fileServer = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "multipart/form-data",
  },
});

const TOAST_ID = "network_error";
fileServer.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    const { response, code } = error as AxiosError<{
      error: string;
    }>;
    if (code === "ERR_NETWORK") {
      toast.error(
        "there seems to be some problem with your network connection, trying again...",
        { id: TOAST_ID },
      );
    }

    return Promise.reject(response);
  },
);

export default fileServer;
