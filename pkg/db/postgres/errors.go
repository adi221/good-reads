package postgres

import (
	"github.com/adi221/good-reads/pkg/model"
	"strings"
)

func mapError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "duplicate key") {
		return model.ErrAlreadyExists
	}

	return err
}
