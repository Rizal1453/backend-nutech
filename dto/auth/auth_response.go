package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Name string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `json:"token"`


}
type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Name string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`

}