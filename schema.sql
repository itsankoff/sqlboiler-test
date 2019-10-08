CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users
(
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY
);

CREATE TABLE messages
(
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY
);

CREATE TABLE read_messages
(
    user_id uuid REFERENCES users NOT NULL,
    message_id uuid REFERENCES messages NOT NULL,

    read_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE read_messages
    ADD CONSTRAINT read_messages_pkey PRIMARY KEY (user_id, message_id);
