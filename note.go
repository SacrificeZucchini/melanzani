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

const (
	CHORD_C1 = iota
	CHORD_C2 = iota
        CHORD_D1 = iota
        CHORD_D2 = iota
        CHORD_E1 = iota
        CHORD_E2 = iota
        CHORD_F1 = iota
        CHORD_F2 = iota
        CHORD_G1 = iota
        CHORD_G2 = iota
        CHORD_A1 = iota
        CHORD_A2 = iota
        CHORD_B1 = iota
        CHORD_B2 = iota
	CHORD_NONE = iota
)

var CHORD_NAME []string = []string {
	"CHORD_C1",
	"CHORD_C2",
	"CHORD_D1",
	"CHORD_D2",
	"CHORD_E1",
	"CHORD_E2",
	"CHORD_F1",
	"CHORD_F2",
	"CHORD_G1",
	"CHORD_G2",
	"CHORD_A1",
	"CHORD_A2",
	"CHORD_B1",
	"CHORD_B2",
}

const (
	CT_BASE = iota
	CT_MAJOR = iota
	CT_MINOR = iota
	CT_MAJOR7 = iota
	CT_MINOR7 = iota
	CT_MINOR7FF = iota
	CT_DOMINANT7 = iota
	CT_DIMINISHED = iota
)

func ChordToTone(chord int) int {
	switch(chord) {
	case CHORD_C1: return TONE_C
	case CHORD_C2: return TONE_CS
	case CHORD_D1: return TONE_D
	case CHORD_D2: return TONE_DS
	case CHORD_E1: return TONE_E
	case CHORD_E2: return TONE_E
	case CHORD_F1: return TONE_F
	case CHORD_F2: return TONE_FS
	case CHORD_G1: return TONE_G
	case CHORD_G2: return TONE_GS
	case CHORD_A1: return TONE_A
	case CHORD_A2: return TONE_AS
	case CHORD_B1: return TONE_B
	case CHORD_B2: return TONE_B
	}
	return TONE_C
}

type Note struct {
	tone int
	octave int
}

func (note Note) ToMidiCode() int {
        return (note.octave + 1) * 12 + note.tone
}


type MidiSequence struct {
	notes []Note
}

type Chord struct {
        tone int
        octave int
        chord int
}
