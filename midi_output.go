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
	"github.com/rakyll/portmidi"
	"fmt"
)

type MidiOutput struct {
	stream *portmidi.Stream
}

func (midi *MidiOutput) StartPlayingNote(note Note) {
	fmt.Println("MIDI: ", note.ToMidiCode())
	midi.stream.WriteShort(0x90, int64(note.ToMidiCode()), 100)
}

func (midi *MidiOutput) StopPlayingNote(note Note) {
	midi.stream.WriteShort(0x80, int64(note.ToMidiCode()), 100)
}
