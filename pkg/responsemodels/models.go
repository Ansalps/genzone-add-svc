package responsemodels
type Address struct {
	//gorm.Model
	ID uint `json:"id"`
	// CreatedAt time.Time    `json:"created_at"`
	// UpdatedAt time.Time    `json:"updated_at"`
	// DeletedAt sql.NullTime `json:"deleted_at"`
	UserID string `validate:"required,numeric" json:"user_id"`
	//User       User   `gorm:"foriegnkey:UserID;references:ID"`
	Country    string `validate:"required" json:"country"`
	State      string `validate:"required" json:"state"`
	District   string `validate:"required" json:"district"`
	StreetName string `validate:"required" json:"street_name"`
	PinCode    string `validate:"required,numeric" json:"pin_code"`
	Phone      string `validate:"required,numeric,len=10" json:"phone"`
	Default    bool   `gorm:"default:false" validate:"required" json:"default"`
}