package injector

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kyawmyintthein/twirp-poc/app/config"
	"github.com/kyawmyintthein/twirp-poc/app/domain/repository"
)

func ProvideAWSSession(globalCfg *config.GlobalCfg) (*session.Session, error) {
	option := session.Options{
		Config: aws.Config{
			Region: &globalCfg.AWS.Region,
		},
		SharedConfigState: session.SharedConfigEnable,
	}
	if globalCfg.AWS.Profile != "" {
		option.Profile = globalCfg.AWS.Profile
	}
	return session.NewSessionWithOptions(option)
}

func ProvideDynamodb(globalCfg *config.GlobalCfg, sess *session.Session) repository.Dynamodb {
	return dynamodb.New(sess)
}
