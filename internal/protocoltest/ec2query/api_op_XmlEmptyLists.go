// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) XmlEmptyLists(ctx context.Context, params *XmlEmptyListsInput, optFns ...func(*Options)) (*XmlEmptyListsOutput, error) {
	if params == nil {
		params = &XmlEmptyListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlEmptyLists", params, optFns, addOperationXmlEmptyListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlEmptyListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEmptyListsInput struct {
}

type XmlEmptyListsOutput struct {
	BooleanList []*bool

	EnumList []types.FooEnum

	FlattenedList []*string

	FlattenedList2 []*string

	IntegerList []*int32

	// A list of lists of strings.
	NestedStringList [][]*string

	RenamedListMembers []*string

	StringList []*string

	StringSet []*string

	StructureList []*types.StructureListMember

	TimestampList []*time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlEmptyListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpXmlEmptyLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpXmlEmptyLists{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEmptyLists(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opXmlEmptyLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlEmptyLists",
	}
}
