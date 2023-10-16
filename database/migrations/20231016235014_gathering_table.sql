-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS gatherings (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    creator VARCHAR(100) NOT NULL,
    member_id BIGINT UNSIGNED,
    type ENUM('family', 'employee', 'customer') NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    schedule_at TIMESTAMP NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gatherings;
-- +goose StatementEnd
