type EnabledStatus = "enabled" | "disabled";

type NewProductCategory = {
  name: string;
  slug: string;
  description: string;
  enabled: EnabledStatus;
};

type EditProductCategory = {
  name?: string;
  slug?: string;
  description?: string;
  enabled?: EnabledStatus;
};

type ProductCategory = {
  id: number;
  name: string;
  slug: string;
  description: string;
  enabled: boolean;
  created_at: string;
  updated_at: string;
  deleted_at: string;
};

type ProductCategoryNames = {
  id: number;
  name: string;
  slug: string;
};
