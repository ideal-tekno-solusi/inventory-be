create table if not exists categories (
	id varchar(50) primary key,
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists branches (
	id varchar(50) primary key,
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists positions (
	code varchar(50) primary key,
	branch_id varchar(50) references branches(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists items (
	id varchar(50) primary key,
	category_id varchar(50) references categories(id),
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists branch_items (
	id varchar(50) primary key,
	item_id varchar(50) references items(id),
	branch_id varchar(50) references branches(id),
	position_code varchar(50) references positions(code),
	qty int not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);