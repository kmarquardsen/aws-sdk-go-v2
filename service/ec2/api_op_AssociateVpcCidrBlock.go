// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a CIDR block with your VPC. You can associate a secondary IPv4 CIDR
// block, an Amazon-provided IPv6 CIDR block, or an IPv6 CIDR block from an IPv6
// address pool that you provisioned through bring your own IP addresses (BYOIP
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)). The IPv6
// CIDR block size is fixed at /56. You must specify one of the following in the
// request: an IPv4 CIDR block, an IPv6 pool, or an Amazon-provided IPv6 CIDR
// block. For more information about associating CIDR blocks with your VPC and
// applicable restrictions, see VPC and Subnet Sizing
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#VPC_Sizing)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) AssociateVpcCidrBlock(ctx context.Context, params *AssociateVpcCidrBlockInput, optFns ...func(*Options)) (*AssociateVpcCidrBlockOutput, error) {
	if params == nil {
		params = &AssociateVpcCidrBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateVpcCidrBlock", params, optFns, addOperationAssociateVpcCidrBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateVpcCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateVpcCidrBlockInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the
	// VPC. You cannot specify the range of IPv6 addresses, or the size of the CIDR
	// block.
	AmazonProvidedIpv6CidrBlock *bool

	// An IPv4 CIDR block to associate with the VPC.
	CidrBlock *string

	// An IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in
	// the request. To let Amazon choose the IPv6 CIDR block for you, omit this
	// parameter.
	Ipv6CidrBlock *string

	// The name of the location from which we advertise the IPV6 CIDR block. Use this
	// parameter to limit the CiDR block to this location. You must set
	// AmazonProvidedIpv6CidrBlock to true to use this parameter. You can have one IPv6
	// CIDR block association per network border group.
	Ipv6CidrBlockNetworkBorderGroup *string

	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	Ipv6Pool *string
}

type AssociateVpcCidrBlockOutput struct {

	// Information about the IPv4 CIDR block association.
	CidrBlockAssociation *types.VpcCidrBlockAssociation

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *types.VpcIpv6CidrBlockAssociation

	// The ID of the VPC.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateVpcCidrBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateVpcCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateVpcCidrBlock{}, middleware.After)
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
	addOpAssociateVpcCidrBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateVpcCidrBlock(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateVpcCidrBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateVpcCidrBlock",
	}
}
