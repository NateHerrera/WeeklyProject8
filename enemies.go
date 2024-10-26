package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemies struct {
	allEnemies  []Enemy
	enemySprite rl.Texture2D
}

func NewEnemies() Enemies {
	sprite := rl.LoadTexture("assets/space_monkey.png")
	return Enemies{
		allEnemies:  make([]Enemy, 0, 100),
		enemySprite: sprite,
	}
}

func (es *Enemies) AddEnemy() {
	es.allEnemies = append(es.allEnemies, SpawnEnemy(es.enemySprite))
}

func (es *Enemies) UpdateEnemies() {
	for i := 0; i < len(es.allEnemies); i++ {
		es.allEnemies[i].Move()
	}
}

func (es *Enemies) DrawEnemies() {
	for _, v := range es.allEnemies {
		v.DrawEnemy()
	}
}
