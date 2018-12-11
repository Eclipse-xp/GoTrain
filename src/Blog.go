package main

////开放博客
//func (p *Person)showBlog(){
//	mux := http.NewServeMux()
//	rh := http.RedirectHandler("http://744722813.iteye.com/", 307)
//	mux.Handle("/blog", rh)
//
//	mux.Handle("/hello", timeHandler(time.RFC1123))
//	log.Println("Listening...")
//
//	http.ListenAndServe(":3000", mux)
//}
//func timeHandler(format string) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		tm := time.Now().Format(format)
//		w.Write([]byte("Hey ****"+ r.FormValue("name") +"**** The time is: " + tm))
//	}
//	return http.HandlerFunc(fn)
//}
