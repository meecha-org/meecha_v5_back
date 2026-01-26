package route

import (
	"gorm.io/gorm"
	"app/application/usecase"
	"app/infrastructure/repository"
	"app/infrastructure/utils"
	"app/presentation"
)

// SetupDependencies は全ての依存関係を解決し、FriendHandlerを返します
func SetupDependencies(db *gorm.DB) *presentation.FriendHandler {
    
    // --- Infrastructure層の実装 --- 
    // Port: FriendRequestPort
    friendRequestRepoImpl := &repository.FriendRequestRepositoryImpl{DB: db} 
    // Port: UserPort
    userRepositoryImpl := &repository.UserRepositoryImpl{DB: db}
    // Port: UUIDGeneratorPort
    uuidGenImpl := &utils.UUIDGenerator{} 
    
    // --- Application層 (ユースケース) ---
    // InteractorにPortを注入
    friendRequestInteractor := &usecase.SendFriendRequestInteractor{
        Repo:  friendRequestRepoImpl,  // 💡 Repository実装をPortとして注入
        Genid: uuidGenImpl,            // 💡 ID生成実装をPortとして注入
        UserRepo: userRepositoryImpl,  // 💡 Userリポジトリ実装をPortとして注入
    }
    
    // --- Presentation層 ---
    // Controller/HandlerにInteractorを注入
    friendHandler := presentation.NewFriendHandler(friendRequestInteractor)  

    return friendHandler
}
