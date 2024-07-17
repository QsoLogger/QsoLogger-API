CREATE TABLE IF NOT EXISTS dev_tab_prefix_system
(
    param           TEXT           NOT NULL,
    value           TEXT           DEFAULT '',
    createTimestamp INTEGER        DEFAULT (strftime('%s','now')),
    updateTimestamp INTEGER        DEFAULT (strftime('%s','now'))
);

CREATE UNIQUE INDEX IF NOT EXISTS sysparam on dev_tab_prefix_system (param);

INSERT OR REPLACE INTO dev_tab_prefix_system(param, value)
VALUES ('data_structure_version', '0'),
       ('install_time', strftime('%s','now')),
       ('release_version', '0.0.0')
ON CONFLICT (param) DO UPDATE
set value = EXCLUDED.value, updateTimestamp=strftime('%s','now')
;

