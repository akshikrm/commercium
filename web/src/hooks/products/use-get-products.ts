import { useEffect } from "react"
import { useQuery } from "@tanstack/react-query"
import { products } from "@api"
import toast from "react-hot-toast"
import { defaultHookOptions } from "@hooks/defaults"

const TOAST_ID = "products_get_toast"

const useGetProducts = (
    filter?: Filter,
    opts: HookOptions = defaultHookOptions
) => {
    const query = useQuery({
        queryKey: ["productsList", filter || {}],
        queryFn: ({ queryKey }) => products.getAll(queryKey[1])
    })

    const { status, error } = query

    useEffect(() => {
        if (opts.showToast) {
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
    }, [status, error, opts.showToast])

    return query
}

export default useGetProducts
