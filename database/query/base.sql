create schema if not exists inventory;

create table if not exists inventory.categories (
	id varchar(50) primary key,
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.branches (
	id varchar(50) primary key,
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.positions (
	code varchar(50) primary key,
	branch_id varchar(50) references inventory.branches(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.items (
	id varchar(50) primary key,
	category_id varchar(50) references inventory.categories(id),
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.branch_items (
	id varchar(50) primary key,
	item_id varchar(50) references inventory.items(id),
	branch_id varchar(50) references inventory.branches(id),
	position_code varchar(50) references inventory.positions(code),
	qty int not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

insert into inventory.categories
	(
		id,
		name,
		description,
		insert_date
	)
values 
	(
		'SMS',
		'SAMSUNG',
		'categories for all samsung phone component',
		now()
	),
	(
		'VIV',
		'VIVO',
		'categories for all vivo phone component',
		now()
	),
	(
		'OTH',
		'OTHER',
		'categories for all global phone component that can be used for many phone',
		now()
	),
	(
		'OPP',
		'OPPO',
		'categories for all oppo phone component',
		now()
	);
	
insert into inventory.branches
	(
		id,
		name,
		address,
		description,
		insert_date
	)
values
	(
		'BSD',
		'Bumi Serpong Damai',
		'tanggerang selatan',
		'branch inventory',
		now()
	),
	(
		'BAT',
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
		'BSD',
		now()
	),
	(
		'BSDA2',
		'BSD',
		now()
	),
	(
		'BATA1',
		'BAT',
		now()
	),
	(
		'BATA2',
		'BAT',
		now()
	);
	
insert into inventory.items
	(
		id,
		category_id,
		name,
		insert_date
	)
values
	(
		'SMS001',
		'SMS',
		'back casing samsung A5x',
		now()
	),
	(
		'SMS002',
		'SMS',
		'main board samsung A5x',
		now()
	),
	(
		'SMS003',
		'SMS',
		'daughter board samsung A5x',
		now()
	),
	(
		'OTH001',
		'OTH',
		'camera 25mp',
		now()
	),
	(
		'OTH002',
		'OTH',
		'type c connector',
		now()
	);

insert into inventory.branch_items
	(
		id,
		item_id,
		qty,
		branch_id,
		position_code,
		insert_date
	)
values
	(
		'BSDSMS001',
		'SMS001',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BSDSMS002',
		'SMS002',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BSDSMS003',
		'SMS003',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BATSMS001',
		'SMS001',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BATOTH001',
		'OTH001',
		1000,
		'BAT',
		'BATA1',
		now()
	);