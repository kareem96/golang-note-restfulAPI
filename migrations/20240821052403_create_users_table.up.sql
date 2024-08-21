create table users
(
    id int auto_increment,
    name       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at bigint       not null,
    updated_at bigint       not null,
    primary key (id)
) engine = InnoDB;