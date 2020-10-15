package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//öncelikle fonksiyonun 5 adet değer döndürüceğini belirtiyorum.
func şehirBilgi(şehirİsmi string) (string, string, float64, string, string) {

	URL := "https://api.weatherapi.com/v1/current.json?key=8421aae0c3c74077b37164423201210&q=" + strings.ToLower(şehirİsmi)
	cevap, _ := http.Get(URL)

	defer cevap.Body.Close()
	jsonVeri, _ := ioutil.ReadAll(cevap.Body)
	var şehirVeri Şehir
	hata := json.Unmarshal(jsonVeri, &şehirVeri)
	if hata != nil {
		log.Println(hata)
	}

	//Komut satırına bastırdığımız tarafların kodları sildim.
	return şehirVeri.Knm.İsim, şehirVeri.Knm.Ülke, şehirVeri.Drm.Sıcaklık, şehirVeri.Drm.HDrumu.Metin, şehirVeri.Drm.HDrumu.İkon
}
