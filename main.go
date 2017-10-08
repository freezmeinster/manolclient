package main

import (
	"fmt"
    "flag"
    "os/exec"
    "time"

	"github.com/gorilla/websocket"
)

func main() {
    
    BASE_URL := flag.String("server", "manol.o-leafs.com", "Manol Server URL")
    KEY := flag.String("key", "", "A Deploy Key")
    PAYLOAD := flag.String("payload", "/tmp/payload.sh", "Deploy Script")
    
    flag.Parse()
    
	URL := fmt.Sprintf("ws://%s/%s/", *BASE_URL, *KEY)

	var dialer *websocket.Dialer
    
    for {
        conn, _, err := dialer.Dial(URL, nil)
        if err != nil {
            fmt.Println(err)
            time.Sleep(time.Second * 2)
        } else {
            fmt.Println("connected to manol server")
            for {
                _, message, err := conn.ReadMessage()
                if err != nil {
                    fmt.Println("read:", err)
                    break
                }
        
                if fmt.Sprintf("%s", message) == "deploy" {
                    go ExecutePayload(*PAYLOAD)
                } else {
                    fmt.Println("unknown command")
                }
            }
        }
    
        
        
    }
	
}

func ExecutePayload(payload string){
    cmd := exec.Command(payload)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + string(output))
        return
    } else {
        fmt.Println(string(output))
    }
}