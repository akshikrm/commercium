type Purchases = {
    id: number
    purchase_price: number
    order_id: string
    product: {
        id: number
        name: string
        price: number
    }
    created_at: string
}
