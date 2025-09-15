-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT feeds.*, users.name AS user_name
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

-- name: GetNextFollowedFeedToFetch :one
SELECT * FROM feeds
WHERE user_id IN (SELECT feed_follows.user_id FROM feed_follows WHERE feed_follows.user_id = $1)
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;