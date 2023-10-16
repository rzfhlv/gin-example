-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invitations (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    member_id BIGINT UNSIGNED,
    gathering_id BIGINT UNSIGNED,
    status ENUM('pending', 'accept', 'reject') DEFAULT 'pending' NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (member_id) REFERENCES members(id),
    FOREIGN KEY (gathering_id) REFERENCES gatherings(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invitations;
-- +goose StatementEnd
