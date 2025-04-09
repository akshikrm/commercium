import { ADMIN_PATHS, USER_PATHS } from "@/paths"

export const USER_MENUS: Paths[] = [
    {
        label: "dashboard",
        path: USER_PATHS.root
    },
    {
        label: "e-commerce",
        path: USER_PATHS.root,
        children: [
            {
                label: "product",
                path: USER_PATHS.store.root
            },
            {
                label: "subscription",
                path: USER_PATHS.subscriptions.root
            },
            {
                label: "orders",
                path: USER_PATHS.orders.root
            }
        ]
    }
]

export const ADMIN_MENUS: Paths[] = [
    {
        label: "dashboard",
        path: ADMIN_PATHS.root
    },
    {
        label: "e-commerce",
        path: ADMIN_PATHS.products.root,
        children: [
            {
                label: "product",
                path: ADMIN_PATHS.products.root
            },
            {
                label: "shipping",
                path: ADMIN_PATHS.shipping.root
            }
        ]
    },
    {
        label: "sales",
        path: ADMIN_PATHS.orders.root,
        children: [
            {
                label: "orders",
                path: ADMIN_PATHS.orders.root
            }
        ]
    }
]
