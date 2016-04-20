package main

type SimpleGuitarMode struct {
	guitar *SimpleGuitar
	midi *MidiOutput
	octave int
	variety int /* +1 = normal | -1 = bass | 0 = equi */
	echo bool
	chordMap *ChordMap
	currentSequence *MidiSequence
}

func (mode *SimpleGuitarMode) EnterMode(guitar *SimpleGuitar) {
	mode.guitar = guitar
	mode.midi = guitar.midi
	mode.variety = 1
	mode.octave = 2
	mode.echo = false
	mode.chordMap = guitar.chordMap
}

func (mode *SimpleGuitarMode) ButtonPressed(buttonType int, buttonState ButtonState) {
	if buttonType == BUTTON_START {
		mode.guitar.ModeSwitch(2)
	} else if buttonType == BUTTON_SELECT {
		mode.echo = !mode.echo
	} else if mode.echo && mode.guitar.match(false, true, true, true, false) {
		mode.stopPlaying()
	}
}

func (mode *SimpleGuitarMode) ButtonReleased(buttonType int, buttonState ButtonState) {
	// nothing
}

func (mode *SimpleGuitarMode) Up() {
	mode.octave += 1
}

func (mode *SimpleGuitarMode) Down() {
	mode.octave -= 1
}

func (mode *SimpleGuitarMode) StrumUp() {
	mode.startPlaying(1)
}

func (mode *SimpleGuitarMode) StrumDown() {
	mode.startPlaying(0)
}

func (mode *SimpleGuitarMode) StrumReleased() {
	if !mode.echo {
		mode.stopPlaying()
	}
}

func (mode *SimpleGuitarMode) startPlaying(octaveMod int) {
	mode.stopPlaying()
	mode.currentSequence = mode.getCurrentSequence(octaveMod)
	mode.midi.StartPlayingSequence(mode.currentSequence)
}

func (mode *SimpleGuitarMode) stopPlaying() {
	if mode.currentSequence != nil {
		mode.midi.StopPlayingSequence(mode.currentSequence)
	}
}

func (mode *SimpleGuitarMode) getCurrentSequence(octaveMod int) *MidiSequence {
	chord := mode.getCurrentChord(octaveMod)
	return mode.chordMap.Get(chord)
}

func (mode *SimpleGuitarMode) getCurrentChord(octaveMod int) Chord {
	chord := mode.guitar.ProduceChord()
	octave := mode.octave + (mode.variety * octaveMod)
	return Chord {tone: 0, chord: chord, octave: octave}
}

func (mode *SimpleGuitarMode) ExitMode() {

}
