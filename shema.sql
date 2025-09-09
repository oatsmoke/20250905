create table subscriptions
(
    id           bigserial primary key,
    service_name varchar(50) not null,
    price        bigint      not null,
    user_id      varchar(50) not null,
    start_date   varchar(50) not null,
    end_date     varchar(50) not null
);