package utils

import(
	"github.com/google/uuid"
)

type UUIDGenerator struct{}

func (g *UUIDGenerator) Genid() (string, error) {
	//UUIDを生成する
	uuid_obj, err := uuid.NewRandom()

	//エラー処理
	if err != nil {
		return "", err
	}

	return uuid_obj.String(),nil
}
