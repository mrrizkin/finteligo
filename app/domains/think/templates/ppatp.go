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
	AccountPurpose       string `json:"account_purpose"`
	Age                  int    `json:"age"`
	BusinessField        string `json:"business_field"`
	Citizenship          string `json:"citizenship"`
	CustomerIdentifier   string `json:"customer_identifier"`
	DebtorClassification string `json:"debtor_classification"`
	Degree               string `json:"degree"`
	Education            string `json:"education"`
	EmploymentLength     int    `json:"employment_length"`
	EmploymentStatus     string `json:"employment_status"`
	Employment           string `json:"employment"`
	Expenditure          int    `json:"expenditure"`
	Gender               string `json:"gender"`
	IncomeRange          string `json:"income_range"`
	IsStaffBank          bool   `json:"is_staff_bank"`
	MaritalStatus        string `json:"marital_status"`
	MonthlyIncome        int    `json:"monthly_income"`
	MonthlyTransaction   string `json:"monthly_transaction"`
	Position             string `json:"position"`
	Residence            string `json:"residence"`
	SourceOfFunds        string `json:"source_of_funds"`
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
		AccountPurpose:       "Simpanan",
		Age:                  22,
		BusinessField:        "Software Development",
		Citizenship:          "WNI",
		CustomerIdentifier:   "KTP",
		DebtorClassification: "Tidak ada",
		Degree:               "Tidak ada",
		Education:            "SMA/SMK",
		Employment:           "Karyawan",
		EmploymentLength:     2,
		EmploymentStatus:     "Tetap",
		Expenditure:          1000000,
		Gender:               "Laki-laki",
		IncomeRange:          "Rp. 2.500.000 - Rp. 5.000.000",
		IsStaffBank:          false,
		MaritalStatus:        "Belum Kawin",
		MonthlyIncome:        4100000,
		MonthlyTransaction:   "< Rp. 10 Juta",
		Position:             "Staff",
		Residence:            "Milik Keluarga",
		SourceOfFunds:        "Pendapatan Gaji",
	})

	exampleResponse := helper.Encode(PPATPResponse{
		RiskLevel: "Rendah",
		Reasoning: `Nasabah ini memiliki profil risiko rendah berdasarkan beberapa faktor:

1) Tujuan rekening untuk simpanan, yang umumnya berisiko rendah.
2) Warga Negara Indonesia dengan identitas jelas (KTP).
3) Pekerjaan tetap di bidang Software Development dengan masa kerja 2 tahun.
4) Pendapatan bulanan stabil (Rp. 4.100.000) yang sesuai dengan range pendapatannya.
5) Pengeluaran (Rp. 1.000.000) jauh di bawah pendapatan, menunjukkan pengelolaan keuangan yang baik.
6) Transaksi bulanan di bawah Rp. 10 juta, sesuai dengan profil pendapatannya.
7) Sumber dana dari gaji, yang mudah diverifikasi.
8) Bukan staf bank, mengurangi risiko konflik kepentingan.

Meskipun usianya relatif muda (22 tahun) dan pendidikannya SMA/SMK, faktor-faktor lain menunjukkan stabilitas finansial dan profil risiko yang rendah.`,
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
	outcome := helper.TrimPromptResultJson(output)
	data := new(PPATPResponse)
	err := json.Unmarshal([]byte(outcome), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
