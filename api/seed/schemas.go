package main

var SCHEMA = map[string]string{
	"roles":              "id SERIAL PRIMARY KEY, name VARCHAR(20) NOT NULL, code VARCHAR(10) UNIQUE NOT NULL, description VARCHAR(120) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	"users":              "id SERIAL PRIMARY KEY, customer_id VARCHAR(30) NOT NULL, password VARCHAR NOT NULL, role_code VARCHAR(10) DEFAULT user NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_role FOREIGN KEY(role_code) REFERENCES roles(code)",
	"profiles":           "id SERIAL PRIMARY KEY, user_id int UNIQUE, first_name VARCHAR(50) DEFAULT '' NOT NULL, last_name VARCHAR(50) DEFAULT '' NOT NULL, email VARCHAR(50) UNIQUE DEFAULT '' NOT NULL, pincode VARCHAR(10) DEFAULT '' NOT NULL, address_one VARCHAR(100) DEFAULT '' NOT NULL, address_two VARCHAR(100) DEFAULT '' NOT NULL, phone_number VARCHAR(15) DEFAULT '' NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)",
	"product_categories": "id SERIAL PRIMARY KEY, name VARCHAR(30) NOT NULL, slug VARCHAR(30) NOT NULL,enabled BOOLEAN DEFAULT true, description VARCHAR(120) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",

	"products": `id SERIAL PRIMARY KEY, product_id VARCHAR(30) NOT NULL,
	name VARCHAR(50),
	type product_type NOT NULL DEFAULT 'one-time',
	status status NOT NULL DEFAULT 'disabled', 
	slug VARCHAR(50), price INTEGER NOT NULL DEFAULT 0,
	price_id VARCHAR(30) NOT NULL, image VARCHAR[],  description VARCHAR(300) NOT NULL,
	category_id INTEGER NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL,
	CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES product_categories(id)`,

	"product_images": "id SERIAL PRIMARY KEY, product_id INTEGER NOT NULL, URI TEXT NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)",
	"prices": `
		id SERIAL PRIMARY KEY,
		product_id INTEGER NOT NULL,
		price INTEGER NOT NULL DEFAULT 0,
		price_id varchar(30),
		label varchar(30),
		interval varchar(30),
		created_at TIMESTAMP DEFAULT NOW() NOT NULL,
		updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
		deleted_at TIMESTAMP DEFAULT NULL
	`,
	"carts": `
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		price_id INTEGER NOT NULL,
		quantity INTEGER DEFAULT 1 NOT NULL,
		created_at TIMESTAMP DEFAULT NOW() NOT NULL,
		updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
		deleted_at TIMESTAMP DEFAULT NULL,
		CONSTRAINT 
			fk_user FOREIGN KEY(user_id) REFERENCES users(id),
		CONSTRAINT 
			fk_price FOREIGN KEY(price_id) REFERENCES prices(id)
	`,
	"orders": `id SERIAL PRIMARY KEY, 
	user_id INTEGER,
	order_id VARCHAR(20),
	transaction_id INT NOT NULL,
	quantity INT NOT NULL,
	price_id varchar(30) NOT NULL,
	product_id varchar(30) NOT NULL,
	amount varchar(30) NOT NULL,
	status varchar(20) NOT NULL DEFAULT 'pending',
	price INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
	deleted_at TIMESTAMP DEFAULT NULL`,

	"purchases": `id SERIAL PRIMARY KEY, 
	product_id INTEGER NOT NULL, order_id INTEGER NOT NULL, 
	quantity INTEGER DEFAULT 1 NOT NULL, price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, 
	CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id), CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)`,
	"uploads": "id SERIAL PRIMARY KEY, path VARCHAR(200) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	"transactions": `id SERIAL PRIMARY KEY,
	transaction_id VARCHAR(30) NOT NULL,
	invoice_number VARCHAR(20) NOT NULL DEFAULT '',
	status VARCHAR(30) NOT NULL,
	customer_id VARCHAR(30) NOT NULL DEFAULT '',
	tax  varchar(20) NOT NULL,
	sub_total  varchar(20) NOT NULL,
	grand_total  varchar(20) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
	`,
}

var KEYS = []string{
	"roles",
	"users",
	"profiles",
	"product_categories",
	"products",
	"prices",
	"product_images",
	"carts",
	"orders",
	"purchases",
	"uploads",
	"transactions",
}
