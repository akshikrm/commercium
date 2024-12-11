import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";
import useDeleteProductCategory from "@hooks/product-categories/use-delete-product-category";

type Props = {
  onClose: () => void;
  selectedID: number | null;
  reload: () => Promise<void>;
};

const DeleteProductCategoryDialog = ({
  selectedID,
  onClose,
  reload,
}: Props) => {
  const { mutate } = useDeleteProductCategory(() => {
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
      <DialogTitle>Delete Product Category</DialogTitle>
      <DialogContent>
        <DialogContentText>
          are you sure you want to continue, this action cannot be reversed
        </DialogContentText>
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

export default DeleteProductCategoryDialog;
