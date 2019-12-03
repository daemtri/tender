package core

type Directive interface {
	Name() string
	Construct(label ...string) error
	RegisterSubDirective(Directive) error
}

type BlockDirective interface {
	Directive
	Configure(def map[string]interface{}) error
}
