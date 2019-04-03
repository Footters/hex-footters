package ibm

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/media"
)

type ibmProvider struct{}

// NewIBMProvider constructor
func NewIBMProvider() media.ProviderRepository {

	return &ibmProvider{}
}

func (ip *ibmProvider) CreateEvent() {
	fmt.Println("Create Live with IBM")
}

func (ip *ibmProvider) GetLive() {
	fmt.Println("Get Live with IBM")
}

func (ip *ibmProvider) GetVOD() {
	fmt.Println("Get VOD with IBM")
}
