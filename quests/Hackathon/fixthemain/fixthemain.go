package main

import "fmt"

type Door struct {
    state string
}

const (
    OPEN  = "OPEN"
    CLOSE = "CLOSE"
)

func PrintStr(s string) {
    for _, r := range s {
        fmt.Println(r)
    }
    fmt.Println('\n')
}

func OpenDoor(ptrDoor *Door) bool {
    PrintStr("Door Opening...")
    ptrDoor.state = OPEN
    return true
}

func CloseDoor(ptrDoor *Door) bool {
    PrintStr("Door Closing...")
    ptrDoor.state = CLOSE
    return true
}

func IsDoorOpen(door *Door) bool {
    PrintStr("is the Door opened ?")
    return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
    PrintStr("is the Door closed ?")
    return ptrDoor.state == CLOSE
}

func main() {
    door := &Door{}

    OpenDoor(door)
    if IsDoorClose(door) {
        OpenDoor(door)
    }
    if IsDoorOpen(door) {
        CloseDoor(door)
    }
    if door.state == OPEN {
        CloseDoor(door)
    }
}