package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const threshold = 0.5
const justThreshold = 0.2

var inputStateVar = &inputState{}

type inputState struct {
	leftLatch  bool
	rightLatch bool
	upLatch    bool
	downLatch  bool
	left       bool
	right      bool
	up         bool
	down       bool
}

func Update() {
	if !(ebiten.GamepadAxisValue(0, 0) < -justThreshold) {
		inputStateVar.leftLatch = true
		inputStateVar.left = false
	} else {
		if inputStateVar.leftLatch {
			inputStateVar.left = true
		} else {
			inputStateVar.left = false
		}
		inputStateVar.leftLatch = false
	}
	if !(ebiten.GamepadAxisValue(0, 0) > justThreshold) {
		inputStateVar.rightLatch = true
		inputStateVar.right = false
	} else {
		if inputStateVar.rightLatch {
			inputStateVar.right = true
		} else {
			inputStateVar.right = false
		}
		inputStateVar.rightLatch = false
	}
	if !(ebiten.GamepadAxisValue(0, 1) < -justThreshold) {
		inputStateVar.upLatch = true
		inputStateVar.up = false
	} else {
		if inputStateVar.upLatch {
			inputStateVar.up = true
		} else {
			inputStateVar.up = false
		}
		inputStateVar.upLatch = false
	}
	if !(ebiten.GamepadAxisValue(0, 1) > justThreshold) {
		inputStateVar.downLatch = true
		inputStateVar.down = false
	} else {
		if inputStateVar.downLatch {
			inputStateVar.down = true
		} else {
			inputStateVar.down = false
		}
		inputStateVar.downLatch = false
	}
}

func IsLeftPressed() bool {
	return ebiten.GamepadAxisValue(0, 0) < -threshold || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
}

func IsRightPressed() bool {
	return ebiten.GamepadAxisValue(0, 0) > threshold || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)
}

func IsUpPressed() bool {
	return ebiten.GamepadAxisValue(0, 1) < -threshold || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)
}

func IsDownPressed() bool {
	return ebiten.GamepadAxisValue(0, 1) > threshold || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)
}

func IsLeftJustPressed() bool {
	return inputStateVar.left || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA)
}

func IsRightJustPressed() bool {
	return inputStateVar.right || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD)
}

func IsUpJustPressed() bool {
	return inputStateVar.up || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW)
}

func IsDownJustPressed() bool {
	return inputStateVar.down || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS)
}

func IsEnterPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton0) || inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

func IsCancelPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton1) || inpututil.IsKeyJustPressed(ebiten.KeyBackspace)
}

func IsMenuPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton3) || inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func IsNextPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyV)
}
