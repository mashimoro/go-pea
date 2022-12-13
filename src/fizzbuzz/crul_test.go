package fizzbuzz

import (
	"encoding/json"
	"fizzbuzz/database"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"gorm.io/gorm"
)

func TestCrul(t *testing.T) {
	reader := strings.NewReader("")
	request, err := http.NewRequest("GET", "https://imf-io.pea.co.th/pea/capex-ex/pst/zbudr085?wbs=P-TDD02.3-B-ACT65.6000", reader)
	var bearer = "Bearer AklgcwG5Hr8QDf2Sefp63ae6zz9Q6nft"
	request.Header.Add("Authorization", bearer)
	request.Header.Add("x-client-id", "adhoc-pst")
	fmt.Println(err)
	// TODO: check err
	client := &http.Client{}
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	data := Request{}
	json.Unmarshal([]byte(sb), &data)
	fmt.Println("---:", data.Wbs, data.Balance, data.Budget)
}

type Request struct {
	Balance string `json:"balance"`
	Budget  float64
	Dept    string
	Gr      float64
	Ir      float64
	NotYet  float64 `json:"not_yet"`
	Pay     float64
	Po      float64
	Pr      float64
	Project string
	Wbs     string
}
type CostSaving struct {
	CostSaving float64
	Wbs        string
}

func TestWithWbs(t *testing.T) {

	var conn database.Database
	db := conn.GetConnectionDB()

	var ps []PurchasingStatus
	db.Find(&ps)
	for _, us := range ps {
		var purchasingStatusDetail PurchasingStatusDetail
		re := getPurchasingStatusUrl(us.Wbs)

		db.Select(
			"sum(purchase_requisition_amount) - sum(purchase_order_amount) as cost_saving", "wbs_id").Where("wbs_id = ?", us.Wbs).Group("wbs_id").Find(&purchasingStatusDetail)
		fmt.Println(purchasingStatusDetail.WbsId, ": costSaving :", purchasingStatusDetail.CostSaving)

		fmt.Println("re>", re.NotYet)
		us.AccumulatedDisbursement = re.Pay
		us.Budget = re.Budget
		us.Organization = re.Dept
		us.BudgetName = re.Project
		us.BudgetRefund = re.NotYet
		us.CostSaving = purchasingStatusDetail.CostSaving

		us.HighLow = us.DisbursementTarget - (us.AccumulatedDisbursement + us.CostSaving + us.BudgetRefund)

		fmt.Println(us.Wbs, "+=>", us.Organization, us.BudgetName, us.AccumulatedDisbursement, us.Budget, us.BudgetRefund, us.HighLow)
		db.Save(&us)
	}
}

func getPurchasingStatusUrl(wbs string) Request {
	reader := strings.NewReader("")
	url := "https://imf-io.pea.co.th/pea/capex-ex/pst/zbudr085?wbs=" + wbs
	request, err := http.NewRequest("GET", url, reader)
	var bearer = "Bearer AklgcwG5Hr8QDf2Sefp63ae6zz9Q6nft"
	request.Header.Add("Authorization", bearer)
	request.Header.Add("x-client-id", "adhoc-pst")

	client := &http.Client{}
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	data := Request{}
	json.Unmarshal([]byte(sb), &data)
	return data
}

type PurchasingStatus struct {
	gorm.Model
	// todo change struct to business domain by data dictionary
	Wbs                     string `gorm:"primarykey"`
	Organization            string
	BudgetName              string
	DisbursementTarget      float64
	AccumulatedDisbursement float64
	CostSaving              float64
	BudgetRefund            float64
	HighLow                 float64
	Budget                  float64
}

func (PurchasingStatus) TableName() string {
	return "purchasing_status"
}
