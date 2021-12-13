package plugs

import (
    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	var MSG string = `
*🇮🇳 Hello, I am a channel spam
detector bot 🇮🇳*.
I can ban the channels which
🔱spams your chat🔱!

*❤️(c) @clanlord7 🇮🇳*
	`
	if ctx.EffectiveChat.Type != "private" {
		ctx.EffectiveMessage.Reply(
			bot,
			"Bot is Alive (:",
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
