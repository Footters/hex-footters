package media

// ContentRepository interface
type ContentRepository interface {
	Create(content *Content) error
	FindByID(id uint) (*Content, error)
	FindAll() ([]Content, error)
	Update(content *Content) error
}

// ProviderRepository of media
type ProviderRepository interface {
	CreateEvent()
	GetLive()
	GetVOD()
}
