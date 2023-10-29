CREATE TABLE user_account (
    uid UUID NOT NULL DEFAULT gen_random_uuid () PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    username VARCHAR(100) DEFAULT NULL,
    email VARCHAR(100) NOT NULL,
    gender CHAR(1) CHECK (gender IN ('M', 'F', 'O')),
    avatar_uri VARCHAR DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE user_login_internal (
    id BIGSERIAL NOT NULL,
    uid UUID REFERENCES user_account(uid) ON DELETE CASCADE,
    grant_type VARCHAR(20) CHECK (grant_type <> ''),
    email VARCHAR(100) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    phone_number VARCHAR(50) NOT NULL,
    hashed_password TEXT NOT NULL,
    salt TEXT NOT NULL,
    PRIMARY KEY (id, uid)
);
CREATE TABLE external_login (
    sub TEXT NOT NULL UNIQUE,
    uid UUID REFERENCES user_account(uid) ON DELETE CASCADE,
    grant_type VARCHAR(20) CHECK (grant_type <> ''),
    provider VARCHAR(20) CHECK (
        provider IN (
            'google',
            'facebook',
            'twitter',
            'apple',
            'github'
        )
    ),
    PRIMARY KEY (sub, uid)
);