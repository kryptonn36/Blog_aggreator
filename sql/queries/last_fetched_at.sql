-- name: MarkFeedFetched :one
UPDATE feeds 
SET last_fetched_at = Now(), updated_at = Now()
WHERE $1 = id
RETURNING *;


-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC 
NULLS FIRST
LIMIT 1;