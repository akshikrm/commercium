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
import {
    getByOrderID,
    getInvoiceURI,
    getOrders,
    getShippingInformation,
    isOrderComplete,
    placeOrder,
    updateShippingStatus
} from "./orders.ts"
import { getCustomerId, profile } from "./user"

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
    profile: profile,
    getCustomerId: getCustomerId
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
    getMyOrders: getOrders,
    gerOrderByID: getByOrderID,
    getInvoiceURI: getInvoiceURI,
    isOrderComplete: isOrderComplete,
    getShippingInformation: getShippingInformation,
    updateShippingStatus: updateShippingStatus
}
