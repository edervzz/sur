package messages

type CreateUrlRequest struct {
	TenantId int32
	Code     string
	IsActive bool
	Url      string
	Hash     uint32
}
