/*
melanzani - converts USB input of guitar devices to MIDI signals
Copyright (C) 2015  Christoph Kober

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"syscall"
)

type ButtonState struct {
	timestamp syscall.Timeval
	pressed bool
}

type SimpleGuitar struct {
	green ButtonState
	red ButtonState
	yellow ButtonState
	blue ButtonState
	orange ButtonState
	start ButtonState
	sel ButtonState
	main ButtonState
	mode GuitarMode
	strum uint
	chordMap *ChordMap
	midi *MidiOutput
}

func (guitar *SimpleGuitar) Reset() {
	var now syscall.Timeval
	syscall.Gettimeofday(&now)
	defaultButtonState := ButtonState {now, false}

	guitar.green = defaultButtonState
	guitar.red = defaultButtonState
	guitar.yellow = defaultButtonState
	guitar.blue = defaultButtonState
	guitar.orange = defaultButtonState
	guitar.start = defaultButtonState
	guitar.sel = defaultButtonState
	guitar.main = defaultButtonState
	guitar.strum = STRUM_NONE

	guitar.chordMap = &ChordMap{}
	guitar.chordMap.Init()

	guitar.mode = &SimpleGuitarMode{}
	guitar.mode.EnterMode(guitar)
}

func now() syscall.Timeval {
	var now syscall.Timeval
	syscall.Gettimeofday(&now)
	return now
}

func (guitar *SimpleGuitar) PressButton(buttonType int) {
	guitar.changeButtonState(buttonType, true)
}

func (guitar *SimpleGuitar) ReleaseButton(buttonType int) {
	guitar.changeButtonState(buttonType, false)
}

/*func (guitar *SimpleGuitar) checkMute() bool {
	return guitar.echo && (
		guitar.match(true, true, true, false, false) ||
		guitar.match(false, true, true, true, false) ||
		guitar.match(false, false, true, true, true))
}*/

func (guitar *SimpleGuitar) changeButtonState(buttonType int, state bool) {
	buttonState := ButtonState{now(), state}
	switch buttonType {
	case BUTTON_GREEN: guitar.green = buttonState
	case BUTTON_RED: guitar.red = buttonState
	case BUTTON_YELLOW: guitar.yellow = buttonState
	case BUTTON_BLUE: guitar.blue = buttonState
	case BUTTON_ORANGE: guitar.orange = buttonState
	case BUTTON_START: guitar.start = buttonState
	case BUTTON_SELECT: guitar.sel = buttonState
	case BUTTON_MAIN: guitar.main = buttonState
	}
	if state {
		guitar.mode.ButtonPressed(buttonType, buttonState)
	} else {
		guitar.mode.ButtonReleased(buttonType, buttonState)
	}
}

func (guitar *SimpleGuitar) StrumDown() {
	guitar.strum = STRUM_DOWN
	guitar.mode.StrumDown()
}

func (guitar *SimpleGuitar) StrumUp() {
	guitar.strum = STRUM_UP
	guitar.mode.StrumUp()
}

func (guitar *SimpleGuitar) ReleaseStrum() {
	guitar.strum = STRUM_NONE
	guitar.mode.StrumReleased()
}

func (guitar *SimpleGuitar) Up() {
	guitar.mode.Up()
}

func (guitar *SimpleGuitar) Down() {
	guitar.mode.Down()
}

func (guitar *SimpleGuitar) ModeSwitch(mode int) {
	if guitar.mode != nil {
		guitar.mode.ExitMode()
	}

	if mode == 0 {
		guitar.mode = &SimpleGuitarMode{}
	} else if mode == 1 {
		guitar.mode = &SoloGuitarMode{}
	} else if mode == 2 {
		guitar.mode = &ConfigurationGuitarMode{}
	}
	guitar.mode.EnterMode(guitar)
}

func (guitar *SimpleGuitar) ProduceChord() int {
	switch {
	case guitar.match(true, false, false, false, false): return CHORD_C1
        case guitar.match(true, true, false, false, false): return CHORD_C2
        case guitar.match(false, true, false, false, false): return CHORD_D1
        case guitar.match(false, true, true, false, false): return CHORD_D2
        case guitar.match(false, false, false, false, false): return CHORD_E1
	case guitar.match(true, false, true, true, false): return CHORD_E2
        case guitar.match(false, false, true, false, false): return CHORD_F1
        case guitar.match(false, false, true, true, false): return CHORD_F2
        case guitar.match(false, false, false, true, false): return CHORD_G1
        case guitar.match(false, false, false, true, true): return CHORD_G2
        case guitar.match(false, false, false, false, true): return CHORD_A1
        case guitar.match(false, false, true, false, true): return CHORD_A2
        case guitar.match(true, false, true, false, false): return CHORD_B1
	case guitar.match(false, true, false, true, false): return CHORD_B2
	case guitar.match(false, false, true, true, true): return CHORD_B1
	}
	return CHORD_C1
}

func (guitar *SimpleGuitar) match(
		green bool, red bool, yellow bool, blue bool, orange bool) bool {

	return guitar.green.pressed == green &&
		guitar.red.pressed == red &&
		guitar.yellow.pressed == yellow &&
		guitar.blue.pressed == blue &&
		guitar.orange.pressed == orange
}
