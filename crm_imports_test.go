package hubspot

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestGetActiveImports(_ *testing.T) {
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN"), os.Getenv("PRIVATE_APP_SECRET")))
	res, err := cli.CRM.Imports.Active(&CrmActiveImportOptions{
		Before: "",
		After:  "",
		Offset: 0,
	})
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}

func TestGetImportByID(_ *testing.T) {
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN"), os.Getenv("PRIVATE_APP_SECRET")))
	res, err := cli.CRM.Imports.Get(32331356)
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}

func TestCancelImportByID(_ *testing.T) {
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN"), os.Getenv("PRIVATE_APP_SECRET")))
	res, err := cli.CRM.Imports.Cancel(32331339)
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}

func TestImportErrorsByID(_ *testing.T) {
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN"), os.Getenv("PRIVATE_APP_SECRET")))
	res, err := cli.CRM.Imports.Errors(12342, &CrmImportErrorsOptions{
		After: "abcd",
		Limit: 1234,
	})
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}

func createTestCsv(count int) *bytes.Buffer {
	buf := &bytes.Buffer{}
	csvwriter := csv.NewWriter(buf)
	csvHeader := []string{"email", "firstname", "lastname"}
	_ = csvwriter.Write(csvHeader)

	for i := 0; i < count; i++ {
		testFirst := fmt.Sprintf("FirstName3%d", i)
		testLast := fmt.Sprintf("LastName%d", i)
		testEmail := fmt.Sprintf("test%d@example.com", i)
		_ = csvwriter.Write([]string{testEmail, testFirst, testLast})
	}
	csvwriter.Flush()
	return buf
}

func createTestMetadataConfig() *CrmImportConfig {
	result := &CrmImportConfig{
		Name: "Example Create Import",
		ImportOperations: map[string]string{
			"0-1": "CREATE",
		},
		Files: []CrmImportFileConfig{
			{
				FileName:   "example.csv",
				FileFormat: "CSV",
				DateFormat: "DAY_MONTH_YEAR",
				Data:       createTestCsv(6),
				FileImportPage: CrmImportFilePageConfig{
					HasHeader: true,
					ColumnMappings: []CrmImportColumnMapping{
						{
							ColumnObjectTypeID: "0-1",
							ColumnName:         "email",
							PropertyName:       "email",
							IDColumnType:       "HUBSPOT_ALTERNATE_ID",
						},
						{
							ColumnObjectTypeID: "0-1",
							ColumnName:         "firstname",
							PropertyName:       "firstname",
						},
						{
							ColumnObjectTypeID: "0-1",
							ColumnName:         "lastname",
							PropertyName:       "lastname",
						},
					},
				},
			},
		},
	}
	return result
}

func TestImportStart(_ *testing.T) {
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN"), os.Getenv("PRIVATE_APP_SECRET")))
	res, err := cli.CRM.Imports.Start(createTestMetadataConfig())
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
	// t.Error(1)
}
