package aws

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type mockSecretsManagerAPI struct {
	secretsmanageriface.SecretsManagerAPI
}

func (m mockSecretsManagerAPI) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	if *input.SecretId == "AWS-SECRET-JSON" {
		return &secretsmanager.GetSecretValueOutput{
			Name:         aws.String(*input.SecretId),
			SecretString: new(string),
		}, nil
	} else {
		return nil, &secretsmanager.ResourceNotFoundException{}
	}

}

func Test_readSecret(t *testing.T) {
	type args struct {
		secret string
		svc    mockSecretsManagerAPI
	}
	tests := []struct {
		name    string
		args    args
		want    *secretsmanager.GetSecretValueOutput
		wantErr bool
	}{
		{
			name: "Run Ok",
			args: args{
				secret: "AWS-SECRET-JSON",
				svc:    mockSecretsManagerAPI{},
			},
			want: &secretsmanager.GetSecretValueOutput{
				Name:         aws.String("AWS-SECRET-JSON"),
				SecretString: new(string),
			},
			wantErr: false,
		},
		{
			name: "Run Error",
			args: args{
				secret: "AWS-SECRET-NON-EXIST",
				svc:    mockSecretsManagerAPI{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSecret(tt.args.secret, tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
