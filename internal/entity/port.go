package entity

import (
	"github.com/lib/pq"
)

type Port struct {
	Unloc       string          `json:"unloc" gorm:"unique"`
	Name        string          `json:"name"`
	City        string          `json:"city"`
	Country     string          `json:"country"`
	Alias       pq.StringArray  `json:"alias" gorm:"type:text[]"`
	Regions     pq.StringArray  `json:"regions" gorm:"type:text[]"`
	Coordinates pq.Float32Array `json:"coordinates" gorm:"type:real[]"`
	Province    string          `json:"province"`
	Timezone    string          `json:"timezone"`
	Unlocs      pq.StringArray  `json:"unlocs" gorm:"type:text[]"`
	Code        string          `json:"code"`
}
