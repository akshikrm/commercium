type ShippingStatus = "delivered" | "pending" | "in-transit"
type ShippingInformation = {
    id: number
    status: ShippingStatus
    amount: number
    quantity: number
    user: {
        id: number
        name: string
        email: string
    }
    product: {
        id: number
        name: string
    }
    created_at: time
}
