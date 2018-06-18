package builtin

import (
	"testing"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/aws/aws-sdk-go/service/sms/smsiface"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

type mockedParameter struct {
	smsiface.SMSAPI
	Resp ssm.GetParameterOutput
}

func (m mockedParameter) GetParameter(in *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	// Only need to return mocked response output
	return &m.Resp, nil
}

func TestSsmParameter(t testing.T) {
	cases := []struct {
		Resp     ssm.GetParameterOutput
		Expected ssm.Parameter
	}{
		{
			Resp: ssm.GetParameterOutput{
				Name: aws.String("/test"),
				Type: aws.String("String"),
				Value: aws.String("test_value"),
			},
			Expected: ssm.Parameter{
				Value: aws.String("test_value"),
			},
		},
	}

	for i, c := range cases {
		p := ssm.Parameter{
			Client: mockedReceiveMsgs{Resp: c.Resp},
			URL:    fmt.Sprintf("mockURL_%d", i),
		}
		msgs, err := q.GetMessages(20)
		if err != nil {
			t.Fatalf("%d, unexpected error", err)
		}
		if a, e := len(msgs), len(c.Expected); a != e {
			t.Fatalf("%d, expected %d messages, got %d", i, e, a)
		}
		for j, msg := range msgs {
			if a, e := msg, c.Expected[j]; a != e {
				t.Errorf("%d, expected %v message, got %v", i, e, a)
			}
		}
	}
}