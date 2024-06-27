// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PriceMultiError, or nil if none found.
func (m *Price) ValidateAll() error {
	return m.validate(true)
}

func (m *Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MarketPrice

	// no validation rules for SalePrice

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriceValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriceValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriceValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PriceMultiError(errors)
	}

	return nil
}

// PriceMultiError is an error wrapping multiple validation errors returned by
// Price.ValidateAll() if the designated constraints aren't met.
type PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PriceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PriceMultiError) AllErrors() []error { return m }

// PriceValidationError is the validation error returned by Price.Validate if
// the designated constraints aren't met.
type PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriceValidationError) ErrorName() string { return "PriceValidationError" }

// Error satisfies the builtin error interface
func (e PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriceValidationError{}

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetAge(); val <= 0 || val >= 120 {
		err := RequestValidationError{
			field:  "Age",
			reason: "value must be inside range (0, 120)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Request_Code_InLookup[m.GetCode()]; !ok {
		err := RequestValidationError{
			field:  "Code",
			reason: "value must be in list [1 2 3]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Request_Score_NotInLookup[m.GetScore()]; ok {
		err := RequestValidationError{
			field:  "Score",
			reason: "value must not be in list [0 99.99]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetState() != true {
		err := RequestValidationError{
			field:  "State",
			reason: "value must equal true",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPath() != "/hello" {
		err := RequestValidationError{
			field:  "Path",
			reason: "value must equal /hello",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPhone()) != 11 {
		err := RequestValidationError{
			field:  "Phone",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetExplain()) < 3 {
		err := RequestValidationError{
			field:  "Explain",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		err := RequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Request_Card_Pattern.MatchString(m.GetCard()) {
		err := RequestValidationError{
			field:  "Card",
			reason: "value does not match regex pattern \"(?i)^[0-9a-f]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

var _Request_Code_InLookup = map[uint32]struct{}{
	1: {},
	2: {},
	3: {},
}

var _Request_Score_NotInLookup = map[float32]struct{}{
	0:     {},
	99.99: {},
}

var _Request_Card_Pattern = regexp.MustCompile("(?i)^[0-9a-f]+$")
