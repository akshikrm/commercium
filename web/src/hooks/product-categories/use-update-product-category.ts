import { useMutation } from "@tanstack/react-query"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"
import { ADMIN_PATHS } from "@/paths"
import { productCategories } from "@api"

const TOAST_ID = "update_product_categories"

const useUpdateProductCategory = (categoryID: number) => {
    const navigate = useNavigate()
    return useMutation({
        mutationFn: (updated: EditProductCategory) => {
            return productCategories.update(categoryID, updated)
        },
        onMutate: () => {
            toast.loading("updating...", { id: TOAST_ID })
        },
        onSuccess: () => {
            toast.success("updated", { id: TOAST_ID })
            navigate(ADMIN_PATHS.products.categories.root)
        },
        onError: () => {
            toast.error("failed to update", { id: TOAST_ID })
        }
    })
}

export default useUpdateProductCategory
