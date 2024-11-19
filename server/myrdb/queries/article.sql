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
