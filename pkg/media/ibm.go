package media

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/content"
)

type ibmProvider struct{}

// NewIBMProvider constructor
func NewIBMProvider() content.MediaProvider {

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
