package entities

type Tentant struct {
	ID         int32  // pk |
	Identifier string //    | idx
	Name       string
	IsActive   bool
}
