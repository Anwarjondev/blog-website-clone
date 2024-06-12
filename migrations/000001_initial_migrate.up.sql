CREATE table if not exists "users" (
    "id" Varchar primary key,
    "first_name" varchar not null,
    "last_name" varchar not null,
    "email" varchar unique,
    "password" varchar not null,
    "created_at" timestamp default current_timestamp
    );
Create table if not exists "posts"(
    "id" varchar primary key,
    "title" varchar not null,
    "body" text not null,
    "published" boolean default false,
    "user_id" varchar not null, references users(id) On delete cascade,
    "created_at" timestamp default current_timestamp
)