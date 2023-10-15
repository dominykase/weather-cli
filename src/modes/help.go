package modes

import (
	"fmt"
	"weather/src/messages"
)

type Help struct {}

func (m *Help) Handle() error { 
    fmt.Printf(messages.CommandsMsg)
    return nil 
}
