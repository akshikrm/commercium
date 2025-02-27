import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"
import { showCommonAmount } from "@components/prefix"

const newProductDefaultValues: NewProduct = {
    name: "",
    image: [],
    primary_image: "",
    slug: "",
    status: "enabled",
    type: "one-time",
    description: "",
    category_id: "",
    price: "",
    subscription_price: {
        month: {
            price: 0,
            label: ""
        }
    }
}

const useProductForm = (defaultValues?: EditProduct) => {
    const methods = useForm({
        resolver: zodResolver(productSchema),
        defaultValues: newProductDefaultValues
    })

    const { reset, setValue } = methods

    useEffect(() => {
        if (defaultValues) {
            const { price, subscription_price, ...rest } = defaultValues
            const temp: NewSubscriptionPrice = {}
            if (subscription_price) {
                Object.entries(subscription_price).forEach(([k, v]) => {
                    temp[k] = {
                        ...v,
                        price: v.price
                    }
                })
            }
            reset({
                ...rest,
                subscription_price: temp,
                price: price ? showCommonAmount(parseInt(price as string)) : ""
            })
        }
    }, [defaultValues, setValue, reset])

    return methods
}

const productSchema = z
    .object({
        name: z
            .string({
                required_error: "name is required",
                invalid_type_error: "name is required"
            })
            .min(1, { message: "name is required" }),
        image: z.array(z.string()).min(1, { message: "image is requred" }),
        primary_image: z.string().optional(),
        slug: z
            .string({
                required_error: "product url is required",
                invalid_type_error: "product url is required"
            })
            .min(1, { message: "product url is required" }),
        description: z
            .string({
                required_error: "description is required",
                invalid_type_error: "description is required"
            })
            .min(1, { message: "description is required" }),
        category_id: z.coerce
            .number({
                required_error: "category is required",
                invalid_type_error: "category is required"
            })
            .gte(1, { message: "category is required" }),
        price: z.coerce
            .number({
                required_error: "price is requred",
                invalid_type_error: "price should be a number"
            })
            .transform(v => v * 100)
            .optional(),
        subscription_price: z
            .any()
            .optional()
            .transform((v: SubscriptionPrice) => {
                const payload: SubscriptionPrice = {}
                Object.entries(v).forEach(([k, v]) => {
                    const price: number = v.price
                        ? parseInt(v.price as string)
                        : 0
                    payload[k] = { ...v, price: price * 100 }
                })
                return payload
            }),
        status: z
            .string({
                required_error: "status is required",
                invalid_type_error: "status is required"
            })
            .min(1, { message: "status is required" }),
        type: z
            .string({
                required_error: "product type is required",
                invalid_type_error: "product type is required"
            })
            .min(1, { message: "product type is required" })
    })
    .refine(
        data => {
            if (data.type === "one-time") {
                return data.price || 0 > 0
            }
            if (data.type === "subscription") {
                const price = data.subscription_price["1_month"]?.price
                return parseInt(price as string) > 0
            }
            return true
        },
        data => {
            const isSubscription = data.type === "subscription"
            return {
                message: isSubscription
                    ? "Atleast One Price is requried for subscription"
                    : "Price is required",
                path: [isSubscription ? "subscription_price" : "price"]
            }
        }
    )

export default useProductForm
