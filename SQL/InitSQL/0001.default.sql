CREATE TABLE IF NOT EXISTS dev_tab_prefix_admin
(
    id              INTEGER        PRIMARY KEY AUTOINCREMENT,
    ssoName         TEXT           NOT NULL,
    email           TEXT           DEFAULT '',
    mobile          TEXT           DEFAULT '',
    displayName     TEXT           DEFAULT '',
    status          INTEGER        DEFAULT 0, -- COMMENT '-1:删除 0:未启用，暂停 1:系统管理员 2:系统高级运营 3:系统普通运营',
    createTimestamp INTEGER        DEFAULT (strftime('%s','now')),
    updateTimestamp INTEGER        DEFAULT (strftime('%s','now'))
);

CREATE UNIQUE INDEX IF NOT EXISTS adminSsoName on dev_tab_prefix_admin (ssoName);
CREATE INDEX IF NOT EXISTS adminEmail on dev_tab_prefix_admin (email);
CREATE INDEX IF NOT EXISTS adminMobile on dev_tab_prefix_admin (mobile);
CREATE INDEX IF NOT EXISTS adminCreateTimestamp on dev_tab_prefix_admin (createTimestamp);
CREATE INDEX IF NOT EXISTS adminUpdateTimestamp on dev_tab_prefix_admin (updateTimestamp);


INSERT INTO dev_tab_prefix_admin(id, ssoName, displayName,status)
VALUES (1, 'cnzhangquan', 'ZhangQuan', 1),
       (10, 'admin-10', 'ADMIN-10', 0)
ON CONFLICT (id) DO UPDATE
set ssoName=EXCLUDED.ssoName, displayName=EXCLUDED.displayName, updateTimestamp=strftime('%s','now')
;

