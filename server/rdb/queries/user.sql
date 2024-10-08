-- name: GetUser :one
SELECT id_ FROM user_
WHERE id_ = $1;

-- name: InsertUser :exec
INSERT INTO user_
(created_at_, id_) 
VALUES ($1, $2);

-- name: InsertUserAuthPassword :exec
INSERT INTO user_auth_password_
(created_at_, updated_at_, user_id_, password_hash_) 
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2);

-- name: InsertUserProfile :exec
INSERT INTO user_profile_
(created_at_, updated_at_, user_id_, username_, email_, bio_, image_url_) 
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2, $3, $4, $5);

-- name: InsertUserProfileMutation :exec
INSERT INTO user_profile_mutation_
(created_at_, user_id_, username_, email_, bio_, image_url_) 
VALUES ($1, $2, $3, $4, $5, $6);
