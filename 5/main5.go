package main

import "net/http"

func search(res http.ResponseWriter, r *http.Request) {

	hl := r.FormValue("hl")
	source := r.FormValue("source")
	ei := r.FormValue("ei")
	q := r.FormValue("q")

	data := "/search?hl=" + hl + "&source=" + source + "&ei=" + ei + "&q=" + q
	res.Write([]byte(data))
}
func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

//https://www.google.com.tr/search?q=cumhur+mesut+&hl=tr&source=hp&ei=CycrYoqTMfSOxc8PtbijKA
