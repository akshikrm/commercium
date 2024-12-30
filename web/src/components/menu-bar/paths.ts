import { USER_PATHS } from "@/paths"

const menus: Paths[] = [
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
                label: "orders",
                path: USER_PATHS.orders.root
            }
        ]
    }
]

export default menus
