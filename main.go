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
	"fmt"
	"github.com/rakyll/portmidi"
	"os"
)

func main() {
	args := os.Args[1:]
	settings := parseCommandLine(args)
	midi := MidiOutput{initPortMidi()}
	guitar := SimpleGuitar{}
	guitar.Reset()
	guitar.midi = &midi
	device := initDevice(&guitar, &midi, settings)
	device.StartListening()
	shutDownPortMidi()
}

func initDevice(instrument Instrument, midi *MidiOutput, settings Settings) Device {
	var device Device
	if settings.deviceType == DEVICE_GUITAR {
		device = CreateGuitarDevice(settings.deviceName, instrument)
	} else {
		fmt.Println("Error: Keyboard input not supported yet.")
		os.Exit(12)
	}
	return device
}

func initPortMidi() *portmidi.Stream {
	portmidi.Initialize()

	deviceCount := portmidi.CountDevices()
	fmt.Println("Number of MIDI devices: ", deviceCount)

	dev := portmidi.DeviceId(3 - 1)

	for i := 0; i < deviceCount; i++ {
	    id := portmidi.DeviceId(i)
	    fmt.Println("Index ", i, "Id", id, " Device ", *portmidi.GetDeviceInfo(id))
	}

	out, err := portmidi.NewOutputStream(dev, 0, 0)
	if err == nil {
		fmt.Println("used device info: ", *portmidi.GetDeviceInfo(dev))
	} else {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	return out
}

func shutDownPortMidi() {
	portmidi.Terminate()
}
