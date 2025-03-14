-- name: FetchInventoryItems :many
select 
    branch_items.id,
    items.name as item_name,
    branch_items.qty,
    categories.name as category_name,
    branches.name as branch_name,
    branch_items.position_code
from 
    branch_items
join
    items
on
    branch_items.item_id = items.id
join
    categories
on
    items.category_id = categories.id
join
    branches
on
    branch_items.branch_id = branches.id
where
    categories.id ilike $1
and
    branches.id ilike $2
order by 
    items.name
desc
limit $3
offset $4;

-- name: CountInventoryItems :one
select 
    count(*)
from 
    branch_items
join
    items
on
    branch_items.item_id = items.id
join
    categories
on
    items.category_id = categories.id
join
    branches
on
    branch_items.branch_id = branches.id
where
    categories.id ilike $1
and
    branches.id ilike $2;

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