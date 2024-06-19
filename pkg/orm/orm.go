package orm

import (
	"github.com/comsma/knead/pkg/domain"
	"io"
)

type Generator interface {
	WriteFile(w io.Writer, table *domain.Table) error
}
