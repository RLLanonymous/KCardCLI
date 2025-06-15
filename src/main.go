package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"KCardCLI/src/ui"
)

type CardType string

const (
	TaskForce  CardType = "KeycardCustomTaskForce"
	Management CardType = "KeycardCustomManagement"
	Site02     CardType = "KeycardCustomSite02"
	MetalCase  CardType = "KeycardCustomMetalCase"
)

var cardOptions = []CardType{
	TaskForce,
	Management,
	Site02,
	MetalCase,
}

type screen int

const (
	screenSelect screen = iota
	screenFieldInput
	screenDone
)

type model struct {
	stage        screen
	cursor       int
	selectedType CardType
	fieldIndex   int
	input        textinput.Model
	answers      map[string]string
	generatedCmd string
	fieldsToAsk  []string
	errorMsg     string
	copied       bool
}

var fieldsByType = map[CardType][]string{
	TaskForce:  {"InventoryItemName", "ContainmentLevel", "ArmoryLevel", "AdminLevel", "PermissionColor", "PrimaryTintColor", "CardHolderName", "SerialNumber", "RankDetailOption"},
	Management: {"InventoryItemName", "ContainmentLevel", "ArmoryLevel", "AdminLevel", "PermissionColor", "PrimaryTintColor", "Label", "LabelTextColor"},
	Site02:     {"InventoryItemName", "ContainmentLevel", "ArmoryLevel", "AdminLevel", "PermissionColor", "PrimaryTintColor", "Label", "LabelTextColor", "CardHolderName", "WearLevel"},
	MetalCase:  {"InventoryItemName", "ContainmentLevel", "ArmoryLevel", "AdminLevel", "PermissionColor", "PrimaryTintColor", "Label", "LabelTextColor", "CardHolderName", "SerialNumber"},
}

var intFields = map[string]bool{
	"ContainmentLevel": true,
	"ArmoryLevel":      true,
	"AdminLevel":       true,
	"WearLevel":        true,
	"RankDetailOption": true,
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.stage {
		case screenSelect:
			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(cardOptions)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.selectedType = cardOptions[m.cursor]
				m.fieldsToAsk = fieldsByType[m.selectedType]
				m.answers = make(map[string]string)
				m.input = textinput.New()
				m.input.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ADEED9"))
				m.input.Prompt = "❯ "
				m.input.Placeholder = m.fieldsToAsk[0]
				m.input.Focus()
				m.stage = screenFieldInput
				return m, nil
			case "ctrl+c":
				return m, tea.Quit
			}

		case screenFieldInput:
			switch msg.String() {
			case "enter":
				field := m.fieldsToAsk[m.fieldIndex]
				value := sanitizeInput(m.input.Value())

				if intFields[field] {
					num, err := strconv.Atoi(value)
					if err != nil || num < 0 || num > 3 {
						m.errorMsg = "Error : The value must be an integer between 0 and 3."
						return m, nil
					}
				}

				m.answers[field] = value
				m.fieldIndex++
				m.errorMsg = ""

				if m.fieldIndex >= len(m.fieldsToAsk) {
					m.generatedCmd = m.generateCommand()
					m.stage = screenDone
					return m, nil
				}

				m.input = textinput.New()
				m.input.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ADEED9"))
				m.input.Prompt = "❯ "
				m.input.Placeholder = m.fieldsToAsk[m.fieldIndex]
				m.input.Focus()
				return m, nil
			case "ctrl+c":
				return m, tea.Quit
			}
			m.input, cmd = m.input.Update(msg)
			return m, cmd

		case screenDone:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "c":
				err := clipboard.WriteAll(m.generatedCmd)
				if err != nil {
					m.errorMsg = "Error : Failed to copy to clipboard."
				} else {
					m.copied = true
					m.errorMsg = ""
				}
				return m, nil
			}
		}
	}

	if m.stage == screenFieldInput {
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var s string

	switch m.stage {
	case screenSelect:
		s = ui.StyleTitle.Render("Choose a card type:") + "\n\n"
		for i, opt := range cardOptions {
			cursor := "  "
			line := string(opt)
			if m.cursor == i {
				cursor = ui.StyleCursor.Render("➔ ")
				line = ui.StyleSelected.Render(line)
			}
			s += fmt.Sprintf("%s%s\n", cursor, line)
		}
		s += "\n(Press Enter to confirm, Ctrl+C to quit)"

	case screenFieldInput:
		err := ""
		if m.errorMsg != "" {
			err = "\n\n" + ui.StyleError.Render(m.errorMsg)
		}
		s = fmt.Sprintf("%s :\n%s%s\n\n(Press Enter to confirm, Ctrl+C to quit)",
			m.fieldsToAsk[m.fieldIndex],
			ui.StyleInput.Render(m.input.View()),
			err)

	case screenDone:
		s = ui.StyleTitle.Render("Command generated:") + "\n\n"
		s += ui.StyleCommand.Render(m.generatedCmd) + "\n\n"
		s += "Ctrl+C to quit. Press 'c' to copy the command."

		if m.copied {
			s += "\n\n" + ui.StyleTitle.Render("Command copied to clipboard!")
		}

		if m.errorMsg != "" {
			s += "\n\n" + ui.StyleError.Render(m.errorMsg)
		}
	}

	return ui.StyleBox.Render(s)
}

func (m model) generateCommand() string {
	args := []string{"customkeycard", "id", string(m.selectedType)}
	for _, field := range m.fieldsToAsk {
		args = append(args, m.answers[field])
	}
	return strings.Join(args, " ")
}

func sanitizeInput(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "_")
	return s
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
