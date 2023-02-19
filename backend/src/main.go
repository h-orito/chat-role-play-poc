package main

import "chat-role-play-poc/src/inject"

func main() {
	r := inject.InjectRouting()
	r.Gin.Run() // 0.0.0.0:8080
}
