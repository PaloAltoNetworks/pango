package interface_management

type Specifier func(*Entry) (any, error)

type Normalizer interface {
	Normalize() ([]*Entry, error)
}
