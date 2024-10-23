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
	DrawTextureEz(p.Sprite, p.Position, 90, 1, rl.White)
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

	 // grab the screen height to make a border
	screenHeight := float32(rl.GetScreenHeight())

	 // take the sprites height, and scale and divide it by 2 to take into account for boundaries
	spriteHalfHeight := float32(p.Sprite.Height) * 1 / 2 // 
 
	 // make the boundaries
	if p.Position.Y - spriteHalfHeight < 0 {

		 p.Position.Y = spriteHalfHeight
	}

	if p.Position.Y + spriteHalfHeight > screenHeight {

		 p.Position.Y = screenHeight - spriteHalfHeight
	} 
}