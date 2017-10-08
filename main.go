package main

import (
	"fmt"
    "flag"
    "os/exec"
    "time"

	"github.com/gorilla/websocket"
)

func main() {
    
    BASE_URL := flag.String("server", "", "Manol Server URL")
    KEY := flag.String("key", "", "A Deploy Key")
    PAYLOAD := flag.String("payload", "", "Deploy Script")
    
    flag.Parse()
    
    if *BASE_URL != "" && *KEY != "" && *PAYLOAD != "" {
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
    } else {
        fmt.Println(`Usage: manolclient -server=<server_url> -key=<deploy_key> -payload=<payload_path>`)
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