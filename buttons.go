package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// defining a button struct based on our lecture, to handle button placement, color, text, and click functions
type Button struct {
	X            int32      // x position of button on screen
	Y            int32      // y position of button on screen
	Width        int32      // button width
	Height       int32      // button height
	Color        rl.Color   // color of the button
	Text         string     // label on the button
	TextSize     int32      // font size of button text
	OnClickFuncs []func()   // list of functions to call when button is clicked
}

// from the lecture, creates a new button with given dimensions and color
func NewButton(x, y, width, height int32, color rl.Color) *Button {
	return &Button{
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Color:    color,
		TextSize: 20, // default text size set
	}
}

// set button text and size for displaying on the button
func (b *Button) SetText(text string, textSize int32) {
	b.Text = text
	b.TextSize = textSize
}

// centers the button horizontally on the screen
func (b *Button) CenterButtonX() {
	b.X = int32(rl.GetScreenWidth()/2) - b.Width/2
}

// adds a click function (from the lecture) that runs when button is clicked
func (b *Button) AddOnClickFunc(onClickFunc func()) {
	b.OnClickFuncs = append(b.OnClickFuncs, onClickFunc)
}

// updates button status - checks for clicks and calls any functions added to OnClickFuncs, then draws the button
func (b *Button) Update() {
	mousePos := rl.GetMousePosition() // get the mouse position
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&
		mousePos.X > float32(b.X) && mousePos.X < float32(b.X+b.Width) &&
		mousePos.Y > float32(b.Y) && mousePos.Y < float32(b.Y+b.Height) {

		// debug message to show when button is clicked
		fmt.Println("Button clicked")

		// calls all functions added with AddOnClickFunc
		for _, onClick := range b.OnClickFuncs {
			onClick()
		}
	}
	
	// draw the button itself
	rl.DrawRectangleLines(b.X, b.Y, b.Width, b.Height, b.Color)
	textWidth := rl.MeasureText(b.Text, b.TextSize) // measures text width to center it
	rl.DrawText(b.Text, b.X+(b.Width/2)-(textWidth/2), b.Y+(b.Height/2)-(b.TextSize/2), b.TextSize, rl.White) // draws the button text in white
}
