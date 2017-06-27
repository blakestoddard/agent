package agent

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

type EC2MetaData struct {
	sess *session.Session
}

func (e EC2MetaData) Get() (map[string]string, error) {
	if e.sess == nil {
		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}
		e.sess = sess
	}

	metaData := make(map[string]string)
	ec2metadataClient := ec2metadata.New(e.sess)

	instanceId, err := ec2metadataClient.GetMetadata("instance-id")
	if err != nil {
		return metaData, err
	}

	metaData["aws:instance-id"] = string(instanceId)

	instanceType, err := ec2metadataClient.GetMetadata("instance-type")
	if err != nil {
		return metaData, err
	}
	metaData["aws:instance-type"] = string(instanceType)

	amiId, err := ec2metadataClient.GetMetadata("ami-id")
	if err != nil {
		return metaData, err
	}
	metaData["aws:ami-id"] = string(amiId)

	return metaData, nil
}
