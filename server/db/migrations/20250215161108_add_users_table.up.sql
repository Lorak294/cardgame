CREATE TABLE "users"(
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password_hash" varchar NOT NULL
) 