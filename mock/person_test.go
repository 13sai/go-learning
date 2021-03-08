package mock

import (
	"mock/mock"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewPerson(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	talker := mock.NewMockTalker(ctl)
	talker.EXPECT().SayHello("hi").Return("13sai said hi")
	t.Log(talker.SayHello("hi"))
}
