package errno

var (
	OK = &Errno{Code:200,Message:"OK"}
	NotFound = &Errno{Code:404,Message:"没有找到您要查找的信息"}
	InternalServerError = &Errno{Code:1001,Message:"Internal server error."}
	ErrBind = &Errno{Code:1002,Message:"Binding request body to the struct failed."}
	ErrUserNotFound = &Errno{Code:1003,Message:"The user not found."}
	ErrPasswordIncorrect = &Errno{Code:1004,Message:"The password not incorrect."}
	ErrTokenInvilid = &Errno{Code:1005,Message:"Error auth token is invilid."}
	ErrToken = &Errno{Code:1006,Message:"Error occurred while signing the JSON web token."}
	ErrDatabase = &Errno{Code:1007,Message:"Database error."}
	ErrValidation = &Errno{Code:1008,Message:"Validation failed."}
	ErrNotMatch = &Errno{Code:1009,Message:"密码不一致."}
	ErrEquipNotFound = &Errno{Code:1010,Message:"设备不匹配."}
	ErrRoomNotFound = &Errno{Code:1011,Message:"会议室不存在."}
	ErrNotAllow = &Errno{Code:1012,Message:"不能修改或删除别人创建的记录"}
	ErrPermission = &Errno{Code:403,Message:"无此权限,禁止访问"}
	ErrExist = &Errno{Code:1013,Message:"记录已存在"}
	ErrOutRange = &Errno{Code:1014,Message:"预定时间超出范围"}
)