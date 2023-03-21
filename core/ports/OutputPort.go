package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type OutputPort interface {
	Execute(message *domain.Message) error
}
