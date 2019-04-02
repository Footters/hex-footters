package mediaProvider

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/media"
)

type pixellotProvider struct{}

// NewPixellotProvider constructor
func NewPixellotProvider() media.ProviderRepository {

	return &pixellotProvider{}
}

func (pp *pixellotProvider) CreateLive() {
	fmt.Println("Create Live")
}

func (pp *pixellotProvider) GetLive() {
	fmt.Println("Get Live")
}

func (pp *pixellotProvider) GetVOD() {
	fmt.Println("Get VOD")
}
