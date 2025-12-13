// SPDX-License-Identifier: AGPL-3.0-or-later
// © 2025 Stephan Hegemann

package main

import (
	"go-keywalker/libkeywalk"
	"os"
	"flag"
	"log"
)

// variables for cli arguments
var (
	cli_keymap string
	cli_min_length int
	cli_max_length int
)

func init() {
	//definition of available cli arguments
	flag.StringVar(&cli_keymap, "k", "", "Mandatory. Keymap file (.nsk) of the keyboard you want to walk.")
	flag.IntVar(&cli_min_length, "m", 1, "Minimum length of the keywalk. Value has to be > 0")
	flag.IntVar(&cli_max_length, "M", 4, "Maximum length of the keywalk. Value has to be > 0")
}

func printHelpAndExit() {
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Parse()

	if len(cli_keymap) <= 0 {
		printHelpAndExit()
	}

	err := libkeywalk.WalkCompleteKeymapFile(cli_keymap, cli_min_length, cli_max_length, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}