// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about a proposal.
func (c *Client) GetProposal(ctx context.Context, params *GetProposalInput, optFns ...func(*Options)) (*GetProposalOutput, error) {
	if params == nil {
		params = &GetProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProposal", params, optFns, addOperationGetProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProposalInput struct {

	// The unique identifier of the network for which the proposal is made.
	//
	// This member is required.
	NetworkId *string

	// The unique identifier of the proposal.
	//
	// This member is required.
	ProposalId *string
}

type GetProposalOutput struct {

	// Information about a proposal.
	Proposal *types.Proposal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProposal{}, middleware.After)
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
	addOpGetProposalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetProposal(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetProposal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "GetProposal",
	}
}
