package urlutil

import (
	"fmt"
	"strings"
	"testing"
)

func Test_urlparse01(t *testing.T) {
	// error string as expected
	errorString := `Get https://aaaaaaaa.blob.core.windows.net/nodeagentpackage-version9-0-0-381/Ubuntu-16.04/batch_config-ubuntu-16.04-1.5.9.tar.gz?sv=2018-03-28&sr=b&sig=a%secret%2Bsecret&st=2019-05-17T01%3A25%3A42Z&se=2021-05-24T01%3A25%3A42Z&sp=r: dial tcp 13.68.165.64:443: i/o timeout`
	inputErr := fmt.Errorf("%s", errorString)
	outputErr := RemoveUrlFromErr(inputErr)
	if strings.Contains(outputErr.Error(), "https://") || strings.Contains(outputErr.Error(), "secret") || !strings.Contains(outputErr.Error(), "[REDACTED]") {
		t.Error("Url removal failed")
	} else {
		fmt.Println(outputErr.Error())
	}
}

func Test_urlparse02(t *testing.T) {
	// error string where scheme is not https because missing space between Get and https://
	errorString := `Gethttps://aaaaaaaa.blob.core.windows.net/nodeagentpackage-version9-0-0-381/Ubuntu-16.04/batch_config-ubuntu-16.04-1.5.9.tar.gz?sv=2018-03-28&sr=b&sig=a%secret%2Bsecret&st=2019-05-17T01%3A25%3A42Z&se=2021-05-24T01%3A25%3A42Z&sp=r: dial tcp 13.68.165.64:443: i/o timeout`
	inputErr := fmt.Errorf("%s", errorString)
	outputErr := RemoveUrlFromErr(inputErr)
	if strings.Contains(outputErr.Error(), "https://") || strings.Contains(outputErr.Error(), "secret") || !strings.Contains(outputErr.Error(), "[REDACTED]") {
		t.Error("Url removal failed")
	} else {
		fmt.Println(outputErr.Error())
	}
}

func Test_urlparse03(t *testing.T) {
	// error string where the uri isn't separated by space with the rest of the error message
	errorString := `Gethttps://aaaaaaaa.blob.core.windows.net/nodeagentpackage-version9-0-0-381/Ubuntu-16.04/batch_config-ubuntu-16.04-1.5.9.tar.gz?sv=2018-03-28&sr=b&sig=a%secret%2Bsecret&st=2019-05-17T01%3A25%3A42Z&se=2021-05-24T01%3A25%3A42Z&sp=r:dial tcp 13.68.165.64:443:i/o timeout`
	inputErr := fmt.Errorf("%s", errorString)
	outputErr := RemoveUrlFromErr(inputErr)
	if strings.Contains(outputErr.Error(), "https://") || strings.Contains(outputErr.Error(), "secret") || !strings.Contains(outputErr.Error(), "[REDACTED]") {
		t.Error("Url removal failed")
	} else {
		fmt.Println(outputErr.Error())
	}
}