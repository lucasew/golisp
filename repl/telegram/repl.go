package main

import (
    "os"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/lucasew/golisp/vm/default"
    "github.com/lucasew/golisp/parser/default"
    "log"
    "fmt"
)

var bot *tgbotapi.BotAPI

const banner = `
Bem-vindo ao golisp, um dialeto de lisp feito em Go!

Como introdução tente realizar operações matemáticas, como por exemplo (+ 2 2)

Outros comandos incluem (env-dump) para despejar o objeto de env que está sendo usado na vm atualmente, que consequentemente expoe todos os elementos visíveis no escopo.

Funções de cálculo, como +, -, * e / só aceitam dois argumentos. Para realizar com mais elementos utilize a função reduce. Ex: (reduce + '(2 3 4)).

Digite /help para reexibir esta mensagem
`



func main() {
    var err error
    if len(os.Args) < 2 {
        panic("Missing telegram bot api key")
    }
    eval := vm_default.NewVM(nil).Eval
    parse := pdefault.Parse
    bot, err = tgbotapi.NewBotAPI(os.Args[1])
    if err != nil {
        panic(err)
    }
    // bot.Debug = true
    log.Printf("Authorized: @%s", bot.Self.UserName)
    updchan, err := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
    if err != nil {
        panic(err)
    }
    for update := range updchan {
        if update.Message == nil {
            continue
        }
        log.Printf("MSG from: @%s: %s", update.Message.From.UserName, update.Message.Text)
        if update.Message.Command() == "help" {
            reply(&update, banner)
            continue
        }
        if update.Message.Command() == "start" {
            reply(&update, banner)
            continue
        }
        ast, err := parse(update.Message.Text)
        if err != nil {
            reply(&update, fmt.Sprintf("🤔 %s", err.Error()))
            continue
        }
        res, err := eval(ast)
        if err != nil {
            reply(&update, fmt.Sprintf("🤯 %s", err.Error()))
            continue
        }
        reply(&update, fmt.Sprintf("👍 %s", res.Repr()))
    }
}

func reply(u *tgbotapi.Update, txt string) {
    m := tgbotapi.NewMessage(int64(u.Message.From.ID), txt)
    m.ReplyToMessageID = u.Message.MessageID
    bot.Send(m)
}
