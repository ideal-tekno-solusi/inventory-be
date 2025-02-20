create table if not exists categories (
	id varchar(50) primary key,
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists locations (
	id varchar(50) primary key,
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists position (
	id varchar(50) primary key,
	location_id varchar(50) references locations(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists global_items (
	id varchar(50) primary key,
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists items (
	id varchar(50) primary key,
	global_item_id varchar(50) references global_items(id),
	category_id varchar(50) references categories(id),
	name varchar(255) not null,
	qty int not null,
	location_id varchar(50) references locations(id),
	position_id varchar(50) references position(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);