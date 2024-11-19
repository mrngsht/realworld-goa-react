-- +goose Up
CREATE TABLE IF NOT EXISTS user_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  id_ UUID NOT NULL,
  PRIMARY KEY (id_)
);

CREATE TABLE IF NOT EXISTS user_profile_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  updated_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  username_ VARCHAR(64) NOT NULL,
  bio_ TEXT NOT NULL,
  image_url_ TEXT NOT NULL,
  PRIMARY KEY (user_id_),
  UNIQUE (username_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS user_profile_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  username_ VARCHAR(64) NOT NULL,
  bio_ TEXT NOT NULL,
  image_url_ TEXT NOT NULL,
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS user_email_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  updated_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  email_ TEXT NOT NULL,
  PRIMARY KEY (user_id_),
  UNIQUE (email_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS user_email_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  email_ TEXT NOT NULL,
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS user_auth_password_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  updated_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  password_hash_ TEXT NOT NULL,
  PRIMARY KEY (user_id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS user_follow_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  followed_user_id_ UUID NOT NULL,
  PRIMARY KEY (user_id_, followed_user_id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_),
  CONSTRAINT fk_followed_user_id_ FOREIGN KEY (followed_user_id_) REFERENCES user_ (id_)
);

CREATE TYPE user_follow_mutation_type_ AS ENUM ('follow', 'unfollow');

CREATE TABLE IF NOT EXISTS user_follow_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  user_id_ UUID NOT NULL,
  followed_user_id_ UUID NOT NULL,
  type_ user_follow_mutation_type_ NOT NULL,
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_),
  CONSTRAINT fk_followed_user_id_ FOREIGN KEY (followed_user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  id_ UUID NOT NULL,
  PRIMARY KEY (id_)
);

CREATE TABLE IF NOT EXISTS article_content_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  updated_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  title_ TEXT NOT NULL,
  description_ TEXT NOT NULL,
  body_ TEXT NOT NULL,
  author_user_id_ UUID NOT NULL,
  PRIMARY KEY (article_id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_author_user_id_ FOREIGN KEY (author_user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_content_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  title_ TEXT NOT NULL,
  description_ TEXT NOT NULL,
  body_ TEXT NOT NULL,
  author_user_id_ UUID NOT NULL,
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_author_user_id_ FOREIGN KEY (author_user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_deleted_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  PRIMARY KEY (article_id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_)
);

CREATE TABLE IF NOT EXISTS article_tag_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  seq_no_ INTEGER NOT NULL,
  tag_ TEXT NOT NULL,
  PRIMARY KEY (article_id_, seq_no_),
  UNIQUE (article_id_, tag_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_)
);

CREATE TABLE IF NOT EXISTS article_tag_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  tags_ JSONB NOT NULL,
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_)
);
COMMENT ON COLUMN article_tag_mutation_.tags_ IS 'json array of tag(string)';

CREATE TABLE IF NOT EXISTS enum_article_tag_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  tag_ TEXT NOT NULL,
  PRIMARY KEY (tag_)
);

CREATE TABLE IF NOT EXISTS article_comment_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  id_ UUID NOT NULL,
  PRIMARY KEY (id_)
);

CREATE TABLE IF NOT EXISTS article_comment_content_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_comment_id_ UUID NOT NULL,
  article_id_ UUID NOT NULL,
  body_ TEXT NOT NULL,
  user_id_ UUID NOT NULL,
  PRIMARY KEY (article_comment_id_),
  CONSTRAINT fk_article_comment_id_ FOREIGN KEY (article_comment_id_) REFERENCES article_comment_ (id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_comment_content_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_comment_id_ UUID NOT NULL,
  article_id_ UUID NOT NULL,
  body_ TEXT NOT NULL,
  user_id_ UUID NOT NULL,
  CONSTRAINT fk_article_comment_id_ FOREIGN KEY (article_comment_id_) REFERENCES article_comment_ (id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_comment_deleted_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_comment_id_ UUID NOT NULL,
  PRIMARY KEY (article_comment_id_),
  CONSTRAINT fk_article_comment_id_ FOREIGN KEY (article_comment_id_) REFERENCES article_comment_ (id_)
);

CREATE TABLE IF NOT EXISTS article_favorite_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  user_id_ UUID NOT NULL,
  PRIMARY KEY (article_id_, user_id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TYPE article_favorite_mutation_type_ AS ENUM ('favorite', 'unfavorite');

CREATE TABLE IF NOT EXISTS article_favorite_mutation_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  user_id_ UUID NOT NULL,
  type_ article_favorite_mutation_type_ NOT NULL,
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_),
  CONSTRAINT fk_user_id_ FOREIGN KEY (user_id_) REFERENCES user_ (id_)
);

CREATE TABLE IF NOT EXISTS article_stats_ (
  created_at_ TIMESTAMPTZ NOT NULL,
  updated_at_ TIMESTAMPTZ NOT NULL,
  article_id_ UUID NOT NULL,
  favorites_count_ BIGINT NULL,
  PRIMARY KEY (article_id_),
  CONSTRAINT fk_article_id_ FOREIGN KEY (article_id_) REFERENCES article_ (id_)
);
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article_stats_;
DROP TABLE IF EXISTS article_favorite_mutation_;
DROP TABLE IF EXISTS article_favorite_;
DROP TABLE IF EXISTS article_comment_deleted_;
DROP TABLE IF EXISTS article_comment_content_mutation_;
DROP TABLE IF EXISTS article_comment_content_;
DROP TABLE IF EXISTS article_comment_;
DROP TABLE IF EXISTS enum_article_tag_;
DROP TABLE IF EXISTS article_tag_mutation_;
DROP TABLE IF EXISTS article_tag_;
DROP TABLE IF EXISTS article_deleted_;
DROP TABLE IF EXISTS article_content_mutation_;
DROP TABLE IF EXISTS article_content_;
DROP TABLE IF EXISTS article_;
DROP TABLE IF EXISTS user_follow_mutation_;
DROP TABLE IF EXISTS user_follow_;
DROP TABLE IF EXISTS user_follow_;
DROP TABLE IF EXISTS user_auth_password_;
DROP TABLE IF EXISTS user_email_mutation_;
DROP TABLE IF EXISTS user_email_;
DROP TABLE IF EXISTS user_profile_mutation_;
DROP TABLE IF EXISTS user_profile_;
DROP TABLE IF EXISTS user_;

DROP TYPE IF EXISTS article_favorite_mutation_type_;
DROP TYPE IF EXISTS user_follow_mutation_type_;
-- +goose StatementEnd
