import HeaderBreadcrumbs from "@components/header";
import ProductForm from "../components/form";
import useGetProductId from "@hooks/products/use-get-product-id";
import useGetProductById from "@hooks/products/use-get-product-by-id";
import useUpdateProduct from "@hooks/products/use-update-product";
import { Card } from "@mui/material";

const EditProduct = () => {
  const productID: number = useGetProductId();
  const { data: product } = useGetProductById(productID);
  const { mutate } = useUpdateProduct(productID);

  return (
    <>
      <HeaderBreadcrumbs
        heading="Edit Product"
        links={[
          {
            label: "home",
            href: "/",
          },
          {
            label: "products",
            href: "/admin/products",
          },
          {
            label: "edit",
          },
        ]}
      />
      <Card>
        <ProductForm
          buttonLabel="update"
          defaultValues={product}
          onSubmit={mutate}
        />
      </Card>
    </>
  );
};

export default EditProduct;
