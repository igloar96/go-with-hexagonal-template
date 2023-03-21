package test

import (
	"errors"
	"testing"

	"github.com/igloar96/hexa-notification/core/domain"
	"github.com/igloar96/hexa-notification/core/ports"
	"github.com/igloar96/hexa-notification/core/useCases"
)

func TestCreateNotification(t *testing.T) {
	t.Run("interfaces", func(t *testing.T) {
		t.Log("mocked useCase should implement UseCase")
		var _ useCases.UseCase = (*MockUseCase)(nil)
	})
	t.Run("interfaces", func(t *testing.T) {
		t.Log("mocked inputPort should implement adapter")
		var _ ports.InputPort = (*InputAdapter)(nil)
	})
	t.Run("interfaces", func(t *testing.T) {
		t.Log("mocked outputPort should implement adapter")
		var _ ports.OutputPort = (*OutputAdapter)(nil)
	})
}

func TestCreateNotificationExecute(t *testing.T) {

	t.Run("TestCreateNotificationExecute_1", func(t *testing.T) {
		t.Log("Expected to adapt Message request body correctly.")
		//arrange
		msg := &InputAdapter{Text: "byli.dev !"}
		useCase := NewMockUseCase(&OutputAdapter{})

		//act
		err := useCase.Execute(msg)
		//assert

		if err != nil {
			t.Errorf("Expected to adapt Message request body correctly but got error: %s", err)
		}
	})
	t.Run("TestCreateNotificationExecute_2", func(t *testing.T) {
		t.Log("Expected to return error if adapt return error.")
		//arrange
		msg := &InputAdapter{Text: ""}
		useCase := NewMockUseCase(&OutputAdapter{})

		//act
		e := useCase.Execute(msg)
		//assert
		if e == nil || e.Error() != "Error adapting message." {
			t.Errorf("Expected to return error if Message text is empty.")
		}

	})
	t.Run("TestCreateNotificationExecute_3", func(t *testing.T) {
		t.Log("Expected to return error if output adapter has an error.")
		//arrange
		msg := &InputAdapter{Text: "byli.dev!"}
		useCase := NewMockUseCase(&OutputAdapter{ErrorMsg: "Error inesperado"})

		//act
		e := useCase.Execute(msg)
		//assert
		if e == nil || e.Error() != "Error inesperado" {
			t.Errorf("Expected to return error if output adapter return error")
		}

	})
}

/*
MOCKS (REPLACE WITH MICROSERVICE IMPLEMENTATION)
*/
type OutputAdapter struct {
	ErrorMsg string
}

func (n *OutputAdapter) Execute(*domain.Message) error {
	if n.ErrorMsg != "" {
		return errors.New(n.ErrorMsg)
	}
	return nil
}

type InputAdapter struct {
	Text string
}

func (n *InputAdapter) Adapt() (*domain.Message, error) {
	if n.Text == "" {
		return nil, errors.New("Error adapting message.")
	}
	return &domain.Message{Text: n.Text}, nil
}

type MockUseCase struct {
	output ports.OutputPort
}

func NewMockUseCase(output ports.OutputPort) *MockUseCase {
	return &MockUseCase{output: output}
}
func (n *MockUseCase) Execute(input ports.InputPort) error {
	msg, err := input.Adapt()
	if err != nil {
		return err
	}
	return n.output.Execute(msg)
}
