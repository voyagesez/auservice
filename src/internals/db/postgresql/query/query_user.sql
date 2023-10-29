-- name: GetUserById :one
SELECT *
FROM user_account
WHERE uid = $1
LIMIT 1;