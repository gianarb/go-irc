package main

import (
    "net"
    "log"
    "bufio"
    "fmt"
    "regexp"
    "net/textproto"
)

type Bot struct{
        server string
        port string
        nick string
        user string
        channel string
        pass string
        pread, pwrite chan string
        conn net.Conn
}

func NewBot() *Bot {
        return &Bot{
            server: "irc.freenode.net",
            port: "6667",
            nick: "SibylBot",
            user: "SibylBot",
            channel: "#gianarb-dev", 
            pass: "",
            conn: nil,
        }
}

func (bot *Bot) Connect() (conn net.Conn, err error){
  conn, err = net.Dial("tcp", bot.server + ":" + bot.port)
  if err != nil{
    log.Fatal("unable to connect to IRC server ", err)
  }
  bot.conn = conn
  log.Printf("Connected to IRC server %s (%s) \n", bot.server, bot.conn.RemoteAddr())
  return bot.conn, nil
}



func main(){
    secretary := NewBot()
    conn, _ := secretary.Connect()

    fmt.Fprintf(conn, "USER %s 8 * :%s \r\n", secretary.nick, secretary.nick)
    fmt.Fprintf(conn, "NICK %s \r\n", secretary.nick)
    fmt.Fprintf(conn, "JOIN %s \r\n", secretary.channel)

    defer conn.Close()

    reader := bufio.NewReader(conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            log.Fatal("unable to connect to IRC server ", err)
        }
        iam, _ := regexp.MatchString(":!iam", line)
        if iam  == true {
            fmt.Fprintf(conn, "PRIVMSG %s :You are beautiful! \r\n", secretary.channel);
        }
        isPing, _ := regexp.MatchString("PING", line)
        if isPing  == true {
            fmt.Fprintf(conn, "PONG \r\n");
        }
        fmt.Printf("%s\n", line)
    }

}
