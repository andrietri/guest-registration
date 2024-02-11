
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.guests (
	id bigserial NOT NULL,
    name varchar(50) NULL,
    id_card_number varchar(50) NULL,
    email varchar(50) NULL,
    phone varchar(50) NULL,
    created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT room_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE public.guests;
