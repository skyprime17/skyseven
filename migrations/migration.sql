create table if not exists user
(
    id         varchar(36)          not null
        primary key,
    username   varchar(36)          not null,
    password   varchar(200)         not null,
    created_at datetime             null,
    enabled    tinyint(1) default 1 null,
    constraint table_name_username_uindex
        unique (username)
);

create table if not exists user_upload
(
    id          varchar(36) not null
        primary key,
    user_id     varchar(36) not null,
    file_id     varchar(36) null,
    file_format varchar(12) null,
    created_at  timestamp   null,
    constraint user_upload_file_id_uindex
        unique (file_id),
    constraint user_upload_user_id_fk
        foreign key (user_id) references user (id)
            on delete cascade
)