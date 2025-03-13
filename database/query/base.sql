create schema if not exists inventory;

create table if not exists inventory.categories (
	id varchar(50) primary key,
	name varchar(255) not null,
	description text not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.locations (
	id varchar(50) primary key,
	name varchar(255) not null,
	address text not null,
	description text,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.position (
	id varchar(50) primary key,
	location_id varchar(50) references inventory.locations(id),
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.global_items (
	id varchar(50) primary key,
	name varchar(255) not null,
	insert_date timestamp not null,
	update_date timestamp,
	delete_date timestamp
);

create table if not exists inventory.items (
	id varchar(50) primary key,
	global_item_id varchar(50) references inventory.global_items(id),
	category_id varchar(50) references inventory.categories(id),
	name varchar(255) not null,
	qty int not null,
	location_id varchar(50) references inventory.locations(id),
	position_id varchar(50) references inventory.position(id),
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
	
insert into inventory.locations
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

insert into inventory.position
	(
		id,
		location_id,
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
	
insert into inventory.global_items
	(
		id,
		name,
		insert_date
	)
values
	(
		'SMS001',
		'back casing samsung A5x',
		now()
	),
	(
		'SMS002',
		'main board samsung A5x',
		now()
	),
	(
		'SMS003',
		'daughter board samsung A5x',
		now()
	),
	(
		'OTH001',
		'camera 25mp',
		now()
	),
	(
		'OTH002',
		'type c connector',
		now()
	);

insert into inventory.items
	(
		id,
		global_item_id,
		category_id,
		name,
		qty,
		location_id,
		position_id,
		insert_date
	)
values
	(
		'BSDSMS001',
		'SMS001',
		'SMS',
		'Back casing for samsung A5x',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BSDSMS002',
		'SMS002',
		'SMS',
		'Main board for samsung A5x',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BSDSMS003',
		'SMS003',
		'SMS',
		'Daughter board for samsung A5x',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BATSMS001',
		'SMS001',
		'SMS',
		'Back casing for samsung A5x',
		1000,
		'BSD',
		'BSDA1',
		now()
	),
	(
		'BATOTH001',
		'OTH001',
		'OTH',
		'Camera 25mp',
		1000,
		'BAT',
		'BATA1',
		now()
	);