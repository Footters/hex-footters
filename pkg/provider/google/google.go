package google

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/media"
)

type googleProvider struct{}

// NewGoogleProvider constructor
func NewGoogleProvider() media.ProviderRepository {

	return &googleProvider{}
}

func (gp *googleProvider) CreateLive() {
	fmt.Println("Create Live with Google")
}

func (gp *googleProvider) GetLive() {
	fmt.Println("Get Live with Google")
}

func (gp *googleProvider) GetVOD() {
	fmt.Println("Get VOD with Google")
}
