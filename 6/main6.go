package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FName = "Cumhur"
	hum.LName = "TOKALAK"
	hum.Age = 24

	//Formu parse etmek için
	r.ParseForm()

	//Sunucudan form bilgisini almak için
	fmt.Println(r.Form)

	//URL nin path bilgisini almak için

	fmt.Println("path", r.URL.Path)

	fmt.Print(w, "<table><tr><td<b>Ad<b></td><td<b>Soyad<b></td><td<b>Yas<b></td></tr><tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr><tr></tr><tr><td><b></td><td>"+r.URL.Path+"</td></tr></table>")

}
func main() {
	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
