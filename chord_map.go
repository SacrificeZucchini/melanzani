package main

type ChordMap struct {
	chords []*MidiSequence
}

func (chordMap *ChordMap) Init() {
	chordMap.chords = []*MidiSequence {
		CHORD_TABLE[TONE_C][CT_BASE],
		CHORD_TABLE[TONE_CS][CT_BASE],
		CHORD_TABLE[TONE_D][CT_BASE],
		CHORD_TABLE[TONE_DS][CT_BASE],
		CHORD_TABLE[TONE_E][CT_BASE],
		CHORD_TABLE[TONE_E][CT_BASE],
		CHORD_TABLE[TONE_F][CT_BASE],
		CHORD_TABLE[TONE_FS][CT_BASE],
		CHORD_TABLE[TONE_G][CT_BASE],
		CHORD_TABLE[TONE_GS][CT_BASE],
		&MidiSequence { []Note{ {TONE_A, 0} }},
		CHORD_TABLE[TONE_AS][CT_BASE],
		CHORD_TABLE[TONE_B][CT_BASE],
		CHORD_TABLE[TONE_B][CT_BASE],/*
		&MidiSequence { []Note{ {TONE_C, 1}, {TONE_F, 0} } },
		&MidiSequence { []Note{ {TONE_CS, 1} } },
		&MidiSequence { []Note{ {TONE_D, 1} } },
		&MidiSequence { []Note{ {TONE_DS, 1} } },
		&MidiSequence { []Note{ {TONE_E, 0}, {TONE_A, -1} } },
		&MidiSequence { []Note{ {TONE_E, 0}, {TONE_A, -1} } },
		&MidiSequence { []Note{ {TONE_F, 1}, {TONE_A, 0 } } },
		&MidiSequence { []Note{ {TONE_FS, 1} } },
		&MidiSequence { []Note{ {TONE_G, 1}, {TONE_C, 1} } },
		&MidiSequence { []Note{ {TONE_GS, 1} } },
		&MidiSequence { []Note{ {TONE_A, 1}, {TONE_D, 1} } },
		&MidiSequence { []Note{ {TONE_AS, 1} } },
		&MidiSequence { []Note{ {TONE_B, 1} } },
		&MidiSequence { []Note{ {TONE_B, 1} } },*/
	}
}

func (chordMap *ChordMap) Get(chord Chord) *MidiSequence {
	sequence := chordMap.chords[chord.chord]
	notes := make([]Note, 0, 4)
	for i := range sequence.notes {
		curr := sequence.notes[i]
		note := Note {curr.tone, curr.octave + chord.octave}
		notes = append(notes, note)
	}
	return &MidiSequence{notes}
}

type ChordTable [][]*MidiSequence

var CHORD_TABLE ChordTable = ChordTable {
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_C, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_E, 0}, {TONE_G, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_DS, 0}, {TONE_G, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_E, 0}, {TONE_G, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_DS, 0}, {TONE_G, 0}, {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_DS, 0}, {TONE_FS, 0}, {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_E, 0}, {TONE_G, 0}, {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_C, 0}, {TONE_DS, 0}, {TONE_FS, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_CS, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_F, 0}, {TONE_GS, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_E, 0}, {TONE_GS, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_F, 0}, {TONE_GS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_E, 0}, {TONE_GS, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_E, 0}, {TONE_G, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_F, 0}, {TONE_GS, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_CS, 0}, {TONE_E, 0}, {TONE_G, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_D, 0} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_FS, 0}, {TONE_A, 0} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_F, 0}, {TONE_A, 0} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_F, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_F, 0}, {TONE_GS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_D, 0}, {TONE_F, 0}, {TONE_GS, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_DS, 0} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_G, 0}, {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_FS, 0}, {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_G, 0}, {TONE_AS, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_FS, 0}, {TONE_AS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_G, 0}, {TONE_AS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_FS, 0}, {TONE_A, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_E, 0} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_GS, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_G, 0}, {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_G, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_G, 0}, {TONE_AS, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_E, 0}, {TONE_G, 0}, {TONE_AS, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_F, 0} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_GS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_GS, 0}, {TONE_C, 1}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_A, 0}, {TONE_C, 1}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_F, 0}, {TONE_GS, 0}, {TONE_AS, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_FS, 0} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_AS, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_AS, 0}, {TONE_C, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_AS, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_FS, 0}, {TONE_A, 0}, {TONE_C, 1} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_G, 0} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_AS, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_AS, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_AS, 0}, {TONE_C, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_G, 0}, {TONE_AS, 0}, {TONE_C, 1} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_GS, 0} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_C, 1}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_C, 1}, {TONE_D, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_C, 1}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_GS, 0}, {TONE_B, 0}, {TONE_D, 1} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_A, 0} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_D, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_E, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_A, 0}, {TONE_C, 1}, {TONE_D, 1} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_AS, 0} } }, 
	&MidiSequence { []Note{ {TONE_DS, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_C, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_A, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_C, 1}, {TONE_F, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_C, 1}, {TONE_E, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_G, 1} } }, 
	&MidiSequence { []Note{ {TONE_AS, 0}, {TONE_C, 1}, {TONE_E, 0} } } },
[]*MidiSequence{ 
	&MidiSequence { []Note{ {TONE_B, 0} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_A, 1} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_A, 0} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_A, 1} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1}, {TONE_A, 1} } }, 
	&MidiSequence { []Note{ {TONE_B, 0}, {TONE_D, 1}, {TONE_F, 1} } } },
}
