package checkers

func DefaultParams() Params {
	return Params{}
}

func (p Params) Validate() error {
	return nil
}
