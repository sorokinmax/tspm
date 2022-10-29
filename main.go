package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	tg "gopkg.in/telebot.v3"
)

var stickersData StickersData
var cfg Config

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found")
	}

	ReadConfigEnv(&cfg)

	ReadDataFile(&stickersData, fmt.Sprintf("%s/stickers.yml", cfg.PathToStickers))
}

func main() {

	pref := tg.Settings{
		Token: cfg.BotToken,
	}

	//create bot
	bot, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	owner := tg.ChatID(cfg.OwnerID)
	packName := fmt.Sprintf("%s_by_%s", stickersData.Packname, bot.Me.Username)

	stickerSet, err := bot.StickerSet(packName)
	if err != nil {
		log.Println(err)
	}
	if stickerSet == nil {
		createStickerPack(owner, *bot, packName)
		err = populateStickers(owner, *bot, "./", packName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// deleting
		for _, sticker := range stickerSet.Stickers[1:] {
			err = bot.DeleteSticker(sticker.FileID)
			if err != nil {
				log.Fatal(err)
			}
		}
		// adding
		err = populateStickers(owner, *bot, "./", packName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
