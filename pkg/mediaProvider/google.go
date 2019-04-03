package mediaProvider

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/media"
)

type googleProvider struct{}

// NewGoogleProvider constructor
func NewGoogleProvider() media.ProviderRepository {

	return &googleProvider{}
}

func (pp *googleProvider) CreateLive() {
	fmt.Println("Create Live")
}

func (pp *googleProvider) GetLive() {
	fmt.Println("Get Live")
}

func (pp *googleProvider) GetVOD() {
	fmt.Println("Get VOD")
}
