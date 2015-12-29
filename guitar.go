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

/*
Standard Notes
  D     C
  GD    C#
  GRD   D
  RD    D#
  RYD   E
  YD    F
  YBD   F#
  BD    G
  BOD   G#
  OD    A
  GYD   A#
  RBD   B
*/
type SimpleGuitar struct {
	green ButtonState
	red ButtonState
	yellow ButtonState
	blue ButtonState
	orange ButtonState
	note Note
	strum uint
	octave int
	echo bool
	variety int /* +1 = normal | -1 = bass | 0 = equi */
	midi *MidiOutput
}

func (guitar *SimpleGuitar) Reset() {
	var now syscall.Timeval
	syscall.Gettimeofday(&now)
	defaultButtonState := ButtonState {now, false}

	guitar.octave = 2
	guitar.variety = 1
	guitar.green = defaultButtonState
	guitar.red = defaultButtonState
	guitar.yellow = defaultButtonState
	guitar.blue = defaultButtonState
	guitar.orange = defaultButtonState
}

func now() syscall.Timeval {
	var now syscall.Timeval
	syscall.Gettimeofday(&now)
	return now
}

func (guitar *SimpleGuitar) PressButton(buttonType int) {
	guitar.changeButtonState(buttonType, true)
	guitar.checkMute();
}

func (guitar *SimpleGuitar) ReleaseButton(buttonType int) {
	guitar.changeButtonState(buttonType, false)
}

func (guitar *SimpleGuitar) checkMute() {
	if guitar.echo && (
		guitar.match(true, true, true, false, false) ||
		guitar.match(false, true, true, true, false) ||
		guitar.match(false, false, true, true, true)) {
		
		guitar.stopPlayingNote()
	}
}

func (guitar *SimpleGuitar) changeButtonState(buttonType int, state bool) {
	switch buttonType {
	case BUTTON_GREEN: guitar.green = ButtonState{now(), state}
	case BUTTON_RED: guitar.red = ButtonState{now(), state}
	case BUTTON_YELLOW: guitar.yellow = ButtonState{now(), state}
	case BUTTON_BLUE: guitar.blue = ButtonState{now(), state}
	case BUTTON_ORANGE: guitar.orange = ButtonState{now(), state}
	}
}

func (guitar *SimpleGuitar) StrumDown() {
	guitar.stopPlayingNote()
	guitar.strum = STRUM_DOWN
	guitar.note = guitar.currentNote()
	guitar.midi.StartPlayingNote(guitar.note)
}

func (guitar *SimpleGuitar) StrumUp() {
	guitar.stopPlayingNote()
	guitar.strum = STRUM_UP
	guitar.note = guitar.currentNote()
	guitar.note.octave += guitar.variety
	guitar.midi.StartPlayingNote(guitar.note)
}

func (guitar *SimpleGuitar) ReleaseStrum() {
	if !guitar.echo {
		guitar.stopPlayingNote()
	}
	guitar.strum = STRUM_NONE
}

func (guitar *SimpleGuitar) stopPlayingNote() {
	guitar.midi.StopPlayingNote(guitar.note)
}

func (guitar *SimpleGuitar) currentNote() Note {
	return Note{guitar.currentTone(), guitar.octave}
}

func (guitar *SimpleGuitar) currentTone() int {
	switch {
	case guitar.match(true, false, false, false, false): return TONE_C
	case guitar.match(true, true, false, false, false): return TONE_CS
	case guitar.match(false, true, false, false, false): return TONE_D
	case guitar.match(false, true, true, false, false): return TONE_DS
	case guitar.match(false, false, false, false, false): return TONE_E
	case guitar.match(false, false, true, false, false): return TONE_F
	case guitar.match(false, false, true, true, false): return TONE_FS
	case guitar.match(false, false, false, true, false): return TONE_G
	case guitar.match(false, false, false, true, true): return TONE_GS
	case guitar.match(false, false, false, false, true): return TONE_A
	case guitar.match(false, false, true, false, true): return TONE_AS
	case guitar.match(true, false, true, false, false): return TONE_B
	}
	return TONE_C
}

func (guitar *SimpleGuitar) match(
		green bool, red bool, yellow bool, blue bool, orange bool) bool {

	return guitar.green.pressed == green &&
		guitar.red.pressed == red &&
		guitar.yellow.pressed == yellow &&
		guitar.blue.pressed == blue &&
		guitar.orange.pressed == orange
}

func (guitar *SimpleGuitar) Start() {
	if guitar.match(true, false, false, false, false) {
		guitar.variety = 1
	} else if guitar.match(false, true, false, false, false) {
		guitar.variety = -1
	} else if guitar.match(false, false, true, false, false) {
		guitar.variety = 0
	} else if guitar.match(false, false, false, false, false) {
		if guitar.echo {
			guitar.stopPlayingNote()
		}
		guitar.echo = !guitar.echo
	}
}

func (guitar *SimpleGuitar) Select() {

}

func (guitar *SimpleGuitar) Main() {

}

func (guitar *SimpleGuitar) OctaveUp() {
	guitar.octave = guitar.octave + 1
}

func (guitar *SimpleGuitar) OctaveDown() {
	guitar.octave = guitar.octave - 1
}

