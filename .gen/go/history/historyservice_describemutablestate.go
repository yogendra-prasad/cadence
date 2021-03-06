// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package history

import (
	errors "errors"
	fmt "fmt"
	shared "github.com/uber/cadence/.gen/go/shared"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// HistoryService_DescribeMutableState_Args represents the arguments for the HistoryService.DescribeMutableState function.
//
// The arguments for DescribeMutableState are sent and received over the wire as this struct.
type HistoryService_DescribeMutableState_Args struct {
	Request *DescribeMutableStateRequest `json:"request,omitempty"`
}

// ToWire translates a HistoryService_DescribeMutableState_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *HistoryService_DescribeMutableState_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeMutableStateRequest_Read(w wire.Value) (*DescribeMutableStateRequest, error) {
	var v DescribeMutableStateRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a HistoryService_DescribeMutableState_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_DescribeMutableState_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_DescribeMutableState_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_DescribeMutableState_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _DescribeMutableStateRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a HistoryService_DescribeMutableState_Args
// struct.
func (v *HistoryService_DescribeMutableState_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("HistoryService_DescribeMutableState_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_DescribeMutableState_Args match the
// provided HistoryService_DescribeMutableState_Args.
//
// This function performs a deep comparison.
func (v *HistoryService_DescribeMutableState_Args) Equals(rhs *HistoryService_DescribeMutableState_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of HistoryService_DescribeMutableState_Args.
func (v *HistoryService_DescribeMutableState_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Args) GetRequest() (o *DescribeMutableStateRequest) {
	if v != nil && v.Request != nil {
		return v.Request
	}

	return
}

// IsSetRequest returns true if Request is not nil.
func (v *HistoryService_DescribeMutableState_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "DescribeMutableState" for this struct.
func (v *HistoryService_DescribeMutableState_Args) MethodName() string {
	return "DescribeMutableState"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *HistoryService_DescribeMutableState_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// HistoryService_DescribeMutableState_Helper provides functions that aid in handling the
// parameters and return values of the HistoryService.DescribeMutableState
// function.
var HistoryService_DescribeMutableState_Helper = struct {
	// Args accepts the parameters of DescribeMutableState in-order and returns
	// the arguments struct for the function.
	Args func(
		request *DescribeMutableStateRequest,
	) *HistoryService_DescribeMutableState_Args

	// IsException returns true if the given error can be thrown
	// by DescribeMutableState.
	//
	// An error can be thrown by DescribeMutableState only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for DescribeMutableState
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// DescribeMutableState into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by DescribeMutableState
	//
	//   value, err := DescribeMutableState(args)
	//   result, err := HistoryService_DescribeMutableState_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from DescribeMutableState: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*DescribeMutableStateResponse, error) (*HistoryService_DescribeMutableState_Result, error)

	// UnwrapResponse takes the result struct for DescribeMutableState
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if DescribeMutableState threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := HistoryService_DescribeMutableState_Helper.UnwrapResponse(result)
	UnwrapResponse func(*HistoryService_DescribeMutableState_Result) (*DescribeMutableStateResponse, error)
}{}

func init() {
	HistoryService_DescribeMutableState_Helper.Args = func(
		request *DescribeMutableStateRequest,
	) *HistoryService_DescribeMutableState_Args {
		return &HistoryService_DescribeMutableState_Args{
			Request: request,
		}
	}

	HistoryService_DescribeMutableState_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.AccessDeniedError:
			return true
		case *ShardOwnershipLostError:
			return true
		case *shared.LimitExceededError:
			return true
		default:
			return false
		}
	}

	HistoryService_DescribeMutableState_Helper.WrapResponse = func(success *DescribeMutableStateResponse, err error) (*HistoryService_DescribeMutableState_Result, error) {
		if err == nil {
			return &HistoryService_DescribeMutableState_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.BadRequestError")
			}
			return &HistoryService_DescribeMutableState_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.InternalServiceError")
			}
			return &HistoryService_DescribeMutableState_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.EntityNotExistError")
			}
			return &HistoryService_DescribeMutableState_Result{EntityNotExistError: e}, nil
		case *shared.AccessDeniedError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.AccessDeniedError")
			}
			return &HistoryService_DescribeMutableState_Result{AccessDeniedError: e}, nil
		case *ShardOwnershipLostError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.ShardOwnershipLostError")
			}
			return &HistoryService_DescribeMutableState_Result{ShardOwnershipLostError: e}, nil
		case *shared.LimitExceededError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_DescribeMutableState_Result.LimitExceededError")
			}
			return &HistoryService_DescribeMutableState_Result{LimitExceededError: e}, nil
		}

		return nil, err
	}
	HistoryService_DescribeMutableState_Helper.UnwrapResponse = func(result *HistoryService_DescribeMutableState_Result) (success *DescribeMutableStateResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.AccessDeniedError != nil {
			err = result.AccessDeniedError
			return
		}
		if result.ShardOwnershipLostError != nil {
			err = result.ShardOwnershipLostError
			return
		}
		if result.LimitExceededError != nil {
			err = result.LimitExceededError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// HistoryService_DescribeMutableState_Result represents the result of a HistoryService.DescribeMutableState function call.
//
// The result of a DescribeMutableState execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type HistoryService_DescribeMutableState_Result struct {
	// Value returned by DescribeMutableState after a successful execution.
	Success                 *DescribeMutableStateResponse `json:"success,omitempty"`
	BadRequestError         *shared.BadRequestError       `json:"badRequestError,omitempty"`
	InternalServiceError    *shared.InternalServiceError  `json:"internalServiceError,omitempty"`
	EntityNotExistError     *shared.EntityNotExistsError  `json:"entityNotExistError,omitempty"`
	AccessDeniedError       *shared.AccessDeniedError     `json:"accessDeniedError,omitempty"`
	ShardOwnershipLostError *ShardOwnershipLostError      `json:"shardOwnershipLostError,omitempty"`
	LimitExceededError      *shared.LimitExceededError    `json:"limitExceededError,omitempty"`
}

// ToWire translates a HistoryService_DescribeMutableState_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *HistoryService_DescribeMutableState_Result) ToWire() (wire.Value, error) {
	var (
		fields [7]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AccessDeniedError != nil {
		w, err = v.AccessDeniedError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.ShardOwnershipLostError != nil {
		w, err = v.ShardOwnershipLostError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.LimitExceededError != nil {
		w, err = v.LimitExceededError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("HistoryService_DescribeMutableState_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeMutableStateResponse_Read(w wire.Value) (*DescribeMutableStateResponse, error) {
	var v DescribeMutableStateResponse
	err := v.FromWire(w)
	return &v, err
}

func _EntityNotExistsError_Read(w wire.Value) (*shared.EntityNotExistsError, error) {
	var v shared.EntityNotExistsError
	err := v.FromWire(w)
	return &v, err
}

func _ShardOwnershipLostError_Read(w wire.Value) (*ShardOwnershipLostError, error) {
	var v ShardOwnershipLostError
	err := v.FromWire(w)
	return &v, err
}

func _LimitExceededError_Read(w wire.Value) (*shared.LimitExceededError, error) {
	var v shared.LimitExceededError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a HistoryService_DescribeMutableState_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_DescribeMutableState_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_DescribeMutableState_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_DescribeMutableState_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _DescribeMutableStateResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.AccessDeniedError, err = _AccessDeniedError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.ShardOwnershipLostError, err = _ShardOwnershipLostError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.LimitExceededError, err = _LimitExceededError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.AccessDeniedError != nil {
		count++
	}
	if v.ShardOwnershipLostError != nil {
		count++
	}
	if v.LimitExceededError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("HistoryService_DescribeMutableState_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a HistoryService_DescribeMutableState_Result
// struct.
func (v *HistoryService_DescribeMutableState_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [7]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.AccessDeniedError != nil {
		fields[i] = fmt.Sprintf("AccessDeniedError: %v", v.AccessDeniedError)
		i++
	}
	if v.ShardOwnershipLostError != nil {
		fields[i] = fmt.Sprintf("ShardOwnershipLostError: %v", v.ShardOwnershipLostError)
		i++
	}
	if v.LimitExceededError != nil {
		fields[i] = fmt.Sprintf("LimitExceededError: %v", v.LimitExceededError)
		i++
	}

	return fmt.Sprintf("HistoryService_DescribeMutableState_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_DescribeMutableState_Result match the
// provided HistoryService_DescribeMutableState_Result.
//
// This function performs a deep comparison.
func (v *HistoryService_DescribeMutableState_Result) Equals(rhs *HistoryService_DescribeMutableState_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.AccessDeniedError == nil && rhs.AccessDeniedError == nil) || (v.AccessDeniedError != nil && rhs.AccessDeniedError != nil && v.AccessDeniedError.Equals(rhs.AccessDeniedError))) {
		return false
	}
	if !((v.ShardOwnershipLostError == nil && rhs.ShardOwnershipLostError == nil) || (v.ShardOwnershipLostError != nil && rhs.ShardOwnershipLostError != nil && v.ShardOwnershipLostError.Equals(rhs.ShardOwnershipLostError))) {
		return false
	}
	if !((v.LimitExceededError == nil && rhs.LimitExceededError == nil) || (v.LimitExceededError != nil && rhs.LimitExceededError != nil && v.LimitExceededError.Equals(rhs.LimitExceededError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of HistoryService_DescribeMutableState_Result.
func (v *HistoryService_DescribeMutableState_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BadRequestError != nil {
		err = multierr.Append(err, enc.AddObject("badRequestError", v.BadRequestError))
	}
	if v.InternalServiceError != nil {
		err = multierr.Append(err, enc.AddObject("internalServiceError", v.InternalServiceError))
	}
	if v.EntityNotExistError != nil {
		err = multierr.Append(err, enc.AddObject("entityNotExistError", v.EntityNotExistError))
	}
	if v.AccessDeniedError != nil {
		err = multierr.Append(err, enc.AddObject("accessDeniedError", v.AccessDeniedError))
	}
	if v.ShardOwnershipLostError != nil {
		err = multierr.Append(err, enc.AddObject("shardOwnershipLostError", v.ShardOwnershipLostError))
	}
	if v.LimitExceededError != nil {
		err = multierr.Append(err, enc.AddObject("limitExceededError", v.LimitExceededError))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetSuccess() (o *DescribeMutableStateResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetBadRequestError returns the value of BadRequestError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetBadRequestError() (o *shared.BadRequestError) {
	if v != nil && v.BadRequestError != nil {
		return v.BadRequestError
	}

	return
}

// IsSetBadRequestError returns true if BadRequestError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetBadRequestError() bool {
	return v != nil && v.BadRequestError != nil
}

// GetInternalServiceError returns the value of InternalServiceError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetInternalServiceError() (o *shared.InternalServiceError) {
	if v != nil && v.InternalServiceError != nil {
		return v.InternalServiceError
	}

	return
}

// IsSetInternalServiceError returns true if InternalServiceError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetInternalServiceError() bool {
	return v != nil && v.InternalServiceError != nil
}

// GetEntityNotExistError returns the value of EntityNotExistError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetEntityNotExistError() (o *shared.EntityNotExistsError) {
	if v != nil && v.EntityNotExistError != nil {
		return v.EntityNotExistError
	}

	return
}

// IsSetEntityNotExistError returns true if EntityNotExistError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetEntityNotExistError() bool {
	return v != nil && v.EntityNotExistError != nil
}

// GetAccessDeniedError returns the value of AccessDeniedError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetAccessDeniedError() (o *shared.AccessDeniedError) {
	if v != nil && v.AccessDeniedError != nil {
		return v.AccessDeniedError
	}

	return
}

// IsSetAccessDeniedError returns true if AccessDeniedError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetAccessDeniedError() bool {
	return v != nil && v.AccessDeniedError != nil
}

// GetShardOwnershipLostError returns the value of ShardOwnershipLostError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetShardOwnershipLostError() (o *ShardOwnershipLostError) {
	if v != nil && v.ShardOwnershipLostError != nil {
		return v.ShardOwnershipLostError
	}

	return
}

// IsSetShardOwnershipLostError returns true if ShardOwnershipLostError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetShardOwnershipLostError() bool {
	return v != nil && v.ShardOwnershipLostError != nil
}

// GetLimitExceededError returns the value of LimitExceededError if it is set or its
// zero value if it is unset.
func (v *HistoryService_DescribeMutableState_Result) GetLimitExceededError() (o *shared.LimitExceededError) {
	if v != nil && v.LimitExceededError != nil {
		return v.LimitExceededError
	}

	return
}

// IsSetLimitExceededError returns true if LimitExceededError is not nil.
func (v *HistoryService_DescribeMutableState_Result) IsSetLimitExceededError() bool {
	return v != nil && v.LimitExceededError != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "DescribeMutableState" for this struct.
func (v *HistoryService_DescribeMutableState_Result) MethodName() string {
	return "DescribeMutableState"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *HistoryService_DescribeMutableState_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
