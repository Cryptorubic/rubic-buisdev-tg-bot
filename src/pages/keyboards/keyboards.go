package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pseudoelement/rubic-buisdev-tg-bot/src/consts"
)

var StartPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"1️⃣ I'm representing a DEX / Bridge / Chain / Aggregator / Intent protocol",
			consts.COLLABORATE,
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"2️⃣ I'm interested in your API / integration docs",
			consts.INTEGRATE,
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"3️⃣ I want to report a bug or product issue",
			consts.SUPPORT,
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"4️⃣ Something else",
			consts.OTHER,
		),
	),
)

var ThanksPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Back to start page.",
			consts.BACK_TO_START,
		),
	))

var SupportPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL(
			"✉️ Egor's Telegram (Lead Support)",
			"https://t.me/eobuhow",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL(
			"✉️ Daniel's Telegram (Support)",
			"https://t.me/@daniel_melgin",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL(
			"✉️ Vasily's Telegram (Support)",
			"https://t.me/whymint",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Describe issue.",
			consts.DESCRIBE_ISSUE,
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Back to start page.",
			consts.BACK_TO_START,
		),
	))

var IntegrationPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	sdkApiLink,
	integrationGuideLink,
	productOverviewLink,
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Back to start page.",
			consts.BACK_TO_START,
		),
	),
)

// admin pages keyboards
var (
	AdminStartPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔍 Show messages",
				consts.SHOW_MESSAGES,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔗 Check links",
				consts.CHECK_LINKS,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"❌ Delete oldest messages",
				consts.DELETE_MESSAGES,
			),
		),
	)

	AdminLinksPageKeyboard = IntegrationPageKeyboard

	AdminListOfLinksPageKeyboard = ThanksPageKeyboard

	AdminInfoAfterDeletionMsgPageKeyboard = ThanksPageKeyboard

	AdminOldOrNewMessagesPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Show only new messages",
				consts.SHOW_NEW_MESSAGES,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Show all messages",
				consts.SHOW_ALL_MESSAGES,
			),
		),
	)
)
