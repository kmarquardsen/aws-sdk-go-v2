// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates up to 100 new attendees for an active Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) BatchCreateAttendee(ctx context.Context, params *BatchCreateAttendeeInput, optFns ...func(*Options)) (*BatchCreateAttendeeOutput, error) {
	if params == nil {
		params = &BatchCreateAttendeeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateAttendee", params, optFns, addOperationBatchCreateAttendeeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateAttendeeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateAttendeeInput struct {

	// The request containing the attendees to create.
	//
	// This member is required.
	Attendees []*types.CreateAttendeeRequestItem

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string
}

type BatchCreateAttendeeOutput struct {

	// The attendee information, including attendees IDs and join tokens.
	Attendees []*types.Attendee

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []*types.CreateAttendeeError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchCreateAttendeeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateAttendee{}, middleware.After)
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
	addOpBatchCreateAttendeeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateAttendee(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchCreateAttendee(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchCreateAttendee",
	}
}
