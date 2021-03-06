// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a static route associated with a VPN connection between an existing
// virtual private gateway and a VPN customer gateway. The static route allows
// traffic to be routed from the virtual private gateway to the VPN customer
// gateway. For more information, see AWS Site-to-Site VPN
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the AWS
// Site-to-Site VPN User Guide.
func (c *Client) CreateVpnConnectionRoute(ctx context.Context, params *CreateVpnConnectionRouteInput, optFns ...func(*Options)) (*CreateVpnConnectionRouteOutput, error) {
	if params == nil {
		params = &CreateVpnConnectionRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpnConnectionRoute", params, optFns, addOperationCreateVpnConnectionRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpnConnectionRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateVpnConnectionRoute.
type CreateVpnConnectionRouteInput struct {

	// The CIDR block associated with the local subnet of the customer network.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the VPN connection.
	//
	// This member is required.
	VpnConnectionId *string
}

type CreateVpnConnectionRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVpnConnectionRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpnConnectionRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpnConnectionRoute{}, middleware.After)
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
	addOpCreateVpnConnectionRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpnConnectionRoute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateVpnConnectionRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpnConnectionRoute",
	}
}
