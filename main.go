package main

import (
	"fmt"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"
)

func main() {
	//Dosya sistemi nesnemizi oluşturalım.
	statikDosya, _ := fs.New()

	//FileServer fonksiyonunun içindekileri silerek
	//statikDosya nesnesini girelim
	klasör := http.FileServer(statikDosya)
	http.Handle("/", http.StripPrefix("/", klasör))
	go http.ListenAndServe(":5555", nil)

	pencere := webview.New(true)
	pencere.SetTitle("Hava Durumu Uygulamam")
	pencere.SetSize(400, 300, webview.HintFixed)
	pencere.Navigate("http://localhost:5555")
	pencere.Bind("sehirSorgu", func(şehirİsmi string) {
		a, b, c, d, e := şehirBilgi(şehirİsmi)
		jsKodumuz := fmt.Sprintf("veriGoster('%s','%s','%.1f','%s','%s')", a, b, c, d, e)
		pencere.Eval(jsKodumuz)
	})
	pencere.Run()
}
