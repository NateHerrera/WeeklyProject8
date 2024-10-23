package main

import rl "github.com/gen2brain/raylib-go/raylib"

// backgroundAnimation struct to manage frame-based animation
type BackgroundAnimation struct {
	Frames       []rl.Texture2D
	CurrentFrame int
	FrameCounter int
	FrameSpeed   int
}

func main() {
	rl.InitWindow(800, 450, "Space Bandit Defender")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// initialize audio device
	rl.InitAudioDevice()
	LoadAudio()

	// create two player instances, one on the left side (normal) and one on the right side (mirrored)
	// the size is 48x48 so offset by 50
	// check if mirrored with true or false
	player1 := NewPlayer(rl.NewVector2(50, float32(rl.GetScreenHeight())/2), false)
	player2 := NewPlayer(rl.NewVector2(float32(rl.GetScreenWidth())-50, float32(rl.GetScreenHeight())/2), true)

	// load frames from a gif that was split using ezgif
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

	// initialize background animation
	// increase the frames with frame speed
	backgroundAnimation := BackgroundAnimation{
		Frames:       frames,
		CurrentFrame: 0,
		FrameCounter: 0,
		FrameSpeed:   6,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// update the frame counter and switch frames
		backgroundAnimation.FrameCounter++
		if backgroundAnimation.FrameCounter >= backgroundAnimation.FrameSpeed {
			backgroundAnimation.FrameCounter = 0
			backgroundAnimation.CurrentFrame++
			if backgroundAnimation.CurrentFrame >= len(backgroundAnimation.Frames) {
				backgroundAnimation.CurrentFrame = 0 // loop back to the first frame
			}
		}

		// get current texture (frame) for the animation
		currentTexture := backgroundAnimation.Frames[backgroundAnimation.CurrentFrame]

		// calculate the destination rectangle to scale the frame to the window size
		windowWidth := float32(rl.GetScreenWidth())
		windowHeight := float32(rl.GetScreenHeight())
		destRect := rl.NewRectangle(0, 0, windowWidth, windowHeight)

		// source rectangle is the full texture
		sourceRect := rl.NewRectangle(0, 0, float32(currentTexture.Width), float32(currentTexture.Height))

		// no rotation, origin at (0,0)
		rl.DrawTexturePro(currentTexture, sourceRect, destRect, rl.NewVector2(0, 0), 0, rl.White)

		// player 1 movement (normal) with W and S keys
		player1.Move(rl.KeyW, rl.KeyS)
		player1.Draw()

		// player 2 mirrors player 1's movement
		player2.Move(rl.KeyW, rl.KeyS)
		player2.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// drawTextureEz draws a texture centered at its origin with rotation and scaling
func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)

	rl.DrawTexturePro(texture, sourceRect, destRect, origin, angle, color)
}
