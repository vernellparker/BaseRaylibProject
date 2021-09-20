package main

type GameFileBase interface {
	Setup()
	Preload()
	Update(dt float32)
	TearDown()
}
