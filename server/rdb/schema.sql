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
CREATE TABLE "user_profile_" (
    "created_at_" timestamp with time zone NOT NULL,
    "updated_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "username_" character varying(24) NOT NULL,
    "email_" "text" NOT NULL,
    "bio_" "text" NOT NULL,
    "image_url_" "text" NOT NULL
);
CREATE TABLE "user_profile_mutation_" (
    "created_at_" timestamp with time zone NOT NULL,
    "user_id_" "uuid" NOT NULL,
    "username_" character varying(24) NOT NULL,
    "email_" "text" NOT NULL,
    "bio_" "text" NOT NULL,
    "image_url_" "text" NOT NULL
);
ALTER TABLE ONLY "user_"
    ADD CONSTRAINT "user__pkey" PRIMARY KEY ("id_");
ALTER TABLE ONLY "user_auth_password_"
    ADD CONSTRAINT "user_auth_password__pkey" PRIMARY KEY ("user_id_");
ALTER TABLE ONLY "user_profile_"
    ADD CONSTRAINT "user_profile__pkey" PRIMARY KEY ("user_id_");
ALTER TABLE ONLY "user_profile_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_profile_mutation_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");
ALTER TABLE ONLY "user_auth_password_"
    ADD CONSTRAINT "fk_user_id_" FOREIGN KEY ("user_id_") REFERENCES "user_"("id_");