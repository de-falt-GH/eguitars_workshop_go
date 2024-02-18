CREATE TABLE IF NOT EXISTS personal_info (
	id SERIAL PRIMARY KEY,
	"name" text NOT NULL,
	phone_number text NOT NULL,
	email text NOT NULL
);

CREATE TABLE IF NOT EXISTS customer_rank (
	id SERIAL PRIMARY KEY,
	"name" text NOT NULL
);

CREATE TABLE IF NOT EXISTS customer (
	id SERIAL PRIMARY KEY,
	personal_info_id int REFERENCES personal_info(id),
	customer_rank_id int REFERENCES customer_rank(id),
	total_purchase int DEFAULT 0,
	notes text
);
-- ALTER TABLE customer ADD COLUMN total_purchase int DEFAULT 0;

CREATE TABLE IF NOT EXISTS master_rank (
	id SERIAL PRIMARY KEY,
	"name" text NOT NULL,
	salary int NOT NULL
);

CREATE TABLE IF NOT EXISTS "master" (
	id SERIAL PRIMARY KEY,
	personal_info_id int REFERENCES personal_info(id),
	master_rank_id int REFERENCES master_rank(id)
);

CREATE TABLE IF NOT EXISTS guitar (
	id SERIAL PRIMARY KEY,
	"name" text NOT NULL,
	condition text NOT NULL,
	serial_number text
);

CREATE TABLE IF NOT EXISTS component (
	id SERIAL PRIMARY KEY,
	type text NOT NULL,
	manufacturer text NOT NULL,
	"name" text NOT NULL,
	quantity int NOT NULL,
	UNIQUE ("type", manufacturer, "name")
);

CREATE TABLE IF NOT EXISTS order_type (
	id SERIAL PRIMARY KEY,
	description text NOT NULL
);

CREATE TABLE IF NOT EXISTS order_status (
	id SERIAL PRIMARY KEY,
	description text NOT NULL
);

CREATE TABLE IF NOT EXISTS "order" (
	id SERIAL PRIMARY KEY,
	customer_id int REFERENCES customer(id) NOT NULL,
	master_id int REFERENCES master(id) NOT NULL,
	order_status_id int REFERENCES order_status(id) NOT NULL,
	guitar_id int REFERENCES guitar(id),
	order_type_id int REFERENCES order_type(id) NOT NULL,
	price int NOT NULL,
	description text NOT NULL,
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS required_components (
	id SERIAL PRIMARY KEY,
	component_id int REFERENCES component(id),
	order_id int REFERENCES "order"(id)
);