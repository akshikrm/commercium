type Order = {
    id: number
    transaction_id: string
    invoice_number: string
    total: number
    payment_status: string
    products: OrderItems[]
    created_at: string
}

type OrderItems = {
    id: number
    product_id: number
    name: string
    price: string
    quantity: number
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
