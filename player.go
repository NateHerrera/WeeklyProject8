package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// create a player struct to handle the player
type Player struct {
	Position        rl.Vector2
	Speed           float32
	Sprite          rl.Texture2D
	IsMirrored      bool
	ShotProjectiles []Projectile
}

// create a new player with an initial position, sprite, and return it
func NewPlayer(pos rl.Vector2, isMirrored bool) Player {
	// load in the sprite using rl.LoadTexture from our asset folder
	sprite := rl.LoadTexture("assets/ship.png")

	// return the player with the starting position, sprite, and mirrored flag
	return Player{Position: pos, Speed: 400, Sprite: sprite, IsMirrored: isMirrored, ShotProjectiles: make([]Projectile, 0, 100)}
}

func (p *Player) ShootProjectile() {
	offset := p.Sprite.Width / 2
	newSpeed := 200
	if p.IsMirrored {
		offset *= -1
		newSpeed *= -1
	}
	newProjectile := CreateProjectile(10, rl.NewVector2(p.Position.X+float32(offset), p.Position.Y), rl.NewVector2(float32(newSpeed), 0))
	p.ShotProjectiles = append(p.ShotProjectiles, newProjectile)
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

func (p *Player) DrawShotProjectiles() {
	for _, v := range p.ShotProjectiles {
		v.DrawProjectile()
	}
}

// move the player, accounting for mirroring if needed
func (p *Player) Move(upKey, downKey int32) {
	// set the speed
	adjustedSpeed := p.Speed * rl.GetFrameTime()

	// if the player is not mirrored (Player 1), move normally
	if !p.IsMirrored {
		// handle normal movement
		if rl.IsKeyDown(upKey) {
			p.Position.Y -= adjustedSpeed
		}
		if rl.IsKeyDown(downKey) {
			p.Position.Y += adjustedSpeed
		}
	} else {
		// handle mirrored movement
		if rl.IsKeyDown(upKey) {
			p.Position.Y += adjustedSpeed // invert movement
		}
		if rl.IsKeyDown(downKey) {
			p.Position.Y -= adjustedSpeed // invert movement
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

func (p *Player) UpdateShotProjectiles() {
	for i := 0; i < len(p.ShotProjectiles); i++ {
		p.ShotProjectiles[i].MoveProjectile()
	}
}

func (p *Player) CheckEnemiesOverlap(es *Enemies) {
	for i, v := range p.ShotProjectiles {
		for j, e := range es.allEnemies {
			if v.CheckEnemyOverlap(e) {
				p.ShotProjectiles[i].Hit = true
				es.allEnemies[j].Damage()
			}
		}
	}
}

func (p *Player) RemoveHitProjectiles() {
	for i := 0; i < len(p.ShotProjectiles); i++ {
		if p.ShotProjectiles[i].Hit {
			p.ShotProjectiles = append(p.ShotProjectiles[:i], p.ShotProjectiles[i+1:]...)
			i--
		}
	}
}
