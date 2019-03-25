package content

// Service Content inteface
type Service interface {
	CreateContent(content *Content) error
	FindContentByID(id uint) (*Content, error)
	FindAllContents() ([]Content, error)
	SetToLive(content *Content) error
}

type contentService struct {
	repo Repository
}

//NewService Constructor
func NewService(repo Repository) Service {
	return &contentService{
		repo: repo,
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
	return c.repo.Update(content)
}
