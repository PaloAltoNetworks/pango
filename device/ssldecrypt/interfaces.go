package ssldecrypt

type Specifier func(*Config) (any, error)

type Normalizer interface {
	Normalize() ([]*Config, error)
}
