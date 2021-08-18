package adapters

import "github.com/jjwoz/assignme/internal/tickets/ports"

type S3Avatar struct {
	uc       ports.AvatarService
	fileName string
	data     byte
}

// NewS3Service accepts an interface and returns a concrete type. within the returned concrete type (S3Avatar)
// other interfaces are embedded allowing for implicitly adhering to an interface allowing for ONLY implementing what is needed
// not javafying the code base
func NewS3Service(service ports.AvatarService) S3Avatar {
	return S3Avatar{
		uc: service,
	}
}

// TODO grab s3 avatar from here and populate image lock down and make not public
func (sa S3Avatar) RetrieveAvatar(path string) (byte, error) {
	//region := "us-east-1"
	//bucketName := "assign-me-avatar"
	return byte(1), nil
}
