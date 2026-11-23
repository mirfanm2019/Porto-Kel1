package models


type Alamatt struct {
	IDAlamat     uint       `gorm:"column:id_alamat;primaryKey;autoIncrement" json:"id_alamat"`
	NamaAlamat    string     `gorm:"column:nama_alamat;not null" json:"nama_alamat"`
	KodePos      string     `gorm:"column:kode_pos;not null" json:"kode_pos"`
	
}

func (Alamatt) TableName() string { return "alamatt" }

type AlamattRequest struct {
	NamaAlamat string `json:"nama_alamat"`
	KodePos string `json:"kode_pos"`
}

