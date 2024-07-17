package httpAPI

// Http输出固定结构
type HttpRes struct {
	Code int         `json:"code" example:"9999"`             // 错误码，0表示正确，其他都为异常
	Msg  string      `json:"msg" example:"The Error Message"` // 错误信息，code非0时填充
	Data interface{} `json:"data"`                            // 业务返回内容，code非0时为nil，code为0时，基于业务填充
}

// Qso LogBook的基本信息字段定义
type QsoLogBook struct {
	LogBookId       int    `json:"logBookId" example:"3"`                // logBook的ID
	UserId          int    `json:"userId" example:"1"`                   // 用户ID
	UserCallSign    string `json:"userCallSign" example:"BI1UJZ"`        // 己方呼号
	UserGrid        string `json:"userGrid" example:"ON80cb"`            // 己方Grid位置
	UserItu         string `json:"userItu" example:"44"`                 // 己方ITU分区
	UserCq          string `json:"userCq" example:"24"`                  // 己方CQ分区
	UserQth         string `json:"userQth" example:"Beijing"`            // 己方Qth
	UserGps         string `json:"userGps" example:""`                   // 己方GPS
	UserComment     string `json:"userComment" example:""`               // 用户备注
	CallTimestamp   int    `json:"callTimestamp" example:"1722441600"`   // 通信时间,2024-08-01表示为：1722441600
	CreateTimestamp int    `json:"createTimestamp" example:"1722441600"` // 记录创建时间，新增时留空
	UpdateTimestamp int    `json:"updateTimestamp" example:"1722441600"` // 记录最后更新时间，新增时留空
}

// Qso Log记录的字段定义
type QsoLog struct {
	LogId           int    `json:"logId" example:"203"`                  // log记录ID，新增时留空
	LogBookId       int    `json:"logBookId" example:"3"`                // logBook的ID
	UserId          int    `json:"userId" example:"1"`                   // 用户ID
	UserCallSign    string `json:"userCallSign" example:"BI1UJZ"`        // 己方呼号
	RemoteCallSign  string `json:"remoteCallSign" example:"BI1UHX"`      // 对方呼号
	UserPwr         int    `json:"userPwr" example:"5"`                  // 己方发射功率
	RemotePwr       int    `json:"remotePwr" example:"5"`                // 对方发射功率
	UserQsl         int    `json:"userQsl" example:"0"`                  // 己方QSL卡 0:不需要 1:待发 2:已发 3:已收
	RemoteQsl       int    `json:"remoteQsl" example:"0"`                // 对方QSL卡 0:不需要 1:待发 2:已发 3:已收
	Band            string `json:"band" example:"WFM"`                   // WFM NFM LSB USB CW
	Freq            string `json:"freq" example:"438500000"`             // 通信频率
	UserRst         string `json:"userRst" example:"59"`                 // 己方信号报告
	RemoteRst       string `json:"remoteRst" example:"59"`               // 对方信号报告
	UserGrid        string `json:"userGrid" example:"ON80cb"`            // 己方Grid位置
	RemoteGrid      string `json:"remoteGrid" example:"ON80db"`          // 对方Grid位置
	UserItu         string `json:"userItu" example:"44"`                 // 己方ITU分区
	RemoteItu       string `json:"remoteItu" example:"44"`               // 对方ITU分区
	UserCq          string `json:"userCq" example:"24"`                  // 己方CQ分区
	RemoteCq        string `json:"remoteCq" example:"24"`                // 对方CQ分区
	UserQth         string `json:"userQth" example:"Beijing"`            // 己方Qth
	RemoteQth       string `json:"remoteQth" example:"Beijing"`          // 对方Qth
	UserGps         string `json:"userGps" example:""`                   // 己方GPS
	RemoteGps       string `json:"remoteGps" example:""`                 // 对方GPS
	UserRig         string `json:"userRig" example:"5RH"`                // 己方电台型号
	RemoteRig       string `json:"remoteRig" example:"K6"`               // 对方电台型号
	UserAnt         string `json:"userAnt" example:"SG-770"`             // 己方天线型号
	RemoteAnt       string `json:"remoteAnt" example:"SG-771"`           // 对方天线型号
	UserComment     string `json:"userComment" example:""`               // 用户备注
	CallTimestamp   int    `json:"callTimestamp" example:"1722441600"`   // 通信时间,2024-08-01表示为：1722441600
	CreateTimestamp int    `json:"createTimestamp" example:"1722441600"` // 记录创建时间，新增时留空
	UpdateTimestamp int    `json:"updateTimestamp" example:"1722441600"` // 记录最后更新时间，新增时留空
}
