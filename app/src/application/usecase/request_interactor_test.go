package usecase_test

import (
	"errors"
	"testing"

	"app/application/usecase"
	"app/domain"
	"app/messages"
)

// MockFriendRequestPort はリポジトリPortのモック
type MockFriendRequestPort struct {
	ExistsResult bool
	ExistsError  error
	CreateError  error
}

func (m *MockFriendRequestPort) Create(req *domain.FriendRequest) error {
	return m.CreateError
}

func (m *MockFriendRequestPort) Exists(senderID, targetID string) (bool, error) {
	return m.ExistsResult, m.ExistsError
}

// MockUUIDGeneratorPort はID生成Portのモック
type MockUUIDGeneratorPort struct {
	GeneratedID string
	GenerateError error
}

func (m *MockUUIDGeneratorPort) Genid() (string, error) {
	return m.GeneratedID, m.GenerateError
}

// TestExecute_Success テスト: 正常にリクエストが送信され、Repo.Createが呼ばれるか
func TestExecute_Success(t *testing.T) {
	// 依存のモック設定
	interactor := createTestInteractor(false, true, true)

	input := &usecase.SendFriendRequestInput{
		SenderID: "A",
		TargetID: "B",
	}

	err := interactor.Execute(*input)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
}

// TestExecute_AlreadySent テスト: 既にリクエスト済みの場合にErrAlreadySentを返すか
func TestExecute_AlreadySent(t *testing.T) {
	// 依存のモック設定
	interactor := createTestInteractor(true, true, true)

	input := usecase.SendFriendRequestInput{SenderID: "A", TargetID: "B"}

	err := interactor.Execute(input)

	if !errors.Is(err, messages.ErrAlreadySent) {
		t.Errorf("Expected ErrAlreadySent, got %v", err)
	}
}

// TestExecute_SelfRequest テスト: 自分自身へのリクエストを拒否するか
func TestExecute_SelfRequest(t *testing.T) {
	// 依存のモック設定
	interactor := createTestInteractor(false, true, true)
	// 送信者とターゲットが同一
	input := usecase.SendFriendRequestInput{SenderID: "A", TargetID: "A"}

	err := interactor.Execute(input)

	if err == nil || !errors.Is(err, messages.ErrSelfRequest) {
		t.Errorf("Expected 'cannot send...' error, got %v", err)
	}
}

// TestExecute_NonExistentUser テスト: 送信者またはターゲットが存在しない場合にエラーを返すか
func TestExecute_NonExistentUser(t *testing.T) {
	interactor := createTestInteractor(false, false, true)

	input := usecase.SendFriendRequestInput{SenderID: "A", TargetID: "B"}

	err := interactor.Execute(input)

	if err == nil || !errors.Is(err, messages.ErrSenderNotFound) {
		t.Errorf("Expected '送信者が存在しません' error, got %v", err)
	}
}

// MockUserPort はUserPortのモック
type MockUserPort struct {
	ExistsByIDResults map[string]bool
}

// ExistsByID はユーザーIDの存在を確認するモック実装
func (m *MockUserPort) ExistsByID(userID string) (bool, error) {
	exists, ok := m.ExistsByIDResults[userID]
	if !ok {
		return false, nil
	}
	return exists, nil
}

// createTestInteractor は、テストしたい状況を bool で指定して Interactor を生成する
func createTestInteractor(isDuplicate bool, senderExists bool, targetExists bool) *usecase.SendFriendRequestInteractor {
	// 1. フレンド申請リポジトリのモック
	mockRepo := &MockFriendRequestPort{
		ExistsResult: isDuplicate, // 重複しているかどうか
	}

	// 2. ID生成のモック（常に固定値を返す設定）
	mockGenid := &MockUUIDGeneratorPort{
		GeneratedID: "test-uuid-001",
	}

	// 3. ユーザーリポジトリのモック
	mockUserRepo := &MockUserPort{
		ExistsByIDResults: map[string]bool{
			"A": senderExists, // 送信者が存在するか
			"B": targetExists, // 相手が存在するか
		},
	}

	return &usecase.SendFriendRequestInteractor{
		Repo:     mockRepo,
		Genid:    mockGenid,
		UserRepo: mockUserRepo,
	}
}
