import { useMutation } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { useNavigate } from "react-router";
import { ADMIN_PATHS } from "@/paths";
import { productCategories } from "@api";

const TOAST_ID = "create_product_category";

const useCreateProductCategory = () => {
  const navigate = useNavigate();
  return useMutation({
    mutationFn: (newProductCategory: NewProductCategory) => {
      return productCategories.create(newProductCategory);
    },
    onMutate: () => {
      toast.loading("creating...", { id: TOAST_ID });
    },
    onSuccess: () => {
      toast.success("created", { id: TOAST_ID });
      navigate(ADMIN_PATHS.products.categories.root);
    },
    onError: () => {
      toast.error("failed to create", { id: TOAST_ID });
    },
  });
};

export default useCreateProductCategory;
