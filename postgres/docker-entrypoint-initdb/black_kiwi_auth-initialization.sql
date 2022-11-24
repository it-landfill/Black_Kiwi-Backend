CREATE SCHEMA "black-kiwi_authentication";

create table "black-kiwi_authentication"."Roles" (
    id serial not null
        constraint "Roles_pk"
            primary key,
    name varchar(5) not null
        constraint "Roles_name_unique"
            unique
);

create table "black-kiwi_authentication"."Users"(
    id serial not null
        constraint "Users_pk"
            primary key,
    username varchar(50) not null
        constraint "Users_username_unique"
            unique,
    password varchar(256) not null,
    role integer not null
        constraint "Users___role_fk"
            references "black-kiwi_authentication"."Roles" (id)
);

CREATE USER "black-kiwi_login" WITH PASSWORD 'vwWTR7sFRw9sh9KA';
GRANT USAGE ON SCHEMA "black-kiwi_authentication" TO "black-kiwi_login";
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE "black-kiwi_authentication"."Roles" TO "black-kiwi_login";
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE "black-kiwi_authentication"."Users" TO "black-kiwi_login";

INSERT INTO "black-kiwi_authentication"."Roles" (id, name) VALUES (1, 'user');
INSERT INTO "black-kiwi_authentication"."Roles" (id, name) VALUES (2, 'admin');

INSERT INTO "black-kiwi_authentication"."Users" VALUES (1, 'testUser', 'testUser', 1);
INSERT INTO "black-kiwi_authentication"."Users" VALUES (2, 'testAdmin', 'testAdmin', 2);