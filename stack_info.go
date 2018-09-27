package main

import "sync"

type ServerStack struct {
	sync.RWMutex
	StackList map[string]*StackInfo
}

type StackInfo struct {
	StackLines []string
	Env        string
	Times      int32
	Level      string
	PanicTime  int64
	UpdateTime int64
	SendFlag   bool
}

var serverStack = &ServerStack{
	StackList: make(map[string]*StackInfo),
}

func (s *StackInfo) GetStackLines() (info string) {
	for _, line := range s.StackLines {
		info += line + "\n"
	}
	return info
}
