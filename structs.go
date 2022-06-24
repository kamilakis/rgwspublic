package rgwspublic

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// XMLResponse is where we parse an http response
type XMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    XMLBody
}

// XMLBody is the body of a response
type XMLBody struct {
	XMLName                       xml.Name `xml:"Body"`
	RGWSPublicAfmMethodResponse   XMLRGWSPublicAfmMethodResponse
	RGWSPublicVersionInfoResponse XMLRGWSPublicVersionInfoResponse
}

// XMLRGWSPublicAfmMethodResponse is the response to a publicAFM method
type XMLRGWSPublicAfmMethodResponse struct {
	XMLName                       xml.Name `xml:"rgWsPublicAfmMethodResponse"`
	RgWsPublicBasicRtOut          AFMData
	ArrayOfRgWsPublicFirmActRtOut ArrRgWsPublicFirmActRtUser
	PCallSeqIDOut                 XMLPCallSeqIDOut
	PErrorRecOut                  XMLPErrorRecOut
}

// XMLRGWSPublicVersionInfoResponse holds version info
type XMLRGWSPublicVersionInfoResponse struct {
	XMLName xml.Name `xml:"rgWsPublicVersionInfoResponse"`
	Result  string   `xml:"result"`
}

// AFMData is the data relative to an entity's VAT search
type AFMData struct {
	XMLName               xml.Name `xml:"RgWsPublicBasicRt_out" json:"-"`
	AFM                   string   `xml:"afm" json:"afm"`                                             // ΑΦΜ
	DOY                   string   `xml:"doy" json:"doy"`                                             // ΚΩΔΙΚΟΣ ΔΟΥ
	DOYDesc               string   `xml:"doyDescr" json:"doy_description"`                            // ΠΕΡΙΓΡΑΦΗ ΔΟΥ
	INiFlagDescr          string   `xml:"INiFlagDescr" json:"ini_flag_description"`                   // ΦΠ /ΜΗ ΦΠ
	DeactivationFlag      string   `xml:"deactivationFlag" json:"deactivation_flag"`                  // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ:1=ΕΝΕΡΓΟΣ ΑΦΜ 2=ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	DeactivationFlagDescr string   `xml:"deactivationFlagDescr" json:"deactivation_flag_description"` // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ(ΠΕΡΙΓΡΑΦΗ): ΕΝΕΡΓΟΣ ΑΦΜ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	FirmFlagDescr         string   `xml:"firmFlagDescr" json:"firm_flag_description"`                 // ΤΙΜΕΣ: ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΜΗ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΠΡΩΗΝ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ
	Onomasia              string   `xml:"onomasia" json:"onomasia"`                                   // ΕΠΩΝΥΜΙΑ
	CommerTitle           string   `xml:"commerTitle" json:"commercial_title"`                        // ΤΙΤΛΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	LegalStatusDescr      string   `xml:"legalStatusDescr" json:"legal_status_descr"`                 // ΠΕΡΙΓΡΑΦΗ ΜΟΡΦΗΣ ΜΗ Φ.Π.
	PostalAddress         string   `xml:"postalAddress" json:"postal_address"`                        // ΟΔΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAddressNo       string   `xml:"postalAddressNo" json:"postal_address_no"`                   // ΑΡΙΘΜΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalZipCode         string   `xml:"postalZipCode" json:"postal_zip_code"`                       // ΤΑΧ. ΚΩΔ. ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAreaDescription string   `xml:"postalAreaDescription" json:"postal_area_description"`       // ΠΕΡΙΟΧΗ ΕΠΙΧΕΙΡΗΣΗΣ
	RegistDate            string   `xml:"registDate" json:"registration_date"`                        // ΗΜ/ΝΙΑ ΕΝΑΡΞΗΣ
	StopDate              string   `xml:"stopDate" json:"stop_date"`                                  // ΗΜ/ΝΙΑ ΔΙΑΚΟΠΗΣ
	Activities            []FirmActivities
	Error                 XMLPErrorRecOut
}

