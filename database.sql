CREATE DATABASE golang_note;

CREATE TABLE notes
(
    id int auto_increment,
    title VARCHAR(100) not NULL,
    description VARCHAR(100) not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP NULL,
    primary key (id),
    user_id    int null,
    foreign key fk_notes_user_id (user_id) references users (id)
) engine = InnoDB;

select * from notes;

DESCRIBE notes;

delete from notes;

drop table notes;


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

select * from users;

DESCRIBE users;

drop table users;
