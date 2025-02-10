import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"

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
            .optional(),
        subscription_price: z
            .any()
            .optional()
            .transform((v: SubscriptionPrice) => {
                return Object.entries(v).map(([k, v]) => {
                    const price = v.price ? parseInt(v.price) : 0
                    return { [k]: { ...v, price } }
                })
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
                return data.subscription_price[0]["1_month"]?.price > 0
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

const newProductDefaultValues: NewProduct = {
    name: "",
    image: ["006_dk6ugp"],
    primary_image: "",
    slug: "",
    status: "enabled",
    type: "one-time",
    description: "",
    category_id: "",
    price: "",
    subscription_price: {
        "1_month": {
            price: "",
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
            reset({
                ...rest,
                subscription_price: subscription_price,
                price: price ? price : ""
            })
        }
    }, [defaultValues, setValue, reset])

    return methods
}

export default useProductForm
