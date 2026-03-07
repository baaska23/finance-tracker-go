package analytics

type Repository interface {
	GetById(id string) (string, error)
}
