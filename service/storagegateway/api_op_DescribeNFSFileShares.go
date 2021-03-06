// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a description for one or more Network File System (NFS) file shares from a
// file gateway. This operation is only supported for file gateways.
func (c *Client) DescribeNFSFileShares(ctx context.Context, params *DescribeNFSFileSharesInput, optFns ...func(*Options)) (*DescribeNFSFileSharesOutput, error) {
	if params == nil {
		params = &DescribeNFSFileSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNFSFileShares", params, optFns, addOperationDescribeNFSFileSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNFSFileSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeNFSFileSharesInput
type DescribeNFSFileSharesInput struct {

	// An array containing the Amazon Resource Name (ARN) of each file share to be
	// described.
	//
	// This member is required.
	FileShareARNList []*string
}

// DescribeNFSFileSharesOutput
type DescribeNFSFileSharesOutput struct {

	// An array containing a description for each requested file share.
	NFSFileShareInfoList []*types.NFSFileShareInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNFSFileSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNFSFileShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNFSFileShares{}, middleware.After)
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
	addOpDescribeNFSFileSharesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNFSFileShares(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeNFSFileShares(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeNFSFileShares",
	}
}
