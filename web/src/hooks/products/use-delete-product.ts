import { useMutation } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { products } from "@api";
import { defaultHookOptions } from "@hooks/defaults";

const TOAST_ID = "delete_product";

const useDeleteProduct = (
  handleSuccess: () => void,
  opts: HookOptions = defaultHookOptions,
) => {
  return useMutation({
    mutationFn: (productID: number) => {
      return products.delete(productID);
    },
    onMutate: () => {
      if (opts.showToast) {
        toast.loading("deleting...", { id: TOAST_ID });
      }
    },
    onSuccess: () => {
      if (opts.showToast) {
        toast.success("deleted", { id: TOAST_ID });
      }
      handleSuccess();
    },
    onError: () => {
      if (opts.showToast) {
        toast.error("failed to delete", { id: TOAST_ID });
      }
    },
  });
};

export default useDeleteProduct;
