package human

import "github.com/zqhhh/gomoku/internal/chessboard"

type ResRoomMsg struct {
	Action     RoomAction      `json:"action"`
	RoomNumber int             `json:"room_number"`
	IsBlack    bool            `json:"is_black"`
	XY         []chessboard.XY `json:"xy"`
	NowWalk    chessboard.XY   `json:"now_walk"`
}

type ResRoomMsgRegret struct {
	Action RoomAction      `json:"action"`
	XY     []chessboard.XY `json:"xy"`
}
