-- auth-keys table
create unique index email_index on auth_keys (email);
create unique index nif_index on auth_keys (nif);
-- users table
create index first_name_index on users (first_name);
create index last_name_index on users (last_name);
create unique index phone_number_index on users (phone_number);
-- burn
create index burn_title_index on burn (title);
-- autarchy
create index autarchy_title_index on autarchy (title);