CREATE TABLE IF NOT EXISTS receivers (
	receiver_id uuid DEFAULT gen_random_uuid(),
	name varchar NOT NULL,
	document varchar NOT NULL,
	email varchar,
	status integer NOT NULL,
	pix_key varchar NOT NULL,
	pix_key_type integer NOT NULL,
	bank varchar,
	office varchar,
	account_number varchar,
	created_at timestamp DEFAULT now(),
	updated_at timestamp DEFAULT now(),
	PRIMARY KEY (receiver_id)
);
