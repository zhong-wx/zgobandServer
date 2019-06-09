package vo

import (
	"./utils"
	"encoding/json"
	"fmt"
	"strings"
)

type GameProcess struct {
	account1 string	// who is black
	account2 string	// who is white
	delCount int	// must be deleted twice
	process *utils.Stack
}

func NewGameProcess(account1, account2 string) *GameProcess {
	return &GameProcess{account1:account1, account2:account2, delCount:0, process:utils.NewStack()}
}

type Pos struct {
	Row int8
	Column int8
}

func (gp *GameProcess) GetAccount1() string {
	return gp.account1
}
func (gp *GameProcess) GetAccount2() string {
	return gp.account2
}

func (gp *GameProcess) AddStep(row, column int8) {
	gp.process.Push(Pos{row, column})
}

func (gp *GameProcess) StepCount() int {
	return gp.process.Count()
}

func (gp *GameProcess) GetLastThreeRmTwo() (Pos, Pos, Pos) {
	p1 := gp.process.Pop()
	p2 := gp.process.Pop()
	p3 := gp.process.Back()
	if p3 == nil {
		return Pos{Row:-1, Column:-1}, p2.(Pos), p1.(Pos)
	}
	return p3.(Pos), p2.(Pos), p1.(Pos)
}

func (gp *GameProcess) GetLastTwoRmOne() (Pos, Pos) {
	p1 := gp.process.Pop()
	p2 := gp.process.Back()
	if p2 == nil {
		return Pos{Row:-1, Column:-1}, p1.(Pos)
	}
	return p2.(Pos), p1.(Pos)
}

func (gp *GameProcess) ToJson() ([]byte, error) {
	var jsonObj map[string]interface{}
	jsonObj = make(map[string]interface{})
	jsonObj["accuont1"] = gp.account1
	jsonObj["accuont2"] = gp.account2
	jsonObj["process"] = gp.process.ToSlice()
	bs, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return bs, nil
}

var processMap map[string]*GameProcess

func init() {
	processMap = make(map[string]*GameProcess)
}

func GetGameProcess(account1, account2 string) *GameProcess {
	gp, ok := processMap[account1+" "+account2]
	if !ok {
		return nil
	}
	return gp
}

func GetGameProcessByAccount(account string) *GameProcess {
	for k, gp := range processMap {
		if strings.Contains(k, account) {
			return gp
		}
	}
	return nil
}

func AddGameProcess(account1, account2 string, gp *GameProcess) {
	processMap[account1 + " " + account2] = gp
}

func DeleteGameProcess(account1, account2 string) {
	gp, ok := processMap[account1+" "+account2];
	if !ok {
		return
	}
	gp.delCount++
	if gp.delCount >= 2 {
		delete(processMap, account1+" "+account2)
	}
}

func DeleteGameProcessByAccount(account string) {
	for k, gp := range processMap {
		if strings.Contains(k, account) {
			gp.delCount++
			if gp.delCount >= 2 {
				delete(processMap, k)
				return
			}
		}
	}
}

func DisplayGameProcess() {
	for k,_ := range processMap {
		fmt.Print(k, " ")
	}
}
