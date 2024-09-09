package templates

import (
	"encoding/json"
	"fmt"

	"github.com/mrrizkin/finteligo/app/domains/think/helper"
	"github.com/mrrizkin/finteligo/app/domains/think/types"
	lctypes "github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
)

type PPATP struct {
	Age                  int    `json:"age"`
	Gender               string `json:"gender"`
	MaritalStatus        string `json:"marital_status"`
	IsStaffBank          bool   `json:"is_staff_bank"`
	CustomerIdentifier   string `json:"customer_identifier"`
	Citizenship          string `json:"citizenship"`
	DebtorClassification string `json:"debtor_classification"`
	EmploymentStatus     string `json:"employment_status"`
	BusinessField        string `json:"business_field"`
	EmploymentLength     int    `json:"employment_length"`
	SourceOfFunds        string `json:"source_of_funds"`
	IncomeRange          string `json:"income_range"`
	MonthlyIncome        int    `json:"monthly_income"`
	Expenditure          int    `json:"expenditure"`
	AccountPurpose       string `json:"account_purpose"`
	MonthlyTransaction   string `json:"monthly_transaction"`
}

type PPATPPayload struct {
	types.PromptPayload
	Token  lctypes.Token `json:"token"`
	Stream bool          `json:"stream"`
	Data   PPATP         `json:"data"`
}

type PPATPResponse struct {
	RiskLevel string `json:"risk_level"`
	Reasoning string `json:"reasoning"`
}

type PPATPTemplates struct {
	systemMessage string
}

func NewPPATPTemplates() types.PromptTemplate {
	payloadPromptSchema := helper.GenerateSchema(PPATP{})
	responsePromptSchema := helper.GenerateSchema(PPATPResponse{})
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
	examplePrompt := helper.Encode(PPATP{
		Age:                  30,
		Gender:               "laki-laki",
		MaritalStatus:        "menikah",
		IsStaffBank:          false,
		CustomerIdentifier:   "KTP",
		Citizenship:          "WNI",
		DebtorClassification: "A",
		EmploymentStatus:     "Karyawan",
		BusinessField:        "IT",
		EmploymentLength:     5,
		SourceOfFunds:        "Gaji",
		IncomeRange:          "10-20 juta",
		MonthlyIncome:        10000000,
		Expenditure:          5000000,
		AccountPurpose:       "Tabungan",
		MonthlyTransaction:   "1-5 juta",
	})

	exampleResponse := helper.Encode(PPATPResponse{
		RiskLevel: "Rendah",
		Reasoning: "Nasabah memiliki pekerjaan tetap dengan penghasilan yang stabil dan pengeluaran yang terkontrol.",
	})

	message := make([]llms.MessageContent, 3)
	message = append(message, llms.TextParts(llms.ChatMessageTypeSystem, p.systemMessage))
	message = append(message, llms.TextParts(llms.ChatMessageTypeGeneric, examplePrompt))
	message = append(message, llms.TextParts(llms.ChatMessageTypeAI, exampleResponse))
	message = append(message, content...)
	return message
}

func (*PPATPTemplates) GenMessage(payload interface{}) string {
	return helper.Encode(payload.(*PPATP))
}

func (*PPATPTemplates) OutputParser(output string) (interface{}, error) {
	data := new(PPATPResponse)
	err := json.Unmarshal([]byte(output), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
