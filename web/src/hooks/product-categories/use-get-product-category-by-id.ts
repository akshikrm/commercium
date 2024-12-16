import { productCategories } from "@api"
import { useQuery } from "@tanstack/react-query"
import { useEffect } from "react"
import toast from "react-hot-toast"

const TOAST_ID = "get_product_category_by_id"

const useGetProductCategoryByID = (productCategoryID: number) => {
    const query = useQuery({
        queryKey: ["productCategoryByID", productCategoryID],
        queryFn: ({ queryKey }) => productCategories.getById(queryKey[1]),
        staleTime: Infinity
    })

    const { status, error } = query

    useEffect(() => {
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
    }, [status, error])

    return query
}

export default useGetProductCategoryByID
