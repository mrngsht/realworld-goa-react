CREATE TYPE "user_follow_mutation_type_" AS ENUM (
    'follow',
    'unfollow'
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