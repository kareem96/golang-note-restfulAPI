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