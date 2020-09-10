// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the firewall port states for a specific Amazon Lightsail instance, the
// IP addresses allowed to connect to the instance through the ports, and the
// protocol.
func (c *Client) GetInstancePortStates(ctx context.Context, params *GetInstancePortStatesInput, optFns ...func(*Options)) (*GetInstancePortStatesOutput, error) {
	stack := middleware.NewStack("GetInstancePortStates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInstancePortStatesMiddlewares(stack)
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
	addOpGetInstancePortStatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstancePortStates(options.Region), middleware.Before)

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
			OperationName: "GetInstancePortStates",
			Err:           err,
		}
	}
	out := result.(*GetInstancePortStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstancePortStatesInput struct {
	// The name of the instance for which to return firewall port states.
	InstanceName *string
}

type GetInstancePortStatesOutput struct {
	// An array of objects that describe the firewall port states for the specified
	// instance.
	PortStates []*types.InstancePortState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInstancePortStatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstancePortStates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstancePortStates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInstancePortStates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstancePortStates",
	}
}