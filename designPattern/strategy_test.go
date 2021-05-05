package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
)

func TestNewTravel(t *testing.T) {
	travel, err := designPattern.NewTravel("car")
	if err != nil {
		t.Log(err)
		return
	}
	travel.RunTo("Yang Zhou")

	travel, err = designPattern.NewTravel("bike")
	if err != nil {
		t.Log(err)
		return
	}
}
