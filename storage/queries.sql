-- name: GetProfiles :many
SELECT * FROM profile ORDER BY name ASC;

-- name: GetProfileByName :one
SELECT * FROM profile WHERE name = ?;

-- name: CreateProfile :exec
INSERT INTO profile (name, emoji) VALUES (?, ?);

-- name: GetConfig :one
SELECT * FROM config WHERE key = ?;


-- name: SetConfig :exec
INSERT INTO config (key, value) VALUES (sqlc.arg(key), sqlc.arg(value))
ON CONFLICT(key) DO UPDATE SET value = sqlc.arg(value);

-- name: GetArticles :many
SELECT * FROM article WHERE profile_id=? ORDER BY created_at DESC;