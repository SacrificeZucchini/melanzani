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

const (
	STRUM_NONE = iota
	STRUM_DOWN = iota
	STRUM_UP = iota
)

const (
	BUTTON_NONE = iota
	BUTTON_GREEN = iota
	BUTTON_RED = iota
	BUTTON_YELLOW = iota
	BUTTON_BLUE = iota
	BUTTON_ORANGE = iota
	BUTTON_START = iota
	BUTTON_SELECT = iota
	BUTTON_MAIN = iota
)

type Instrument interface {
	PressButton(buttonType int)
	ReleaseButton(buttonType int)
	Up()
	Down()	
	StrumUp()
	StrumDown()
	ReleaseStrum()
}
