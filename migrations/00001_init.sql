-- +goose Up
-- +goose StatementBegin
CREATE TABLE slot (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE
);
CREATE TABLE banner (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE
);
CREATE TABLE user_group (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE
);

CREATE TABLE banner_slot (
    banner_id integer CONSTRAINT banner_slot_fk_banner_id REFERENCES banner(id) ON UPDATE CASCADE ON DELETE CASCADE,
    slot_id integer CONSTRAINT banner_slot_fk_slot_id REFERENCES slot(id) ON UPDATE CASCADE ON DELETE CASCADE,
    UNIQUE (banner_id, slot_id)
);

CREATE TABLE click_banner_event (
    id bigserial PRIMARY KEY,
    banner_id bigint CONSTRAINT click_event_fk_banner_id REFERENCES banner(id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_group_id bigint CONSTRAINT click_event_fk_user_group_id REFERENCES user_group(id) ON UPDATE CASCADE ON DELETE CASCADE,
    slot_id bigint CONSTRAINT show_banner_event_fk_click_id REFERENCES slot(id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_dt timestamptz
);

CREATE TABLE show_banner_event (
    id bigserial PRIMARY KEY,
    banner_id bigint CONSTRAINT show_banner_event_fk_banner_id REFERENCES banner(id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_group_id bigint CONSTRAINT show_banner_event_fk_user_group_id REFERENCES user_group(id) ON UPDATE CASCADE ON DELETE CASCADE,
    slot_id bigint CONSTRAINT show_banner_event_fk_click_id REFERENCES slot(id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_dt timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
---DROP TABLE event;
DROP TABLE show_banner_event;
DROP TABLE click_banner_event;
DROP TABLE banner_slot;
DROP TABLE user_group;
DROP TABLE banner;
DROP TABLE slot;
-- +goose StatementEnd
