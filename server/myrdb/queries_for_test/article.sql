-- name: GetArticleByID :one
SELECT * FROM article_
WHERE id_ = $1
LIMIT 1;

-- name: GetArticleContentByArticleID :one
SELECT * FROM article_content_
WHERE article_id_ = $1
LIMIT 1;

-- name: ListArticleContentMutationByArticleID :many
SELECT * FROM article_content_mutation_
WHERE article_id_ = $1
ORDER BY created_at_ ASC;

-- name: ListArticleTagByArticleID :many
SELECT * FROM article_tag_
WHERE article_id_ = $1
ORDER BY seq_no_ ASC;

-- name: ListArticleTagMutationByArticleID :many
SELECT * FROM article_tag_mutation_
WHERE article_id_ = $1
ORDER BY created_at_ ASC;

-- name: GetArticleStatsByArticleID :one
SELECT * FROM article_stats_
WHERE article_id_ = $1
LIMIT 1;

-- name: GetArticleDeletedByArticleID :one
SELECT * FROM article_deleted_
WHERE article_id_ = $1
LIMIT 1;


-- name: GetArticleCommentDeletedByCommentID :one
SELECT 
  created_at_,
  article_comment_id_
FROM article_comment_deleted_
WHERE article_comment_id_ = $1;
