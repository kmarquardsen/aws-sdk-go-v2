// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more scheduled actions for the specified Auto Scaling group.
func (c *Client) BatchDeleteScheduledAction(ctx context.Context, params *BatchDeleteScheduledActionInput, optFns ...func(*Options)) (*BatchDeleteScheduledActionOutput, error) {
	if params == nil {
		params = &BatchDeleteScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteScheduledAction", params, optFns, addOperationBatchDeleteScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteScheduledActionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The names of the scheduled actions to delete. The maximum number allowed is 50.
	//
	// This member is required.
	ScheduledActionNames []*string
}

type BatchDeleteScheduledActionOutput struct {

	// The names of the scheduled actions that could not be deleted, including an error
	// message.
	FailedScheduledActions []*types.FailedScheduledUpdateGroupActionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchDeleteScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchDeleteScheduledAction{}, middleware.After)
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
	addOpBatchDeleteScheduledActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteScheduledAction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeleteScheduledAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "BatchDeleteScheduledAction",
	}
}
