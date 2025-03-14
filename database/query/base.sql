create schema if not exists inventory;

create sequence inventory.categories_id_seq
increment 1
minvalue 1
maxvalue 9999999999999999
start 1;

create table if not exists inventory.categories (
	id varchar(20) primary key check (id ~ '^CAT[0-9]+$') default 'CAT' || nextval('inventory.categories_id_seq'),
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create sequence inventory.branches_id_seq
increment 1
minvalue 1
maxvalue 9999999999999999
start 1;

create table if not exists inventory.branches (
	id varchar(20) primary key check (id ~ '^BRA[0-9]+$') default 'BRA' || nextval('inventory.branches_id_seq'),
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.positions (
	code varchar(50) primary key,
	branch_id varchar(20) references inventory.branches(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create sequence inventory.items_id_seq
increment 1
minvalue 1
maxvalue 9999999999999999
start 1;

create table if not exists inventory.items (
	id varchar(20) primary key check (id ~ '^ITE[0-9]+$') default 'ITE' || nextval('inventory.items_id_seq'),
	category_id varchar(20) references inventory.categories(id),
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create sequence inventory.branch_items_id_seq
increment 1
minvalue 1
maxvalue 9999999999999999
start 1;

create table if not exists inventory.branch_items (
	id varchar(20) primary key check (id ~ '^BIT[0-9]+$') default 'BIT' || nextval('inventory.branch_items_id_seq'),
	item_id varchar(20) references inventory.items(id),
	branch_id varchar(20) references inventory.branches(id),
	position_code varchar(50) references inventory.positions(code),
	qty int not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

insert into inventory.categories
	(
		name,
		description,
		insert_date
	)
values 
	(
		
		'SAMSUNG',
		'categories for all samsung phone component',
		now()
	),
	(
		'VIVO',
		'categories for all vivo phone component',
		now()
	),
	(
		'OTHER',
		'categories for all global phone component that can be used for many phone',
		now()
	),
	(
		'OPPO',
		'categories for all oppo phone component',
		now()
	);
	
insert into inventory.branches
	(
		name,
		address,
		description,
		insert_date
	)
values
	(
		'Bumi Serpong Damai',
		'tanggerang selatan',
		'branch inventory',
		now()
	),
	(
		'Batam',
		'riau',
		'main inventory for indonesian branch',
		now()
	);

insert into inventory.positions
	(
		code,
		branch_id,
		insert_date
	)
values
	(
		'BSDA1',
		'BRA1',
		now()
	),
	(
		'BSDA2',
		'BRA1',
		now()
	),
	(
		'BATA1',
		'BRA2',
		now()
	),
	(
		'BATA2',
		'BRA2',
		now()
	);
	
insert into inventory.items
	(
		category_id,
		name,
		insert_date
	)
values
	(
		'CAT1',
		'back casing samsung A5x',
		now()
	),
	(
		'CAT1',
		'main board samsung A5x',
		now()
	),
	(
		'CAT1',
		'daughter board samsung A5x',
		now()
	),
	(
		'CAT3',
		'camera 25mp',
		now()
	),
	(
		'CAT3',
		'type c connector',
		now()
	);

insert into inventory.branch_items
	(
		item_id,
		qty,
		branch_id,
		position_code,
		insert_date
	)
values
	(
		'ITE1',
		1000,
		'BRA1',
		'BSDA1',
		now()
	),
	(
		'ITE2',
		1000,
		'BRA1',
		'BSDA1',
		now()
	),
	(
		'ITE3',
		1000,
		'BRA1',
		'BSDA1',
		now()
	),
	(
		'ITE1',
		1000,
		'BRA2',
		'BATA1',
		now()
	),
	(
		'ITE5',
		1000,
		'BRA2',
		'BATA2',
		now()
	);