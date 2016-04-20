package main

import "fmt"

type ConfigurationGuitarMode struct {
	guitar *SimpleGuitar
	chordMap *ChordMap
	currentChord int
}

func (mode *ConfigurationGuitarMode) EnterMode(guitar *SimpleGuitar) {
	mode.guitar = guitar
	mode.chordMap = guitar.chordMap
	mode.currentChord = -1
	fmt.Println("ENTER configuration mode")
}

func (mode *ConfigurationGuitarMode) ButtonPressed(buttonType int, buttonState ButtonState) {
	if buttonType == BUTTON_START {
		mode.guitar.ModeSwitch(0)
	} else if buttonType == BUTTON_SELECT {
		mode.currentChord = -1	
	} else if mode.currentChord != -1 {
		mode.configure()
	}
}

func (mode *ConfigurationGuitarMode) ButtonReleased(buttonType int, buttonState ButtonState) {
	// nothing
}

func (mode *ConfigurationGuitarMode) configure() {
	guitar := mode.guitar
	val := 0
	switch {
	case guitar.match(true, false, false, false, false): val = 1
	case guitar.match(false, true, false, false, false): val = 2
	case guitar.match(false, false, true, false, false): val = 3
	case guitar.match(false, false, false, true, false): val = 4
	case guitar.match(false, false, false, false, true): val = 5
	}
	mode.setValue(val)
}

func (mode *ConfigurationGuitarMode) Up() {
	mode.setValue(6)
}

func (mode *ConfigurationGuitarMode) Down() {
	mode.setValue(7)
}

func (mode *ConfigurationGuitarMode) setValue(val int) {
	chordMap := mode.chordMap
	tone := ChordToTone(mode.currentChord)
	fmt.Println(chordMap.chords)
	newVal := CHORD_TABLE[tone][val]
	chordMap.chords[mode.currentChord] = newVal
	fmt.Println("DONE configuring chord ", CHORD_NAME[mode.currentChord])
	mode.currentChord = -1
}

func (mode *ConfigurationGuitarMode) StrumUp() {

}

func (mode *ConfigurationGuitarMode) StrumDown() {
	if mode.currentChord == -1 {
		mode.currentChord = mode.guitar.ProduceChord()
		fmt.Println("CONFIGURE chord ", CHORD_NAME[mode.currentChord])
	} else {
		if mode.guitar.match(false, false, false, false, false) {
			mode.setValue(0)		
		}
	}
}

func (mode *ConfigurationGuitarMode) StrumReleased() {

}

func (mode *ConfigurationGuitarMode) ExitMode() {
	fmt.Println("EXIT configuration mode")
}
