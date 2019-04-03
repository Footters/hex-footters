package mediaProvider

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/media"
)

type ibmProvider struct{}

// NewIBMProvider constructor
func NewIBMProvider() media.ProviderRepository {

	return &ibmProvider{}
}

func (ip *ibmProvider) CreateLive() {
	fmt.Println("Create Live")
}

func (ip *ibmProvider) GetLive() {
	fmt.Println("Get Live")
}

func (ip *ibmProvider) GetVOD() {
	fmt.Println("Get VOD")
}
