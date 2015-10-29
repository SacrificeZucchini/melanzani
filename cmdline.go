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
	"strings"
	"fmt"
	"os"
)

const (
	DEVICE_GUITAR = iota
	DEVICE_KEYBOARD = iota
)

type Settings struct {
	deviceType uint
	deviceName string
}

func printUsage() {
	msg := "Usage: melanzani <DEVICE_TYPE>\n\n" +
		"Possible Values for DEVICE_TYPE:\n" +
		"  --keyboard\tUse your computer's keyboard as input device.\n" +
		"  --guitar\tUse a USB guitar device. Optionally, pass the device's " +
		"name after this parameter.\n\n" +
		"Other options:\n" +
		"  --help\tPrint this help message.\n"
	fmt.Println(msg)
}

func parseCommandLine(args []string) Settings {
	settings := Settings{}
	i := 0
	if len(args) < 1 {
		printUsage()
		os.Exit(0)
	}
	for i < len(args) {
		i = parseCommand(i, args, &settings)
	}
	return settings
}

func parseCommand(index int, args []string, settings *Settings) int {
	cmd := args[index]
	index++
	if cmd == "--help" || cmd == "help" || cmd == "-h" {
		printUsage()
		os.Exit(0)
	} else if cmd == "--keyboard" {
		settings.deviceType = DEVICE_KEYBOARD
	} else if cmd == "--guitar" {
		settings.deviceType = DEVICE_GUITAR
		if len(args) > index && !strings.HasPrefix(args[index], "--") {
			settings.deviceName = args[index]
			index++
		}
	} else {
		fmt.Println("Unknown command: " + cmd)
		os.Exit(1)
	}
	return index
}
