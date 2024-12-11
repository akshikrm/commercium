import RHFProvider from "@/components/rhf/provider";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Box, Card, Stack } from "@mui/material";
import RHFDatePicker from "@/components/rhf/date-picker";
import { LoadingButton } from "@mui/lab";
import ProductCategoryNames from "@components/product-category-names";

const schema = z.object({
  category_id: z
    .string()
    .optional()
    .transform((v) => (v === "all" ? null : v)),
  start_date: z.string(),
  end_date: z.string(),
});

type Props = {
  filter: (inputData: Filter) => Promise<void>;
  defaultValues: Filter;
};

const ProductFilter = ({ filter, defaultValues }: Props) => {
  const methods = useForm<Filter>({
    defaultValues,
    resolver: zodResolver(schema),
  });

  return (
    <Card sx={{ mb: 5 }}>
      <RHFProvider methods={methods} onSubmit={filter}>
        <Stack direction="row" alignItems="center">
          <RHFDatePicker label="start date" name="start_date" />
          <RHFDatePicker label="end date" name="end_date" />
          <ProductCategoryNames
            customOption={<option value="all">All</option>}
          />
          <Box>
            <LoadingButton
              loading={methods.formState.isSubmitting}
              type="submit"
              variant="contained"
            >
              filter
            </LoadingButton>
          </Box>
        </Stack>
      </RHFProvider>
    </Card>
  );
};
export default ProductFilter;
