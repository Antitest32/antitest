package plugs

import (
    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	var MSG string = `
*๐ฎ๐ณ Hello dear, I am a channel spam
detector bot ๐ฎ๐ณ*.
I can ban the channels which
๐ฑspams your chat๐ฑ!

*โค๏ธ(c) @Soul_4 ๐ฎ๐ณ*
	`
	if ctx.EffectiveChat.Type != "private" {
		ctx.EffectiveMessage.Reply(
			bot,
			"โ๐ฎ๐ณ[๐ฑ๐พ๐ ๐ธ๐ ๐ฐ๐ป๐ธ๐๐ด] ๐โ",
		        &gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)
	} else {
		ctx.EffectiveMessage.Reply(
			bot,
			MSG,
			&gotgbot.SendMessageOpts{
				ParseMode: "markdown",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "My Source Code", Url: "https://youtu.be/XqLQ4oKh-0w"},
					}},
				},
			},
		)
	}
	return nil
}
