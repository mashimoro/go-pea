package fizzbuzz

import (
	"bufio"
	"bytes"
	"fizzbuzz/database"
	"fmt"
	"strconv"
	"strings"
	"testing"
	_ "testing"

	"os"
	"time"

	"github.com/secsy/goftp"
)

func TestFstp(t *testing.T) {
	main()
}
func DownloadTextFileFromFtp() []string {
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
		na, naB := getFileImport(file.Name())
		fmt.Println("----->", name, na, naB)

	}

	buf := new(bytes.Buffer)
	err = client.Retrieve("/pea/20WBS.txt", buf)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}

// func TestCreateFileFromFtp(t *testing.T) {
// 	DownloadTextFileFromFtp2()
// }

// func DownloadTextFileFromFtp2() {
// 	config := goftp.Config{
// 		User:               "ftpuser",
// 		Password:           "P@ssw0rdFTPUser",
// 		ConnectionsPerHost: 10,
// 		Timeout:            10 * time.Second,
// 		Logger:             os.Stderr,
// 	}

// 	client, err := goftp.DialConfig(config, "159.138.239.35:21")
// 	if err != nil {
// 		panic(err)
// 	}
// 	files, err := client.ReadDir("pea")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, file := range files {

// 		name := strings.ReplaceAll(file.Name(), ".txt", "")
// 		typeFile, isPCI := getFileImport(file.Name())
// 		fmt.Println("----->", name, typeFile, isPCI)
// 		if isPCI {
// 			buf := new(bytes.Buffer)
// 			err = client.Retrieve("/pea/"+file.Name(), buf)
// 			if err != nil {
// 				panic(err)
// 			}

// 			scanner := bufio.NewScanner(buf)
// 			scanner.Split(bufio.ScanLines)
// 			var txtlines []string

// 			for scanner.Scan() {
// 				txtlines = append(txtlines, scanner.Text())
// 			}
// 			purchasingStatusDetail := getPurchasingStatusDetails(txtlines)
// 			CreatePurchasingStatusDetailByType(purchasingStatusDetail, typeFile)
// 		}
// 	}

// }

// func CreatePurchasingStatusDetailByType(purchasingStatusDetails []purchasing_status_detail.PurchasingStatusDetail, typeFile string) error {
// 	var conn database.Database
// 	db := conn.GetConnectionDB()
// 	db.Exec("delete from purchasing_status_detail wbs_id like ?", typeFile+"%")
// 	fmt.Println("Success delete purchasing_status_detail")
// 	for _, propurchasingStatusDetail := range purchasingStatusDetails {
// 		db.Create(&propurchasingStatusDetail)
// 	}
// 	return nil
// }

//	func getFileImport(fileName string) (string, bool) {
//		name := strings.ReplaceAll(fileName, ".txt", "")
//		if strings.HasSuffix(name, "C") {
//			return "C", true
//		} else if strings.HasSuffix(name, "P") {
//			return "P", true
//		} else if strings.HasSuffix(name, "I") {
//			return "I", true
//		}
//		return "", false
//	}
func main() {
	txtlines := DownloadTextFileFromFtp()
	purchasingStatusDetail := GetPurchasingStatusDetails(txtlines)
	CreatePurchasingStatusDetail(purchasingStatusDetail)
}

func CreatePurchasingStatusDetail(purchasingStatusDetails []PurchasingStatusDetail) error {
	var conn database.Database
	db := conn.GetConnectionDB()
	db.Exec("delete from purchasing_status_detail")
	fmt.Println("Success delete purchasing_status_detail")
	for _, propurchasingStatusDetail := range purchasingStatusDetails {
		db.Create(&propurchasingStatusDetail)
	}
	return nil
}

func GetPurchasingStatusDetails(txtlines []string) []PurchasingStatusDetail {
	var purchasingStatusDetails []PurchasingStatusDetail
	for i, eachline := range txtlines {
		s := strings.Split(eachline, "\t")
		if len(s) > 30 && i != 0 {
			wbs := s[3]
			pr := s[17]
			prq := strings.ReplaceAll(strings.ReplaceAll(s[19], ",", ""), " ", "")
			prpu := strings.ReplaceAll(strings.ReplaceAll(s[21], ",", ""), " ", "")
			companyName := s[22]
			po := s[23]
			poq := strings.ReplaceAll(strings.ReplaceAll(s[25], ",", ""), " ", "")
			popu := strings.ReplaceAll(strings.ReplaceAll(s[27], ",", ""), " ", "")
			bidNo := s[31]
			var fprq = 0.00
			var fprpu = 0.00

			var fpoq = 0.0
			var fpopu = 0.0
			if len(companyName) > 0 {

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

			// fmt.Println("[", i, "]", wbs, companyName, pr, fprq*fprpu, po, fpoq*fpopu, bidNo)
			purchasingStatusDetails = append(purchasingStatusDetails, PurchasingStatusDetail{
				WbsId:                     wbs,
				BidId:                     bidNo,
				PurchaseRequisitionId:     pr,
				PurchaseOrderId:           po,
				PurchaseRequisitionAmount: fprq * fprpu,
				CompanyCode:               companyName,
				PurchaseOrderAmount:       fpoq * fpopu,
				CostSaving:                fprq*fprpu - fpoq*fpopu,
			})
		}
	}
	return purchasingStatusDetails
}
