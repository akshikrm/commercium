import { useMutation } from "@tanstack/react-query"
import toast from "react-hot-toast"
import { cart } from "@api"
import { defaultHookOptions } from "@hooks/defaults"
import { useNavigate } from "react-router"
import { USER_PATHS } from "@/paths"
import { useState } from "react"

const TOAST_ID = "add_cart"

const useAddToCart = (opts: HookOptions = defaultHookOptions) => {
    const [isBuyNow, setIsBuyNow] = useState(false)
    const navigate = useNavigate()

    const mutation = useMutation({
        mutationFn: (payload: NewCartEntry) => cart.create(payload),
        onMutate: () => {
            if (opts.showToast) {
                toast.loading("adding to cart...", { id: TOAST_ID })
            }
        },
        onSuccess: () => {
            if (opts.showToast) {
                toast.success("added to cart", { id: TOAST_ID })
            }
            if (isBuyNow) {
                navigate(USER_PATHS.cart.root)
            }
        },
        onError: () => {
            if (opts.showToast) {
                toast.error("failed to add to cart", { id: TOAST_ID })
            }
        }
    })

    const addToCart = (payload: NewCartEntry) => mutation.mutate(payload)
    const buyNow = (payload: NewCartEntry) => {
        setIsBuyNow(true)
        mutation.mutate(payload)
    }

    return { addToCart, buyNow }
}

export default useAddToCart
