// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stop an Automation that is currently running.
func (c *Client) StopAutomationExecution(ctx context.Context, params *StopAutomationExecutionInput, optFns ...func(*Options)) (*StopAutomationExecutionOutput, error) {
	if params == nil {
		params = &StopAutomationExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopAutomationExecution", params, optFns, addOperationStopAutomationExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopAutomationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopAutomationExecutionInput struct {

	// The execution ID of the Automation to stop.
	//
	// This member is required.
	AutomationExecutionId *string

	// The stop request type. Valid types include the following: Cancel and Complete.
	// The default type is Cancel.
	Type types.StopType
}

type StopAutomationExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopAutomationExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopAutomationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopAutomationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpStopAutomationExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopAutomationExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopAutomationExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "StopAutomationExecution",
	}
}
