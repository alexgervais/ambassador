// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/lds.proto

package envoy_api_v3alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"

	envoy_api_v3alpha_core "github.com/datawire/ambassador/go/apis/envoy/api/v3alpha/core"
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
	_ = types.DynamicAny{}

	_ = envoy_api_v3alpha_core.TrafficDirection(0)
)

// Validate checks the field values on Listener with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if m.GetAddress() == nil {
		return ListenerValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetAddress()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetFilterChains() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenerValidationError{
						field:  fmt.Sprintf("FilterChains[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetUseOriginalDst()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "UseOriginalDst",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetPerConnectionBufferLimitBytes()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "PerConnectionBufferLimitBytes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMetadata()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetDeprecatedV1()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "DeprecatedV1",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for DrainType

	for idx, item := range m.GetListenerFilters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenerValidationError{
						field:  fmt.Sprintf("ListenerFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetListenerFiltersTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "ListenerFiltersTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for ContinueOnListenerFiltersTimeout

	{
		tmp := m.GetTransparent()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Transparent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetFreebind()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Freebind",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetSocketOptions() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenerValidationError{
						field:  fmt.Sprintf("SocketOptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetTcpFastOpenQueueLength()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "TcpFastOpenQueueLength",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for TrafficDirection

	{
		tmp := m.GetUdpListenerConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "UdpListenerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListenerValidationError is the validation error returned by
// Listener.Validate if the designated constraints aren't met.
type ListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerValidationError) ErrorName() string { return "ListenerValidationError" }

// Error satisfies the builtin error interface
func (e ListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerValidationError{}

// Validate checks the field values on Listener_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Listener_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetBindToPort()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return Listener_DeprecatedV1ValidationError{
					field:  "BindToPort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// Listener_DeprecatedV1ValidationError is the validation error returned by
// Listener_DeprecatedV1.Validate if the designated constraints aren't met.
type Listener_DeprecatedV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Listener_DeprecatedV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Listener_DeprecatedV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Listener_DeprecatedV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Listener_DeprecatedV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Listener_DeprecatedV1ValidationError) ErrorName() string {
	return "Listener_DeprecatedV1ValidationError"
}

// Error satisfies the builtin error interface
func (e Listener_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener_DeprecatedV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Listener_DeprecatedV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Listener_DeprecatedV1ValidationError{}
