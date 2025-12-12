package port

// IDGeneratorPort はID生成のためのポートインターフェースです。
type UUIDGeneratorPort interface {
	Genid() (string, error)
}
