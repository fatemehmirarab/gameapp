-- +migrate Up

CREATE TABLE users (
    id int not null primary key auto_increment,
    phone_number varchar(255) not null unique,
    name varchar(255) not null,
    password varchar(255) not null
    /*created_at timestamp default current_timestamp*/

);

-- +migrate Down
DROP TABLE users