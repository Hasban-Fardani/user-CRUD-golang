use user_crud_go;

create table if not exists Users (
	id int auto_increment ,
    name varchar(30) unique,
    email varchar(30) unique,
    password varchar(15),
    create_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    primary key(id)
);

create table if not exists ChangePasswordHistory (
	id_user int,
    password_lama varchar(15),
    password_baru varchar(15),
    change_at timestamp default current_timestamp,
    foreign key (id_user) references Users(id)
);
