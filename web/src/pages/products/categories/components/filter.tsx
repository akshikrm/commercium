import RHFProvider from "@/components/rhf/provider";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Box, Button, Card, Stack } from "@mui/material";
import RHFDatePicker from "@/components/rhf/date-picker";

const schema = z.object({
  start_date: z.string(),
  end_date: z.string(),
});

type Props = {
  filter: (inputData: Filter) => Promise<void>;
  defaultFilter: Filter;
};

const CategoryFilter = ({ filter, defaultFilter }: Props) => {
  const methods = useForm<Filter>({
    defaultValues: defaultFilter,
    resolver: zodResolver(schema),
  });

  return (
    <Card sx={{ mb: 5 }}>
      <RHFProvider methods={methods} onSubmit={filter}>
        <Stack direction="row" alignItems="center">
          <RHFDatePicker label="start date" name="start_date" />
          <RHFDatePicker label="end date" name="end_date" />
          <Box>
            <Button type="submit" variant="contained">
              filter
            </Button>
          </Box>
        </Stack>
      </RHFProvider>
    </Card>
  );
};
export default CategoryFilter;
