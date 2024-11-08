package main

import (
	"fmt"
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Time Zone Changer")
	w.Resize(fyne.NewSize(300, 200))

	timeZones := []string{"America/Guyana", "Europe/Rome"}
	selectTZ := widget.NewSelect(timeZones, func(value string) {
		fmt.Println("Selected:", value)
	})
	selectTZ.PlaceHolder = "Select Time Zone"

	applyButton := widget.NewButton("Apply", func() {
		selectedTZ := selectTZ.Selected
		if selectedTZ == "" {
			log.Println("No time zone selected")
			return
		}
		err := setTimeZone(selectedTZ)
		if err != nil {
			log.Println("Failed to set time zone:", err)
		} else {
			log.Println("Time zone set to", selectedTZ)
		}
	})

	w.SetContent(
		container.NewVBox(
			selectTZ,
			applyButton,
		),
	)

	w.ShowAndRun()
}

func setTimeZone(timeZone string) error {
	cmd := exec.Command("pkexec", "timedatectl", "set-timezone", timeZone)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}
	return nil
}