// ArrRgWsPublicFirmActRtUser holds an array of the entity's actitivies (ΚΑΔ)
type ArrRgWsPublicFirmActRtUser struct {
	XMLName                 xml.Name `xml:"arrayOfRgWsPublicFirmActRt_out"`
	RgWsPublicFirmActRtUser []FirmActivities
}

// FirmActivities is the activities of the entity
type FirmActivities struct {
	XMLName          xml.Name `xml:"RgWsPublicFirmActRtUser" json:"-"`
	FirmActCode      string   `xml:"firmActCode" json:"firm_activity_code"`                  // ΚΩΔΙΚΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActDescr     string   `xml:"firmActDescr" json:"firm_activity_description"`          // ΠΕΡΙΓΡΑΦΗ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActKind      string   `xml:"firmActKind" json:"firm_activity_kind"`                  // ΕΙΔΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: 1=ΚΥΡΙΑ, 2=ΔΕΥΤΕΡΕΥΟΥΣΑ, 3=ΛΟΙΠΗ, 4=ΒΟΗΘΗΤΙΚΗ
	FirmActKindDescr string   `xml:"firmActKindDescr" json:"firm_activity_kind_description"` // ΠΕΡΙΓΡΑΦΗ ΕΙΔΟΥΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: ΚΥΡΙΑ, ΔΕΥΤΕΡΕΥΟΥΣΑ, ΛΟΙΠΗ, ΒΟΗΘΗΤΙΚΗ
}

// XMLPCallSeqIDOut is a number returned with a successful service call
type XMLPCallSeqIDOut struct {
	XMLName xml.Name `xml:"pCallSeqId_out"`
}

// XMLPErrorRecOut holds error info
type XMLPErrorRecOut struct {
	XMLName    xml.Name `xml:"pErrorRec_out" json:"-"`
	ErrorDescr string   `xml:"errorDescr" json:"error_description"`
	ErrorCode  string   `xml:"errorCode" json:"error_code"`
}

const (
	// Endpoint is the url for WSDL service
	Endpoint                                       = "https://www1.gsis.gr/webtax2/wsgsis/RgWsPublic/RgWsPublicPort"
	RG_WS_PUBLIC_AFM_CALLED_BY_BLOCKED             = "Ο χρήστης που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_AFM_CALLED_BY_NOT_FOUND           = "Ο Α.Φ.Μ. για τον οποίο γίνεται η κλήση δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_EPIT_NF                           = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν ανήκει και δεν ανήκε ποτέ σε νομικό πρόσωπο, νομική οντότητα, ή φυσικό πρόσωπο με εισόδημα από επιχειρηματική δραστηριότητα."
	RG_WS_PUBLIC_FAILURES_TOLERATED_EXCEEDED       = "Υπέρβαση μέγιστου επιτρεπτού ορίου πρόσφατων αποτυχημένων κλήσεων. Προσπαθήστε εκ νέου σε μερικές ώρες."
	RG_WS_PUBLIC_MAX_DAILY_USERNAME_CALLS_EXCEEDED = "Υπέρβαση μέγιστου επιτρεπτού ορίου ημερήσιων κλήσεων ανά χρήστη (ανεξαρτήτως εξουσιοδοτήσεων)."
	RG_WS_PUBLIC_MONTHLY_LIMIT_EXCEEDED            = "Υπέρβαση του Μέγιστου Επιτρεπτού Μηνιαίου Ορίου Κλήσεων."
	RG_WS_PUBLIC_MSG_TO_TAXISNET_ERROR             = "Δημιουργήθηκε πρόβλημα κατά την ενημέρωση των εισερχόμενων μηνυμάτων στο MyTAXISnet."
	RG_WS_PUBLIC_NO_INPUT_PARAMETERS               = "Δε δόθηκαν υποχρεωτικές παράμετροι εισόδου για την κλήση της υπηρεσίας."
	RG_WS_PUBLIC_SERVICE_NOT_ACTIVE                = "Η υπηρεσία δεν είναι ενεργή."
	RG_WS_PUBLIC_TAXPAYER_NF                       = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_BLOCKED                 = "Ο χρήστης (ή ο εξουσιοδοτημένος τρίτος) που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_AUTHORIZED          = "Ο τρέχον χρήστης δεν έχει εξουσιοδοτηθεί από τον Α.Φ.Μ. για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_FOUND               = "Ο Α.Φ.Μ. του τρέχοντος χρήστη δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_REGISTERED          = "Ο τρέχον χρήστης δεν έχει εγγραφεί για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_ACTIVE         = "Ο κωδικός χρήστη (username) που χρησιμοποιήθηκε έχει ανακληθεί."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED  = "Ο συνδυασμός χρήστη/κωδικού πρόσβασης που δόθηκε δεν είναι έγκυρος."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_DEFINED        = "Δεν ορίσθηκε ο χρήστης που καλεί την υπηρεσία."
	RG_WS_PUBLIC_TOKEN_USERNAME_TOO_LONG           = "Διαπιστώθηκε υπέρβαση του μήκους του ονόματος του χρήστη (username) της υπηρεσίας"
	RG_WS_PUBLIC_WRONG_AFM                         = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν είναι έγκυρος."
)

