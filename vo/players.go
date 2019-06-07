package vo

import "fmt"

type empty struct {}
var EMPTY empty = empty{}
var loginMap map[string]empty

func init() {
	loginMap = make(map[string]empty)
}

func Logout(account string) {
	fmt.Println(account, " logout")
	delete(loginMap, account)
}

func Logined(account string) bool {
	_, ok := loginMap[account]
	return ok
}

func Login(account string) {
	loginMap[account] = EMPTY
}

func DisplayPlayers() {
	for k, _ := range loginMap {
		fmt.Print(k, " ")
	}
	fmt.Println()
}
