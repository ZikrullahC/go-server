package main

import (
	"fmt"
	"log"
	"net/http"
)

// Bu formHandler fonksiyonu, bir HTTP POST isteğini alır, form verilerini ayrıştırır ve name ve address alanlarının değerlerini
// HTTP yanıtına yazar. Aşamalar şunlardır:

// Form verilerini ayrıştırmak.
// Başarılı bir isteği bildiren bir mesaj yazmak.
// Form verilerindeki name ve address değerlerini almak ve yanıt olarak yazmak.
// Bu fonksiyon, bir formdan gelen verileri işlemek ve kullanıcıya geri bildirim sağlamak için tipik bir işleyici örneğidir.

func formHandler(w http.ResponseWriter, r *http.Request) {
	//ayristirma
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
