// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
)

type awsAwsjson11_deserializeOpDescribeDimensionKeys struct {
}

func (*awsAwsjson11_deserializeOpDescribeDimensionKeys) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDescribeDimensionKeys) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeDimensionKeys(response)
	}
	output := &DescribeDimensionKeysOutput{}
	out.Result = output

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeOpDocumentDescribeDimensionKeysOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorDescribeDimensionKeys(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalServiceError", errorCode):
		return awsAwsjson11_deserializeErrorInternalServiceError(response, errorBody)

	case strings.EqualFold("InvalidArgumentException", errorCode):
		return awsAwsjson11_deserializeErrorInvalidArgumentException(response, errorBody)

	case strings.EqualFold("NotAuthorizedException", errorCode):
		return awsAwsjson11_deserializeErrorNotAuthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpGetResourceMetrics struct {
}

func (*awsAwsjson11_deserializeOpGetResourceMetrics) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpGetResourceMetrics) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorGetResourceMetrics(response)
	}
	output := &GetResourceMetricsOutput{}
	out.Result = output

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeOpDocumentGetResourceMetricsOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorGetResourceMetrics(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalServiceError", errorCode):
		return awsAwsjson11_deserializeErrorInternalServiceError(response, errorBody)

	case strings.EqualFold("InvalidArgumentException", errorCode):
		return awsAwsjson11_deserializeErrorInvalidArgumentException(response, errorBody)

	case strings.EqualFold("NotAuthorizedException", errorCode):
		return awsAwsjson11_deserializeErrorNotAuthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson11_deserializeErrorInternalServiceError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.InternalServiceError{}
	err := awsAwsjson11_deserializeDocumentInternalServiceError(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeErrorInvalidArgumentException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.InvalidArgumentException{}
	err := awsAwsjson11_deserializeDocumentInvalidArgumentException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeErrorNotAuthorizedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.NotAuthorizedException{}
	err := awsAwsjson11_deserializeDocumentNotAuthorizedException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeDocumentDataPoint(v **types.DataPoint, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.DataPoint
	if *v == nil {
		sv = &types.DataPoint{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Timestamp":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Timestamp = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Value":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected Double to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Value = &f64
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentDataPointsList(v *[]*types.DataPoint, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.DataPoint
	if *v == nil {
		cv = []*types.DataPoint{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.DataPoint
		if err := awsAwsjson11_deserializeDocumentDataPoint(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionKeyDescription(v **types.DimensionKeyDescription, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.DimensionKeyDescription
	if *v == nil {
		sv = &types.DimensionKeyDescription{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, decoder); err != nil {
				return err
			}

		case "Partitions":
			if err := awsAwsjson11_deserializeDocumentMetricValuesList(&sv.Partitions, decoder); err != nil {
				return err
			}

		case "Total":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected Double to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Total = &f64
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionKeyDescriptionList(v *[]*types.DimensionKeyDescription, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.DimensionKeyDescription
	if *v == nil {
		cv = []*types.DimensionKeyDescription{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.DimensionKeyDescription
		if err := awsAwsjson11_deserializeDocumentDimensionKeyDescription(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionMap(v *map[string]*string, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var mv map[string]*string
	if *v == nil {
		mv = map[string]*string{}
	} else {
		mv = *v
	}

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		key, ok := token.(string)
		if !ok {
			return fmt.Errorf("expected map-key of type string, found type %T", token)
		}

		var parsedVal *string
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", val)
			}
			parsedVal = &jtv
		}
		mv[key] = parsedVal

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = mv
	return nil
}

func awsAwsjson11_deserializeDocumentInternalServiceError(v **types.InternalServiceError, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.InternalServiceError
	if *v == nil {
		sv = &types.InternalServiceError{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentInvalidArgumentException(v **types.InvalidArgumentException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.InvalidArgumentException
	if *v == nil {
		sv = &types.InvalidArgumentException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentMetricKeyDataPoints(v **types.MetricKeyDataPoints, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.MetricKeyDataPoints
	if *v == nil {
		sv = &types.MetricKeyDataPoints{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "DataPoints":
			if err := awsAwsjson11_deserializeDocumentDataPointsList(&sv.DataPoints, decoder); err != nil {
				return err
			}

		case "Key":
			if err := awsAwsjson11_deserializeDocumentResponseResourceMetricKey(&sv.Key, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentMetricKeyDataPointsList(v *[]*types.MetricKeyDataPoints, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.MetricKeyDataPoints
	if *v == nil {
		cv = []*types.MetricKeyDataPoints{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.MetricKeyDataPoints
		if err := awsAwsjson11_deserializeDocumentMetricKeyDataPoints(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentMetricValuesList(v *[]*float64, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*float64
	if *v == nil {
		cv = []*float64{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *float64
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(json.Number)
			if !ok {
				return fmt.Errorf("expected Double to be json.Number, got %T instead", val)
			}
			f64, err := jtv.Float64()
			if err != nil {
				return err
			}
			col = &f64
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentNotAuthorizedException(v **types.NotAuthorizedException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.NotAuthorizedException
	if *v == nil {
		sv = &types.NotAuthorizedException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentResponsePartitionKey(v **types.ResponsePartitionKey, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.ResponsePartitionKey
	if *v == nil {
		sv = &types.ResponsePartitionKey{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentResponsePartitionKeyList(v *[]*types.ResponsePartitionKey, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.ResponsePartitionKey
	if *v == nil {
		cv = []*types.ResponsePartitionKey{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.ResponsePartitionKey
		if err := awsAwsjson11_deserializeDocumentResponsePartitionKey(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentResponseResourceMetricKey(v **types.ResponseResourceMetricKey, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.ResponseResourceMetricKey
	if *v == nil {
		sv = &types.ResponseResourceMetricKey{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, decoder); err != nil {
				return err
			}

		case "Metric":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Metric = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentDescribeDimensionKeysOutput(v **DescribeDimensionKeysOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *DescribeDimensionKeysOutput
	if *v == nil {
		sv = &DescribeDimensionKeysOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "AlignedEndTime":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedEndTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "AlignedStartTime":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedStartTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Keys":
			if err := awsAwsjson11_deserializeDocumentDimensionKeyDescriptionList(&sv.Keys, decoder); err != nil {
				return err
			}

		case "NextToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.NextToken = &jtv
			}

		case "PartitionKeys":
			if err := awsAwsjson11_deserializeDocumentResponsePartitionKeyList(&sv.PartitionKeys, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentGetResourceMetricsOutput(v **GetResourceMetricsOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *GetResourceMetricsOutput
	if *v == nil {
		sv = &GetResourceMetricsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "AlignedEndTime":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedEndTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "AlignedStartTime":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", val)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedStartTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Identifier":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Identifier = &jtv
			}

		case "MetricList":
			if err := awsAwsjson11_deserializeDocumentMetricKeyDataPointsList(&sv.MetricList, decoder); err != nil {
				return err
			}

		case "NextToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.NextToken = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}