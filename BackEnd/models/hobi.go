package models

type Hobi struct {
	IdHobi   uint   `gorm:"column:id_hobi;primaryKey;autoIncrement" json:"id_hobi"`
	NamaHobi string `gorm:"column:nama_hobi;not null" json:"nama_hobi"`
	Kategori string `gorm:"column:kategori_hobi;type:enum('makan','tidur','bermain','mancing');not null;index" json:"kategori_hobi"`
}

func (Hobi) TableName() string { return "hobi" }

type HobiRequest struct {
	NamaHobi string `json:"nama_hobi"`
	Kategori string `json:"kategori_hobi"`
}
 