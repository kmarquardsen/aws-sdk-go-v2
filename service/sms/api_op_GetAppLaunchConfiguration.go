// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the application launch configuration associated with an application.
func (c *Client) GetAppLaunchConfiguration(ctx context.Context, params *GetAppLaunchConfigurationInput, optFns ...func(*Options)) (*GetAppLaunchConfigurationOutput, error) {
	stack := middleware.NewStack("GetAppLaunchConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAppLaunchConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAppLaunchConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetAppLaunchConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetAppLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppLaunchConfigurationInput struct {
	// ID of the application launch configuration.
	AppId *string
}

type GetAppLaunchConfigurationOutput struct {
	// List of launch configurations for server groups in this application.
	ServerGroupLaunchConfigurations []*types.ServerGroupLaunchConfiguration
	// Name of the service role in the customer's account that Amazon CloudFormation
	// uses to launch the application.
	RoleName *string
	// ID of the application associated with the launch configuration.
	AppId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAppLaunchConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAppLaunchConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAppLaunchConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAppLaunchConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetAppLaunchConfiguration",
	}
}