import { Card } from "@mui/material";
import HeaderBreadcrumbs from "@components/header";
import ProductForm from "../components/form";
import useCreateProduct from "@hooks/products/use-create-product";

const AddProduct = () => {
  const { mutate } = useCreateProduct();
  return (
    <>
      <HeaderBreadcrumbs
        heading="Add Product"
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
            label: "add",
          },
        ]}
      />
      <Card>
        <ProductForm onSubmit={mutate} buttonLabel="create" />
      </Card>
    </>
  );
};

export default AddProduct;
