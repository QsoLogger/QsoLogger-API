CREATE TABLE IF NOT EXISTS dev_tab_prefix_system
(
    param           VARCHAR(255)   NOT NULL,
    value           TEXT           DEFAULT '',
    createTimestamp INTEGER        DEFAULT (CAST(CURRENT_TIMESTAMP() AS INT)),
    updateTimestamp INTEGER        DEFAULT (CAST(CURRENT_TIMESTAMP() AS INT))
);

CREATE UNIQUE INDEX IF NOT EXISTS sysparam on dev_tab_prefix_system (param);

INSERT INTO dev_tab_prefix_system(param, value)
VALUES ('data_structure_version', '0'),
       ('install_time', CAST(CURRENT_TIMESTAMP() AS INT)::varchar(255)),
       ('release_version', '0.0.0')
ON CONFLICT (param) DO UPDATE
set value = EXCLUDED.value, updateTimestamp=CAST(CURRENT_TIMESTAMP() AS INT)
;

