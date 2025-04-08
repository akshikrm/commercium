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
    price: number
    label: string
    interval: BillingInterval
    frequency: number
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
    price: string
    subscription_price: NewSubscriptionPrice[]
}

type OneProductResponse = {
    id: number
    category_id: number
    product_id: number
    status: string
    type: string
    name: string
    slug: string
    prices: NewSubscriptionPrice[]
    image: string[]
    description: string
    created_at: string
    updated_at: string
    deleted_at: string
}

type EditProduct = {
    name?: string
    image?: string[]
    primary_image: string
    slug?: string
    status?: ProductStatus
    type?: ProductType
    description?: string
    prices?: NewSubscriptionPrice[]
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
