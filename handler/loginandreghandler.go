package handler

import (
	"../models"
	"context"
	"../gen-go/zgobandRPC"
	"../vo"
	"fmt"
)

type LoginAndRegHandler struct{}



func NewLoginAndRegHandler() *LoginAndRegHandler {
	return &LoginAndRegHandler{}
}

func (*LoginAndRegHandler) Login(ctx context.Context, account string, password string) (player *zgobandRPC.PlayerInfo, err error) {
	// #TODO 区分已经登陆和密码错误
	// #TODO 是否要记录登陆状态，是保存在内存中还是数据库中

	if(models.IsBlock(account)) {
		return nil, fmt.Errorf("is blocked")
	}
	if(vo.Logined(account)) {
		return nil, nil
	}

	playerInfo := models.QueryPlayer(account)
	if playerInfo != nil && playerInfo.Password == password{
		player = &zgobandRPC.PlayerInfo{}
		player.Nickname = playerInfo.Nickname
		player.Core = int32(playerInfo.Score)
		player.DrawRound = int32(playerInfo.Drawround)
		player.WinRound = int32(playerInfo.Winround)
		player.LoseRound = int32(playerInfo.Loseround)
		player.EscapeRound = int32(playerInfo.Escaperound)

		vo.Login(account)
		return player, nil
	}
	return nil, nil
}

func (*LoginAndRegHandler) Reg(ctx context.Context, account string, password string, nickname string) (r bool, err error) {
	e := models.CreatePlayer(account, password, nickname)
	if e == nil {
		return true, nil
	} else {
		return false, nil
	}
}