package content

// Service Content inteface
type Service interface {
	CreateContent(content *Content) error
	FindContentByID(id uint) (*Content, error)
	FindAllContents() ([]Content, error)
	SetToLive(content *Content) error
}

type contentService struct {
	repo  Repository
	media MediaProvider
}

//NewService Constructor
func NewService(repo Repository, media MediaProvider) Service {
	return &contentService{
		repo:  repo,
		media: media,
	}
}

func (c *contentService) CreateContent(content *Content) error {
	content.Status = "pending"
	return c.repo.Create(content)
}

func (c *contentService) FindContentByID(id uint) (*Content, error) {

	return c.repo.FindByID(id)
}

func (c *contentService) FindAllContents() ([]Content, error) {

	return c.repo.FindAll()
}
func (c *contentService) SetToLive(content *Content) error {
	content.Status = "live"
	c.media.CreateLive()
	return c.repo.Update(content)
}
