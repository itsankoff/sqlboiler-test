CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE messages
(
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY
);

CREATE TABLE attendees
(
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY
);

CREATE TABLE received_messages
(
    message_id uuid NOT NULL REFERENCES messages ON DELETE CASCADE,
    attendee_id uuid NOT NULL REFERENCES attendees ON DELETE CASCADE,

    read_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE received_messages
    ADD CONSTRAINT received_messages_pkey PRIMARY KEY (attendee_id, message_id);

ALTER TABLE received_messages
    ADD UNIQUE (attendee_id, message_id);
