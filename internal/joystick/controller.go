package joystick

import (
	input "JoyGo/internal/Input"
	"encoding/json"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
)

type JoystickController struct {
	Buttons  []Button
	Executor input.Executor
}

type Button struct {
	Name          string `json:"name"`
	Value         uint8  `json:"value"`
	KeyDown       string `json:"keyDown"`
	KeyUp         string `json:"keyUp"`
	MouseFunction bool   `json:"mouseFunction"`
}

func NewJoystickController() *JoystickController {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	path := filepath.Join(dir, "config", "mapButtons.json")

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Erro ao abrir o arquivo: %v", err))
	}
	defer file.Close()

	var buttons []Button
	if err := json.NewDecoder(file).Decode(&buttons); err != nil {
		panic(fmt.Sprintf("Erro ao fazer conversão do JSON: %v", err))
	}

	var joystick JoystickController

	joystick = JoystickController{Buttons: buttons, Executor: input.DefaultExecutor}

	joystick.initSDL()

	fmt.Println("Controle conectado!")
	return &joystick
}

func (j JoystickController) initSDL() {

	if err := sdl.Init(sdl.INIT_GAMECONTROLLER); err != nil {
		panic(fmt.Sprintf("Erro ao inicializar SDL: %v", err))
	}

	for sdl.NumJoysticks() < 1 {
		sdl.Quit()
		j.initSDL()
	}

	controller := sdl.GameControllerOpen(0)
	if controller == nil {
		panic(fmt.Sprintf("Não foi possível ler os comandos do controle"))
	}
}
