// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DecodeAuthorizationMessageInput struct {
	_ struct{} `type:"structure"`

	// The encoded message that was returned with the response.
	//
	// EncodedMessage is a required field
	EncodedMessage *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DecodeAuthorizationMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecodeAuthorizationMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecodeAuthorizationMessageInput"}

	if s.EncodedMessage == nil {
		invalidParams.Add(aws.NewErrParamRequired("EncodedMessage"))
	}
	if s.EncodedMessage != nil && len(*s.EncodedMessage) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EncodedMessage", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A document that contains additional information about the authorization status
// of a request from an encoded message that is returned in response to an AWS
// request.
type DecodeAuthorizationMessageOutput struct {
	_ struct{} `type:"structure"`

	// An XML document that contains the decoded message.
	DecodedMessage *string `type:"string"`
}

// String returns the string representation
func (s DecodeAuthorizationMessageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDecodeAuthorizationMessage = "DecodeAuthorizationMessage"

// DecodeAuthorizationMessageRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Decodes additional information about the authorization status of a request
// from an encoded message returned in response to an AWS request.
//
// For example, if a user is not authorized to perform an operation that he
// or she has requested, the request returns a Client.UnauthorizedOperation
// response (an HTTP 403 response). Some AWS operations additionally return
// an encoded message that can provide details about this authorization failure.
//
// Only certain AWS operations return an encoded authorization message. The
// documentation for an individual operation indicates whether that operation
// returns an encoded message in addition to returning an HTTP code.
//
// The message is encoded because the details of the authorization status can
// constitute privileged information that the user who requested the operation
// should not see. To decode an authorization status message, a user must be
// granted permissions via an IAM policy to request the DecodeAuthorizationMessage
// (sts:DecodeAuthorizationMessage) action.
//
// The decoded message includes the following type of information:
//
//    * Whether the request was denied due to an explicit deny or due to the
//    absence of an explicit allow. For more information, see Determining Whether
//    a Request is Allowed or Denied (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-denyallow)
//    in the IAM User Guide.
//
//    * The principal who made the request.
//
//    * The requested action.
//
//    * The requested resource.
//
//    * The values of condition keys in the context of the user's request.
//
//    // Example sending a request using DecodeAuthorizationMessageRequest.
//    req := client.DecodeAuthorizationMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/DecodeAuthorizationMessage
func (c *Client) DecodeAuthorizationMessageRequest(input *DecodeAuthorizationMessageInput) DecodeAuthorizationMessageRequest {
	op := &aws.Operation{
		Name:       opDecodeAuthorizationMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecodeAuthorizationMessageInput{}
	}

	req := c.newRequest(op, input, &DecodeAuthorizationMessageOutput{})

	return DecodeAuthorizationMessageRequest{Request: req, Input: input, Copy: c.DecodeAuthorizationMessageRequest}
}

// DecodeAuthorizationMessageRequest is the request type for the
// DecodeAuthorizationMessage API operation.
type DecodeAuthorizationMessageRequest struct {
	*aws.Request
	Input *DecodeAuthorizationMessageInput
	Copy  func(*DecodeAuthorizationMessageInput) DecodeAuthorizationMessageRequest
}

// Send marshals and sends the DecodeAuthorizationMessage API request.
func (r DecodeAuthorizationMessageRequest) Send(ctx context.Context) (*DecodeAuthorizationMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecodeAuthorizationMessageResponse{
		DecodeAuthorizationMessageOutput: r.Request.Data.(*DecodeAuthorizationMessageOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecodeAuthorizationMessageResponse is the response type for the
// DecodeAuthorizationMessage API operation.
type DecodeAuthorizationMessageResponse struct {
	*DecodeAuthorizationMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DecodeAuthorizationMessage request.
func (r *DecodeAuthorizationMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}