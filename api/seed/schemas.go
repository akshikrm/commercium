package main

var SCHEMA = map[string]string{
	"roles": "id SERIAL PRIMARY KEY, name VARCHAR(20) NOT NULL, code VARCHAR(10) UNIQUE NOT NULL, description VARCHAR(120) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	// "resources":          "id SERIAL PRIMARY KEY, name VARCHAR(10) NOT NULL, code VARCHAR(10) UNIQUE NOT NULL, description VARCHAR(120) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	// "permissions":        "id SERIAL PRIMARY KEY, role_code VARCHAR(10) NOT NULL, resource_code VARCHAR(10) NOT NULL, r BOOLEAN DEFAULT false NOT NULL, w BOOLEAN DEFAULT false NOT NULL, u BOOLEAN DEFAULT false NOT NULL, d BOOLEAN DEFAULT false NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL,deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_role FOREIGN KEY(role_code) REFERENCES roles(code), CONSTRAINT fk_resource FOREIGN KEY(role_code) REFERENCES resources(code)",
	"users":              "id SERIAL PRIMARY KEY, password VARCHAR NOT NULL, role_code VARCHAR(10) DEFAULT user NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_role FOREIGN KEY(role_code) REFERENCES roles(code)",
	"profiles":           "id SERIAL PRIMARY KEY, user_id int UNIQUE, first_name VARCHAR(50) DEFAULT '' NOT NULL, last_name VARCHAR(50) DEFAULT '' NOT NULL, email VARCHAR(50) UNIQUE DEFAULT '' NOT NULL, pincode VARCHAR(10) DEFAULT '' NOT NULL, address_one VARCHAR(100) DEFAULT '' NOT NULL, address_two VARCHAR(100) DEFAULT '' NOT NULL, phone_number VARCHAR(15) DEFAULT '' NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)",
	"product_categories": "id SERIAL PRIMARY KEY, name VARCHAR(30) NOT NULL, slug VARCHAR(30) NOT NULL,enabled BOOLEAN DEFAULT true, description VARCHAR(120) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	"products":           "id SERIAL PRIMARY KEY, name VARCHAR(50), slug VARCHAR(50), price INTEGER NOT NULL DEFAULT 0, image VARCHAR(100),  description VARCHAR(300) NOT NULL, category_id INTEGER NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES product_categories(id)",
	"carts":              "id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL, product_id INTEGER NOT NULL, quantity INTEGER DEFAULT 1 NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id), CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)",
	"orders":             "id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL, order_id VARCHAR(20) NOT NULL, price INTEGER NOT NULL DEFAULT 0, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
	"purchases": `id SERIAL PRIMARY KEY, product_id INTEGER NOT NULL,
	order_id INTEGER NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL, 
	CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id),
	CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)`,
	"uploads": "id SERIAL PRIMARY KEY, path VARCHAR(200) NOT NULL, created_at TIMESTAMP DEFAULT NOW() NOT NULL, updated_at TIMESTAMP DEFAULT NOW() NOT NULL, deleted_at TIMESTAMP DEFAULT NULL",
}

var KEYS = []string{
	"roles",
	// "resources",
	// "permissions",
	"users",
	"profiles",
	"product_categories",
	"products",
	"carts",
	"orders",
	"purchases",
	"uploads",
}
