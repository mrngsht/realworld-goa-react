-- name: GetUserProfile :one
SELECT user_id_, username_ FROM user_profile_
WHERE user_id_ = $1;

