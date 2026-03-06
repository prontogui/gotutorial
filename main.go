// Copyright 2024-2026 ProntoGUI, LLC
// ProntoGUI™ is a trademark of ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	// Import the Go library for ProntoGUI
	pg "github.com/prontogui/golib"
)

func main() {

	// Initialize ProntoGUI
	pgui := pg.NewProntoGUI()
	err := pgui.StartServing("127.0.0.1", 50053)

	if err != nil {
		fmt.Printf("Error trying to start server:  %s", err.Error())
		return
	}

	// Build the GUI using primitives

	boldTextEmbodiment := "fontSize:30.0, marginAll: 20, fontWeight:bold"
	normalTextEmbodiment := "fontSize:30.0, marginAll: 20"

	titleText := pg.TextWith{
		Content:    "International Characters Example",
		Embodiment: "fontSize: 20, color: #FF00AAAA, marginAll: 10",
	}.Make()

	helloText := pg.TextWith{
		Content:    "Hello, world!",
		Embodiment: normalTextEmbodiment,
	}.Make()

	languageChoice := pg.ChoiceWith{
		Choices:    []string{"English", "Japanese / 日本語", "German / Deutsche", "Russian / Русский"},
		Choice:     "English",
		Embodiment: "width: 200, horizontalAlignment: center",
	}.Make()

	boldCheck := pg.CheckWith{
		Label:      "Bold Text",
		Embodiment: "width: 200, horizontalAlignment: center",
	}.Make()

	pgui.SetGUI(titleText, helloText, languageChoice, boldCheck)

	for {
		// Wait for something to happen in the GUI
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}

		// Update fields of primitives based on user interactions

		// Set the content of Hello, World based on user selected language choice
		switch languageChoice.Choice() {
		case "English":
			helloText.SetContent("Hello, world!")
		case "Japanese / 日本語":
			helloText.SetContent("こんにちは世界！")
		case "German / Deutsche":
			helloText.SetContent("Hallo Welt!")
		case "Russian / Русский":
			helloText.SetContent("Привет, мир!")
		}

		// Make the text bold or normal based on user selection
		if boldCheck.Checked() {
			helloText.SetEmbodiment(boldTextEmbodiment)
		} else {
			helloText.SetEmbodiment(normalTextEmbodiment)
		}
	}
}
