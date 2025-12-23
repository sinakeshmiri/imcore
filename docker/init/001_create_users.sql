CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(64) PRIMARY KEY,
    email VARCHAR(320) NOT NULL UNIQUE,
    fullname VARCHAR(320),
    password_hash TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS roles (
    rolename VARCHAR(64) PRIMARY KEY,
    description VARCHAR(640),
    owner_username VARCHAR(64) REFERENCES users(username),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS user_roles (
    username VARCHAR(64) NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    rolename VARCHAR(64) NOT NULL REFERENCES roles(rolename) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (username, rolename)
    );

CREATE INDEX IF NOT EXISTS idx_user_roles_username ON user_roles(username);
CREATE INDEX IF NOT EXISTS idx_user_roles_rolename ON user_roles(rolename);

CREATE TABLE IF NOT EXISTS applications (
    application_id UUID PRIMARY KEY,
    rolename VARCHAR(64) NOT NULL REFERENCES roles(rolename) ON DELETE CASCADE,
    applicant_username VARCHAR(64) NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    owner_username VARCHAR(64) NOT NULL,

    status VARCHAR(16) NOT NULL,
    reason VARCHAR(640),
    decision_note VARCHAR(640),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    decided_at TIMESTAMPTZ
    );


CREATE INDEX IF NOT EXISTS idx_applications_applicant ON applications(applicant_username);
CREATE INDEX IF NOT EXISTS idx_applications_owner ON applications(owner_username);
CREATE INDEX IF NOT EXISTS idx_applications_status ON applications(status);
CREATE INDEX IF NOT EXISTS idx_applications_role ON applications(rolename);
