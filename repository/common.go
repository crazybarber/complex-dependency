package repository

const (
	TypeNamePostgres string = "postgres"
)

type Repository interface {
	SetupSchema() error
	GetSourceSystems() ([]string, error)
}


