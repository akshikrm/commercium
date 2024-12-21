type Order = {
    id: number
    purchase_price: number
    order_id: string
    products: {
        id: number
        name: string
        price: number
    }[]
    created_at: string
}

type OrderView = {
    id: number
    purchase_price: number
    order_id: string
    product: {
        id: number
        name: string
        slug: string
        image: string
        description: string
        category: {
            id: number
            name: string
            slug: string
        }
    }
    created_at: string
}
