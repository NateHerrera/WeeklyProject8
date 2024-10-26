package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// setting up a "game" struct to hold info about the game state and menu buttons
type Game struct {
	State       GameState
	States      GameStates
	StartButton *Button
	QuitButton  *Button
}

// making a type for different game states so we can track where the player is in the game
type GameState int

// holding different game state options like menu, playing, and game over
type GameStates struct {
	Menu     GameState
	Playing  GameState
	GameOver GameState
}

// struct to handle the background animation frames
type BackgroundAnimation struct {
	Frames       []rl.Texture2D // all the animation frames for the background
	CurrentFrame int            // keeps track of which frame is being shown
	FrameCounter int            // counts frames to control speed of the animation
	FrameSpeed   int            // sets speed of frame switching
}

type Lives struct {
	Sprite rl.Texture2D
	Number int
}

func main() {
	// start up the game window
	rl.InitWindow(1600, 900, "Space Bandit Defender")
	defer rl.CloseWindow() // close the window when the game ends

	// set the frames per second (fps) so it runs smoothly
	rl.SetTargetFPS(60)

	// set up audio stuff for the game
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	LoadAudio()              // load in-game sounds (defined elsewhere)
	game := InitializeGame() // initialize game states and menu buttons

	// create two players, one on each side
	player1 := NewPlayer(rl.NewVector2(50, float32(rl.GetScreenHeight())/2), false)
	player2 := NewPlayer(rl.NewVector2(float32(rl.GetScreenWidth())-50, float32(rl.GetScreenHeight())/2), true)
	playerLives := NewLives()
	playerEnemies := NewEnemies()
	playerEnemies.AddEnemy()

	var spawnTimer float32 = 0

	// loading animation frames from image files for the background
	frames := []rl.Texture2D{
		rl.LoadTexture("assets/frame_0_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_1_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_2_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_3_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_4_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_5_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_6_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_7_delay-0.1s.gif"),
		rl.LoadTexture("assets/frame_8_delay-0.1s.gif"),
	}

	// setting up the background animation with frames, starting at frame 0
	backgroundAnimation := BackgroundAnimation{
		Frames:       frames,
		CurrentFrame: 0,
		FrameCounter: 0,
		FrameSpeed:   6, // controls how fast frames switch
	}

	// making a semi-transparent black color for the hud background
	hudColor := rl.NewColor(0, 0, 0, 128)
	mutedHudColor := rl.NewColor(0, 0, 0, 64)

	// main game loop - this keeps running till the window closes
	for !rl.WindowShouldClose() {
		// start drawing everything to the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black) // set a black background

		// update music stream each frame so it keeps playing
		rl.UpdateMusicStream(gameAudio.BackgroundMusic)

		// update background animation frame counter and change frame if needed
		backgroundAnimation.FrameCounter++
		if backgroundAnimation.FrameCounter >= backgroundAnimation.FrameSpeed {
			backgroundAnimation.FrameCounter = 0 // reset frame counter
			backgroundAnimation.CurrentFrame++   // move to next frame
			if backgroundAnimation.CurrentFrame >= len(backgroundAnimation.Frames) {
				backgroundAnimation.CurrentFrame = 0 // loop back to first frame if needed
			}
		}

		// get the current texture frame for the background animation
		currentTexture := backgroundAnimation.Frames[backgroundAnimation.CurrentFrame]

		// set up the rectangle for background size and placement
		windowWidth := float32(rl.GetScreenWidth())
		windowHeight := float32(rl.GetScreenHeight())
		destRect := rl.NewRectangle(0, 0, windowWidth, windowHeight) // fit the whole window
		sourceRect := rl.NewRectangle(0, 0, float32(currentTexture.Width), float32(currentTexture.Height))
		rl.DrawTexturePro(currentTexture, sourceRect, destRect, rl.NewVector2(0, 0), 0, rl.White) // draw it
		if playerLives.Number < 1 {
			game.State = game.States.GameOver
		}
		// if we're on the menu screen, add a semi-transparent overlay for the hud
		if game.State == game.States.Menu {
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), hudColor)
		}

		// check game state to see what to draw
		switch game.State {
		case game.States.Menu:
			// on menu screen, draw the start and quit buttons
			game.StartButton.Update()
			game.QuitButton.Update()

		case game.States.Playing:
			// on playing screen, move players and draw them
			spawnTimer += rl.GetFrameTime()
			if spawnTimer > 5 {
				playerEnemies.AddEnemy()
				spawnTimer = 0
			}
			if rl.IsKeyPressed(rl.KeySpace) {
				player1.ShootProjectile()
				player2.ShootProjectile()
				rl.PlaySound(gameAudio.ShootingSound)
			}

			player1.Move(rl.KeyW, rl.KeyS)
			player2.Move(rl.KeyW, rl.KeyS)

			player1.UpdateShotProjectiles()
			player2.UpdateShotProjectiles()
			playerEnemies.UpdateEnemies()
			player1.CheckEnemiesOverlap(&playerEnemies)
			player2.CheckEnemiesOverlap(&playerEnemies)
			player1.RemoveHitProjectiles()
			player2.RemoveHitProjectiles()
			playerEnemies.CheckOffScreen(&playerLives)
			playerEnemies.RemoveDeadEnemies()
			playerEnemies.DrawEnemies()
			player1.Draw()
			player2.Draw()
			player1.DrawShotProjectiles()
			player2.DrawShotProjectiles()

			rl.DrawRectangle(0, 0, 250, playerLives.Sprite.Height*2, mutedHudColor)
			playerLives.DrawLives()

		case game.States.GameOver:
			// on game over screen, just show "Game Over" text
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), hudColor)
			rl.DrawText("Game Over", 350, 200, 30, rl.Red)
		}

		// finish drawing everything for this frame
		rl.EndDrawing()
	}

	// when the loop ends, close the game window
	rl.CloseWindow()
}

