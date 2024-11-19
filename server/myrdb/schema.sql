CREATE TYPE "article_favorite_mutation_type_" AS ENUM (
    'favorite',
    'unfavorite'
);
CREATE TYPE "user_follow_mutation_type_" AS ENUM (
    'follow',
    'unfollow'
);
CREATE TABLE "article_" (
    "created_at_" timestamp with time zone NOT NULL,
    "id_" "uuid" NOT NULL
);
CREATE TABLE "article_comment_" (
    "created_at_" timestamp with time zone NOT NULL,
    "id_" "uuid" NOT NULL
);
CREATE TABLE "article_comment_content_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_comment_id_" "uuid" NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "body_" "text" NOT NULL,
    "user_id_" "uuid" NOT NULL
);
CREATE TABLE "article_comment_content_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_comment_id_" "uuid" NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "body_" "text" NOT NULL,
    "user_id_" "uuid" NOT NULL
);
CREATE TABLE "article_comment_deleted_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_comment_id_" "uuid" NOT NULL
);
CREATE TABLE "article_content_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "title_" "text" NOT NULL,
    "description_" "text" NOT NULL,
    "body_" "text" NOT NULL,
    "author_user_id_" "uuid" NOT NULL
);
CREATE TABLE "article_content_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "title_" "text" NOT NULL,
    "description_" "text" NOT NULL,
    "body_" "text" NOT NULL,
    "author_user_id_" "uuid" NOT NULL
);
CREATE TABLE "article_deleted_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL
);
CREATE TABLE "article_favorite_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "user_id_" "uuid" NOT NULL
);
CREATE TABLE "article_favorite_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "type_" "article_favorite_mutation_type_" NOT NULL
);
CREATE TABLE "article_stats_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "favorites_count_" bigint
);
CREATE TABLE "article_tag_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "seq_no_" integer NOT NULL,
    "tag_" "text" NOT NULL
);
CREATE TABLE "article_tag_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "article_id_" "uuid" NOT NULL,
    "tags_" "jsonb" NOT NULL
);
CREATE TABLE "enum_article_tag_" (
    "created_at_" timestamp with time zone NOT NULL,
    "tag_" "text" NOT NULL
);
CREATE TABLE "user_" (
    "created_at_" timestamp with time zone NOT NULL,
    "id_" "uuid" NOT NULL
);
CREATE TABLE "user_auth_password_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "password_hash_" "text" NOT NULL
);
CREATE TABLE "user_email_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "email_" "text" NOT NULL
);
CREATE TABLE "user_email_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "email_" "text" NOT NULL
);
CREATE TABLE "user_follow_" (
    "created_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "followed_user_id_" "uuid" NOT NULL
);
CREATE TABLE "user_follow_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "followed_user_id_" "uuid" NOT NULL,
    "type_" "user_follow_mutation_type_" NOT NULL
);
CREATE TABLE "user_profile_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "username_" character varying(64) NOT NULL,
    "bio_" "text" NOT NULL,
    "image_url_" "text" NOT NULL
);
CREATE TABLE "user_profile_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "username_" character varying(64) NOT NULL,
    "bio_" "text" NOT NULL,
    "image_url_" "text" NOT NULL
);
ALTER TABLE ONLY "article_"
    ADD CONSTRAINT "article__pkey" PRIMARY KEY ("id_");
ALTER TABLE ONLY "article_comment_"
    ADD CONSTRAINT "article_comment__pkey" PRIMARY KEY ("id_");
ALTER TABLE ONLY "article_comment_content_"
    ADD CONSTRAINT "article_comment_content__pkey" PRIMARY KEY ("article_comment_id_");
ALTER TABLE ONLY "article_comment_deleted_"
    ADD CONSTRAINT "article_comment_deleted__pkey" PRIMARY KEY ("article_comment_id_");
ALTER TABLE ONLY "article_content_"
    ADD CONSTRAINT "article_content__pkey" PRIMARY KEY ("article_id_");
ALTER TABLE ONLY "article_deleted_"
    ADD CONSTRAINT "article_deleted__pkey" PRIMARY KEY ("article_id_");
ALTER TABLE ONLY "article_favorite_"
    ADD CONSTRAINT "article_favorite__pkey" PRIMARY KEY ("article_id_", "user_id_");
ALTER TABLE ONLY "article_stats_"
    ADD CONSTRAINT "article_stats__pkey" PRIMARY KEY ("article_id_");
ALTER TABLE ONLY "article_tag_"
    ADD CONSTRAINT "article_tag__article_id__tag__key" UNIQUE ("article_id_", "tag_");
ALTER TABLE ONLY "article_tag_"
    ADD CONSTRAINT "article_tag__pkey" PRIMARY KEY ("article_id_", "seq_no_");
ALTER TABLE ONLY "enum_article_tag_"
    ADD CONSTRAINT "enum_article_tag__pkey" PRIMARY KEY ("tag_");
ALTER TABLE ONLY "user_"
    ADD CONSTRAINT "user__pkey" PRIMARY KEY ("id_");
ALTER TABLE ONLY "user_auth_password_"
    ADD CONSTRAINT "user_auth_password__pkey" PRIMARY KEY ("user_id_");
ALTER TABLE ONLY "user_email_"
    ADD CONSTRAINT "user_email__email__key" UNIQUE ("email_");
ALTER TABLE ONLY "user_email_"
    ADD CONSTRAINT "user_email__pkey" PRIMARY KEY ("user_id_");
ALTER TABLE ONLY "user_follow_"
    ADD CONSTRAINT "user_follow__pkey" PRIMARY KEY ("user_id_", "followed_user_id_");
ALTER TABLE ONLY "user_profile_"
    ADD CONSTRAINT "user_profile__pkey" PRIMARY KEY ("user_id_");
ALTER TABLE ONLY "user_profile_"
    ADD CONSTRAINT "user_profile__username__key" UNIQUE ("username_");
ALTER TABLE ONLY "article_comment_content_"
    ADD CONSTRAINT "fk_article_comment_id_" FOREIGN KEY ("article_comment_id_") REFERENCES "article_comment_"("id_");
ALTER TABLE ONLY "article_comment_content_mutation_"
    ADD CONSTRAINT "fk_article_comment_id_" FOREIGN KEY ("article_comment_id_") REFERENCES "article_comment_"("id_");
ALTER TABLE ONLY "article_comment_deleted_"
    ADD CONSTRAINT "fk_article_comment_id_" FOREIGN KEY ("article_comment_id_") REFERENCES "article_comment_"("id_");
ALTER TABLE ONLY "article_content_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_content_mutation_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_deleted_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_tag_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_tag_mutation_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_comment_content_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_comment_content_mutation_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_favorite_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_favorite_mutation_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_stats_"
    ADD CONSTRAINT "fk_article_id_" FOREIGN KEY ("article_id_") REFERENCES "article_"("id_");
ALTER TABLE ONLY "article_content_"
    ADD CONSTRAINT "fk_author_user_id_" FOREIGN KEY ("author_user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "article_content_mutation_"
    ADD CONSTRAINT "fk_author_user_id_" FOREIGN KEY ("author_user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_follow_"
    ADD CONSTRAINT "fk_followed_user_id_" FOREIGN KEY ("followed_user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_follow_mutation_"
    ADD CONSTRAINT "fk_followed_user_id_" FOREIGN KEY ("followed_user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_profile_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_profile_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_email_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_email_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_auth_password_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_follow_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_follow_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "article_comment_content_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "article_comment_content_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "article_favorite_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "article_favorite_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");