package serializer

import "github.com/datty258/gorm_project/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	CreatedAt int64  `json:"created_at"`
	Birthday  int64  `json:"birthday"`
	Age       int    `json:"age"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	tel, _ := model.GetTel(user.ID)
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Tel:       tel.Tel,
		CreatedAt: user.CreatedAt.Unix(),
		Age:       user.Age,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
