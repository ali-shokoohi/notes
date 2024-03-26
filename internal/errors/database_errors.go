package errors

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Predefined error variables
var (
	ErrRecordNotFound                = errors.New("record not found")
	ErrInvalidTransaction            = errors.New("invalid transaction")
	ErrNotImplemented                = errors.New("not implemented")
	ErrMissingWhereClause            = errors.New("missing WHERE clause")
	ErrUnsupportedRelation           = errors.New("unsupported relations")
	ErrPrimaryKeyRequired            = errors.New("primary key required")
	ErrModelValueRequired            = errors.New("model value required")
	ErrModelAccessibleFieldsRequired = errors.New("model accessible fields required")
	ErrSubQueryRequired              = errors.New("subquery required")
	ErrInvalidData                   = errors.New("invalid data")
	ErrUnsupportedDriver             = errors.New("unsupported driver")
	ErrRegistered                    = errors.New("already registered")
	ErrInvalidField                  = errors.New("invalid field")
	ErrEmptySlice                    = errors.New("empty slice found")
	ErrDryRunModeUnsupported         = errors.New("dry run mode unsupported")
	ErrInvalidDB                     = errors.New("invalid DB")
	ErrInvalidValue                  = errors.New("invalid value")
	ErrInvalidValueOfLength          = errors.New("invalid value of length")
	ErrPreloadNotAllowed             = errors.New("preload not allowed")
	ErrDuplicatedKey                 = errors.New("duplicated key not allowed")
	ErrForeignKeyViolated            = errors.New("foreign key constraint violated")
)

// ConvertGormError checks the type of Gorm error and converts it to a predefined application error
func ConvertGormError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrRecordNotFound
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return ErrInvalidTransaction
	case errors.Is(err, gorm.ErrNotImplemented):
		return ErrNotImplemented
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return ErrMissingWhereClause
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return ErrUnsupportedRelation
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return ErrPrimaryKeyRequired
	case errors.Is(err, gorm.ErrModelValueRequired):
		return ErrModelValueRequired
	case errors.Is(err, gorm.ErrModelAccessibleFieldsRequired):
		return ErrModelAccessibleFieldsRequired
	case errors.Is(err, gorm.ErrSubQueryRequired):
		return ErrSubQueryRequired
	case errors.Is(err, gorm.ErrInvalidData):
		return ErrInvalidData
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return ErrUnsupportedDriver
	case errors.Is(err, gorm.ErrRegistered):
		return ErrRegistered
	case errors.Is(err, gorm.ErrInvalidField):
		return ErrInvalidField
	case errors.Is(err, gorm.ErrEmptySlice):
		return ErrEmptySlice
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return ErrDryRunModeUnsupported
	case errors.Is(err, gorm.ErrInvalidDB):
		return ErrInvalidDB
	case errors.Is(err, gorm.ErrInvalidValue):
		return ErrInvalidValue
	case errors.Is(err, gorm.ErrInvalidValueOfLength):
		return ErrInvalidValueOfLength
	case errors.Is(err, gorm.ErrPreloadNotAllowed):
		return ErrPreloadNotAllowed
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrDuplicatedKey
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return ErrForeignKeyViolated
	default:
		// For errors not explicitly handled, return a generic error with the original error's message
		return fmt.Errorf("unknown database error: %v", err)
	}
}
