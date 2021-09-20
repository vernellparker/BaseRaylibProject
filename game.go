package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	windowWidth int32 = 800
	windowHeight int32 = 450
	fps int32 = 60
)

type Game struct {

}

func NewGame() GameFileBase{
	return &Game{}
}

func (g *Game) Setup()  {
	rl.InitWindow(windowWidth,windowHeight, "Dapper Dasher")
	rl.SetTargetFPS(fps)
}

func (g *Game) Preload(){

}

func (g *Game) Update(dt float32)  {

}

func (g *Game) TearDown()  {

}