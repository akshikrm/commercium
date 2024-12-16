import { useEffect } from "react"
import { useQuery } from "@tanstack/react-query"
import toast from "react-hot-toast"
import { productCategories } from "@api"
import { defaultHookOptions } from "@hooks/defaults"

const TOAST_ID = "product_categories_get_toast"

const useGetProductCategories = (
    filter: Filter,
    opts: HookOptions = defaultHookOptions
) => {
    const query = useQuery({
        queryKey: ["productCategoryList", filter],
        queryFn: ({ queryKey }) => productCategories.getAll(queryKey[1])
    })

    opts = { ...defaultHookOptions, ...opts }
    const { status, error } = query
    useEffect(() => {
        if (opts.showToast) {
            if (opts.onlyErrorToast) {
                if (status === "error") {
                    toast.error(error.message, { id: TOAST_ID })
                }
                return
            }
            switch (status) {
                case "error": {
                    toast.error(error.message, { id: TOAST_ID })
                    break
                }
                case "success": {
                    toast.dismiss(TOAST_ID)
                    break
                }
                case "pending": {
                    toast.loading("loading...", { id: TOAST_ID })
                    break
                }
            }
        }
    }, [status, error, opts.showToast, opts.onlyErrorToast])

    return query
}
export default useGetProductCategories