func (a *AFMData) JSON() (string, error) {

	var j []byte
	j, err := json.MarshalIndent(&a, "", "\t")
	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (a *AFMData) String() string {

	var s string

	s = fmt.Sprintf("XMLName:%s\n", a.XMLName)
	s += fmt.Sprintf("AFM:%s\n", a.AFM)
	s += fmt.Sprintf("DOY:%s\n", a.DOY)
	s += fmt.Sprintf("DOYDesc:%s\n", a.DOYDesc)
	s += fmt.Sprintf("INiFlagDescr:%s\n", a.INiFlagDescr)
	s += fmt.Sprintf("DeactivationFlag:%s\n", a.DeactivationFlag)
	s += fmt.Sprintf("DeactivationFlagDescr:%s\n", a.DeactivationFlagDescr)
	s += fmt.Sprintf("FirmFlagDescr:%s\n", a.FirmFlagDescr)
	s += fmt.Sprintf("Onomasia:%s\n", a.Onomasia)
	s += fmt.Sprintf("CommerTitle:%s\n", a.CommerTitle)
	s += fmt.Sprintf("LegalStatusDescr:%s\n", a.LegalStatusDescr)
	s += fmt.Sprintf("PostalAddress:%s\n", a.PostalAddress)
	s += fmt.Sprintf("PostalAddressNo:%s\n", a.PostalAddressNo)
	s += fmt.Sprintf("PostalZipCode:%s\n", a.PostalZipCode)
	s += fmt.Sprintf("PostalAreaDescription:%s\n", a.PostalAreaDescription)
	s += fmt.Sprintf("RegistDate:%s\n", a.RegistDate)
	s += fmt.Sprintf("StopDate:%s\n", a.StopDate)

	s += fmt.Sprintf("ACTIVITIES:--------------------\n")
	for k, v := range a.Activities {
		s += fmt.Sprintf("ACTIVITY #%d\n", k)
		s += fmt.Sprintf("FirmActCode: %s\n", v.FirmActCode)
		s += fmt.Sprintf("FirmActDescr: %s\n", v.FirmActDescr)
		s += fmt.Sprintf("FirmActKind: %s\n", v.FirmActKind)
		s += fmt.Sprintf("FirmActKindDescr: %s\n", v.FirmActKindDescr)
	}

	s += fmt.Sprintf("ErrorDescr: %s\n", a.Error.ErrorDescr)
	s += fmt.Sprintf("ErrorCode: %s\n", a.Error.ErrorCode)

	return s

}
