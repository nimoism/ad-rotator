-- +goose Up
-- +goose StatementBegin
INSERT INTO slot (name) VALUES
('top'),
('right'),
('bottom');

INSERT INTO banner (name) VALUES
('banner0'),
('banner1'),
('banner2'),
('banner3'),
('banner4');

INSERT INTO user_group (name) VALUES
('0-18'),
('18-25'),
('25-40'),
('40-65'),
('65+');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM slot WHERE name IN ('top', 'right', 'bottom');
DELETE FROM banner WHERE name IN ('banner0', 'banner1', 'banner2', 'banner3', 'banner4');
DELETE FROM user_group WHERE name IN ('0-18', '18-25', '25-40', '40-65', '65+');
-- +goose StatementEnd
