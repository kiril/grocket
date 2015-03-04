package client

import "fmt"
import core "github.com/kiril/grocket/core"

func StartClient(port int) {
    fmt.Printf("Starting Grocket client on port %d\n", port)
}

func Help(command string) string {
    return ""
}

func Status() string {
    return ""
}

func Schedule() bool {
    return false
}

func Get(id string) *core.Event {
    return nil
}

func Next() *core.Event {
    return nil
}
