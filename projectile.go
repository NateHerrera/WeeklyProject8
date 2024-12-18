package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Radius float32
	Pos    rl.Vector2
	Vel    rl.Vector2
	Hit    bool
}

func CreateProjectile(newRadius float32, newPos rl.Vector2, newVel rl.Vector2) Projectile {
	return Projectile{
		Radius: newRadius,
		Pos:    newPos,
		Vel:    newVel,
		Hit:    false,
	}
}

func (p *Projectile) MoveProjectile() {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Vel, rl.GetFrameTime()))
}

func (p Projectile) DrawProjectile() {
	rl.DrawCircle(int32(p.Pos.X), int32(p.Pos.Y), p.Radius, rl.White)
}

func (p Projectile) CheckEnemyOverlap(e Enemy) bool {
	return rl.Vector2Distance(p.Pos, e.Position) <= p.Radius+float32(e.Sprite.Width)
}
