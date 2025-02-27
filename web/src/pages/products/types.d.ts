type ProductStatus = "enabled" | "disabled"
type ProductType = "one-time" | "subscription"

type SubscriptionPrice = {
    [key: string]: {
        id?: number
        price: number
        label: string
        price_id?: string
    }
}

type BillingInterval = "day" | "week" | "month" | "year"

type NewSubscriptionPrice = {
    [interval: BillingInterval]: {
        price: number
        label: string
    }
}

type NewProduct = {
    name: string
    primary_image: string
    image: string[]
    slug: string
    status: ProductStatus
    type: ProductType
    description: string
    category_id: string
    price?: string
    subscription_price?: NewSubscriptionPrice
}

type EditProduct = {
    name?: string
    image?: string[]
    primary_image: string
    slug?: string
    status?: ProductStatus
    type?: ProductType
    description?: string
    price?: string
    subscription_price?: SubscriptionPrice
    category_id?: string
}

type Filter = {
    [key: string]: string | null
    start_date?: string | null
    end_date?: string | null
    category_id?: string | null
}

type Prices = {
    id: number
    price: number
    label: string
    price_id: string
}

type Product = {
    id: number
    name: string
    image: string
    slug: string
    type: ProductType
    description: string
    prices: Prices[]
}
