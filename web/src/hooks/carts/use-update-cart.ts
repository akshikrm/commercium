import { cart } from "@api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

const TOAST_ID = "update_cart";
const useUpdateCart = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (payload: UpdateCart) => cart.update(payload),

    onMutate: () => {
      toast.loading("updating....", { id: TOAST_ID });
    },
    onError: (e) => {
      toast.error(e.message, { id: TOAST_ID });
    },
    onSuccess: (data) => {
      toast.success("updated", { id: TOAST_ID });
      queryClient.setQueryData(["cartList"], (prevData: Cart[]) => {
        const temp = [...prevData];
        const itemIndex = temp.findIndex(({ id }) => id === data.id);
        if (itemIndex > -1) {
          temp.splice(itemIndex, 1, data);
        }
        return temp;
      });
    },
  });

  return mutation;
};

export default useUpdateCart;
