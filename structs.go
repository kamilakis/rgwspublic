package rgwspublic

import (
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
	XMLName               xml.Name `xml:"RgWsPublicBasicRt_out"`
	AFM                   string   `xml:"afm"`                   // ΑΦΜ
	DOY                   string   `xml:"doy"`                   // ΚΩΔΙΚΟΣ ΔΟΥ
	DOYDesc               string   `xml:"doyDescr"`              // ΠΕΡΙΓΡΑΦΗ ΔΟΥ
	INiFlagDescr          string   `xml:"INiFlagDescr"`          // ΦΠ /ΜΗ ΦΠ
	DeactivationFlag      string   `xml:"deactivationFlag"`      // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ:1=ΕΝΕΡΓΟΣ ΑΦΜ 2=ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	DeactivationFlagDescr string   `xml:"deactivationFlagDescr"` // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ(ΠΕΡΙΓΡΑΦΗ): ΕΝΕΡΓΟΣ ΑΦΜ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	FirmFlagDescr         string   `xml:"firmFlagDescr"`         // ΤΙΜΕΣ: ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΜΗ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΠΡΩΗΝ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ
	Onomasia              string   `xml:"onomasia"`              // ΕΠΩΝΥΜΙΑ
	CommerTitle           string   `xml:"commerTitle"`           // ΤΙΤΛΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	LegalStatusDescr      string   `xml:"legalStatusDescr"`      // ΠΕΡΙΓΡΑΦΗ ΜΟΡΦΗΣ ΜΗ Φ.Π.
	PostalAddress         string   `xml:"postalAddress"`         // ΟΔΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAddressNo       string   `xml:"postalAddressNo"`       // ΑΡΙΘΜΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalZipCode         string   `xml:"postalZipCode"`         // ΤΑΧ. ΚΩΔ. ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAreaDescription string   `xml:"postalAreaDescription"` // ΠΕΡΙΟΧΗ ΕΠΙΧΕΙΡΗΣΗΣ
	RegistDate            string   `xml:"registDate"`            // ΗΜ/ΝΙΑ ΕΝΑΡΞΗΣ
	StopDate              string   `xml:"stopDate"`              // ΗΜ/ΝΙΑ ΔΙΑΚΟΠΗΣ
	Activities            []FirmActivities
}

// ArrRgWsPublicFirmActRtUser holds an array of the entity's actitivies (ΚΑΔ)
type ArrRgWsPublicFirmActRtUser struct {
	XMLName                 xml.Name `xml:"arrayOfRgWsPublicFirmActRt_out"`
	RgWsPublicFirmActRtUser []FirmActivities
}

// FirmActivities is the activities of the entity
type FirmActivities struct {
	XMLName          xml.Name `xml:"RgWsPublicFirmActRtUser"`
	FirmActCode      string   `xml:"firmActDescr"`     // ΚΩΔΙΚΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActDescr     string   `xml:"firmActKind"`      // ΠΕΡΙΓΡΑΦΗ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActKind      string   `xml:"firmActKindDescr"` // ΕΙΔΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: 1=ΚΥΡΙΑ, 2=ΔΕΥΤΕΡΕΥΟΥΣΑ, 3=ΛΟΙΠΗ, 4=ΒΟΗΘΗΤΙΚΗ
	FirmActKindDescr string   `xml:"firmActCode"`      // ΠΕΡΙΓΡΑΦΗ ΕΙΔΟΥΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: ΚΥΡΙΑ, ΔΕΥΤΕΡΕΥΟΥΣΑ, ΛΟΙΠΗ, ΒΟΗΘΗΤΙΚΗ
}

// XMLPCallSeqIDOut is a number returned with a successful service call
type XMLPCallSeqIDOut struct {
	XMLName xml.Name `xml:"pCallSeqId_out"`
}

// XMLPErrorRecOut holds error info
type XMLPErrorRecOut struct {
	XMLName    xml.Name `xml:"pErrorRec_out"`
	ErrorDescr string   `xml:"errorDescr"`
	ErrorCode  string   `xml:"errorCode"`
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

func (a *AFMData) String() {
	fmt.Printf("XMLName:%s\n", a.XMLName)
	fmt.Printf("AFM:%s\n", a.AFM)
	fmt.Printf("DOY:%s\n", a.DOY)
	fmt.Printf("DOYDesc:%s\n", a.DOYDesc)
	fmt.Printf("INiFlagDescr:%s\n", a.INiFlagDescr)
	fmt.Printf("DeactivationFlag:%s\n", a.DeactivationFlag)
	fmt.Printf("DeactivationFlagDescr:%s\n", a.DeactivationFlagDescr)
	fmt.Printf("FirmFlagDescr:%s\n", a.FirmFlagDescr)
	fmt.Printf("Onomasia:%s\n", a.Onomasia)
	fmt.Printf("CommerTitle:%s\n", a.CommerTitle)
	fmt.Printf("LegalStatusDescr:%s\n", a.LegalStatusDescr)
	fmt.Printf("PostalAddress:%s\n", a.PostalAddress)
	fmt.Printf("PostalAddressNo:%s\n", a.PostalAddressNo)
	fmt.Printf("PostalZipCode:%s\n", a.PostalZipCode)
	fmt.Printf("PostalAreaDescription:%s\n", a.PostalAreaDescription)
	fmt.Printf("RegistDate:%s\n", a.RegistDate)
	fmt.Printf("StopDate:%s\n", a.StopDate)

	fmt.Println("ACTIVITIES:--------------------")
	for k, v := range a.Activities {
		fmt.Printf("ACTIVITY #%d\n", k)
		fmt.Printf("FirmActCode: %s\n", v.FirmActCode)
		fmt.Printf("FirmActDescr: %s\n", v.FirmActDescr)
		fmt.Printf("FirmActKind: %s\n", v.FirmActKind)
		fmt.Printf("FirmActKindDescr: %s\n", v.FirmActKindDescr)
	}

}
