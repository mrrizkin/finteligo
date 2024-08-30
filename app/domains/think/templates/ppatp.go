package templates

import (
	"encoding/json"
	"fmt"

	"github.com/mrrizkin/finteligo/app/domains/think/types"
	"github.com/tmc/langchaingo/llms"
)

type PPATP struct{}

func NewPPATP() types.PromptTemplate {
	return &PPATP{}
}

func (*PPATP) GenContent(content ...llms.MessageContent) []llms.MessageContent {
	message := make([]llms.MessageContent, 0)

	message = append(
		message,
		llms.TextParts(
			llms.ChatMessageTypeSystem,
			`Anda adalah seorang pegawai bank. Anda harus menilai tingkat risiko nasabah berdasarkan skema berikut:
{
  "age": integer,
  "gender": string,
  "occupation": string,
  "monthly_income": integer,
  "location": string,
  "monthly_transaction_count": integer,
  "total_monthly_transaction_value": integer,
  "source_of_funds": string
  "account_purpose": string
  "financial_status": string
  "credit_history": string
  "legal_history": string
}

Anda harus menentukan tingkat risiko nasabah ini (Rendah, Sedang, atau Tinggi) dan memberikan alasannya.

Mohon berikan jawaban Anda dalam format berikut:
{
  "risk_level": string,
  "reasoning": string
}`,
		),
	)

	message = append(message, llms.TextParts(llms.ChatMessageTypeGeneric, `{
  "age": 30,
  "gender": "laki-laki",
  "occupation": "engineer",
  "monthly_income": 10000000,
  "location": "Jakarta",
  "monthly_transaction_count": 10,
  "total_monthly_transaction_value": 10000000,
  "source_of_funds": "gaji"
  "account_purpose": "tabungan"
  "financial_status": "baik"
  "credit_history": "baik"
  "legal_history": "bersih"
}`))

	message = append(message, llms.TextParts(llms.ChatMessageTypeAI, `{
  "risk_level": "Rendah",
  "reasoning": "Nasabah memiliki status keuangan yang baik, riwayat kredit yang baik, dan riwayat hukum yang bersih."
}`))

	message = append(message, content...)
	return message
}

func (*PPATP) GenMessage(payload interface{}) string {
	data := payload.(*types.PPATPPayload)
	return fmt.Sprintf(`{
  "age": %d,
  "gender": "%s",
  "occupation": "%s",
  "monthly_income": %d,
  "location": "%s",
  "monthly_transaction_count": %d,
  "total_monthly_transaction_value": %d,
  "source_of_funds": "%s",
  "account_purpose": "%s",
  "financial_status": "%s",
  "credit_history": "%s",
  "legal_history": "%s",
}`, data.Age,
		data.Gender,
		data.Occupation,
		data.MonthlyIncome,
		data.Location,
		data.MonthlyTransactionCount,
		data.TotalMonthlyTransactionValue,
		data.SourceOfFunds,
		data.AccountPurpose,
		data.FinancialStatus,
		data.CreditHistory,
		data.LegalHistory,
	)
}

func (*PPATP) OutputParser(output string) (interface{}, error) {
	data := new(types.PPATPResponse)
	err := json.Unmarshal([]byte(output), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
