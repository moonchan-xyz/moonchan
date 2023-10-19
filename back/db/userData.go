package db

type Userdata struct {
	Email  string
	Passwd string
	Reason string
}

func CreateUserdata(userdata *Userdata) error {
	tx := db.Create(&userdata)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
