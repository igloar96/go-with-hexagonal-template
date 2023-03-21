package driven

import (
	"fmt"

	"github.com/igloar96/hexa-notification/core/domain"
)

type ConsoleAdapter struct {
}

func NewConsoleAdapter() *ConsoleAdapter {
	return &ConsoleAdapter{}
}

func (s *ConsoleAdapter) Execute(message *domain.Message) error {
	fmt.Printf("[ConsoleAdapter] message: %s", message.Text)
	return nil
}
