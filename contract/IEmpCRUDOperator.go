package contract

type IEmpCRUDOperator interface {
	ICRUDOperator
	RetrieveByUserName() (any, bool)
	RetrieveByEmpId() (any, bool)
}
