package main

type GuitarMode interface {
	EnterMode(guitar *SimpleGuitar)
	ButtonPressed(buttonType int, buttonState ButtonState)
	ButtonReleased(buttonType int, buttonState ButtonState)
	Up()
	Down()	
	StrumUp()
	StrumDown()
	StrumReleased()
	ExitMode()
}
