-- name: FetchInventoryItems :many
select 
    items.id,
    items.name,
    items.qty,
    items.global_item_id,
    items.category_id,
    items.location_id,
    items.position_id
from 
    items
where
    items.category_id ilike $1
and
    items.location_id ilike $2
order by 
    items.name
desc
limit $3
offset $4;

-- name: CountInventoryItems :one
select 
    count(*)
from 
    items
where
    items.category_id ilike $1
and
    items.location_id ilike $2;

-- name: CountCategory :one
select
    count(*)
from
    categories
where
    categories.name ilike $1;

-- name: FetchCategory :many
select
    categories.id,
    categories.name,
    categories.description,
    categories.insert_date,
    categories.update_date
from
    categories
where
    categories.name ilike $1
order by
    categories.id
desc
limit $2
offset $3;