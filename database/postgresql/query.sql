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
    items.category_id like $1
and
    items.location_id like $2
order by 
    items.name
desc;