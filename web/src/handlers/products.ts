import server from "@/utils/server";
import { AxiosResponse } from "axios";

export const getProducts = async (
  params: string | Filter,
): Promise<Product[]> => {
  try {
    const { data } = await server.get("/products", { params });
    return data.data;
  } catch (err) {
    const { data } = err as AxiosResponse;
    console.error(data);
    return Promise.reject({ message: "failed to get products" });
  }
};

export const getProductById = async (
  id: number | string,
): Promise<EditProduct | undefined> => {
  try {
    const { data } = await server.get(`/products/${id}`);
    return data.data;
  } catch (error) {
    const { status, data } = error as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product not found" });
      }
      default: {
        return Promise.reject({
          message: `something went wrong while getting product with id ${id}`,
        });
      }
    }
  }
};

export const createProduct = async (reqData: NewProduct) => {
  try {
    const { data } = await server.post("/products", reqData);
    return data.data;
  } catch (err) {
    const { data } = err as AxiosResponse;
    console.error(data);
    return Promise.reject({ messate: "failed to create product" });
  }
};

export const updateProduct = async (
  productID: number,
  inputData: EditProduct,
) => {
  try {
    const { data } = await server.put(`/products/${productID}`, inputData);
    return data.data;
  } catch (error) {
    const { status, data } = error as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product not found" });
      }
      default: {
        return Promise.reject({ message: "failed to update product" });
      }
    }
  }
};

export const productDelete = async (
  id: number,
  params: object = {},
): Promise<NewProduct | undefined> => {
  try {
    const { data } = await server.delete(`/products/${id}`, { params });
    return data.data;
  } catch (error) {
    const { status, data } = error as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product not found" });
      }
      default: {
        return Promise.reject({ message: "failed to delete product" });
      }
    }
  }
};
