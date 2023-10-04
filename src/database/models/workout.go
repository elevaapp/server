package models

type Workout struct {
	UserId      string             `json:"user_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Components  []WorkoutComponent `json:"components" gorm:"many2many:workout_components"`
}

type WorkoutComponent struct {
	WorkoutId   string `json:"workout_id" gorm:"primaryKey"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Sets        int8   `json:"sets"`
	Repetitions int8   `json:"repetitions"`
	Time        int16  `json:"time"`
}
