CREATE TABLE IF NOT EXISTS dev_tab_prefix_user
(
    id              BIGSERIAL      PRIMARY KEY,
    name            VARCHAR(255)   NOT NULL,
    email           VARCHAR(255)   DEFAULT '',
    mobile          VARCHAR(255)   DEFAULT '',
    displayName     VARCHAR(255)   DEFAULT '',
    status          INTEGER        DEFAULT 0, -- COMMENT '-1:删除 0:新增，未验证 1:已验证',
    createTimestamp INTEGER        DEFAULT (CAST(CURRENT_TIMESTAMP() AS INT)),
    updateTimestamp INTEGER        DEFAULT (CAST(CURRENT_TIMESTAMP() AS INT))
);

CREATE UNIQUE INDEX IF NOT EXISTS userName on dev_tab_prefix_user (name);
CREATE INDEX IF NOT EXISTS userEmail on dev_tab_prefix_user (email);
CREATE INDEX IF NOT EXISTS userMobile on dev_tab_prefix_user (mobile);
CREATE INDEX IF NOT EXISTS userCreateTimestamp on dev_tab_prefix_user (createTimestamp);
CREATE INDEX IF NOT EXISTS userUpdateTimestamp on dev_tab_prefix_user (updateTimestamp);


INSERT INTO dev_tab_prefix_user(id, name, displayName)
VALUES (1, 'user-1', 'USER-1'),
       (2, 'user-2', 'USER-2'),
       (3, 'user-3', 'USER-3'),
       (4, 'user-4', 'USER-4'),
       (5, 'user-5', 'USER-5'),
       (6, 'user-6', 'USER-6'),
       (7, 'user-7', 'USER-7'),
       (8, 'user-8', 'USER-8'),
       (9, 'user-9', 'USER-9'),
       (10, 'user-10', 'USER-10')
ON CONFLICT (id) DO UPDATE
set name=EXCLUDED.name, displayName=EXCLUDED.displayName, updateTimestamp=CAST(CURRENT_TIMESTAMP() AS INT)
;

