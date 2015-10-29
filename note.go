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
	TONE_C = 0
	TONE_CS = 1
	TONE_D = 2
	TONE_DS = 3
	TONE_E = 4
	TONE_F = 5
	TONE_FS = 6
	TONE_G = 7
	TONE_GS = 8
	TONE_A = 9
	TONE_AS = 10
	TONE_B = 11
)

type Note struct {
	tone int
	octave int
}

func (note Note) ToMidiCode() int {
	return (note.octave + 1) * 12 + note.tone
}
