package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// create a player struct to handle the player
type Player struct {
	Position   rl.Vector2
	Speed      float32
	Sprite     rl.Texture2D
	IsMirrored bool
}

// create a new player with an initial position, sprite, and return it
func NewPlayer(pos rl.Vector2, isMirrored bool) Player {
	// load in the sprite using rl.LoadTexture from our asset folder
	sprite := rl.LoadTexture("assets/ship.png")

	// return the player with the starting position, sprite, and mirrored flag
	return Player{Position: pos, Sprite: sprite, IsMirrored: isMirrored}
}

// draw the player using the ship sprite
func (p *Player) Draw() {
	// rotate player 1 by 90 degrees, and player 2 by -90 degrees
	rotation := float32(90)
	if p.IsMirrored {
		rotation = -90
	}
	DrawTextureEz(p.Sprite, p.Position, rotation, 1, rl.White)
}

// move the player, accounting for mirroring if needed
func (p *Player) Move(upKey, downKey int32) {
	// set the speed
	p.Speed = 100 * rl.GetFrameTime()

	// if the player is not mirrored (Player 1), move normally
	if !p.IsMirrored {
		// handle normal movement
		if rl.IsKeyDown(upKey) {
			p.Position.Y -= 5
		}
		if rl.IsKeyDown(downKey) {
			p.Position.Y += 5
		}
	} else {
		// handle mirrored movement
		if rl.IsKeyDown(upKey) {
			p.Position.Y += 5 // invert movement
		}
		if rl.IsKeyDown(downKey) {
			p.Position.Y -= 5 // invert movement
		}
	}

	// grab the screen height to make a border
	screenHeight := float32(rl.GetScreenHeight())

	// calculate half of the sprite's height to adjust for boundaries
	spriteHalfHeight := float32(p.Sprite.Height) / 2

	// find the size of the sprite which is 48
	// fmt.Printf("Player sprite size: Width = %d, Height = %d\n", p.Sprite.Width, p.Sprite.Height)

	// enforce the boundaries
	if p.Position.Y-spriteHalfHeight < 0 {
		p.Position.Y = spriteHalfHeight
	}

	if p.Position.Y+spriteHalfHeight > screenHeight {
		p.Position.Y = screenHeight - spriteHalfHeight
	}
}
