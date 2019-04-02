package media

import (
	"fmt"

	"github.com/Footters/hex-footters/pkg/content"
)

type pixellotProvider struct{}

// NewPixellotProvider constructor
func NewPixellotProvider() content.MediaProvider {

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
