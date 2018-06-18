package builtin

import (
	"testing"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
)


type fakeSSM struct {
	ssmiface.SSMAPI
	payload ssm.GetParameterOutput
	err error
}

func (fs fakeSSM) GetParameter(in *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	// Only need to return mocked response output
	return &fs.payload, nil
}

func TestSsmParameter(t *testing.T) {
	expected_key := "/test"
	expected_value := "test value"

	ssmapi := fakeSSM{
			payload:ssm.GetParameterOutput{
				Parameter: &ssm.Parameter{
					Name: aws.String(expected_key),
					Value: aws.String(expected_value),
				},
			},
	}
	resp, err := GetSsmParameter(ssmapi, expected_key)

	if err != nil {
		t.Fatalf("%d, unexpected error", err)
	}
	if resp != expected_value {
		t.Fatalf("%v != %v", resp, expected_value)
	}

}
