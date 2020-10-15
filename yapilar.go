package main

type Şehir struct {
	Knm Konum `json:"location"`
	Drm Durum `json:"current"`
}
type Konum struct {
	İsim string `json:"name"`
	Ülke string `json:"country"`
}
type Durum struct {
	Sıcaklık float64    `json:"temp_c"`
	HDrumu   HavaDurumu `json:"condition"`
}
type HavaDurumu struct {
	Metin string `json:"text"`
	İkon  string `json:"icon"`
}
