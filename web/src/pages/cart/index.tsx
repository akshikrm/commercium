import { Currency } from "@components/prefix";
import RenderList from "@components/render-list";
import { BASE_URL_FILE } from "@config";
import useDeleteCart from "@hooks/carts/use-delete";
import useGetCart from "@hooks/carts/use-get-cart";
import useUpdateCart from "@hooks/carts/use-update-cart";
import {
  Button,
  Card,
  CardContent,
  Container,
  Grid2 as Grid,
  IconButton,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import icons from "@/icons";
import RenderIcon from "@components/render-icon";
import HeaderBreadcrumbs from "@components/header";

const Cart = () => {
  const { data: carts } = useGetCart();
  const { mutate: update } = useUpdateCart();
  const { mutate: deleteCart } = useDeleteCart();

  const handleUpdate = (payload: UpdateCart) => {
    update(payload);
  };

  const handleDelete = (payload: number) => {
    deleteCart(payload);
  };

  return (
    <>
      <HeaderBreadcrumbs
        heading="Cart"
        links={[{ label: "home", href: "/" }]}
      />

      <Container maxWidth="md" component={Grid} container spacing={4}>
        <Grid
          size={{ md: 8 }}
          sx={{
            maxHeight: "60vh",
            overflow: "scroll",
          }}
        >
          <Stack>
            <RenderList
              list={carts}
              render={(cart) => {
                const { id, product, quantity } = cart;
                return (
                  <Card key={id}>
                    <Stack direction="row" alignItems="center">
                      <img
                        src={[BASE_URL_FILE, product.image].join("/")}
                        width={100}
                      />
                      <Stack spacing={0}>
                        <Typography component="div" variant="subtitle2">
                          {product.name}
                        </Typography>
                        <Typography variant="caption">
                          {product.description.slice(0, 20)}...
                        </Typography>
                      </Stack>
                      <Typography variant="body2">
                        <TextField
                          size="small"
                          value={quantity}
                          type="number"
                          onChange={(e) => {
                            handleUpdate({
                              quantity: parseInt(e.target.value),
                              cartID: id,
                            });
                          }}
                        />
                      </Typography>
                      <Typography variant="body2">
                        <Currency>{product.price * quantity}</Currency>
                      </Typography>
                      <IconButton
                        size="small"
                        color="error"
                        onClick={() => handleDelete(id)}
                      >
                        <RenderIcon icon={icons.delete} />
                      </IconButton>
                    </Stack>
                  </Card>
                );
              }}
            />
          </Stack>
        </Grid>

        <Grid size={{ md: 4 }}>
          <Card sx={{ mb: 3 }}>
            <CardContent>
              <Typography variant="h6" fontWeight="bold">
                Total
              </Typography>
              <Typography variant="body1">
                <Currency>{total}</Currency>
              </Typography>
            </CardContent>
          </Card>

          <Button
            color="secondary"
            fullWidth
            startIcon={<RenderIcon icon="mdi:cash-fast" />}
          >
            complete your order
          </Button>
        </Grid>
      </Container>
    </>
  );
};

export default Cart;
