CREATE TABLE IF NOT EXISTS users
(
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    full_name   VARCHAR(45)  NOT NULL,
    phone       VARCHAR(15)  NOT NULL,
    gender      VARCHAR(15)  NOT NULL,
    email       VARCHAR(100) NOT NULL,
    password    VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    CONSTRAINT UQ_user_email UNIQUE (email),
    CONSTRAINT UQ_user_phone UNIQUE (phone)
    );