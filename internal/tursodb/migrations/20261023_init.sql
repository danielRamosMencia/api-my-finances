CREATE TABLE tb_user (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    pid        VARCHAR(64)  NOT NULL,
    active     BOOLEAN      NOT NULL DEFAULT 1,
    name       VARCHAR(255) NOT NULL,
    username   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    token      VARCHAR(255),
    password   VARCHAR(255) NOT NULL,
    type       VARCHAR(20)  NOT NULL CHECK (type IN ('OPERATOR', 'CLIENT')),
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at TEXT,
    deleted_at TEXT
);

CREATE TABLE tb_movement_type (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    pid         VARCHAR(64)  NOT NULL,
    active      BOOLEAN      NOT NULL DEFAULT 1,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    code        VARCHAR(255) NOT NULL,
    created_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at  TEXT,
    deleted_at  TEXT
);

CREATE TABLE tb_account (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    pid        VARCHAR(64)  NOT NULL,
    active     BOOLEAN      NOT NULL DEFAULT 1,
    alias      VARCHAR(255) NOT NULL,
    amount     NUMERIC      NOT NULL DEFAULT 0,
    user_id    INTEGER      NOT NULL,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at TEXT,
    deleted_at TEXT,
    FOREIGN KEY (user_id) REFERENCES tb_user (id)
);

CREATE TABLE tb_movement (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    pid              VARCHAR(64)  NOT NULL,
    active           BOOLEAN      NOT NULL DEFAULT 1,
    code             VARCHAR(255) NOT NULL,
    amount           NUMERIC      NOT NULL,
    account_id       INTEGER      NOT NULL,
    movement_type_id INTEGER      NOT NULL,
    user_id          INTEGER      NOT NULL,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at TEXT,
    deleted_at TEXT,
    FOREIGN KEY (account_id)       REFERENCES tb_account (id),
    FOREIGN KEY (movement_type_id) REFERENCES tb_movement_type (id),
    FOREIGN KEY (user_id)          REFERENCES tb_user (id)
);

CREATE TABLE tb_movement_history (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    pid         VARCHAR(64) NOT NULL,
    amount      NUMERIC     NOT NULL,
    movement_id INTEGER     NOT NULL,
    created_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    FOREIGN KEY (movement_id) REFERENCES tb_movement (id)
);

-- pid: unique and indexed (public identifier)
CREATE UNIQUE INDEX uq_user_pid             ON tb_user (pid);
CREATE UNIQUE INDEX uq_movement_type_pid    ON tb_movement_type (pid);
CREATE UNIQUE INDEX uq_account_pid          ON tb_account (pid);
CREATE UNIQUE INDEX uq_movement_pid         ON tb_movement (pid);
CREATE UNIQUE INDEX uq_movement_history_pid ON tb_movement_history (pid);

-- username and email uniques between not deleted users
CREATE UNIQUE INDEX uq_user_username ON tb_user (username) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX uq_user_email    ON tb_user (email)    WHERE deleted_at IS NULL;

-- deleted_at: to filter by "not deleted"
CREATE INDEX idx_user_deleted_at          ON tb_user (deleted_at);
CREATE INDEX idx_movement_type_deleted_at ON tb_movement_type (deleted_at);
CREATE INDEX idx_account_deleted_at       ON tb_account (deleted_at);
CREATE INDEX idx_movement_deleted_at      ON tb_movement (deleted_at);

-- Foreign keys
CREATE INDEX idx_account_user_id         ON tb_account (user_id);
CREATE INDEX idx_movement_account_id     ON tb_movement (account_id);
CREATE INDEX idx_movement_type_id        ON tb_movement (movement_type_id);
CREATE INDEX idx_movement_user_id        ON tb_movement (user_id);
CREATE INDEX idx_movement_history_mov_id ON tb_movement_history (movement_id);
