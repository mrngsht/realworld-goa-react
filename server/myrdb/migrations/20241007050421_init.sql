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

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_auth_password_;
DROP TABLE IF EXISTS user_email_mutation_;
DROP TABLE IF EXISTS user_email_;
DROP TABLE IF EXISTS user_profile_mutation_;
DROP TABLE IF EXISTS user_profile_;
DROP TABLE IF EXISTS user_;
-- +goose StatementEnd
