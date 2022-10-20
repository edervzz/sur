package entities

type Url struct {
	TenantId  int32  // pk
	Code      string // pk
	IsActive  bool
	Url       string
	Hash      uint32
	Count     int32
	IPAddress string
	CreateAt  string
	CreateBy  string
}
