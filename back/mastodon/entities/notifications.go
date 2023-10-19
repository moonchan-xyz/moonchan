package entities

type Notification struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	CreateAt  DATEstring `json:"create_at"`
	AccountID string     `json:"-"`
	Account   *Account   `json:"account" gorm:"-"`
	StatusID  string     `json:"-"`
	Status    *Status    `json:"status" gorm:"-"`
	ReportID  string     `json:"-"`
	Report    *Report    `json:"report,omitempty" gorm:"-"`
	Read      bool       `json:"-"`
}
