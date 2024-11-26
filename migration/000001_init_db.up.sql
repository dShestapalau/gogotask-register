CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE user_profile(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    last_name varchar(50),
    first_name varchar(50),
    middle_name varchar(50),
    locale varchar(50),
    status varchar(50),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE user_credentials(
    user_id UUID primary key references user_profile(id),
    login varchar(50),
    password varchar(50),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
)
