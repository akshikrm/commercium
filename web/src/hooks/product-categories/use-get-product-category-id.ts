import { useParams } from "react-router"
import { useMemo } from "react"

const useGetProductCategoryID = () => {
    const { category_id } = useParams<{ category_id: string }>()

    const productID: number = useMemo(() => {
        if (category_id) {
            return parseInt(category_id)
        }
        return 0
    }, [category_id])
    return productID
}

export default useGetProductCategoryID
