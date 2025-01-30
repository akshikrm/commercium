import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"

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
        .gte(1, { message: "price should be greater than zero" }),
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

const newProductDefaultValues: NewProduct = {
    name: "",
    image: [],
    primary_image: "",
    slug: "",
    status: "enabled",
    type: "one-time",
    description: "",
    category_id: "",
    price: ""
}

const useProductForm = (defaultValues?: EditProduct) => {
    const methods = useForm({
        resolver: zodResolver(productSchema),
        defaultValues: newProductDefaultValues
    })

    const { reset } = methods

    useEffect(() => {
        if (defaultValues) {
            reset(defaultValues)
        }
    }, [defaultValues, reset])

    return methods
}

export default useProductForm
