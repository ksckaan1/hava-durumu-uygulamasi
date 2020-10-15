var onaylaButon = document.querySelector("#onayla")

onaylaButon.addEventListener("click",()=>{
    var sehirIsmi = document.querySelector("#sehir").value;
    sehirSorgu(sehirIsmi)
})

function veriGoster(sehirBilgi, ulkeBilgi, sicaklikBilgi, durumBilgi, ikonBilgi) {
    //Öncelikle index.html'deki elementleri tanıtalım.
    var sonucDiv = document.querySelector("#sonuc")
    var sehir = document.querySelector("#sehirIsmi")
    var ulke = document.querySelector("#ulkeIsmi")
    var sicaklik = document.querySelector("#sicaklik")
    var durum = document.querySelector("#durum")
    var ikon = document.querySelector("#ikon")

     //Ve yerleştirme işlemleri
    sehir.textContent = sehirBilgi;
    ulke.textContent = ulkeBilgi;
    sicaklik.textContent = sicaklikBilgi;
    durum.textContent = durumBilgi;

    var ikonAdres="https:"+ikonBilgi;
    //img etiketindeki src
    ikon.setAttribute("src",ikonAdres);

    //Daha sonra hava durumu verisi bu fonksiyona geldiği için index.html'deki
    //"sonuc" id'li elementi gizli halden çıkaralım.
    sonucDiv.style.display = "block";
}