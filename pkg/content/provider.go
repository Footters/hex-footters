package content

// MediaProvider of media
type MediaProvider interface {
	CreateLive()
	GetLive()
	GetVOD()
}
