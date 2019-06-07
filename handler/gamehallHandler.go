package handler

import (
	"context"
	"../gen-go/zgobandRPC"
	"../models"
	"../msgPush"
	"../vo"
)

type GameHallHandler struct{}

func NewGameHallHandler() *GameHallHandler {
	return &GameHallHandler{}
}

func (*GameHallHandler) SitDown(ctx context.Context, account string, deskID int32, seat int32) (r bool, err error) {
	succ, err := models.AddSeat(deskID, seat, account)
	if(!succ) {
		return succ, err
	}

	models.CreateOrUpdatePlayerStat(account, 1, int(deskID), int(seat))

	player := models.QueryPlayer(account)
	// #TODO 推送失败是否影响入座成功？
	msgPush.PushSitDownMessage(account, int(deskID), int(seat), player)

	return succ, err
}

func (*GameHallHandler) AutoMatch(ctx context.Context, account string) (r int32, err error) {
	return 0, nil
}

func (*GameHallHandler) GetSavedGame(ctx context.Context, account string, gameName string) (r string, err error) {
	return "fds", nil
}

func (*GameHallHandler) GetSavedGameList(ctx context.Context, account string) (r []string, err error) {
	return []string{"1"}, nil
}

func (*GameHallHandler) GetDeskList(ctx context.Context) (r []*zgobandRPC.Desk, err error) {
	desks, err := models.QueryAllDesk()
	for i:=0; i<len(desks); i++ {
		var ready1 bool = true
		var ready2 bool = true
		if desks[i].Ready1 == nil || desks[i].Ready1[0] == 0 {
			ready1 = false
		}
		if desks[i].Ready2 == nil || desks[i].Ready2[0] == 0 {
			ready2 = false
		}
		d := &zgobandRPC.Desk{DeskID:desks[i].Deskid, Player1:desks[i].Player1, Player2:desks[i].Player2, Ready1:ready1, Ready2:ready2}
		r = append(r, d)
	}
	return r, err
}
// #TODO update ready
func (*GameHallHandler) LeaveSeat(ctx context.Context, account string, deskID int32, seatID int32) (r int32, err error) {
	desk, err := models.GetDesk(deskID)
	if err != nil {
		return -1, err
	}
	err = models.DelSeat(deskID, seatID)
	if err != nil {
		return -1, err
	}

	ret := 0
	if desk.Player1!="" && desk.Player2!="" && desk.Ready2[0]==1 && desk.Ready1[0]==1 {
		var otherSide string
		if seatID == 1 {
			otherSide = desk.Player2
		} else if(seatID == 2) {
			otherSide = desk.Player1
		}

		// 计算输赢积分
		models.EscapeUpdate(account)
		vo.DeleteChessBoard(desk.Player1, desk.Player2)
		models.WinUpdate(otherSide)
		ret = 1
	}
	models.CreateOrUpdatePlayerStat(account, 0, int(deskID), int(seatID))

	msgPush.PushSitDownMessage("", int(deskID), int(seatID), nil)
	return int32(ret), nil
}

func (*GameHallHandler) SetReady(ctx context.Context, account string, deskID int32, seatID int32, isReady bool) (err error) {
	err = models.SetReady(deskID, seatID, isReady)
	if err != nil {
		return err
	}
	var anotherSeat int32
	if seatID == 1 {
		anotherSeat = 2
	} else {
		anotherSeat = 1
	}

	_account, _isReady, _ := models.GetSeatAccountInfo(deskID, anotherSeat)
	var player1 string
	var player2 string
	if seatID == 1 {
		player1 = account
		player2 = _account
	} else if seatID == 2 {
		player1 = _account
		player2 = account
	}
	if isReady && _isReady{
		// 清空上一局下棋过程
		models.CreateOrUpdatePlayerStat(account, 3, int(deskID), int(seatID))
		vo.DeleteGameProcess(player1, player2)
	} else if(isReady) {
		models.CreateOrUpdatePlayerStat(account, 2, int(deskID), int(seatID))
	}
	msgPush.PushReadyChangeMessage(account, int(deskID), int(seatID), isReady)
	return err
}

func (*GameHallHandler) GetSeatInfo(ctx context.Context, deskID int32, seatID int32) (r *zgobandRPC.PlayerInfo, err error) {
	account, isReady, err := models.GetSeatAccountInfo(deskID, seatID)
	if err != nil {
		return nil, err
	}
	if account == "" {
		return nil, nil
	}
	player := models.QueryPlayer(account)
	r = &zgobandRPC.PlayerInfo{}
	r.Nickname = player.Nickname
	r.EscapeRound = int32(player.Escaperound)
	r.LoseRound = int32(player.Loseround)
	r.WinRound = int32(player.Winround)
	r.Core = int32(player.Score)
	r.DrawRound = int32(player.Drawround)
	r.Account = account
	r.IsReady = &isReady

	return r, nil
}