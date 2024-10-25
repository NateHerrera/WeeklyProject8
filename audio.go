package main

import rl "github.com/gen2brain/raylib-go/raylib"

// creating a struct to hold the audio
type Audio struct {
	BackgroundMusic rl.Music
	ShootingSound   rl.Sound
	MonkeyDefeat    rl.Sound
	BananaTree      rl.Sound
	PlayerPoint     rl.Sound
}

// create a global instance to hold all sounds
var gameAudio Audio

// load in the audio
func LoadAudio() {
	gameAudio.BackgroundMusic = rl.LoadMusicStream("assets/retrobackgroundmusic.wav")
	gameAudio.ShootingSound = rl.LoadSound("assets/pewpew.wav")
	gameAudio.BananaTree = rl.LoadSound("assets/banana_eat.flac")
	// going to add a monkey defeat sound later
	gameAudio.MonkeyDefeat = rl.LoadSound("assets/monkey_noise.wav")
	gameAudio.PlayerPoint = rl.LoadSound("assets/playercollect.wav")

	// start playing the music in a loop
	rl.PlayMusicStream(gameAudio.BackgroundMusic)

	// set the music volume
	rl.SetMusicVolume(gameAudio.BackgroundMusic, 0.5)
}
