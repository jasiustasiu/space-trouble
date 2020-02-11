create table bookings
(
    id             serial      not null
        constraint bookings_pk
        primary key,
    first_name     varchar(64) not null,
    last_name      varchar(64) not null,
    gender         varchar(1)  not null,
    birthday       date        not null,
    launchpad_id   varchar(32) not null,
    destination_id varchar(32) not null,
    launch_date    date        not null
);

create unique index bookings_launchpad_id_uindex
    on bookings (launchpad_id, launch_date);

GRANT ALL PRIVILEGES ON TABLE bookings TO space_trouble_user;
GRANT ALL PRIVILEGES ON SEQUENCE bookings_id_seq TO space_trouble_user;