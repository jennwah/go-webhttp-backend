create table tasks
(
    id      int auto_increment primary key,
    title   varchar(60)                         not null,
    created timestamp default CURRENT_TIMESTAMP null,
    updated timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);