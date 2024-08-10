package dto

type Food struct {
	ID          string        `json:"id" readonly:"true"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Category    string        `json:"category"`
	Price       string        `json:"price"`
	IsNew       bool          `json:"is_new"`
	IsSpicy     bool          `json:"is_spicy"`
	Description string        `json:"description"`
	Composition []Composition `json:"compositions"`
	ErrorReason string        `json:"error_reason"`
}

type Foods []Food
