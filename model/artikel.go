package model

type Artikel struct {
	Id                     string
	Judul_jurnal           string
	Penulis_jurnal         string
	Tentang_penulis_jurnal string
	Penerbit_jurnal        string
	Tanggal_publikasi      int64
	Bidang_subjek          string
	Abstrak                string
	Created_at             int64
	Updated_at             int64
}

type ArtikelView struct {
	Judul_jurnal           string `json:"judul_jurnal" example:""`
	Penulis_jurnal         string `json:"penulis_jurnal" example:""`
	Tentang_penulis_jurnal string `json:"tentang_penulis_jurnal" example:""`
	Penerbit_jurnal        string `json:"penerbit_jurnal" example:""`
	Bidang_subjek          string `json:"bidang_subjek" example:""`
	Abstrak                string `json:"abstrak" example:""`
}
