package joystick

import (
	"encoding/json"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"path/filepath"
)

type JoystickController struct {
	Buttons []Button
}

func NewJoystickController() *JoystickController {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(exePath)

	file, err := os.Open(filepath.Join(dir, "config", "mapButtons.json"))
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	var buttons []Button
	if err := json.NewDecoder(file).Decode(&buttons); err != nil {
		log.Fatal("Erro ao fazer conversão do JSON:", err)
	}

	j := &JoystickController{Buttons: buttons}
	j.initSDL()

	return j
}

func (j JoystickController) initSDL() {
	if err := sdl.Init(sdl.INIT_GAMECONTROLLER); err != nil {
		log.Fatalf("Erro ao inicializar SDL: %v", err)
	}

	for sdl.NumJoysticks() < 1 {
		sdl.Quit()
		j.initSDL()
	}

	controller := sdl.GameControllerOpen(0)
	if controller == nil {
		log.Fatal("Não foi possível ler os comandos do controle")
	}

	fmt.Println("Controle conectado!")
}
