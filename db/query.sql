CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar NOT NULL,
  image varchar NOT NULL,
  category varchar NOT NULL,
  description text,
  rating int NOT NULL,
  num_reviews int NOT NULL DEFAULT 0,
  price decimal(10,2) NOT NULL,
  count_in_stock int NOT NULL,
  created_at bigint NOT NULL DEFAULT EXTRACT (EPOCH FROM CURRENT_TIMESTAMP),
  updated_at bigint NOT NULL DEFAULT EXTRACT (EPOCH FROM CURRENT_TIMESTAMP)
);

CREATE TABLE orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_method varchar NOT NULL,
  tax_price decimal(10,2) NOT NULL,
  shipping_price decimal(10,2) NOT NULL,
  total_price decimal(10,2) NOT NULL,
  created_at bigint NOT NULL DEFAULT EXTRACT (EPOCH FROM CURRENT_TIMESTAMP),
  updated_at bigint NOT NULL DEFAULT EXTRACT (EPOCH FROM CURRENT_TIMESTAMP)
);

CREATE TABLE order_items (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id UUID NOT NULL,
  product_id UUID NOT NULL,
  name varchar NOT NULL,
  quantity int NOT NULL,
  image varchar NOT NULL,
  price int NOT NULL
);

ALTER TABLE order_items ADD FOREIGN KEY (order_id) REFERENCES orders (id);
ALTER TABLE order_items ADD FOREIGN KEY (product_id) REFERENCES products (id);