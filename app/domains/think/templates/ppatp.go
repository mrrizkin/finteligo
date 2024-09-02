package templates

import (
	"encoding/json"
	"fmt"

	"github.com/mrrizkin/finteligo/app/domains/think/helper"
	"github.com/mrrizkin/finteligo/app/domains/think/types"
	"github.com/tmc/langchaingo/llms"
)

type PPATPTemplates struct {
	systemMessage string
}

func NewPPATPTemplates() types.PromptTemplate {
	payloadPromptSchema := helper.GenerateSchema(types.PPATP{})
	responsePromptSchema := helper.GenerateSchema(types.PPATPResponse{})
	systemMessage :=
		fmt.Sprintf(
			`Anda adalah seorang pegawai bank. Anda harus menilai tingkat risiko nasabah berdasarkan skema berikut:
%s

Anda harus menentukan tingkat risiko nasabah ini (Rendah, Sedang, atau Tinggi) dan memberikan alasannya.

Mohon berikan jawaban Anda dalam format berikut:
%s`,
			payloadPromptSchema,
			responsePromptSchema,
		)

	return &PPATPTemplates{
		systemMessage: systemMessage,
	}
}

func (p *PPATPTemplates) GenContent(content ...llms.MessageContent) []llms.MessageContent {
	examplePrompt := helper.Encode(types.PPATP{
		Age:                          30,
		Gender:                       "laki-laki",
		Occupation:                   "engineer",
		MonthlyIncome:                10000000,
		Location:                     "Jakarta",
		MonthlyTransactionCount:      10,
		TotalMonthlyTransactionValue: 10000000,
		SourceOfFunds:                "gaji",
		AccountPurpose:               "tabungan",
		FinancialStatus:              "baik",
		CreditHistory:                "baik",
		LegalHistory:                 "bersih",
	})

	exampleResponse := helper.Encode(types.PPATPResponse{
		RiskLevel: "Rendah",
		Reasoning: "Nasabah memiliki status keuangan yang baik, riwayat kredit yang baik, dan riwayat hukum yang bersih.",
	})

	message := make([]llms.MessageContent, 3)
	message = append(message, llms.TextParts(llms.ChatMessageTypeSystem, p.systemMessage))
	message = append(message, llms.TextParts(llms.ChatMessageTypeGeneric, examplePrompt))
	message = append(message, llms.TextParts(llms.ChatMessageTypeAI, exampleResponse))
	message = append(message, content...)
	return message
}

func (*PPATPTemplates) GenMessage(payload interface{}) string {
	return helper.Encode(payload.(*types.PPATP))
}

func (*PPATPTemplates) OutputParser(output string) (interface{}, error) {
	data := new(types.PPATPResponse)
	err := json.Unmarshal([]byte(output), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
