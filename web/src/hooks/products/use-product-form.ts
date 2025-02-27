import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"
import { convertToCommonAmount } from "@components/prefix"

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
    subscription_price: []
}

const useProductForm = (defaultValues?: EditProduct) => {
    const methods = useForm({
        resolver: zodResolver(productSchema),
        defaultValues: newProductDefaultValues
    })

    const { reset } = methods

    useEffect(() => {
        if (defaultValues) {
            const { prices, type, subscription_price, ...rest } = defaultValues
            console.log(prices)
            const temp = { ...rest }
            if (type === "one-time") {
                temp.price = convertToCommonAmount(prices[0].price)
            }

            reset(temp)
        }
    }, [defaultValues, reset])

    return methods
}

const productSchema = z.object({
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
        .array(
            z.object({
                price: z.coerce
                    .number({
                        required_error: "price is required",
                        invalid_type_error: "price is required"
                    })
                    .gte(100, "price should be atleast 100"),
                frequency: z.coerce
                    .number({
                        required_error: "frequency is required",
                        invalid_type_error: "frequency is required"
                    })
                    .gte(1, "frequency should be atleast 1"),
                label: z
                    .string({ required_error: "label is required" })
                    .min(1, "label is required"),
                interval: z
                    .string({
                        required_error: "interval is required"
                    })
                    .min(1, "interval is required")
            })
        )
        .optional(),
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
// .refine(
//     data => {
//         if (data.type === "one-time") {
//             return data.price || 0 > 0
//         }
//         // if (data.type === "subscription") {
//         //     const price = data.subscription_price["1_month"]?.price
//         //     return parseInt(price as string) > 0
//         // }
//         return true
//     },
//     data => {
//         const isSubscription = data.type === "subscription"
//         return {
//             message: isSubscription
//                 ? "Atleast One Price is requried for subscription"
//                 : "Price is required",
//             path: [isSubscription ? "subscription_price" : "price"]
//         }
//     }
// )

export default useProductForm
