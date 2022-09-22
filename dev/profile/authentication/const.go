package authentication

const (
	UsernameModifierInput       = `%USERINPUT%`
	UsernameModifierInputDomain = `%USERINPUT%@%USERDOMAIN%`
	UsernameModifierDomainInput = `%USERDOMAIN%\%USERINPUT%`
)

const (
	UsernameAttributeDefault = "username"
)

const (
	singular = "authentication profile"
	plural   = "authentication profiles"
)
