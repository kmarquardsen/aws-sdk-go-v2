// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the customer gateway or the target gateway of an AWS Site-to-Site VPN
// connection. To modify the target gateway, the following migration options are
// available:
//
//     * An existing virtual private gateway to a new virtual private
// gateway
//
//     * An existing virtual private gateway to a transit gateway
//
//     *
// An existing transit gateway to a new transit gateway
//
//     * An existing transit
// gateway to a virtual private gateway
//
// Before you perform the migration to the
// new gateway, you must configure the new gateway. Use CreateVpnGateway () to
// create a virtual private gateway, or CreateTransitGateway () to create a transit
// gateway. This step is required when you migrate from a virtual private gateway
// with static routes to a transit gateway. You must delete the static routes
// before you migrate to the new gateway.  <p>Keep a copy of the static route
// before you delete it. You will need to add back these routes to the transit
// gateway after the VPN connection migration is complete.</p> <p>After you migrate
// to the new gateway, you might need to modify your VPC route table. Use
// <a>CreateRoute</a> and <a>DeleteRoute</a> to make the changes described in <a
// href="https://docs.aws.amazon.com/vpn/latest/s2svpn/modify-vpn-target.html#step-update-routing">VPN
// Gateway Target Modification Required VPC Route Table Updates</a> in the <i>AWS
// Site-to-Site VPN User Guide</i>.</p> <p> When the new gateway is a transit
// gateway, modify the transit gateway route table to allow traffic between the VPC
// and the AWS Site-to-Site VPN connection. Use <a>CreateTransitGatewayRoute</a> to
// add the routes.</p> <p> If you deleted VPN static routes, you must add the
// static routes to the transit gateway route table.</p> <p>After you perform this
// operation, the AWS VPN endpoint's IP addresses on the AWS side and the tunnel
// options remain intact. Your AWS Site-to-Site VPN connection will be temporarily
// unavailable for a brief period while we provision the new endpoints.</p>
func (c *Client) ModifyVpnConnection(ctx context.Context, params *ModifyVpnConnectionInput, optFns ...func(*Options)) (*ModifyVpnConnectionOutput, error) {
	stack := middleware.NewStack("ModifyVpnConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyVpnConnectionMiddlewares(stack)
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
	addOpModifyVpnConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnConnection(options.Region), middleware.Before)

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
			OperationName: "ModifyVpnConnection",
			Err:           err,
		}
	}
	out := result.(*ModifyVpnConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnConnectionInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the customer gateway at your end of the VPN connection.
	CustomerGatewayId *string
	// The ID of the VPN connection.
	VpnConnectionId *string
	// The ID of the transit gateway.
	TransitGatewayId *string
	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	VpnGatewayId *string
}

type ModifyVpnConnectionOutput struct {
	// Describes a VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyVpnConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnConnection{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyVpnConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpnConnection",
	}
}