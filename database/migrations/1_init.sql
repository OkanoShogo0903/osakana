-- +goose Up
CREATE TABLE users (
    ip CHAR(15) NOT NULL,
    background INTEGER UNSIGNED NOT NULL default 0,
    plant00 INTEGER UNSIGNED default 0,
    plant01 INTEGER UNSIGNED default 0,
    plant02 INTEGER UNSIGNED default 0,
    plant03 INTEGER UNSIGNED default 0,
    plant04 INTEGER UNSIGNED default 0,
    plant05 INTEGER UNSIGNED default 0,
    plant06 INTEGER UNSIGNED default 0,
    plant07 INTEGER UNSIGNED default 0,
    plant08 INTEGER UNSIGNED default 0,
    plant09 INTEGER UNSIGNED default 0,
    creature00 INTEGER UNSIGNED default 0,
    creature01 INTEGER UNSIGNED default 0,
    creature02 INTEGER UNSIGNED default 0,
    creature03 INTEGER UNSIGNED default 0,
    creature04 INTEGER UNSIGNED default 0,
    creature05 INTEGER UNSIGNED default 0,
    creature06 INTEGER UNSIGNED default 0,
    creature07 INTEGER UNSIGNED default 0,
    creature08 INTEGER UNSIGNED default 0,
    creature09 INTEGER UNSIGNED default 0,
    creature10 INTEGER UNSIGNED default 0,
    creature11 INTEGER UNSIGNED default 0,
    creature12 INTEGER UNSIGNED default 0,
    creature13 INTEGER UNSIGNED default 0,
    creature14 INTEGER UNSIGNED default 0,
    creature15 INTEGER UNSIGNED default 0,
    creature16 INTEGER UNSIGNED default 0,
    creature17 INTEGER UNSIGNED default 0,
    creature18 INTEGER UNSIGNED default 0,
    creature19 INTEGER UNSIGNED default 0,
    PRIMARY KEY(ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE users;

