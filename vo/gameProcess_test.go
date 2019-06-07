package vo

import (
	"fmt"
	"testing"
)

func TestGameProcess_ToJson(t *testing.T) {
	gp := NewGameProcess("a", "b")
	gp.AddStep(1, 2)
	bs, err := gp.ToJson()
	if err == nil {
		fmt.Println(string(bs))
	}
}

func TestGameProcess_GetLastThreeRmTwo(t *testing.T) {
	gp := NewGameProcess("a","b")
	gp.AddStep(1, 2)
	gp.AddStep(123, 1)
	gp.AddStep(11, 23)
	gp.AddStep(12, 2)
	gp.AddStep(0, 2)
	p1, p2, p3 := gp.GetLastThreeRmTwo()
	fmt.Println(p1, p2, p3)
	fmt.Println("count:", gp.StepCount())
	p1, p2 = gp.GetLastTwoRmOne()
	fmt.Println(p1, p2)
	fmt.Println("count:", gp.StepCount())
}