package irc

import (
    "net"
    "log"
    "fmt"
)

type BotInterface interface{
    Connect(conn net.Conn, err error)
    Send(string)
}

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

func (bot *Bot) Send(command string){
    fmt.Fprintf(bot.conn, "%s \r\n", command)
} 

func NewBot(server string, port string, nick string, user string, channel string, pass string) *Bot {
    return &Bot{
        server: server,
        port: port,
        nick: nick,
        user: user,
        channel: channel, 
        pass: pass,
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
    
    bot.Send(fmt.Sprintf("USER %s 8 * :%s", bot.nick, bot.nick));
    bot.Send(fmt.Sprintf("NICK %s \r\n", bot.nick));
    bot.Send(fmt.Sprintf("JOIN %s \r\n", bot.channel));
    
    return bot.conn, err
}

