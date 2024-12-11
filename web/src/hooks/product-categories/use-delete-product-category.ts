import { useMutation } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { productCategories } from "@api";

const TOAST_ID = "delete_product_categories";

const useDeleteProductCategory = (handleSuccess: () => void) => {
  return useMutation({
    mutationFn: (categoryID: number) => {
      return productCategories.delete(categoryID);
    },
    onMutate: () => {
      toast.loading("deleting...", { id: TOAST_ID });
    },
    onSuccess: () => {
      toast.success("deleted", { id: TOAST_ID });
      handleSuccess();
    },
    onError: () => {
      toast.error("failed to delete", { id: TOAST_ID });
    },
  });
};

export default useDeleteProductCategory;
