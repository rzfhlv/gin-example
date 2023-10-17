-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS attendee (
    member_id BIGINT UNSIGNED,
    gathering_id BIGINT UNSIGNED,

    FOREIGN KEY (member_id) REFERENCES members(id),
    FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attendee;
-- +goose StatementEnd
