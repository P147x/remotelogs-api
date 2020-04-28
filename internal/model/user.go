package model


//User is a model
type User struct {
	ID      	 int       `gorm:"primary_key, AUTO_INCREMENT"`
	Username     string    `form:"username" binding:"required" gorm:"unique_index" json:”username”`
	Password     string    `form:"password" binding:"required" json:”password”`
}

//IsInitialized is used to check if the structur is not empty
func (u User) IsInitialized() bool {
	if (User{}) == u {
		return false
	}
	return true
}