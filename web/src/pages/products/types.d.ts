type NewProduct = {
  name: string;
  image: string;
  slug: string;
  description: string;
  category_id: string;
  price: string;
};

type EditProduct = {
  name?: string;
  image?: string;
  slug?: string;
  description?: string;
  price?: string;
  category_id?: string;
};

type Filter = {
  [key: string]: string | null;
  start_date?: string | null;
  end_date?: string | null;
  category_id?: string | null;
};

type Product = {
  id: number;
  name: string;
  image: string;
  slug: string;
  description: string;
  price: number;
};
