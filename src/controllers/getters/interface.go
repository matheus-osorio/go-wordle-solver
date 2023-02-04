package getters

// Interface for getters, made to be used with composition with a higher level controller
type GetterInterface interface {
	GetWords() []string
}
