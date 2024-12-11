/* eslint-disable  @typescript-eslint/no-explicit-any */
import { FieldValues, FormProvider, UseFormReturn } from "react-hook-form";
import { ReactNode } from "react";

type OnSubmit<T> = (inputData: T) => void;
type AsyncOnSubmit<T> = (inputData: T) => Promise<void>;

type Props<T extends FieldValues> = {
  methods: UseFormReturn<T, any, undefined>;
  onSubmit: AsyncOnSubmit<T> | OnSubmit<T>;
  children: ReactNode;
};

const RHFProvider = <T extends FieldValues>({
  methods,
  onSubmit,
  children,
}: Props<T>) => {
  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>{children}</form>
    </FormProvider>
  );
};

export default RHFProvider;
