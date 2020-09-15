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

// [VPC only] Adds the specified egress rules to a security group for use with a
// VPC. An outbound rule permits instances to send traffic to the specified IPv4 or
// IPv6 CIDR address ranges, or to the instances associated with the specified
// destination security groups. You specify a protocol for each rule (for example,
// TCP). For the TCP and UDP protocols, you must also specify the destination port
// or port range. For the ICMP protocol, you must also specify the ICMP type and
// code. You can use -1 for the type or code to mean all types or all codes. Rule
// changes are propagated to affected instances as quickly as possible. However, a
// small delay might occur. For more information about VPC security group limits,
// see Amazon VPC Limits
// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
func (c *Client) AuthorizeSecurityGroupEgress(ctx context.Context, params *AuthorizeSecurityGroupEgressInput, optFns ...func(*Options)) (*AuthorizeSecurityGroupEgressOutput, error) {
	stack := middleware.NewStack("AuthorizeSecurityGroupEgress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAuthorizeSecurityGroupEgressMiddlewares(stack)
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
	addOpAuthorizeSecurityGroupEgressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeSecurityGroupEgress(options.Region), middleware.Before)

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
			OperationName: "AuthorizeSecurityGroupEgress",
			Err:           err,
		}
	}
	out := result.(*AuthorizeSecurityGroupEgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeSecurityGroupEgressInput struct {
	// Not supported. Use a set of IP permissions to specify the port.
	ToPort *int32
	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupName *string
	// Not supported. Use a set of IP permissions to specify the port.
	FromPort *int32
	// The ID of the security group.
	GroupId *string
	// Not supported. Use a set of IP permissions to specify the protocol name or
	// number.
	IpProtocol *string
	// The sets of IP permissions. You can't specify a destination security group and a
	// CIDR IP address range in the same set of permissions.
	IpPermissions []*types.IpPermission
	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupOwnerId *string
	// Not supported. Use a set of IP permissions to specify the CIDR.
	CidrIp *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AuthorizeSecurityGroupEgressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAuthorizeSecurityGroupEgressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAuthorizeSecurityGroupEgress{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAuthorizeSecurityGroupEgress{}, middleware.After)
}

func newServiceMetadataMiddleware_opAuthorizeSecurityGroupEgress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AuthorizeSecurityGroupEgress",
	}
}