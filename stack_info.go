package main

import "sync"

type ServerStack struct {
	sync.RWMutex
	StackList map[string]*StackInfo
}

type StackInfo struct {
	StackLines []string
	Times      int32
	Level      string
	PanicTime  int64
	UpdateTime int64
	SendFlag   bool
}

var serverStack = &ServerStack{
	StackList: make(map[string]*StackInfo),
}