// helper function to draw a texture with rotation and scaling
func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	// set up the source and destination rectangles for the texture
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.NewVector2(float32(texture.Width)/2*scale, float32(texture.Height)/2*scale) // center point for rotation

	// draw the texture with rotation and scaling
	rl.DrawTexturePro(texture, sourceRect, destRect, origin, angle, color)
}

// initialize the game, set up menu buttons and game states
func InitializeGame() *Game {
	// define game states
	states := GameStates{
		Menu:     0,
		Playing:  1,
		GameOver: 2,
	}

	// create the game struct with menu state as default
	game := &Game{
		State:  states.Menu,
		States: states,
	}

	// set button color (white)
	buttonColor := rl.NewColor(255, 255, 255, 255)

	// define button size and spacing
	buttonWidth := int32(300)
	buttonHeight := int32(100)
	buttonSpacing := int32(20)

	// calculate starting y-position for the "Start Game" button
	startY := (int32(rl.GetScreenHeight()) - (buttonHeight*2 + buttonSpacing)) / 2

	// set up "Start Game" button, center horizontally
	game.StartButton = NewButton(0, startY, buttonWidth, buttonHeight, buttonColor)
	game.StartButton.SetText("Start Game", 20)
	game.StartButton.CenterButtonX() // only center horizontally
	game.StartButton.AddOnClickFunc(func() {
		// when "Start Game" is clicked, switch to playing state
		game.State = game.States.Playing
	})

	// set up "Quit Game" button below the "Start Game" button, center horizontally
	game.QuitButton = NewButton(0, startY+buttonHeight+buttonSpacing, buttonWidth, buttonHeight, buttonColor)
	game.QuitButton.SetText("Quit Game", 20)
	game.QuitButton.CenterButtonX()                // only center horizontally
	game.QuitButton.AddOnClickFunc(rl.CloseWindow) // exits game when clicked

	return game // return game struct with initialized states and buttons
}

func NewLives() Lives {
	return Lives{
		Sprite: rl.LoadTexture("assets/banana_tree.png"),
		Number: 3,
	}
}

func (l Lives) DrawLives() {
	rl.DrawText("Lives:", 10, 20, 24, rl.White)
	for i := 0; i < l.Number; i++ {
		DrawTextureEz(l.Sprite, rl.NewVector2(float32(100+(50*i)), float32(l.Sprite.Height)), 0, 1, rl.White)
	}
}
