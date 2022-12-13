package fizzbuzz

import (
	"bufio"
	"bytes"
	"fizzbuzz/database"
	"fizzbuzz/domain/purchasing_status_detail"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/secsy/goftp"
)

func TestCreateFileFromFtp2(t *testing.T) {
	fmt.Print("007")
	DownloadTextFileFromFtp2()
}

func DownloadTextFileFromFtp2() {
	config := goftp.Config{
		User:               "ftpuser",
		Password:           "P@ssw0rdFTPUser",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, "159.138.239.35:21")
	if err != nil {
		panic(err)
	}
	files, err := client.ReadDir("pea")
	if err != nil {
		panic(err)
	}
	for _, file := range files {

		name := strings.ReplaceAll(file.Name(), ".txt", "")
		typeFile, isPCI := getFileImport(file.Name())
		fmt.Println("----->", name, typeFile, isPCI)
		if isPCI {
			buf := new(bytes.Buffer)
			err = client.Retrieve("/pea/"+file.Name(), buf)
			if err != nil {
				panic(err)
			}

			scanner := bufio.NewScanner(buf)
			scanner.Split(bufio.ScanLines)
			var txtlines []string

			for scanner.Scan() {
				txtlines = append(txtlines, scanner.Text())
			}
			purchasingStatusDetail := getPurchasingStatusDetails(txtlines)
			CreatePurchasingStatusDetailByType(purchasingStatusDetail, typeFile)

			err = client.Store("/pea/complate/"+file.Name(), buf)
			if err != nil {
				fmt.Println("----Move Eror fff----", err)
			}
			err = client.Delete("/pea/" + file.Name())
			if err != nil {
				fmt.Println("Delete Eror", err)
			}
		}

	}

}

func CreatePurchasingStatusDetailByType(purchasingStatusDetails []purchasing_status_detail.PurchasingStatusDetail, typeFile string) error {
	var conn database.Database
	db := conn.GetConnectionDB()
	db.Exec("delete from purchasing_status_detail wbs_id like ?", typeFile+"%")
	fmt.Println("Success delete purchasing_status_detail")
	for _, propurchasingStatusDetail := range purchasingStatusDetails {
		db.Create(&propurchasingStatusDetail)
	}
	return nil
}

func getFileImport(fileName string) (string, bool) {
	name := strings.ReplaceAll(fileName, ".txt", "")
	if strings.HasSuffix(name, "C") {
		return "C", true
	} else if strings.HasSuffix(name, "P") {
		return "P", true
	} else if strings.HasSuffix(name, "I") {
		return "I", true
	}
	return "", false
}
func getPurchasingStatusDetails(txtlines []string) []purchasing_status_detail.PurchasingStatusDetail {
	var purchasingStatusDetails []purchasing_status_detail.PurchasingStatusDetail
	for i, eachline := range txtlines {
		s := strings.Split(eachline, "\t")
		var row = -1
		if len(s) > 30 && i != 0 {
			wbs := s[3+row]
			materialId := s[8+row]
			material := s[9+row]
			pr := s[17+row]
			prq := strings.ReplaceAll(strings.ReplaceAll(s[19+row], ",", ""), " ", "")
			prpu := strings.ReplaceAll(strings.ReplaceAll(s[21+row], ",", ""), " ", "")
			companyCode := s[22+row]
			po := s[23+row]
			poq := strings.ReplaceAll(strings.ReplaceAll(s[25+row], ",", ""), " ", "")
			popu := strings.ReplaceAll(strings.ReplaceAll(s[27+row], ",", ""), " ", "")

			grIrDoc := s[34+row]
			pmtDoc := s[40+row]
			dateString := convertStringToDateFormat(s[41+row])
			pmtDate, _ := time.Parse("2006-01-02", dateString)

			var fprq = 0.00
			var fprpu = 0.00

			var fpoq = 0.0
			var fpopu = 0.0
			if len(companyCode) > 0 {
				if s, err := strconv.ParseFloat(prq, 64); err == nil {
					fprq = s
				}
				if s, err := strconv.ParseFloat(prpu, 64); err == nil {
					fprpu = s
				}
				if s, err := strconv.ParseFloat(poq, 64); err == nil {
					fpoq = s
				}
				if s, err := strconv.ParseFloat(popu, 64); err == nil {
					fpopu = s
				}
			}

			purchasingStatusDetails = append(purchasingStatusDetails, purchasing_status_detail.PurchasingStatusDetail{
				WbsId:                     wbs,
				MaterialId:                materialId,
				Material:                  material,
				PurchaseRequisitionId:     pr,
				PurchaseOrderId:           po,
				PurchaseRequisitionAmount: fprq * fprpu,
				CompanyCode:               companyCode,
				PurchaseOrderAmount:       fpoq * fpopu,
				CostSaving:                fprq*fprpu - fpoq*fpopu,
				GrIrDoc:                   grIrDoc,
				PmtDoc:                    pmtDoc,
				PmtDate:                   pmtDate,
			})
		}
	}
	return purchasingStatusDetails
}
func convertStringToDateFormat(dateString string) string {
	res := strings.Split(dateString, ".")
	res = revert(res)
	if len(res) == 3 {
		return res[0] + "-" + res[1] + "-" + res[2]
	}
	return ""
}

func revert(res []string) []string {
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
