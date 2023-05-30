CREATE DATABASE company_logistics;


CREATE TABLE IF NOT EXISTS clients (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  street_address VARCHAR(255) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS truck_deliveries (
  id SERIAL PRIMARY KEY,
  client_id INTEGER REFERENCES clients(id) ON DELETE CASCADE,
  product_type VARCHAR(255) NOT NULL,
  quantity INTEGER NOT NULL CHECK (quantity >= 1),
  registration_date TIMESTAMPTZ NOT NULL,
  delivery_date TIMESTAMPTZ NOT NULL,
  warehouse VARCHAR(255) NOT NULL,
  shipping_price DECIMAL NOT NULL CHECK (shipping_price >= 0),
  discounted_price DECIMAL NOT NULL CHECK (discounted_price >= 0),
  vehicle_plate VARCHAR(6) NOT NULL CHECK (vehicle_plate ~ '^[A-Z]{3}[0-9]{3}$'),
  guide_number VARCHAR(10) NOT NULL CHECK (LENGTH(guide_number) = 10)
);

CREATE TABLE IF NOT EXISTS ship_deliveries (
  id SERIAL PRIMARY KEY,
  client_id INTEGER REFERENCES clients(id) ON DELETE CASCADE,
  product_type VARCHAR(255) NOT NULL,
  quantity INTEGER NOT NULL CHECK (quantity >= 1),
  registration_date TIMESTAMPTZ NOT NULL,
  delivery_date TIMESTAMPTZ NOT NULL,
  port VARCHAR(255) NOT NULL,
  shipping_price DECIMAL NOT NULL CHECK (shipping_price >= 0),
  discounted_price DECIMAL NOT NULL CHECK (discounted_price >= 0),
  fleet_number VARCHAR(10) NOT NULL CHECK (LENGTH(fleet_number) = 10),
  guide_number VARCHAR(10) NOT NULL CHECK (LENGTH(guide_number) = 10)
);
