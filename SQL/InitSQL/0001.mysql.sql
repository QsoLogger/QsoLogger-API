CREATE TABLE IF NOT EXISTS dev_tab_prefix_admin
(
    id              INT            PRIMARY KEY AUTO_INCREMENT,
    ssoName         VARCHAR(255)   NOT NULL,
    email           VARCHAR(255)   DEFAULT '',
    mobile          VARCHAR(255)   DEFAULT '',
    displayName     VARCHAR(255)   DEFAULT '',
    status          INT            DEFAULT 0     COMMENT '-1:删除 0:未启用，暂停 1:系统管理员 2:系统高级运营 3:系统普通运营',
    createTimestamp INT            DEFAULT UNIX_TIMESTAMP(),
    updateTimestamp INT            DEFAULT UNIX_TIMESTAMP()
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE UNIQUE INDEX IF NOT EXISTS adminSsoName on dev_tab_prefix_admin (ssoName);
CREATE INDEX IF NOT EXISTS adminEmail on dev_tab_prefix_admin (email);
CREATE INDEX IF NOT EXISTS adminMobile on dev_tab_prefix_admin (mobile);
CREATE INDEX IF NOT EXISTS adminCreateTimestamp on dev_tab_prefix_admin (createTimestamp);
CREATE INDEX IF NOT EXISTS adminUpdateTimestamp on dev_tab_prefix_admin (updateTimestamp);


INSERT INTO dev_tab_prefix_admin(id, ssoName, displayName,status)
VALUES (1, 'cnzhangquan', 'ZhangQuan', 1),
       (10, 'admin-10', 'ADMIN-10', 0)
ON DUPLICATE KEY    UPDATE
    ssoName=VALUES(ssoName), displayName=VALUES(displayName), updateTimestamp=UNIX_TIMESTAMP()
;

