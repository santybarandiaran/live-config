package domain

type Property struct {
	Id          ReadOnlyId `json:"id" gorm:"primaryKey"`
	Application string     `json:"application" gorm:"index"`
	Profile     string     `json:"profile" gorm:"index"`
	Label       string     `json:"label" gorm:"index"`
	Key         string     `json:"key"`
	Value       string     `json:"value"`
}

type ReadOnlyId uint64

// UnmarshalJSON ignores the field value completely when reading.
func (ReadOnlyId) UnmarshalJSON([]byte) error {
	return nil
}
