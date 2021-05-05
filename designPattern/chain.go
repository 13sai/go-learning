package designPattern

import "fmt"

type handle interface {
	Handle(content string)
	next(h handle, content string)
}

// 广告过滤
type AdHandle struct {
	H handle
}

func (ad *AdHandle) Handle(content string) {
	fmt.Println("ad ad ad !!!")
	ad.next(ad.H, content)
}

func (ad *AdHandle) next(h handle, content string) {
	if ad.H != nil {
		ad.H.Handle(content)
	}
}

type SensitiveHandle struct {
	H handle
}

func (s *SensitiveHandle) Handle(content string) {
	fmt.Println("sensitive !!!")
	s.next(s.H, content)
}

func (s *SensitiveHandle) next(h handle, content string) {
	if s.H != nil {
		s.H.Handle(content)
	}
}
