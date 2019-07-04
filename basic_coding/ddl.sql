create table user_reviews
(
    id         serial not null
        constraint user_reviews_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id   integer,
    product_id integer,
    user_id    integer,
    rating     integer,
    review     text
);

alter table user_reviews
    owner to postgres;

create index idx_user_reviews_deleted_at
    on user_reviews (deleted_at);

