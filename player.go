package main

import rl "github.com/gen2brain/raylib-go/raylib"

// create a player struct to handle the player
type Player struct {
	Position rl.Vector2
	Speed float32
	Sprite rl.Texture2D
}

// create a new player with an initial position, sprite, and return it
func NewPlayer() Player {

	// load in the sprite using rl.LoadTexture from our asset folder
	sprite := rl.LoadTexture("assets/ship.png")

	// get sprite dimensions to adjust initial position, because origin point is 0,0, it needs spawn in the screen not half way
	spriteWidth := float32(sprite.Width)

	posX := spriteWidth / 2
	// return it with the starting pos
	return Player{Position: rl.NewVector2(posX, 300), Sprite: sprite}
}


// draw the player using the ship sprite
// use the function from the slides
func (p *Player) Draw() {

	// need the sprite to face the right so rotate by 90 degs
	DrawTextureEz(p.Sprite, p.Position, 90, 1.5, rl.White)
}


// create the basic movement and borders
func (p *Player) Move() {

	// set a speed
	p.Speed = 100 * rl.GetFrameTime()
	// grab the width

	if rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= 5
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += 5
	}

	 // Define boundaries: the top and bottom edges of the screen
	 screenHeight := float32(rl.GetScreenHeight())

	 // The Y boundaries need to account for the sprite's half-height, since its origin is in the center
	 spriteHalfHeight := float32(p.Sprite.Height) * 1.5 / 2 // 1.5 is the scaling factor
 
	 // Clamp the player's Y position so they stay on screen
	 if p.Position.Y - spriteHalfHeight < 0 {
		 p.Position.Y = spriteHalfHeight
	 } else if p.Position.Y + spriteHalfHeight > screenHeight {
		 p.Position.Y = screenHeight - spriteHalfHeight
	 }
}