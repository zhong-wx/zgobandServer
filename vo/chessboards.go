package vo

import (
	"../gen-go/zgobandRPC"
	"fmt"
	"strings"
)

type chessBoardStat struct {
	empty int
	chessBoard [15][15]int8
}
func NewChessBoardStat() *chessBoardStat{
	return &chessBoardStat{empty:15*15, chessBoard:[15][15]int8{}}
}

func (p *chessBoardStat) get(row, column int8) int8 {
	return p.chessBoard[row][column]
}

func (p *chessBoardStat) set(row, column, value int8) {
	p.chessBoard[row][column] = value
}

var chessBoardMap map[string]*chessBoardStat

func init() {
	chessBoardMap = make(map[string]*chessBoardStat)
}

func DisplayChessboard() {
	for k,_ := range chessBoardMap {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

func checkWinner(key string, which, row, column int8) int8 {
	if _, ok := chessBoardMap[key]; !ok {
		return -1
	}

	count := 1
	for c := column+1; c<15; c++ {
		if chessBoardMap[key].get(row, c)==which {
			count++
		} else {
			break
		}
	}
	fmt.Println("1, count:",count)
	for c := column-1; c>=0; c-- {
		if chessBoardMap[key].get(row, c)==which {
			count++
		} else {
				break
		}
	}
	fmt.Println("2, count:",count)
	if count>=5 {
		return 1
	}

	count = 1
	for r := row+1; r<15; r++ {
		if chessBoardMap[key].get(r, column)==which {
			count++
		} else {
			break
		}
	}
	fmt.Println("3, count:",count)
	for r := row-1; r>=0; r-- {
		if chessBoardMap[key].get(r, column)==which {
			count++
		} else {
			break
		}
	}
	fmt.Println("4, count:",count)
	if count>=5 {
		return 1
	}

	count = 1
	r := row+1
	c := column+1
	for r<15&&c<15{
		if chessBoardMap[key].get(r, c)==which {
			count++
		} else {
			break
		}
		r++;c++
	}
	fmt.Println("5, count:",count)
	r = row-1
	c = column-1
	for r>=0&&c>=0{
		if chessBoardMap[key].get(r, c)==which {
			count++
		} else {
			break
		}
		r--;c--
	}
	fmt.Println("6, count:",count)
	if count>=5 {
		return 1
	}

	count = 1
	r = row+1
	c = column-1
	for r<15&&c>=0{
		if chessBoardMap[key].get(r, c)==which {
			count++
		} else {
			break
		}
		r++;c--
	}
	fmt.Println("7, count:",count)
	r = row-1
	c = column+1
	for r>=0&&c<15{
		if chessBoardMap[key].get(r, c)==which {
			count++
		} else {
			break
		}
		r--;c++
	}
	fmt.Println("8, count:",count)
	if count>=5 {
		return 1
	}

	return 0
}

func DeleteChessBoard(account1, account2 string) {
	delete(chessBoardMap, account1 + " " + account2)
}
func DeleteChessBoardByAccount(account string) {
	for k, _ := range chessBoardMap {
		if strings.Contains(k, account) {
			delete(chessBoardMap, k)
		}
	}
}

// return 1 means win, 2 mean draw, >=0 means put chess succeed
func PutChess(player1, player2 string, row, column, seatID int8) (int8, error){
	key := player1 + " " + player2
	if _, exist := chessBoardMap[key]; !exist {
		chessBoardMap[key] = NewChessBoardStat()
	}
	if chessBoardMap[key].get(row-1, column-1) != 0 {
		e := zgobandRPC.NewInvalidOperation()
		e.Type = "服务器发送错误"
		e.Why = "该位置已有棋子"
		return -1, e
	}

	if chessBoardMap[key].empty <= 0 {
		return 2, nil
	}

	if seatID == 1 {
		chessBoardMap[key].set(row-1, column-1, -1)
	} else if(seatID == 2) {
		chessBoardMap[key].set(row-1, column-1, 1)
	} else {
		return -1, fmt.Errorf("unkown seatID")
	}
	chessBoardMap[key].empty--

	var t int8
	if seatID==1 {
		t = -1
	} else if(seatID==2) {
		t = 1
	}
	result := checkWinner(key, t, row-1, column-1)
	fmt.Println("isWin:", result)
	return result, nil
}

func TakeBack(player1, player2 string, row, column int8) {
	key := player1 + " " + player2
	if _, exist := chessBoardMap[key]; !exist {
		fmt.Println("TakeBack:may be a error, no chess in this pos")
		return
	}
	chessBoardMap[key].set(row-1, column-1, 0)
}