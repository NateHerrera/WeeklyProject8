package main

import rl "github.com/gen2brain/raylib-go/raylib"

// create a player struct to handle the player
type Player struct {
	Position rl.Vector2
	Speed float32

}

func (p Player) Move() {

	p.Speed = 100 * rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyW) {

		

	}

}