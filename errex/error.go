package errex

var (
	ErrFail           = Create(-1, "系统错误")
	ErrDupCreateRoom  = Create(-2, "请勿重复创建房间")
	ErrNotExistedRoom = Create(-3, "不存在的房间")
	ErrNotInRoom      = Create(-4, "不在该房间")
	ErrNotOnline      = Create(-5, "对方已掉线")
	ErrNoEnemy        = Create(-6, "请等待对手加入")
	ErrInRoom         = Create(-7, "您已在房间")
	ErrReconnect      = Create(-8, "重连错误")
	ErrNotRoomMaster  = Create(-9, "您不是房主")
	ErrNotCurrentYou  = Create(-10, "还没有轮到您")
	ErrGameOver       = Create(-11, "已分出胜负了")
	ErrInvalidPos     = Create(-12, "无效的位置")
)