package main

import (
	"fmt"
	"log"

	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/config"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/wire"
)

//	@title			Totality-Crew
//	@version		2.0
//	@description	MICROSERVICES BUILD USING GOLANG following Clean-Code Architecture.

//	@contact
// name: Ganesh
// url: https://github.com/rganes5
// email: ganeshraveendranit@gmail.com

//	@license
// name: MIT
// url: https://opensource.org/licenses/MIT

//	@host	localhost:3000

// @Basepath	/
// @Accept		json
// @Produce	json
// @Router		/ [get]

func main() {
	//swag init -g cmd/main.go
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatalln("Failed to load config:", configErr)
	}

	server, wireErr := wire.InitializeAPI(&config)
	if wireErr != nil {
		log.Fatal("Failed to initialize server", wireErr)
	}
	server.Start()
	fmt.Println("Running on PORT 3000")

}
