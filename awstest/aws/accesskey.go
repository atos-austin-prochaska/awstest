package aws

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IAMListAccessKeysAPI interface {
	ListAccessKeys(ctx context.Context,
		params *iam.ListAccessKeysInput,
		optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)
}

func GetAccessKeys(c context.Context, api IAMListAccessKeysAPI, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	return api.ListAccessKeys(c, input)
}

func ShowKeys() {
	maxItems := flag.Int("m", 10, "The maximum number of access keys to show")
	userName := flag.String("u", "austin.prochaska@atos.net", "The name of the user")
	flag.Parse()

	if *userName == "" {
		fmt.Println("You must supply the name of a user (-u USER)")
		return
	}

	if *maxItems < 0 {
		*maxItems = 10
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := iam.NewFromConfig(cfg)
	input := &iam.ListAccessKeysInput{
		MaxItems: aws.Int32(int32(*maxItems)),
		UserName: userName,
	}

	result, err := GetAccessKeys(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving user access keys:")
		fmt.Println(err)
		return
	}

	for _, key := range result.AccessKeyMetadata {
		fmt.Println("Status for access key " + *key.AccessKeyId + ": " + string(key.Status))
	}
}
