import { useMutation } from "@tanstack/react-query"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"
import { ADMIN_PATHS } from "@/paths"
import { products } from "@api"
import { defaultHookOptions } from "@hooks/defaults"

const TOAST_ID = "update_product"

const useUpdateProduct = (
    categoryID: number,
    opts: HookOptions = defaultHookOptions
) => {
    const navigate = useNavigate()
    return useMutation({
        mutationFn: (updated: EditProduct) => {
            return products.update(categoryID, updated)
        },
        onMutate: () => {
            if (opts.showToast) {
                toast.loading("updating...", { id: TOAST_ID })
            }
        },
        onSuccess: () => {
            if (opts.showToast) {
                toast.success("updated", { id: TOAST_ID })
            }
            navigate(ADMIN_PATHS.products.root)
        },
        onError: () => {
            if (opts.showToast) {
                toast.error("failed to update", { id: TOAST_ID })
            }
        }
    })
}

export default useUpdateProduct
