package sample

import "net/http"

func SampleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
	Tmp string `json:"hoge"`
}
