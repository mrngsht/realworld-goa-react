-- name: InsertArticle :exec
INSERT INTO article_
(created_at_, id_) 
VALUES ($1, $2);

-- name: InsertArticleContent :exec
INSERT INTO article_content_
(created_at_, updated_at_, article_id_, title_, description_, body_, author_user_id_) 
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2, $3, $4, $5);

-- name: InsertArticleContentMutation :exec
INSERT INTO article_content_mutation_
(created_at_, article_id_, title_, description_, body_, author_user_id_) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: InsertArticleTag :copyfrom
INSERT INTO article_tag_
(created_at_, article_id_, seq_no_, tag_) 
VALUES ($1, $2, $3, $4);

-- name: InsertArticleTagMutation :exec
INSERT INTO article_tag_mutation_
(created_at_, article_id_, tags_) 
VALUES ($1, $2, $3);

-- name: InsertArticleStats :exec
INSERT INTO article_stats_
(created_at_, updated_at_, article_id_, favorites_count_)
VALUES (sqlc.arg(created_at), sqlc.arg(created_at), $1, $2);

-- name: InsertArticleFavorite :exec
INSERT INTO article_favorite_
(created_at_, article_id_, user_id_)
VALUES ($1, $2, $3);

-- name: InsertArticleFavoriteMutation :exec
INSERT INTO article_favorite_mutation_
(created_at_, article_id_, user_id_, type_)
VALUES ($1, $2, $3, $4);

-- name: UpdateArticleStatsFavoritesCount :exec
UPDATE article_stats_
SET favorites_count_ = $1
WHERE article_id_ = $2;

-- name: GetArticleContentByArticleID :one
SELECT 
  created_at_,
  updated_at_,
  article_id_,
  title_,
  description_,
  body_,
  author_user_id_
FROM article_content_ 
WHERE article_id_ = $1
LIMIT 1;

-- name: GetArticleStatsByArticleID :one
SELECT 
  article_id_,
  favorites_count_
FROM article_stats_ 
WHERE article_id_ = $1
LIMIT 1;

-- name: GetArticleStatsByArticleIDForUpdate :one
SELECT 
  article_id_,
  favorites_count_
FROM article_stats_ 
WHERE article_id_ = $1
LIMIT 1
FOR UPDATE;

-- name: ListArticleTagByArticleID :many
SELECT 
  tag_
FROM article_tag_
WHERE article_id_ = $1
ORDER BY seq_no_ ASC;

-- name: IsArticleFavoritedByArticleIDAndUserID :one
SELECT EXISTS (
  SELECT 1
  FROM article_favorite_ 
  WHERE article_id_ = $1 AND user_id_ = $2
);

