CREATE TABLE IF NOT EXISTS dev_tab_prefix_system
(
    param           VARCHAR(255)   NOT NULL,
    value           TEXT           DEFAULT '',
    createTimestamp BIGINT         DEFAULT (UNIX_TIMESTAMP()),
    updateTimestamp BIGINT         DEFAULT (UNIX_TIMESTAMP())
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE UNIQUE INDEX IF NOT EXISTS sysparam on dev_tab_prefix_system (param);

INSERT INTO dev_tab_prefix_system(param, value)
VALUES ('data_structure_version', '0'),
       ('install_time', UNIX_TIMESTAMP()),
       ('release_version', '0.0.0')
ON DUPLICATE KEY    UPDATE
    value = VALUES(value), updateTimestamp=UNIX_TIMESTAMP()
;

