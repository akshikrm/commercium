import server from "@utils/server";
import { AxiosResponse } from "axios";

export const getProductCategories = async (
  params: string | Filter,
): Promise<ProductCategory[]> => {
  try {
    const { data } = await server.get("/products/categories", { params });
    return data.data;
  } catch (err) {
    const { data } = err as AxiosResponse;
    console.error(data);

    return Promise.reject({ message: "failed to get categories" });
  }
};

export const createProductCategory = async (
  reqData: NewProductCategory,
): Promise<EditProductCategory | undefined> => {
  try {
    const { data } = await server.post("/products/categories", reqData);
    return data.data;
  } catch (err) {
    const { status, data } = err as AxiosResponse;
    console.error(data);

    switch (status) {
      case 404: {
        return Promise.reject({ message: "product category not found" });
      }
      default: {
        return Promise.reject({ message: "failed to create category" });
      }
    }
  }
};

export const getProductCategoryByID = async (
  id: number | string,
): Promise<EditProductCategory> => {
  try {
    const { data } = await server.get(`/products/categories/${id}`);
    const serializedResponse: EditProductCategory = {
      ...data.data,
      enabled: data.data.enabled ? "enabled" : "disabled",
    };

    return serializedResponse;
  } catch (error) {
    const { status, data } = error as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product category not found" });
      }
      default: {
        return Promise.reject({
          message: `someting went wrong while getting category with id ${id}`,
        });
      }
    }
  }
};

export const updateProductCategory = async (
  id: number,
  inputData: EditProductCategory,
) => {
  try {
    const { data } = await server.put(`/products/categories/${id}`, inputData);
    return data.data;
  } catch (error) {
    const { status, data } = error as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product category not found" });
      }
      default: {
        return Promise.reject({ message: "failed to update product category" });
      }
    }
  }
};

export const productCategoryDelete = async (
  id: number,
  params: object = {},
): Promise<{ status: boolean; message: string }> => {
  try {
    const { data } = await server.delete(`/products/categories/${id}`, {
      params,
    });
    return data.data;
  } catch (err) {
    const { status, data } = err as AxiosResponse;
    console.error(data);
    switch (status) {
      case 404: {
        return Promise.reject({ message: "product category not found" });
      }
      default: {
        return Promise.reject({
          message: "someting went wrong while deleting category",
        });
      }
    }
  }
};
