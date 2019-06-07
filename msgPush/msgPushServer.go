package msgPush

import (
	"fmt"
	"net"
	"../vo"
	"../models"
)

var accountConnMap map[string]net.Conn
var connAccountMap map[net.Conn]string
var msgPushServer *MsgPushServer

func init() {
	accountConnMap = make(map[string]net.Conn)
	connAccountMap = make(map[net.Conn]string)
	msgPushServer = NewMsgPushServer("localhost:9092")
	msgPushServer.Start()
}

func deleteAccount(account string) {
	conn := accountConnMap[account]
	delete(accountConnMap, account)
	delete(connAccountMap, conn)
}

func deleteConn(conn net.Conn) {
	account := connAccountMap[conn]
	delete(accountConnMap, account)
	delete(connAccountMap, conn)
}

type MsgPushServer struct{
	addr string
	writeChan chan string
}

func NewMsgPushServer(addr string) *MsgPushServer {
	writeChan := make(chan string)
	return &MsgPushServer{addr: addr, writeChan:writeChan}
}

func(p *MsgPushServer) Start() {
	l, err := net.Listen("tcp", p.addr)//
	if err != nil {
		panic("net.Listen fail")
	}

	go func() {
		for {
			conn, err:= l.Accept()
			if err != nil {
				fmt.Println("Accept error:", err.Error())
				continue
			}
			fmt.Println("accept a conn")
			go p.handleConn(conn)
		}
	}()
}

func(p *MsgPushServer) Send(strData string) {
	p.writeChan <- strData
}

func readConn(conn net.Conn, readChan chan<- string, stopChan chan<- int) {
	for{
		// #TODO how to read a string end by \0
		bytes := make([]byte, 1024)
		//conn.SetReadDeadline(time.Time{})
		n, err := conn.Read(bytes);
		bytes = bytes[0:n]
		if err != nil {
			// #TODO timeout may be received
			fmt.Println("net.Conn read error:", err.Error())
			account := connAccountMap[conn]
			vo.Logout(account)
			// #TODO 目前策略断线就删除棋盘记录
			vo.DeleteChessBoardByAccount(account)
			// #TODO 离线时自动离座 ？
			stat, err := models.QueryPlayerStat(account)
			if err == nil {
				models.DelSeat(int32(stat.DeskID), int32(stat.SeatID))
			}

			// #TODO 如果双方正在对弈 通知对方逃跑

			// 删除状态记录
			models.DeletePlayerStat(account)
			// 删除上局下棋过程
			vo.DeleteGameProcessByAccount(account)
			deleteConn(conn)

			fmt.Print("logined:")
			vo.DisplayPlayers()
			fmt.Print("accountMap:")
			for k, _ := range accountConnMap {
				fmt.Print(k, " ")
			}
			fmt.Println()
			fmt.Print("connMap:")
			for _, v := range connAccountMap {
				fmt.Print(v, " ")
			}
			fmt.Println()
			fmt.Print("chessboard:")
			vo.DisplayChessboard()
			fmt.Println()
			fmt.Print("gameProcess:")
			vo.DisplayGameProcess()
			fmt.Println()

			stopChan <- 1
			break
		}
		//recv account
		strData := string(bytes)
		readChan <- strData
	}
}

func writeConn(conn net.Conn, writeChan <-chan string, stopChan <-chan int) {
	for {
		select {
			case strData := <-writeChan:
				leftToSend := len(strData)
				for leftToSend>0 {
					n, err := conn.Write([]byte(strData))
					if err != nil {
						// #TODO timeout may be received
						fmt.Println("net.Conn write error:", err.Error())
						deleteConn(conn)
						break
					}
					leftToSend -= n
				}
				case <-stopChan:
					fmt.Println("writeConn end")
					break
		}
	}
}

// recv account #TODO 可以扩展支持更多业务
func handleRead(conn net.Conn, writeChan <-chan string	, readData string) {
	account := readData
	fmt.Println("handleRead account:", account)
	accountConnMap[account] = conn
	connAccountMap[conn] = account
}

func (*MsgPushServer) handleConn(conn net.Conn) {
	readChan := make(chan string)
	writeChan := make(chan string)
	stopChan := make(chan int)

	go readConn(conn, readChan, stopChan)
	go writeConn(conn, readChan, stopChan)

	for {
		select {
			case readStr :=<- readChan:
				handleRead(conn, writeChan, readStr)
			case <- stopChan:
				break
		}
	}
}