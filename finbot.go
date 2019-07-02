package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {

	var (
		bot       *tgbotapi.BotAPI
		err       error
		updChan   tgbotapi.UpdatesChannel
		update    tgbotapi.Update //struct with ubt
		updConfig tgbotapi.UpdateConfig
	)

	const FinApiKey = "ApyKey"

	// Here is buttons.

	var mainMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Prices: "),
			tgbotapi.NewKeyboardButton("Bitcoin 🇸🇴"),
			tgbotapi.NewKeyboardButton("Usd 🤮"),
			tgbotapi.NewKeyboardButton("Eterius💱 "),
		),
	)

	//Bot initialization.
	bot, err = tgbotapi.NewBotAPI(FinApiKey)

	if err != nil {
		log.Panic("bot init error:", err)
		return
	}

	updConfig.Timeout = 60
	updConfig.Limit = 10 //how many msg i can get
	updConfig.Offset = 0

	updChan, err = bot.GetUpdatesChan(updConfig)

	if err != nil {

		log.Panic("update error:", err)

	}
	//Starting gettin updates.
	for {
		update = <-updChan

		if update.Message != nil {

			fmt.Printf("from: %s; chatID: %v; message: %s\n",
				update.Message.From.UserName,
				update.Message.Text,
				update.Message.Chat.ID,
				update.Message.ForwardFromMessageID,
			)

			if update.Message.IsCommand() {
				//Here is first cuple commads i made.
				cmdText := update.Message.Command()
				if cmdText == "help" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"/version - showing curent version, /menu - go to main menu, /close - to close menu ")
					bot.Send(msgConfig)
				}
				if cmdText == "version" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID,
						"version alphan 0.07")

					bot.Send(msg)
				}
				if cmdText == "menu" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There menu:") //obyavlanie peremen
					msg.ReplyMarkup = mainMenu                                        //peredacha znachenia
					bot.Send(msg)
				}
				if cmdText == "close" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Menu is closed")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)
				} else {
					//I need to find out how to improve it.
					//It's should show message if user wrote invalid command.
					fmt.Printf(update.Message.Text, "comand: %s\n")

					msg := tgbotapi.NewMessage(update.Message.Chat.ID,
						" /help to see full comand list") //eto kostil
					bot.Send(msg)
				}

			} else {
				if update.Message.Text == mainMenu.Keyboard[0][1].Text {
					fmt.Printf("message; %s\n",
						update.Message.Text)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID, "13 000 usd")
					bot.Send(msgConfig)
				}
				if update.Message.Text == mainMenu.Keyboard[0][2].Text {
					fmt.Printf("message; %s\n",
						update.Message.Text)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID, "500  usd")
					bot.Send(msgConfig)
				}
				//Here is a problem number 3.
				//This buttom dont work i don't know why.
				//It is should work because it's third in the line, ut it's doesn't
				if update.Message.Text == mainMenu.Keyboard[0][3].Text {
					fmt.Printf("message: %s\n", update.Message.Text)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID, "0,0001 btc")
					bot.Send(msgConfig)
				}
			}
		}
	}

	//and i tryed to use io and bufio to scann document for ApiKey but smthnk went wrong
	//i will find out how to fix that problem

}
