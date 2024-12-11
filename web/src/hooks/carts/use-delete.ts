import { cart } from "@api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

const TOAST_ID = "delete_cart";

const useDeleteCart = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (cartId: number) => cart.delete(cartId),

    onMutate: () => {
      toast.loading("deleting....", { id: TOAST_ID });
    },
    onError: (e) => {
      toast.error(e.message, { id: TOAST_ID });
    },
    onSuccess: (data, vars) => {
      toast.success(data, { id: TOAST_ID });
      queryClient.setQueryData(["cartList"], (prevData: Cart[]) => {
        const temp = [...prevData];
        const itemIndex = temp.findIndex(({ id }) => id === vars);
        if (itemIndex > -1) {
          temp.splice(itemIndex, 1);
        }
        return temp;
      });
    },
  });

  return mutation;
};

export default useDeleteCart;
