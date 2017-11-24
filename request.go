package rgwspublic

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Version gets web service version
// returns a string or an error
func Version() (string, error) {

	var version string

	env := `<?xml version="1.0" encoding="UTF-8"?>
			<env:Envelope
			xmlns:env="http://schemas.xmlsoap.org/soap/envelope/"
			xmlns:ns="http://gr/gsis/rgwspublic/RgWsPublic.wsdl"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xmlns:xsd="http://www.w3.org/2001/XMLSchema">
			<env:Header/>
			<env:Body>
			<ns:rgWsPublicVersionInfo/>
			</env:Body>
			</env:Envelope>`

	header := http.Header{}
	header.Set("Content-Type", "text/xml")
	header.Set("Connection", "Close")
	header.Set("Content-Length", string(len(env)))

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

		return version, fmt.Errorf("error code: %s, description: %s", s.ErrorCode, s.ErrorDescr)
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
func AFMInfo(calledby, calledfor, user, pass string) (*AFMData, error) {

	info := &AFMData{}

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

	env := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
		<env:Envelope
		xmlns:env="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:ns="http://gr/gsis/rgwspublic/RgWsPublic.wsdl"
		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:xsd="http://www.w3.org/2001/XMLSchema"
		xmlns:ns1="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-
		1.0.xsd">
		<env:Header>
		<ns1:Security>
		<ns1:UsernameToken>
		<ns1:Username>%v</ns1:Username>
		<ns1:Password>%v</ns1:Password>
		</ns1:UsernameToken>
		</ns1:Security>
		</env:Header>
		<env:Body>
		<ns:rgWsPublicAfmMethod>
		<RgWsPublicInputRt_in xsi:type="ns:RgWsPublicInputRtUser">
		<ns:afmCalledBy>%v</ns:afmCalledBy>
		<ns:afmCalledFor>%v</ns:afmCalledFor>
		</RgWsPublicInputRt_in>
		<RgWsPublicBasicRt_out xsi:type="ns:RgWsPublicBasicRtUser">
		<ns:afm xsi:nil="true"/>
		<ns:stopDate xsi:nil="true"/>
		<ns:postalAddressNo xsi:nil="true"/>
		<ns:doyDescr xsi:nil="true"/>
		<ns:doy xsi:nil="true"/>
		<ns:onomasia xsi:nil="true"/>
		<ns:legalStatusDescr xsi:nil="true"/>
		<ns:registDate xsi:nil="true"/>
		<ns:deactivationFlag xsi:nil="true"/>
		<ns:deactivationFlagDescr xsi:nil="true"/>
		<ns:postalAddress xsi:nil="true"/>
		<ns:firmFlagDescr xsi:nil="true"/>
		<ns:commerTitle xsi:nil="true"/>
		<ns:postalAreaDescription xsi:nil="true"/>
		<ns:INiFlagDescr xsi:nil="true"/>
		<ns:postalZipCode xsi:nil="true"/>
		</RgWsPublicBasicRt_out>
		<arrayOfRgWsPublicFirmActRt_out xsi:type="ns:RgWsPublicFirmActRtUserArray"/>
		<pCallSeqId_out xsi:type="xsd:decimal">0</pCallSeqId_out>
		<pErrorRec_out xsi:type="ns:GenWsErrorRtUser">
		<ns:errorDescr xsi:nil="true"/>
		<ns:errorCode xsi:nil="true"/>
		</pErrorRec_out>
		</ns:rgWsPublicAfmMethod>
		</env:Body>
		</env:Envelope>`, user, pass, calledby, calledfor)

	header := http.Header{}
	header.Set("Content-Type", "text/xml")
	header.Set("Connection", "Close")
	header.Set("Content-Length", string(len(env)))

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
		// decode error
		s := parseXMLError(resp)

		return info, fmt.Errorf("error code: %s, description: %s", s.ErrorCode, s.ErrorDescr)
	}

	// parse response
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

	err = xml.Unmarshal(rbody, &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml", err)
	}

	version := xmlr.Body.RGWSPublicVersionInfoResponse.Result

	return version, nil
}

// helper function to get AFM data from an xml response
func parseAFMInfo(r *http.Response) (*AFMData, error) {

	// read response Body
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
	}

	// create an empty struct to unmarshal xml into
	xmlr := XMLResponse{}

	// parse response body into struct
	err = xml.Unmarshal(rbody, &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml", err)
	}

	// can't decide whether to change those horrific names or not
	data := xmlr.Body.RGWSPublicAfmMethodResponse.RgWsPublicBasicRtOut
	data.Activities = xmlr.Body.RGWSPublicAfmMethodResponse.ArrayOfRgWsPublicFirmActRtOut.RgWsPublicFirmActRtUser

	return &data, nil
}

// helper function to get error from an xml response
func parseXMLError(r *http.Response) XMLPErrorRecOut {

	// read response body
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
	}

	// create an empty struct to unmarshal response body into
	xmlr := XMLResponse{}

	// parse response body
	err = xml.Unmarshal(rbody, &xmlr)
	if err != nil {
		fmt.Println("error unmarshaling xml:", err)
	}

	// return the child element that contains the actual error
	return xmlr.Body.RGWSPublicAfmMethodResponse.PErrorRecOut
}
