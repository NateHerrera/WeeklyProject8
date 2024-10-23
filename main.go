package main

import rl "github.com/gen2brain/raylib-go/raylib"

// backgroundAnimation struct to manage frame-based animation
type BackgroundAnimation struct {
    Frames      []rl.Texture2D
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

    // create a player instance using newPlayer
    player := NewPlayer()

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
    backgroundAnimation := BackgroundAnimation{
        Frames:      frames,
        CurrentFrame: 0,
        FrameCounter: 0,
        FrameSpeed:   10,
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

        // source rectangle is the full texture (no cropping)
        sourceRect := rl.NewRectangle(0, 0, float32(currentTexture.Width), float32(currentTexture.Height))

        // no rotation, origin at (0,0)
        rl.DrawTexturePro(currentTexture, sourceRect, destRect, rl.NewVector2(0, 0), 0, rl.White)

        // draw the player on top of the background
        player.Draw()
        player.Move()

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
