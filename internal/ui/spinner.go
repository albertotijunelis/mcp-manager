// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package ui

import (
	"fmt"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SpinnerModel is a Bubble Tea model for displaying a spinner.
type SpinnerModel struct {
	spinner  spinner.Model
	label    string
	done     bool
	err      error
	quitting bool
}

type spinnerDoneMsg struct {
	err error
}

// NewSpinnerModel creates a new spinner with the given label.
func NewSpinnerModel(label string) SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCyan))

	return SpinnerModel{
		spinner: s,
		label:   label,
	}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}

	case spinnerDoneMsg:
		m.done = true
		m.err = msg.err
		return m, tea.Quit

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m SpinnerModel) View() string {
	if m.done {
		if m.err != nil {
			return ErrorStyle.Render("  ✗ " + m.label) + "\n"
		}
		return SuccessStyle.Render("  ✓ " + m.label) + "\n"
	}
	return fmt.Sprintf("  %s %s\n", m.spinner.View(), m.label)
}

// RunWithSpinner runs the given function while showing a spinner with the given label.
// Returns the error from the function, if any.
func RunWithSpinner(label string, fn func() error) error {
	var fnErr error
	var once sync.Once

	model := NewSpinnerModel(label)

	p := tea.NewProgram(model)

	go func() {
		err := fn()
		once.Do(func() {
			fnErr = err
			p.Send(spinnerDoneMsg{err: err})
		})
	}()

	finalModel, err := p.Run()
	if err != nil {
		return err
	}

	if m, ok := finalModel.(SpinnerModel); ok {
		if m.quitting && !m.done {
			return fmt.Errorf("interrupted")
		}
		if m.err != nil {
			return m.err
		}
	}

	return fnErr
}
