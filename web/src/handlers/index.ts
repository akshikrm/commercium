import { login } from "./auth"
import { addToCart, deleteCart, getCart, updateCart } from "./cart"
import { uploadSingleFile } from "./file"
import {
    createProductCategory,
    getProductCategories,
    getProductCategoryByID,
    productCategoryDelete,
    updateProductCategory
} from "./product-categories"

import {
    createProduct,
    getProductById,
    getProducts,
    productDelete,
    updateProduct
} from "./products"
import { getPurcahses, placeOrder } from "./purchase"
import { profile } from "./user"

export const products = {
    getAll: getProducts,
    getById: getProductById,
    create: createProduct,
    update: updateProduct,
    delete: productDelete
}

export const productCategories = {
    getAll: getProductCategories,
    getById: getProductCategoryByID,
    create: createProductCategory,
    update: updateProductCategory,
    delete: productCategoryDelete
}

export const auth = {
    login: login
}

export const user = {
    profile: profile
}

export const files = {
    single: uploadSingleFile
}

export const cart = {
    create: addToCart,
    getAll: getCart,
    update: updateCart,
    delete: deleteCart
}

export const order = {
    placeOrder: placeOrder,
    getMyOrders: getPurcahses
}
