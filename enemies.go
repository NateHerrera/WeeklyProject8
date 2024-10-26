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

func (es *Enemies) CheckOffScreen(l *Lives) {
	for i, v := range es.allEnemies {
		if v.Position.X > float32(rl.GetScreenWidth()) || v.Position.X < 0 || v.Position.Y > float32(rl.GetScreenHeight()) || v.Position.Y < 0 {
			es.allEnemies[i].Health = 0
			rl.PlaySound(gameAudio.BananaTree)
			l.Number--
		}
	}
}

func (es *Enemies) RemoveDeadEnemies() {
	for i := 0; i < len(es.allEnemies); i++ {
		if es.allEnemies[i].Health <= 0 {
			es.allEnemies = append(es.allEnemies[:i], es.allEnemies[i+1:]...)
		}
	}
}
