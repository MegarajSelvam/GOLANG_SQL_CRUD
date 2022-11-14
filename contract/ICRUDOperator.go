package contract

type any = interface{}

type ICRUDOperator interface {
	Create() (any, bool)
	Retrieve() (any, bool)
	Update() (any, bool)
	Delete() (any, bool)
}
