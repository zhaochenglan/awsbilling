package Session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetSession() *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")}))
	return sess
}
