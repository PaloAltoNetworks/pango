package dynamicupdates

type Specifier func(*Config) (any, error)

type Normalizer interface {
	Normalize() ([]*Config, error)
}
