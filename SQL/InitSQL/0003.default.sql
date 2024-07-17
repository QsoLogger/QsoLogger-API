CREATE TABLE IF NOT EXISTS dev_tab_prefix_qso_log
(
    id              INTEGER        PRIMARY KEY AUTOINCREMENT,
    userId          INTEGER        NOT NULL, -- COMMENT '用户ID',
    logBookId       INTEGER        NOT NULL, -- COMMENT '用户logBook的ID',
    userCallSign    TEXT           NOT NULL, -- COMMENT '己方呼号',
    remoteCallSign  TEXT           NOT NULL, -- COMMENT '对方呼号',
    userPwr         INTEGER        DEFAULT 5, -- COMMENT '己方发射功率',
    remotePwr       INTEGER        DEFAULT 5, -- COMMENT '对方发射功率',
    userQsl         INTEGER        DEFAULT 0, -- COMMENT '己方QSL卡 0:不需要 1:待发 2:已发 3:已收',
    remoteQsl       INTEGER        DEFAULT 0, -- COMMENT '对方QSL卡 0:不需要 1:待发 2:已发 3:已收',
    band            TEXT           DEFAULT 'WFM', -- COMMENT 'WFM NFM LSB USB CW',
    freq            TEXT           DEFAULT '', -- COMMENT '通信频率',
    userRst         TEXT           DEFAULT '', -- COMMENT '己方信号报告',
    remoteRst       TEXT           DEFAULT '', -- COMMENT '对方信号报告',
    userGrid        TEXT           DEFAULT '', -- COMMENT '己方Grid位置',
    remoteGrid      TEXT           DEFAULT '', -- COMMENT '对方Grid位置',
    userItu         TEXT           DEFAULT '', -- COMMENT '己方ITU分区',
    remoteItu       TEXT           DEFAULT '', -- COMMENT '对方ITU分区',
    userCq          TEXT           DEFAULT '', -- COMMENT '己方CQ分区',
    remoteCq        TEXT           DEFAULT '', -- COMMENT '对方CQ分区',
    userQth         TEXT           DEFAULT '', -- COMMENT '己方Qth',
    remoteQth       TEXT           DEFAULT '', -- COMMENT '对方Qth',
    userGps         TEXT           DEFAULT '', -- COMMENT '己方GPS',
    remoteGps       TEXT           DEFAULT '', -- COMMENT '对方GPS',
    userRig         TEXT           DEFAULT '', -- COMMENT '己方电台型号',
    remoteRig       TEXT           DEFAULT '', -- COMMENT '对方电台型号',
    userAnt         TEXT           DEFAULT '', -- COMMENT '己方天线型号',
    remoteAnt       TEXT           DEFAULT '', -- COMMENT '对方天线型号',
    userComment     TEXT           DEFAULT '', -- COMMENT '用户备注',
    callTimestamp   INTEGER        DEFAULT (strftime('%s','now')),
    createTimestamp INTEGER        DEFAULT (strftime('%s','now')),
    updateTimestamp INTEGER        DEFAULT (strftime('%s','now'))
);


CREATE INDEX IF NOT EXISTS logUID on dev_tab_prefix_qso_log (userId);
CREATE INDEX IF NOT EXISTS logUCallSign on dev_tab_prefix_qso_log (userCallSign);
CREATE INDEX IF NOT EXISTS logRCallSign on dev_tab_prefix_qso_log (remoteCallSign);
CREATE INDEX IF NOT EXISTS logUQsl on dev_tab_prefix_qso_log (userQsl);
CREATE INDEX IF NOT EXISTS logRQsl on dev_tab_prefix_qso_log (remoteQsl);
CREATE INDEX IF NOT EXISTS logCreateTimestamp on dev_tab_prefix_qso_log (createTimestamp);
CREATE INDEX IF NOT EXISTS logUpdateTimestamp on dev_tab_prefix_qso_log (updateTimestamp);



