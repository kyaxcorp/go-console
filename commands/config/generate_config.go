package config

// add this command for generating the config file... and after that to use it!

// import (
// 	"log"

// 	"github.com/kyaxcorp/go-console/command"
// 	"github.com/kyaxcorp/go-core/core/config/autoloader"
// )

// var GenerateDefaultConfig = &command.AddCmd{
// 	GeneratePID: false,
// 	LockProcess: false,
// 	ProcessName: "generate_config",
// 	Cmd:         "generate:config",
// 	Name:        "Generate Default Config",
// 	OnExecute: func(cmd *command.AddCmd) {
// 		// Print information !
// 		log.Println("generating config file...")
// 		if _err := autoloader.GenerateConfigFromMemory(); _err != nil {
// 			log.Println("failed to generate config -> ", _err.Error())
// 		} else {
// 			log.Println("config has been generated successfully")
// 		}
// 	},
// }
