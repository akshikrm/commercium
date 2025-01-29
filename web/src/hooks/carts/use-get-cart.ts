import { useEffect, useMemo } from "react"
import { useQuery } from "@tanstack/react-query"
import { cart } from "@api"
import toast from "react-hot-toast"
import { defaultHookOptions } from "@hooks/defaults"

const TOAST_ID = "carts_get_toast"

const useGetCart = (opts: HookOptions = defaultHookOptions) => {
    const query = useQuery({
        initialData: [],
        queryKey: ["cartList"],
        queryFn: () => cart.getAll()
    })

    const { status, error, data: carts } = query

    const total: number = useMemo(() => {
        if (!carts) return 0
        if (carts.length === 0) {
            return 0
        }

        return carts.reduce((acc, curr) => {
            acc += parseFloat(curr.product.price) * curr.quantity
            return acc
        }, 0)
    }, [carts])

    const paddlePurchaseItems: PaddlePurchaseItem[] = useMemo(() => {
        if (carts?.length) {
            return carts.map(cart => {
                return {
                    priceId: cart.price_id,
                    quantity: cart.quantity
                }
            })
        }

        return []
    }, [carts])

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

    return { ...query, total, paddlePurchaseItems }
}

export default useGetCart
