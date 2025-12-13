// SPDX-License-Identifier: AGPL-3.0-or-later
// © 2025 Stephan Hegemann

package libkeywalk

import (
	"fmt"
	"os"
	"unicode/utf8"
	"strings"
	"bufio"
	"slices"
	"errors"
	"log"
)

type Keymap struct {
	keys []rune
	key_to_neighbours map[rune][]rune
}



func ParseKeymapFile(keymap_file_path string) (Keymap, error) {
	keymap, err := os.ReadFile(keymap_file_path)
	if err != nil {
		log.Fatal(err)
	}

	keymap_slice := strings.Split(string(keymap), string('\n'))
	last_keymap_slice_index := len(keymap_slice) - 1
	if keymap_slice[last_keymap_slice_index] == "" {
		keymap_slice = keymap_slice[:last_keymap_slice_index]
	}
	
	var keymap_struct Keymap
	keymap_struct.keys = make([]rune, len(keymap_slice))
	keymap_struct.key_to_neighbours = make(map[rune][]rune)
	
	for i, m := range(keymap_slice) {
		m_runes := make([]rune, utf8.RuneCountInString(m))

		// Convert the line of the keymap file into a rune array
		rune_width := 0
		for l := 0; l < len(m_runes); l++ {
			m_runes[l], rune_width = utf8.DecodeRuneInString(m)
			m = m[rune_width:]
		}
		keymap_struct.keys[i] = m_runes[0]
		keymap_struct.key_to_neighbours[m_runes[0]] = m_runes[1:]
	}

	return keymap_struct, nil
}



func reallyWalkKeyAndOutput(key rune, keymap *Keymap, out *bufio.Writer, min_length int, max_length int, current_length int, keywalk_runes []rune) error {
	if current_length >= min_length {
		_, err := fmt.Fprintln(out, string(keywalk_runes[:current_length]))
		if err != nil {
			return err
		}
	}
	if current_length >= max_length {
		return nil
	}

	neighbours := keymap.key_to_neighbours[key]
	for _, neighbour := range(neighbours) {
		keywalk_runes[current_length] = neighbour
		key = neighbour
		err := reallyWalkKeyAndOutput(key, keymap, out, min_length, max_length, current_length + 1, keywalk_runes)
		if err != nil {
			return err
		}
	}
	return nil
}

// This is a wrapper to set up the basis for doing the recursions in reallyWalkKeyAndOutput
func WalkKeyAndOutput(key rune, keymap *Keymap, out *os.File, min_length int, max_length int)  error {
	if min_length <= 0 {
		return errors.New("min_length has to be > 0")
	}
	if max_length <= 0 {
		return errors.New("max_length has to be > 0")
	}
	if max_length < min_length {
		return errors.New("max_length has to be >= min_length")
	}
	keywalk_runes := make([]rune, max_length)
	keywalk_runes[0] = key	// with this we already made the first step
	current_length := 1		// therefore current_length is already 1
	writer := bufio.NewWriter(out)
	
	err := reallyWalkKeyAndOutput(key, keymap, writer, min_length, max_length, current_length, keywalk_runes)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}



func WalkCompleteKeymapFile(keymap_file_path string, min_length int, max_length int, out *os.File) error {
	keymap, err := ParseKeymapFile(keymap_file_path)
	if err != nil {
		return err
	}

	for _, k := range(keymap.keys){
		err := WalkKeyAndOutput(k, &keymap, out, min_length, max_length)
		if err != nil {
			return err
		}
	}
	return nil
}



func CheckIfKeywalk (line string, keymap Keymap) bool {
	is_keywalk := true
	line_runes := make([]rune, utf8.RuneCountInString(line))

	// convert line into rune array
	rune_width := 0
	for i := 0; i < len(line_runes); i++ {
			line_runes[i], rune_width = utf8.DecodeRuneInString(line)
			line = line[rune_width:]
	}

	for i := 0; i < len(line_runes) - 1; i++ {
		rune_neighbours := keymap.key_to_neighbours[line_runes[i]]

		if slices.Contains(rune_neighbours, line_runes[i + 1]) == false {
			is_keywalk = false
		}
	}

	return is_keywalk
}