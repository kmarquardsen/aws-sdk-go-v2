// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the groups to which the thing belongs.
func (c *Client) UpdateThingGroupsForThing(ctx context.Context, params *UpdateThingGroupsForThingInput, optFns ...func(*Options)) (*UpdateThingGroupsForThingOutput, error) {
	if params == nil {
		params = &UpdateThingGroupsForThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThingGroupsForThing", params, optFns, addOperationUpdateThingGroupsForThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThingGroupsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThingGroupsForThingInput struct {

	// Override dynamic thing groups with static thing groups when 10-group limit is
	// reached. If a thing belongs to 10 thing groups, and one or more of those groups
	// are dynamic thing groups, adding a thing to a static group removes the thing
	// from the last dynamic group.
	OverrideDynamicGroups *bool

	// The groups to which the thing will be added.
	ThingGroupsToAdd []*string

	// The groups from which the thing will be removed.
	ThingGroupsToRemove []*string

	// The thing whose group memberships will be updated.
	ThingName *string
}

type UpdateThingGroupsForThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateThingGroupsForThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThingGroupsForThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThingGroupsForThing{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThingGroupsForThing(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateThingGroupsForThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateThingGroupsForThing",
	}
}
