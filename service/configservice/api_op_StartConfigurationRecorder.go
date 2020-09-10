// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts recording configurations of the AWS resources you have selected to record
// in your AWS account. You must have created at least one delivery channel to
// successfully start the configuration recorder.
func (c *Client) StartConfigurationRecorder(ctx context.Context, params *StartConfigurationRecorderInput, optFns ...func(*Options)) (*StartConfigurationRecorderOutput, error) {
	stack := middleware.NewStack("StartConfigurationRecorder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartConfigurationRecorderMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpStartConfigurationRecorderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartConfigurationRecorder(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "StartConfigurationRecorder",
			Err:           err,
		}
	}
	out := result.(*StartConfigurationRecorderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the StartConfigurationRecorder () action.
type StartConfigurationRecorderInput struct {
	// The name of the recorder object that records each configuration change made to
	// the resources.
	ConfigurationRecorderName *string
}

type StartConfigurationRecorderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartConfigurationRecorderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartConfigurationRecorder{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartConfigurationRecorder{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartConfigurationRecorder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "StartConfigurationRecorder",
	}
}