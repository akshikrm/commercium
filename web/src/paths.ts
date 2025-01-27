export const ADMIN_PATHS = {
    root: "/admin",
    products: {
        root: "/admin/products",
        add: "/admin/products/add",
        edit: (slug: number | string) => `/admin/products/${slug}`,
        categories: {
            root: "/admin/products/categories",
            add: "/admin/products/categories/add",
            edit: (slug: number | string) =>
                `/admin/products/categories/${slug}`
        }
    },
    orders: {
        root: "/admin/orders",
        view: (orderID: string) => `admin/orders/${orderID}`
    }
}

export const USER_PATHS = {
    root: "/",
    store: {
        root: "/stores"
    },
    cart: { root: "/carts" },
    orders: {
        root: "/orders",
        view: (orderID: string) => `/orders/${orderID}`
    }
}

export const AUTH_PATHS = {
    login: {
        root: "/auth/login"
    }
}
