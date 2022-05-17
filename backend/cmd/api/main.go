package main

import "secondChance/internal/app"

func main() {
	app.Run()
}

// func main() {
// 	handler := http.HandlerFunc(handleRequest)
// 	http.Handle("/photo", handler)
// 	http.ListenAndServe(":3000", nil)
// }

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	fileBytes, err := ioutil.ReadFile("./images/shop/13/742b1d2d58244ebab48641c969b14302.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(fileBytes)
// 	return
// }
