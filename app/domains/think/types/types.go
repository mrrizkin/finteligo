package types

import (
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
)

type PromptTemplate interface {
	GenContent(content ...llms.MessageContent) []llms.MessageContent
	GenMessage(payload interface{}) string
	OutputParser(output string) (interface{}, error)
}

type PromptPayload struct {
	types.PromptPayload
	Token  types.Token `json:"token"`
	Stream bool        `json:"stream"`
}

type PPATPPayload struct {
	types.PromptPayload
	Token                        types.Token `json:"token"`
	Stream                       bool        `json:"stream"`
	Age                          int         `json:"age"                             validate:"required"`
	Gender                       string      `json:"gender"                          validate:"required"`
	Occupation                   string      `json:"occupation"`
	MonthlyIncome                int         `json:"monthly_income"                  validate:"required"`
	Location                     string      `json:"location"                        validate:"required"`
	MonthlyTransactionCount      int         `json:"monthly_transaction_count"       validate:"required"`
	TotalMonthlyTransactionValue int         `json:"total_monthly_transaction_value" validate:"required"`
	SourceOfFunds                string      `json:"source_of_funds"`
	AccountPurpose               string      `json:"account_purpose"`
	FinancialStatus              string      `json:"financial_status"`
	CreditHistory                string      `json:"credit_history"`
	LegalHistory                 string      `json:"legal_history"`
}

type PPATPResponse struct {
	RiskLevel string `json:"risk_level"`
	Reasoning string `json:"reasoning"`
}
