create table if not exists categories (
	id varchar(20) primary key check (id ~ '^CAT[0-9]+$') default 'CAT' || nextval('inventory.categories_id_seq'),
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists branches (
	id varchar(20) primary key check (id ~ '^BRA[0-9]+$') default 'BRA' || nextval('inventory.branches_id_seq'),
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists positions (
	id varchar(20) primary key check (id ~ '^POS[0-9]+$') default 'POS' || nextval('inventory.positions_id_seq'),
	code varchar(50),
	branch_id varchar(20) references branches(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists items (
	id varchar(20) primary key check (id ~ '^ITE[0-9]+$') default 'ITE' || nextval('inventory.items_id_seq'),
	category_id varchar(20) references categories(id),
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists branch_items (
	id varchar(20) primary key check (id ~ '^BIT[0-9]+$') default 'BIT' || nextval('inventory.branch_items_id_seq'),
	item_id varchar(20) references items(id),
	branch_id varchar(20) references branches(id),
	position_id varchar(20) references positions(id),
	qty int not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists challenges (
	code_verifier text,
	code_challenge text,
	code_challenge_method varchar(5),
	insert_date timestamp
);