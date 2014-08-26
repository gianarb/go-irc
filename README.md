## go-irc
Little lib for manage irc bot


### PACKAGE DOCUMENTATION
```go
package irc
```
```go
type Bot struct {
    // contains filtered or unexported fields
}
```
Implement BotInterface 
```go
type BotInterface interface {
    Connect(conn net.Conn, err error)
    Send(string)
}
```
Is the Bot struct contain host, post, nickm user, channel..

```go
func NewBot(server string, port string, nick string, user string, channel string, pass string) *Bot
```
Return new Bot


Bot Methods
* ``` func (bot *Bot) Connect() (conn net.Conn, err error)``` : Connect bot to irc server
* ``` func (bot *Bot) Send(command string) ``` : Send message to channel after connection

```
type Message struct {
    Draft string
}
```
Best prative for manage simple update of lib

### Use Case
```go
package main

import (
    "log"
    "fmt"
    "regexp"
    "bufio"
    "net/textproto"
    "github.com/gianarb/go-irc"
)
 
func main(){
    secretary := NewBot(
        "irc.freenode.net",
        "6667",
        "SybilBot",
        "SybilBot",
        "#channel-name",
        "",
    )
    conn, _ := secretary.Connect()
    defer conn.Close()
 
    reader := bufio.NewReader(bot.conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            log.Fatal("unable to connect to IRC server ", err)
        }
 
        isPing, _ := regexp.MatchString("PING", line)
        if isPing  == true {
            bot.Send("PONG");
        }
        
 
        fmt.Printf("%s\n", line)
    }
}
```
