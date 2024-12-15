import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"

const newProductDefaultValues: NewProductCategory = {
    name: "",
    slug: "",
    description: "",
    enabled: "enabled"
}

const schema = z.object({
    name: z.string().min(1, { message: "name is required" }),
    slug: z.string().min(1, { message: "slug is required" }),
    description: z.string().min(1, { message: "description is required" }),
    enabled: z.string().transform(v => {
        return v === "enabled"
    })
})

const useProductCategoryForm = (defaultValues?: EditProductCategory) => {
    const methods = useForm<NewProductCategory>({
        defaultValues: newProductDefaultValues,
        resolver: zodResolver(schema)
    })

    const { reset } = methods
    useEffect(() => {
        if (defaultValues) {
            reset(defaultValues)
        }
    }, [defaultValues, reset])

    return methods
}

export default useProductCategoryForm
