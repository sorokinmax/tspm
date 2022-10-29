package main

import (
	"fmt"

	tg "gopkg.in/telebot.v3"
)

// Populate Stickers
func populateStickers(owner tg.Recipient, bot tg.Bot, path string, stickerPackName string) error {

	for _, sticker := range stickersData.Stickers {
		file := tg.FromDisk(fmt.Sprintf("%s/%s", cfg.PathToStickers, sticker.File))
		png, err := bot.UploadSticker(owner, &file)
		if err != nil {
			return err
		}

		stickerSet := tg.StickerSet{
			Name:   stickerPackName,
			PNG:    png,
			Emojis: sticker.Emojis,
		}

		err = bot.AddSticker(owner, stickerSet)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create StickerPack
func createStickerPack(owner tg.Recipient, bot tg.Bot, stickerPackName string) error {
	file := tg.FromDisk(fmt.Sprintf("%s/%s", cfg.PathToStickers, stickersData.Headsticker.File))
	png, err := bot.UploadSticker(owner, &file)
	if err != nil {
		return err
	}
	stickerSet := tg.StickerSet{
		Name:      stickerPackName,
		Title:     stickersData.Packtitle,
		Thumbnail: &tg.Photo{File: tg.FromDisk(fmt.Sprintf("%s/%s", cfg.PathToStickers, stickersData.Thumbnail.File)), Width: 100, Height: 100},
		PNG:       png,
		Emojis:    fmt.Sprintf("%s/%s", cfg.PathToStickers, stickersData.Headsticker.Emojis),
	}

	err = bot.CreateStickerSet(owner, stickerSet)
	if err != nil {
		return err
	}
	bot.AddSticker(owner, stickerSet)

	return nil
}
