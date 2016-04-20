package main

type SoloGuitarMode struct {
	guitar *SimpleGuitar
	midi *MidiOutput
	octave int
	variety int /* +1 = normal | -1 = bass | 0 = equi */
	chordMap *ChordMap
	currentSequence *MidiSequence	
}

func (mode *SoloGuitarMode) EnterMode(guitar *SimpleGuitar) {
	mode.guitar = guitar
	mode.midi = guitar.midi
	mode.variety = 1
	mode.octave = 2
	mode.chordMap = guitar.chordMap
	mode.chordMap.Init()
}

func (mode *SoloGuitarMode) ButtonPressed(buttonType int, buttonState ButtonState) {
	mode.stopPlaying()
	mode.startPlaying(0)
}

func (mode *SoloGuitarMode) ButtonReleased(buttonType int, buttonState ButtonState) {

}

func (mode *SoloGuitarMode) Up() {
	mode.octave += 1
}

func (mode *SoloGuitarMode) Down() {
	mode.octave -= 1
}

func (mode *SoloGuitarMode) StrumUp() {
	mode.stopPlaying()
}

func (mode *SoloGuitarMode) StrumDown() {
	mode.stopPlaying()
}

func (mode *SoloGuitarMode) StrumReleased() {
	// noting
}

func (mode *SoloGuitarMode) startPlaying(octaveMod int) {
	mode.currentSequence = mode.getCurrentSequence(octaveMod)
	mode.midi.StartPlayingSequence(mode.currentSequence)
}

func (mode *SoloGuitarMode) stopPlaying() {
	if mode.currentSequence != nil {	
		mode.midi.StopPlayingSequence(mode.currentSequence)
	}
}

func (mode *SoloGuitarMode) getCurrentSequence(octaveMod int) *MidiSequence {
	chord := mode.getCurrentChord(octaveMod)
	return mode.chordMap.Get(chord)
}

func (mode *SoloGuitarMode) getCurrentChord(octaveMod int) Chord {
	chord := mode.guitar.ProduceChord()
	octave := mode.octave + (mode.variety * octaveMod)
	return Chord {tone: 0, chord: chord, octave: octave}
}

func (mode *SoloGuitarMode) ExitMode() {

}
