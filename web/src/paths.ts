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
    }
}

export const AUTH_PATHS = {
    login: {
        root: "/auth/login"
    }
}
