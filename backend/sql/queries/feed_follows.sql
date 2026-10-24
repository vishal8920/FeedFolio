-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id=$1;

/*
    one===>Returns one
    many===>Returns many
    exec===>Doesn't return just executes the query
*/ 
-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id=$1 and user_id=$2;

