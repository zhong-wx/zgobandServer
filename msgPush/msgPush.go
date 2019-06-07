package msgPush

import (
	"../models"
	"./utils"
	"fmt"
	"../vo"
)

func pushMessageToEveryOne(account string, messageMap map[string]interface{}) {
	bs, err := utils.CreatMessage(messageMap)
	if err != nil {
		fmt.Println("CreatMessage:", err.Error())
		panic("CreatMessage:"+ err.Error())
	}
	for k, v := range accountConnMap {
		if err := utils.Writen(v, bs); err != nil {
			fmt.Println("messageType:", messageMap["messageType"], " writen fail:", err.Error())
			v.Close()
			delete(accountConnMap, k)
			return
		}
		fmt.Println(string(bs), "is sended to ", k)
	}
}

func pushMessageToOne(account string, messageMap map[string]interface{}) {
	bs, err := utils.CreatMessage(messageMap)
	if err != nil {
		fmt.Println("CreatMessage:", err.Error())
		panic("CreatMessage:"+ err.Error())
	}
	if conn, ok := accountConnMap[account]; ok {
		err := utils.Writen(conn, bs);
		if err != nil {
			fmt.Println("messageType:", messageMap["messageType"], " writen fail:", err.Error())
			conn.Close()
			delete(accountConnMap, account)
			return
		}
		fmt.Println(string(bs), "is sended")
	}
}

func PushSitDownMessage(account string, deskID, seatID int, player *models.Player) {
	sitDownMsg := map[string]interface{}{"messageType":0, "account":account, "deskID": deskID, "seatID":seatID}
	if account != "" {
		sitDownMsg["playerInfo"] = map[string]interface{}{"account":player.Account, "nickname":player.Nickname, "score":player.Score, "winRound":player.Winround, "loseRound":player.Loseround, "drawRound":player.Drawround, "escapeRound":player.Escaperound}
	}
	pushMessageToEveryOne(account, sitDownMsg)
}

func PushReadyChangeMessage(account string, deskID, seatID int, isReady bool) {
	readyChangeMessage := map[string]interface{}{"messageType":1, "deskID": deskID, "seatID":seatID, "isReady":isReady}
	pushMessageToEveryOne(account, readyChangeMessage)
}

func PushPutChessMassage(account string, row, column int, result int8) {
	putChessMessage := map[string]interface{}{"messageType":2, "row":row, "column":column, "result":result}
	pushMessageToOne(account, putChessMessage)
}

func PushTakeBackReq(account string) {
	takeBackReqMessage := map[string]interface{}{"messageType":3}
	pushMessageToOne(account, takeBackReqMessage)
}

func PushTakeBackMessage(whoReq string, whoResp string, lastSteps []vo.Pos, resp bool) {
	takeBackMessage := map[string]interface{} {"messageType":4, "resp":resp, "whoReq":whoReq, "whoResp":whoResp, "lastSteps": lastSteps}
	pushMessageToOne(whoReq, takeBackMessage)
	if resp {
		pushMessageToOne(whoResp, takeBackMessage)
	}
}

func PushReqLoseMessage(account string) {
	reqLoseMessage := map[string]interface{}{"messageType":5}
	pushMessageToOne(account, reqLoseMessage)
}

func PushDrawReqMessage(account string) {
	drawReqMessage := map[string]interface{}{"messageType":6}
	pushMessageToOne(account, drawReqMessage)
}

func PushDrawRespondMessage(account string, resp bool) {
	drawRespondMessage := map[string]interface{}{"messageType":7, "resp":resp}
	pushMessageToOne(account, drawRespondMessage)
}

func PushChatTextMessage(account string, text string) {
	chatTextMessage := map[string]interface{}{"messageType":8, "text":text}
	pushMessageToOne(account, chatTextMessage)
}