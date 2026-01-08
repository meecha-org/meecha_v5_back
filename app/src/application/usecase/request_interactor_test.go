package usecase_test

import (
	"errors"
	"testing"

	"app/application/usecase"
	"app/domain"
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
	mockRepo := &MockFriendRequestPort{
		ExistsResult: false, // 存在しない
		CreateError:  nil,
	}
	mockGenid := &MockUUIDGeneratorPort{
		GeneratedID: "test-uuid-001",
	}

	interactor := usecase.SendFriendRequestInteractor{
		Repo:  mockRepo,  // モック注入
		Genid: mockGenid, // モック注入
	}
    
	input := usecase.SendFriendRequestInput{
		SenderID: "A",
		TargetID: "B",
	}

	err := interactor.Execute(input)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	// モックが期待通りに動作したことを確認するアサーションを追加するのが理想的
}

// TestExecute_AlreadySent テスト: 既にリクエスト済みの場合にErrAlreadySentを返すか
func TestExecute_AlreadySent(t *testing.T) {
	// 依存のモック設定
	mockRepo := &MockFriendRequestPort{
		ExistsResult: true, // 既に存在する
	}
	mockGenid := &MockUUIDGeneratorPort{
		GeneratedID: "test-uuid-001",
	}

	interactor := usecase.SendFriendRequestInteractor{
		Repo:  mockRepo,
		Genid: mockGenid,
	}
	input := usecase.SendFriendRequestInput{SenderID: "A", TargetID: "B"}

	err := interactor.Execute(input)

	if !errors.Is(err, usecase.ErrAlreadySent) {
		t.Errorf("Expected ErrAlreadySent, got %v", err)
	}
}

// TestExecute_SelfRequest テスト: 自分自身へのリクエストを拒否するか
func TestExecute_SelfRequest(t *testing.T) {
	interactor := usecase.SendFriendRequestInteractor{}
	input := usecase.SendFriendRequestInput{SenderID: "A", TargetID: "A"}

	err := interactor.Execute(input)

	if err == nil || err.Error() != "cannot send friend request to oneself" {
		t.Errorf("Expected 'cannot send...' error, got %v", err)
	}
}
