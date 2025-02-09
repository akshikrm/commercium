type ProductStatus = "enabled" | "disabled"
type ProductType = "one-time" | "subscription"

type SubscriptionPrice = {
    [key: string]: {
        id: number
        price: number
        label: string
        price_id: string
    }
}
type Price = { id: string; price: number; label: string; label: string }

type NewProduct = {
    name: string
    primary_image: string
    image: string[]
    slug: string
    status: ProductStatus
    type: ProductType
    description: string
    category_id: string
    price?: string | null
    prices: Price[]
    subscription_price?: SubscriptionPrice | null
}

type EditProduct = {
    name?: string
    image?: string[]
    primary_image: string
    slug?: string
    status?: ProductStatus
    type?: ProductType
    description?: string
    price?: string | null
    subscription_price?: SubscriptionPrice | null
    category_id?: string
}

type Filter = {
    [key: string]: string | null
    start_date?: string | null
    end_date?: string | null
    category_id?: string | null
}

type Product = {
    id: number
    name: string
    image: string
    slug: string
    description: string
    price: number
}
