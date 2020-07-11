package namespace

// Namer returns the names of objects returned from PAN-OS.
type Namer interface {
	Names() []string
}

// MovePather returns an xpath given the name.
type MovePather func(string) []string

// MoveLister returns a list of current rules.
type MoveLister func() ([]string, error)
