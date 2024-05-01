module example.com/call

go 1.22.2

require github.com/gorilla/mux v1.8.1

require example.com/throttle v0.0.0-00010101000000-000000000000 // indirect

replace example.com/throttle => ../throttle
