CREATE TABLE IF NOT EXISTS dev_tab_prefix_user
(
    id              INT            PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(255)   NOT NULL,
    email           VARCHAR(255)   DEFAULT '',
    mobile          VARCHAR(255)   DEFAULT '',
    displayName     VARCHAR(255)   DEFAULT '',
    status          INT            DEFAULT 0     COMMENT '-1:删除 0:新增，未验证 1:已验证',
    createTimestamp INT            DEFAULT UNIX_TIMESTAMP(),
    updateTimestamp INT            DEFAULT UNIX_TIMESTAMP()
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

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
ON DUPLICATE KEY    UPDATE
    name=VALUES(name), displayName=VALUES(displayName), updateTimestamp=UNIX_TIMESTAMP()
;

