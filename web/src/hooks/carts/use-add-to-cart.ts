import { useMutation } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { cart } from "@api";
import { defaultHookOptions } from "@hooks/defaults";

const TOAST_ID = "add_cart";

const useAddToCart = (opts: HookOptions = defaultHookOptions) => {
  return useMutation({
    mutationFn: (newCart: NewCart) => {
      return cart.create(newCart);
    },
    onMutate: () => {
      if (opts.showToast) {
        toast.loading("adding to cart...", { id: TOAST_ID });
      }
    },
    onSuccess: () => {
      if (opts.showToast) {
        toast.success("added to cart", { id: TOAST_ID });
      }
    },
    onError: () => {
      if (opts.showToast) {
        toast.error("failed to add to cart", { id: TOAST_ID });
      }
    },
  });
};

export default useAddToCart;
