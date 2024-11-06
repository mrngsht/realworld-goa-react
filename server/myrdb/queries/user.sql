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
(created_at_, updated_at_, user_id_, username_, bio_, image_url_) 
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2, $3, $4);

-- name: InsertUserProfileMutation :exec
INSERT INTO user_profile_mutation_
(created_at_, user_id_, username_, bio_, image_url_) 
VALUES ($1, $2, $3, $4, $5);

-- name: InsertUserEmail :exec
INSERT INTO user_email_
(created_at_, updated_at_, user_id_, email_) 
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2);

-- name: InsertUserEmailMutation :exec
INSERT INTO user_email_mutation_
(created_at_, user_id_, email_) 
VALUES ($1, $2, $3);

-- name: GetPasswordHashByUserID :one
SELECT password_hash_ FROM user_auth_password_ 
WHERE user_id_ = $1
LIMIT 1;

-- name: GetUserIDByEmail :one
SELECT user_id_ FROM user_email_ 
WHERE email_ = $1
LIMIT 1;

-- name: GetUserProfileByUserID :one
SELECT 
  username_, 
  bio_, 
  image_url_ 
FROM user_profile_
WHERE user_id_ = $1
LIMIT 1;

-- name: GetUserEmailByUserID :one
SELECT email_ FROM user_email_ 
WHERE user_id_ = $1
LIMIT 1;

-- name: UpdateUserEmail :exec
UPDATE user_email_
SET updated_at_ = $2, email_ = $3
WHERE user_id_ = $1;

-- name: UpdateUserAuthPasswordHash :exec
UPDATE user_auth_password_
SET updated_at_ = $2, password_hash_ = $3
WHERE user_id_ = $1;

-- name: UpdateUserProfile :exec
UPDATE user_profile_
SET updated_at_ = $2, username_ = $3, bio_ = $4, image_url_ = $5 
WHERE user_id_ = $1;


