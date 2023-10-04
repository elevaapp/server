package models

type User struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token" gorm:"unique"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime"`
	Workouts  []Workout `json:"workouts" gorm:"foreignKey:UserId;References:Id"`
}
