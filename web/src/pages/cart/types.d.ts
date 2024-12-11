type NewCart = {
  product_id: number;
  quantity: number;
};

type UpdateCart = {
  cartID: number;
  quantity: number;
};

type Cart = {
  id: number;
  user_id: number;
  quantity: number;
  created_at: string;
  product: {
    id: number;
    name: string;
    slug: string;
    price: number;
    description: string;
    image: string;
  };
};
