package main

import cmd "github.com/SoulChildTc/soul/cmd/app"

// @title						Swagger Example API
// @version						0.0.1
// @description					This is a sample server celler server.
// @contact.name				SoulChild
// @contact.url					http://soulchild.cn
// @host						localhost:8080
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @BasePath					/
func main() {
	cmd.Execute()
}
