import { useMutation } from "@tanstack/react-query"
import toast from "react-hot-toast"
import { useNavigate } from "react-router"
import { ADMIN_PATHS } from "@/paths"
import { products } from "@api"
import { defaultHookOptions } from "@hooks/defaults"

const TOAST_ID = "create_product"

const useCreateProduct = (opts: HookOptions = defaultHookOptions) => {
    const navigate = useNavigate()
    return useMutation({
        mutationFn: (newProduct: NewProduct) => products.create(newProduct),
        onMutate: () => {
            if (opts.showToast) {
                toast.loading("creating...", { id: TOAST_ID })
            }
        },
        onSuccess: () => {
            if (opts.showToast) {
                toast.success("created", { id: TOAST_ID })
            }
            navigate(ADMIN_PATHS.products.root)
        },
        onError: () => {
            if (opts.showToast) {
                toast.error("failed to create", { id: TOAST_ID })
            }
        }
    })
}

export default useCreateProduct
