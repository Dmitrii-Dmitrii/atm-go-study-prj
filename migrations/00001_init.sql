create type operation_type as enum
    (
        'withdraw',
        'refill',
        'incoming_transfer',
        'outgoing_transfer'
);

create table accounts
(
    account_id uuid primary key,
    account_number text not null unique,
    account_pin int not null,
    account_balance money not null
);

create table admins
(
    admin_id uuid primary key,
    admin_number text not null unique,
    admin_password text not null
);

create table transactions
(
    transaction_id uuid primary key,
    account_id uuid not null references accounts(account_id),
    transaction_type operation_type not null,
    transaction_value money not null,
    transaction_date timestamp
);

SELECT datname FROM pg_database;

insert into accounts(account_id, account_number, account_pin, account_balance)
values (gen_random_uuid(), 1111, 1111, 10000)