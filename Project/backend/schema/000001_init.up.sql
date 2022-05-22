-- Product
create SEQUENCE product_id_seq;

create table if not exists product
(
    product_id  integer   default nextval('product_id_seq'::regclass) not null
        constraint product_pk
            primary key,
    shop_id     integer                                               not null,
    price       numeric                                               not null,
    name        varchar                                               not null,
    description varchar,
    is_auction  bool      default False                               not null,
    discount    decimal   default 0                                   not null,
    created_at  timestamp default now()                               not null,
    selled_at   timestamp
);

alter table product owner to qyrtpash;

-- Shop owner
create SEQUENCE users_customer_id_seq;

create table if not exists owner
(
    shop_id             integer default nextval('users_customer_id_seq'::regclass) not null
        constraint users_pk
            primary key,
    phone               bigint                                                     not null,
    email               varchar                                                    not null,
    name                varchar                                                    not null,
    password            varchar                                                    not null,
    address             varchar                                                    not null,
    social_network_uuid varchar                                                    not null
);

alter table owner
    owner to qyrtpash;

create unique index users_email_uindex
    on owner (email);

create unique index users_phone_uindex
    on owner (phone);

-- Customer
create table if not exists customer
(
    id        serial
        constraint customer_pk
            primary key,
    phone     bigint  not null,
    email     varchar not null,
    name      varchar not null,
    password  varchar not null,
    card_uuid varchar
);

alter table customer
    owner to qyrtpash;

create unique index customer_email_uindex
    on customer (email);

create unique index customer_phone_uindex
    on customer (phone);

-- Card
create table if not exists card
(
    card_uuid  varchar   not null,
    name       varchar       not null,
    pan        bigint,
    expiration timestamp not null
);

comment on column card.name is 'Cardholder name';

comment on column card.pan is 'Primary Account Number - the 16 digit number on the front of the card';

-- Social network
create table if not exists social_network
(
    social_network_uuid varchar not null,
    telegram            varchar,
    whatsapp            varchar,
    instagram           varchar
);

-- Order
create table if not exists "order"
(
    order_id    serial,
    customer_id int not null,
    product_id  int not null,
    shop_id     int not null
);