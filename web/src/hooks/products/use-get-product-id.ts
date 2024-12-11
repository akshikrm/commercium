import { useParams } from "react-router";
import { useMemo } from "react";

const useGetProductId = () => {
  const { product_id } = useParams<{ product_id: string }>();

  const productID: number = useMemo(() => {
    if (product_id) {
      return parseInt(product_id);
    }
    return 0;
  }, [product_id]);
  return productID;
};

export default useGetProductId;
