// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example/greeter/v1/api.proto

package greeterv1

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

// Validate checks the field values on GreetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GreetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GreetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GreetRequestMultiError, or
// nil if none found.
func (m *GreetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GreetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _GreetRequest_Name_NotInLookup[m.GetName()]; ok {
		err := GreetRequestValidationError{
			field:  "Name",
			reason: "value must not be in list [Jack Smith Jackity Jack]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := GreetRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSomeValue() == nil {
		err := GreetRequestValidationError{
			field:  "SomeValue",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSomeValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GreetRequestValidationError{
					field:  "SomeValue",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GreetRequestValidationError{
					field:  "SomeValue",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSomeValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GreetRequestValidationError{
				field:  "SomeValue",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GreetRequestMultiError(errors)
	}

	return nil
}

// GreetRequestMultiError is an error wrapping multiple validation errors
// returned by GreetRequest.ValidateAll() if the designated constraints aren't met.
type GreetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GreetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GreetRequestMultiError) AllErrors() []error { return m }

// GreetRequestValidationError is the validation error returned by
// GreetRequest.Validate if the designated constraints aren't met.
type GreetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GreetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GreetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GreetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GreetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GreetRequestValidationError) ErrorName() string { return "GreetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GreetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGreetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GreetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GreetRequestValidationError{}

var _GreetRequest_Name_NotInLookup = map[string]struct{}{
	"Jack Smith":   {},
	"Jackity Jack": {},
}

// Validate checks the field values on GreetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GreetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GreetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GreetResponseMultiError, or
// nil if none found.
func (m *GreetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GreetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return GreetResponseMultiError(errors)
	}

	return nil
}

// GreetResponseMultiError is an error wrapping multiple validation errors
// returned by GreetResponse.ValidateAll() if the designated constraints
// aren't met.
type GreetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GreetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GreetResponseMultiError) AllErrors() []error { return m }

// GreetResponseValidationError is the validation error returned by
// GreetResponse.Validate if the designated constraints aren't met.
type GreetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GreetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GreetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GreetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GreetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GreetResponseValidationError) ErrorName() string { return "GreetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GreetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGreetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GreetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GreetResponseValidationError{}