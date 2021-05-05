package designPattern

import (
	"errors"
	"fmt"
)

// 定义策略接口
type TravelStrategy interface {
	RunTo(dist string) string
}

/**
 * 具体的策略Train，实现 RunTo 方法
 */
type Train struct{}

func (t *Train) RunTo(dist string) string {
	str := "run to " + dist + " by train"
	fmt.Println(str)
	return str
}

type Car struct{}

func (t *Car) RunTo(dist string) string {
	str := "run to " + dist + " by small car"
	fmt.Println(str)
	return str
}

var strategyMap = map[string]TravelStrategy{
	"train": &Train{},
	"car":   &Car{},
}

/**
 * NewTravel 生成具体策略类，其实这里的实现就是工厂模式
 */
func NewTravel(tool string) (TravelStrategy, error) {
	t, exist := strategyMap[tool]
	if !exist {
		return nil, errors.New("unknown travel tool")
	}

	return t, nil
}
