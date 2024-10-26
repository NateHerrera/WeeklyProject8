package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Position rl.Vector2
	Speed    float32
	Sprite   rl.Texture2D
	Health   int
}

func NewEnemy(newPos rl.Vector2, newSpeed float32, newSprite rl.Texture2D) Enemy {
	return Enemy{
		Position: newPos,
		Speed:    newSpeed,
		Sprite:   newSprite,
		Health:   3,
	}
}

func SpawnEnemy(newSprite rl.Texture2D) Enemy {
	var newSpeed float32
	switch rl.GetRandomValue(0, 1) {
	case 0:
		newSpeed = -50

	case 1:
		newSpeed = 50
	}

	row := int(rl.GetRandomValue(1, 5))

	rowSpacing := rl.GetScreenHeight() / 5

	newYPos := (rowSpacing * (row - 1)) + rowSpacing/2
	return NewEnemy(rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(newYPos)), newSpeed, newSprite)
}

func (e *Enemy) Move() {
	adjustedSpeed := e.Speed * rl.GetFrameTime()
	e.Position.X += adjustedSpeed
}

func (e *Enemy) DrawEnemy() {
	DrawTextureEz(e.Sprite, e.Position, 0, 1, rl.White)
}
