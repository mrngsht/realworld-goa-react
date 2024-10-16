-- name: GetUserProfileByUsername :one
SELECT * FROM user_profile_
WHERE username_ = $1
LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM user_
WHERE id_ = $1
LIMIT 1;

-- name: ListUserProfileMutationByUserID :many
SELECT * FROM user_profile_mutation_
WHERE user_id_ = $1;

-- name: GetUserAuthPasswordByUserID :one
SELECT * FROM user_auth_password_
WHERE user_id_ = $1
LIMIT 1;

-- name: GetUserEmailByID :one
SELECT * FROM user_email_
WHERE user_id_ = $1
LIMIT 1;

-- name: GetUserEmailByEmail :one
SELECT * FROM user_email_
WHERE email_ = $1
LIMIT 1;

-- name: ListUserEmailMutationByUserID :many
SELECT * FROM user_email_mutation_
WHERE user_id_ = $1;
