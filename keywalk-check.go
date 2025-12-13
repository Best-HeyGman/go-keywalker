// SPDX-License-Identifier: AGPL-3.0-or-later
// © 2025 Stephan Hegemann

package main

import (
	"go-keywalker/libkeywalk"
	"os"
	"flag"
	"fmt"
	"bufio"
	"strings"
	"log"
)

// variables for cli arguments
var (
	cli_keymap string
	cli_reverse_mode bool
)

func init() {
	//definition of available cli arguments
	flag.StringVar(&cli_keymap, "k", "", "Mandatory. Keymap file (.nsk) of the keyboard you want to walk.")
	flag.BoolVar(&cli_reverse_mode, "r", false, "Reverse. Instead of only showing the keywalks, show anything but the keywalks from the input.")
}

func printHelpAndExit() {
	fmt.Fprintln(os.Stderr, "Pipe your input into keywalk-check\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Parse()

	if len(cli_keymap) <= 0 {
		printHelpAndExit()
	}

	keymap, err := libkeywalk.ParseKeymapFile(cli_keymap)
	if err != nil {
		log.Fatal(err)
	}

	stdout := bufio.NewWriter(os.Stdout)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		in_line := stdin.Text()
		line := strings.ToValidUTF8(string(in_line), "")
		if cli_reverse_mode == true {
			if libkeywalk.CheckIfKeywalk(line, keymap) == false {
				fmt.Fprintln(stdout, line)
			}
		} else {
			if libkeywalk.CheckIfKeywalk(line, keymap) == true {
				fmt.Fprintln(stdout, line)
			}
		}
	}
	stdout.Flush()
	
}