package main

import rl "github.com/gen2brain/raylib-go/raylib"

// creating a struct to hold the audio
type Audio struct {
	BackgroundMusic rl.Music
	ShootingSound rl.Sound
	// MonkeyDefeat rl.Sound
	PlayerPoint rl.Sound
}

// create a global instance to hold all sounds
var gameAudio Audio
// load in the audio
func LoadAudio () {
	gameAudio.BackgroundMusic = rl.LoadMusicStream("assets/retrobackgroundmusic.wav")
	gameAudio.ShootingSound = rl.LoadSound("assets/pewpew.wav")
	// going to add a monkey defeat sound later
	// gameAudio.MonkeyDefeat = rl.LoadSound("assets/")
	gameAudio.PlayerPoint = rl.LoadSound("assets/playercollect.wav")
}
