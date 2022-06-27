package rgwspublic

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Version gets web service version
// returns a string or an error
func Version() (string, error) {

	var version string

	env := `<?xml version="1.0" encoding="UTF-8"?>
	<soap:Envelope 
	xmlns:soap="http://www.w3.org/2003/05/soap-envelope" 
	xmlns:rgw="http://rgwspublic2/RgWsPublic2Service">
	<soap:Header/>
	<soap:Body>
	   <rgw:rgWsPublic2VersionInfo/>
	</soap:Body>
 </soap:Envelope>`

	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(env)))

	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(env))
	if err != nil {
		fmt.Println(err)
	}

	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// decode error
		s := parseXMLError(resp)

		return version, fmt.Errorf("error code: %s, description: %s", s.ErrorCode, s.ErrorDescription)
	}

	// parse response
	version, err = parseVersion(resp)
	if err != nil {
		return version, err
	}

	return version, nil

}

// AFMInfo gets info associated with a VAT number
// accepts a called by VAT and a called for VAT, username and password
// returns AFMData or an error
func AFMInfo(calledby, calledfor, user, pass string) (*ResultTypeData, error) {

	info := &ResultTypeData{}

	// vat number must be between 9 and 12 chars
	if len(calledfor) < 9 || len(calledfor) > 12 {
		return info, fmt.Errorf("non valid afm given: %s", calledfor)
	}

	// calledby can be an empty string
	if calledby != "" {
		// duplicate code from above
		// vat number must be between 9 and 12 chars
		if len(calledby) < 9 || len(calledby) > 12 {
			return info, fmt.Errorf("non valid afm given: %s", calledby)
		}
	}

	// same for username/password
	if len(user) < 6 || len(pass) < 6 {
		return info, fmt.Errorf("username or password cannot be less than 6 chars")
	}

	env := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
	<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:ns1="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:ns2="http://rgwspublic2/RgWsPublic2Service" xmlns:ns3="http://rgwspublic2/RgWsPublic2">
	<env:Header>
	   <ns1:Security>
		  <ns1:UsernameToken>
			 <ns1:Username>%v</ns1:Username>
			 <ns1:Password>%v</ns1:Password>
		  </ns1:UsernameToken>
	   </ns1:Security>
	</env:Header>
	<env:Body>
	   <ns2:rgWsPublic2AfmMethod>
		  <ns2:INPUT_REC>
			 <ns3:afm_called_by/>
			 <ns3:afm_called_for>%v</ns3:afm_called_for>
		  </ns2:INPUT_REC>
	   </ns2:rgWsPublic2AfmMethod>
	</env:Body>
 </env:Envelope> `, user, pass, calledfor)

	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(env)))

	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(env))
	if err != nil {
		return info, err
	}

	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return info, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return info, fmt.Errorf("HTTP Status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	// parse response
	// if errors are present it's the caller's responsibility to print
	info, err = parseAFMInfo(resp)
	if err != nil {
		return info, err
	}

	return info, nil
}

// helper function to get service Version from an xml response
func parseVersion(r *http.Response) (string, error) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
	}

	xmlr := XMLResponse{}

	err = xml.Unmarshal([]byte(rbody), &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml", err)
	}

	// version := xmlr.Body.AFMMethodResponse.AFMCalledByRec.AFMCalledBy
	version := xmlr.Body.PublicVersionInfoResponse.Result

	return version, nil
}

// helper function to get AFM data from an xml response
// func parseAFMInfo(r *http.Response) (*ResultTypeData, error) {
func parseAFMInfo(r *http.Response) (*ResultTypeData, error) {

	// read response Body
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
	}

	// print response body
	// fmt.Println(string(rbody))

	// create an empty struct to unmarshal xml into
	xmlr := XMLResponse{}

	// parse response body into struct
	err = xml.Unmarshal([]byte(rbody), &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml", err)
		return nil, err
	}

	// can't decide whether to change those horrific names or not
	data := xmlr.Body.AFMMethodResponse.Result.ResultType
	data.BasicRec.Activities = xmlr.Body.AFMMethodResponse.Result.ResultType.Activities.Activities
	// data.Error = xmlr.Body.RGWSPublicAfmMethodResponse.PErrorRecOut

	return &data, nil
}

// helper function to get error from an xml response
func parseXMLError(r *http.Response) ErrorRecData {

	// read response body
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
	}

	xmlFile, err := os.Open(string(rbody))
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	// create an empty struct to unmarshal response body into
	xmlr := XMLResponse{}

	// parse response body
	err = xml.Unmarshal(byteValue, &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml:", err)
	}

	// return the child element that contains the actual error
	// return xmlr.Body.AFMMethodResponse.Result.ResultType.ErrorRec
	test := ErrorRecData{}
	return test
}
