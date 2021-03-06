// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the GET operation uses the acl subresource to return the
// access control list (ACL) of a bucket. To use GET to return the ACL of the
// bucket, you must have READ_ACP access to the bucket. If READ_ACP permission is
// granted to the anonymous user, you can return the ACL of the bucket without
// using an authorization header. Related Resources
func (c *Client) GetBucketAcl(ctx context.Context, params *GetBucketAclInput, optFns ...func(*Options)) (*GetBucketAclOutput, error) {
	if params == nil {
		params = &GetBucketAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketAcl", params, optFns, addOperationGetBucketAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketAclInput struct {

	// Specifies the S3 bucket whose ACL is being requested.
	//
	// This member is required.
	Bucket *string
}

type GetBucketAclOutput struct {

	// A list of grants.
	Grants []*types.Grant

	// Container for the bucket owner's display name and ID.
	Owner *types.Owner

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketAcl{}, middleware.After)
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
	addOpGetBucketAclValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketAcl(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketAcl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketAcl",
	}
}
