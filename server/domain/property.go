package domain

import "encoding/json"

type Property struct {
	Id          ReadOnlyId  `json:"id" gorm:"primaryKey"`
	Application string      `json:"application" gorm:"index"`
	Profile     string      `json:"profile" gorm:"index"`
	Label       string      `json:"label" gorm:"index"`
	Key         string      `json:"key"`
	Value       interface{} `json:"value"`
}

type ReadOnlyId uint64

// UnmarshalJSON ignores the field value completely when reading.
func (ReadOnlyId) UnmarshalJSON([]byte) error {
	return nil
}

// MarshalBinary encodes the struct into a binary blob
func (p *Property) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalBinary decodes the struct into a Property
func (p *Property) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	return nil
}
