create table subscriptions
(
    id           bigserial primary key,
    service_name varchar(50) not null,
    price        bigint      not null,
    user_id      varchar(50) not null,
    start_date   date        not null,
    end_date     date
);

create index idx_subscriptions_user_service_date on subscriptions (user_id, service_name, start_date, end_date);