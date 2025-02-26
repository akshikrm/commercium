type NewCartEntry = {
    price_id: number
    quantity: number
    price: number
}

type CreateCartPayload = {
    price_id: number
    quantity: number
}

type UpdateCart = {
    cartID: number
    quantity: number
}

type Cart = {
    id: number
    user_id: number
    quantity: number
    created_at: string
    price_id: string
    price: number
    product: {
        id: number
        name: string
        slug: string
        description: string
        image: string
    }
}

type PaddlePurchaseItem = {
    priceId: string
    quantity: number
}
