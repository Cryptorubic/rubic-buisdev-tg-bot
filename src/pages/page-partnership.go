package pages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pseudoelement/rubic-buisdev-tg-bot/src/consts"
	"github.com/pseudoelement/rubic-buisdev-tg-bot/src/models"
	query_builder "github.com/pseudoelement/rubic-buisdev-tg-bot/src/query-builder"
)

type PartnershipPage struct {
	*AbstrUserInputPage
}

func NewPartnershipPage(db models.IDatabase, bot *tgbotapi.BotAPI, adminQueryBuilder *query_builder.AdminQueryBuilder) *PartnershipPage {
	basePage := NewPage(db, bot, adminQueryBuilder)
	p := &PartnershipPage{
		AbstrUserInputPage: NewAbstrUserInputPage(basePage),
	}
	p.setCurrenPage(p)

	return p
}

func (this *PartnershipPage) Name() string {
	return consts.PARTNERSHIP_PAGE
}

func (this *PartnershipPage) RespText(update tgbotapi.Update) string {
	if this.errResp != "" {
		return this.errResp
	}

	return `Awesome!🙌 Let's explore a potential collaboration.

Can you share the following:
- Project name
- Website
- Your role
- Your main goal with us? (integration / liquidity aggregation / mutual routing / co-marketing / other)
- Screenshot (no more than 1 image per request)

Once you're done, I’ll share this with our BD team and we’ll follow up fast.  `
}

var _ models.IPage = (*PartnershipPage)(nil)
var _ models.IPageWithActionOnDestroy = (*PartnershipPage)(nil)
