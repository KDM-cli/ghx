package components

import (
	"fmt"
	"strings"

	"github.com/utkarshpatrikar/ghx/internal/styles"
)

type StatusBar struct {
	Provider string
	Model    string
	Help     string
	Message  string
}

func (s StatusBar) View() string {
	ai := fmt.Sprintf("AI: %s", s.Provider)
	if s.Model != "" {
		ai = fmt.Sprintf("%s (%s)", ai, s.Model)
	}

	parts := []string{styles.Subtle.Render(ai), styles.Subtle.Render(s.Help)}
	if s.Message != "" {
		parts = append(parts, styles.Info.Render(s.Message))
	}
	return styles.Bar.Render(strings.Join(parts, "  |  "))
}
