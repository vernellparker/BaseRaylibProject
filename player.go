package main

type Player struct {
	
}

func (p Player) Setup() {
	panic("implement me")
}

func (p Player) Preload() {
	panic("implement me")
}

func (p Player) Update(dt float32) {
	panic("implement me")
}

func (p Player) TearDown() {
	panic("implement me")
}

func NewPlayer() GameFileBase{
	return &Player{}
}