package security

type Specifier func(Entry) (any, error)

type Normalizer interface {
	Normalize() ([]Entry, error)
}
