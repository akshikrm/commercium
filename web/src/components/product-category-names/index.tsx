import { FunctionComponent, ReactNode } from "react"
import RHFSelect from "@/components/rhf/select"
import useGetProductCategories from "@hooks/product-categories/use-get-product-categories"

type Props = {
    customOption?: ReactNode | null
}

export const ProductCategoryNames: FunctionComponent<Props> = ({
    customOption
}) => {
    const { data } = useGetProductCategories({ type: "name" })

    return (
        <RHFSelect label='Category' name='category_id'>
            {customOption ? customOption : <option value='' />}
            {data?.map(({ id, name }) => {
                return (
                    <option value={id} key={id}>
                        {name}
                    </option>
                )
            })}
        </RHFSelect>
    )
}

export default ProductCategoryNames
