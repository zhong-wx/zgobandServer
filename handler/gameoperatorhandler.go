package handler

import (
	"../msgPush"
	"../vo"
	"context"
	"../models"
	"fmt"
)



type GameOperatorHandler struct {
}

func NewGameOperatorHandler() *GameOperatorHandler {
	gameOperatorHandler := &GameOperatorHandler{}
	return gameOperatorHandler
}

func (*GameOperatorHandler) Ready(ctx context.Context, account string, deskID int32) (err error) {
	return nil
}

func (p *GameOperatorHandler) PutChess(ctx context.Context, player1 string, player2 string, deskID int32, seatID int8, row int8, column int8) (r int8, err error) {
	var whichToPushMsg string
	var me string
	var other string
	var otherSeatID int
	if seatID==1 {
		me = player1
		other = player2
		whichToPushMsg = player2
		otherSeatID = 2
	} else if(seatID==2) {
		me = player2
		other = player1
		whichToPushMsg = player1
		otherSeatID = 1
	}

	result, err := vo.PutChess(player1, player2, row, column, seatID)
	if err != nil {
		return -1, err
	}

	// put chess succeed
	if result >=0 {
		var gp *vo.GameProcess
		gp = vo.GetGameProcess(player1, player2)
		if gp==nil {
			gp = vo.NewGameProcess(player1, player2)
			vo.AddGameProcess(player1, player2, gp)
		}
		gp.AddStep(row, column)

		msgPush.PushPutChessMassage(whichToPushMsg, int(row), int(column), result)

		// #TODO 统计用时， 判断是否超时
		// #TODO 可以添加观战功能

		// win or draw
		if result > 0 {
			// 取消准备，更新状态，删除棋盘信息,
			models.SetReady(deskID, 1, false)
			models.SetReady(deskID, 2, false)
			models.CreateOrUpdatePlayerStat(me, 1, int(deskID), int(seatID))
			models.CreateOrUpdatePlayerStat(other, 1, int(deskID), int(otherSeatID))
			vo.DeleteChessBoard(player1, player2)
		}

		// 积分统计
		if result == 1 {
			// win
			models.WinUpdate(me)
		} else if result == 2 {
			// draw
			models.DrawUpdate(me)
		}
	}
	return result, nil
}

func (*GameOperatorHandler) TakeBackReq(ctx context.Context, account string, otherSide string, seatID int8) (r bool, err error) {
	var gp *vo.GameProcess
	if seatID == 1 {
		gp = vo.GetGameProcess(account, otherSide)
		if gp.StepCount()<1 {
			return false, nil
		}
	} else if(seatID == 2) {
		gp = vo.GetGameProcess(otherSide, account)
		if gp.StepCount()<2 {
			return false, nil
		}
	} else {
		fmt.Println("error, seatID !=1 || 2")
		return false, fmt.Errorf("error, seatID !=1 || 2")
	}
	msgPush.PushTakeBackReq(otherSide)
	return true, nil
}

func (*GameOperatorHandler) TakeBackRespond(ctx context.Context, player1 string, player2 string, seatID int8, resp bool) (r bool, err error) {
	gp := vo.GetGameProcess(player1, player2)
	var whoReq string
	var whoResp string
	if seatID == 1 {
		whoReq = player2
		whoResp = player1
	} else if seatID == 2 {
		whoReq = player1
		whoResp = player2
	}
	var lastStep []vo.Pos
	if !resp {
		msgPush.PushTakeBackMessage(whoReq, whoResp, nil, false)
		return true, nil
	}
	if (seatID==1 && gp.StepCount()%2==1) || (seatID==2 && gp.StepCount()%2==0) {
		p1, p2, p3 := gp.GetLastThreeRmTwo()
		lastStep = append(lastStep, p1)
		lastStep = append(lastStep, p2)
		lastStep = append(lastStep, p3)
		// 从棋盘拿回最后两个棋子
		vo.TakeBack(player1, player2, p2.Row, p2.Column)
		vo.TakeBack(player1, player2, p3.Row, p3.Column)
	} else if (seatID==1 && gp.StepCount()%2==0) || (seatID==2 && gp.StepCount()%2==1) {
		p1, p2 := gp.GetLastTwoRmOne()
		lastStep = append(lastStep, p1)
		lastStep = append(lastStep, p2)
		// 从棋盘拿回最后一个棋子
		vo.TakeBack(player1, player2, p2.Row, p2.Column)
	}

	msgPush.PushTakeBackMessage(whoReq, whoResp, lastStep, true)
	return true, nil
}

func (*GameOperatorHandler) LoseReq(ctx context.Context, player1 string, player2 string, deskID int32, seatID int8) (err error) {
	var me string
	var otherSide string
	var otherSeatID int8
	if seatID == 1 {
		me = player1
		otherSide = player2
		otherSeatID = 2
	}else {
		me = player2
		otherSide = player1
		otherSeatID = 1
	}
	//统计分数，修改状态，删除棋盘
	models.WinUpdate(otherSide)
	models.LoseUpdate(me)
	vo.DeleteChessBoard(player1, player2)
	models.SetReady(deskID, 1, false)
	models.SetReady(deskID, 2, false)
	models.CreateOrUpdatePlayerStat(me, 1, int(deskID), int(seatID))
	models.CreateOrUpdatePlayerStat(otherSide, 1, int(deskID), int(otherSeatID))

	// 认输消息推送
	msgPush.PushReqLoseMessage(otherSide)
	return nil
}

// #TODO noly require otherSide
func (*GameOperatorHandler) DrawReq(ctx context.Context, account string, otherSide string, seatID int8) (err error) {
	msgPush.PushDrawReqMessage(otherSide)
	return nil
}

func (*GameOperatorHandler) DrawResponse(ctx context.Context, player1 string, player2 string, deskID int32, seatID int8, resp bool) (err error) {
	var me string
	var otherSide string
	var otherSeatID int8
	if seatID == 1 {
		me = player1
		otherSide = player2
		otherSeatID = 2
	}else {
		me = player2
		otherSide = player1
		otherSeatID = 1
	}
	if resp {
		models.DrawUpdate(player1)
		models.DrawUpdate(player1)
		vo.DeleteChessBoard(player1, player2)
		models.SetReady(deskID, 1, false)
		models.SetReady(deskID, 2, false)
		models.CreateOrUpdatePlayerStat(me, 1, int(deskID), int(seatID))
		models.CreateOrUpdatePlayerStat(otherSide, 1, int(deskID), int(otherSeatID))

		msgPush.PushDrawRespondMessage(otherSide, true)
	} else {
		msgPush.PushDrawRespondMessage(otherSide, false)
	}
	return nil
}

func (*GameOperatorHandler) SaveGame(ctx context.Context, account string) (err error) {
	return nil
}

func (*GameOperatorHandler) SendChatText(ctx context.Context, toAccount string, account string, text string) (err error) {
	msgPush.PushChatTextMessage(toAccount, text)
	return nil
}