-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS(
    INSERT INTO feed_follows(id, created_at,updated_at,user_id,feed_id)
    VALUES(
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.user_id = users.id

INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowForUser :many
SELECT
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows AS ff
INNER JOIN users
ON ff.user_id = users.id
INNER JOIN feeds
ON ff.feed_id = feeds.id
WHERE $1 = ff.user_id;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE $1 = user_id and
$2 = feed_id;