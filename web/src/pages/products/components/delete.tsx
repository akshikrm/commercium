import useDeleteProduct from "@hooks/products/use-delete-product";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Typography,
} from "@mui/material";

type Props = {
  selectedID: number | null;
  onClose: () => void;
  reload: () => Promise<void>;
};

const DeleteProduct = ({ selectedID, onClose, reload }: Props) => {
  const { mutate } = useDeleteProduct(() => {
    onClose();
    reload();
  });

  const handleDelete = async () => {
    if (selectedID) {
      mutate(selectedID);
    }
  };

  return (
    <Dialog open={Boolean(selectedID)} onClose={onClose}>
      <DialogTitle>delete product</DialogTitle>
      <DialogContent>
        <Typography>
          are you sure you want to continue this action cannot be reversed
        </Typography>
      </DialogContent>
      <DialogActions>
        <Button color="error" onClick={handleDelete}>
          confirm
        </Button>
        <Button color="warning" onClick={onClose}>
          cancel
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default DeleteProduct;
