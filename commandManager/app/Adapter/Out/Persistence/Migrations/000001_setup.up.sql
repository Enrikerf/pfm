create table go_project.tasks
(
    id             bigint auto_increment primary key,
    uuid           varchar(255)                        not null,
    host           varchar(30)                         not null,
    port           varchar(30)                         not null,
    mode           varchar(30)                         not null,
    status         varchar(30)                         not null,
    created_at     timestamp default CURRENT_TIMESTAMP null,
    updated_at     timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    execution_mode varchar(30)                         not null,
    constraint uuid unique (uuid)
);

create table go_project.batches
(
    id         bigint auto_increment
        primary key,
    uuid       varchar(255)                        not null,
    task_id    bigint                              not null,
    task_uuid  varchar(255)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint uuid unique (uuid),
    index IDX_BATCHES_C_TASK_ID (task_id),
    constraint FK_BATCHES_C_TASK_ID
        foreign key (task_id) references go_project.tasks (id)
            on delete cascade
);

create table go_project.commands
(
    id         bigint auto_increment
        primary key,
    uuid       varchar(255)                        not null,
    task_id    bigint                              not null,
    task_uuid  varchar(255)                        not null,
    name       varchar(30)                         not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint uuid unique (uuid),
    index IDX_COMMANDS_C_TASK_ID (task_id),
    constraint FK_COMMANDS_C_TASK_ID
        foreign key (task_id) references go_project.tasks (id)
            on delete cascade
);

create table go_project.results
(
    id         bigint auto_increment
        primary key,
    uuid       varchar(255)                        not null,
    content    text                                null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    batch_id   bigint                              not null,
    batch_uuid varchar(255)                        not null,
    constraint uuid unique (uuid),
    index IDX_RESULTS_C_BATCH_ID (batch_id),
    constraint FK_RESULTS_C_BATCH_ID
        foreign key (batch_id) references go_project.batches (id)
            on delete cascade
);

