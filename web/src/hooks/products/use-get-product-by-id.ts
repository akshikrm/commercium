import { products } from "@api";
import { defaultHookOptions } from "@hooks/defaults";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import toast from "react-hot-toast";

const TOAST_ID = "get_product_by_id";

const useGetProductById = (
  productID: number,
  opts: HookOptions = defaultHookOptions,
) => {
  const query = useQuery({
    queryKey: ["productByID", productID],
    queryFn: ({ queryKey }) => products.getById(queryKey[1]),
  });

  const { status, error } = query;
  useEffect(() => {
    if (opts.showToast) {
      switch (status) {
        case "error": {
          toast.error(error.message, {
            id: TOAST_ID,
          });
          break;
        }
        case "success": {
          toast.success("product retrieved", {
            id: TOAST_ID,
          });
          break;
        }
        case "pending": {
          toast.loading("loading...", { id: TOAST_ID });
        }
      }
    }
  }, [status, error, opts.showToast]);

  return query;
};

export default useGetProductById;
