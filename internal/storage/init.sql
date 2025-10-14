CREATE TABLE if not exists users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  password TEXT
);

CREATE TABLE if not exists peoples (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE if not exists products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price NUMERIC(10,2) NOT NULL,
  count INT NOT NULL
);

CREATE TABLE if not exists bills (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  created_by SERIAL REFERENCES users(id)
);

CREATE TABLE if not exists bill_people (
  id SERIAL PRIMARY KEY,
  bill_id SERIAL REFERENCES bills(id),
  person_id INT
);

CREATE TABLE if not exists bill_products (
  id SERIAL PRIMARY KEY,
  bill_id SERIAL REFERENCES bills(id),
  product_id SERIAL REFERENCES products(id)
);

CREATE TABLE if not exists product_assignments (
  id SERIAL PRIMARY KEY,
  product_id SERIAL REFERENCES products(id),
  person_id SERIAL REFERENCES peoples(id),
  count INT
);

CREATE TABLE if not exists debts (
  id SERIAL PRIMARY KEY,
  from_person_id SERIAL REFERENCES peoples(id),
  to_person_id SERIAL REFERENCES peoples(id),
  amount NUMERIC(10,2) NOT NULL
);

CREATE TABLE if not exists password_reset_tokens (
   token TEXT PRIMARY KEY,
   user_id SERIAL REFERENCES users(id),
   expires_at TIMESTAMP,
   used BOOLEAN DEFAULT FALSE
);
