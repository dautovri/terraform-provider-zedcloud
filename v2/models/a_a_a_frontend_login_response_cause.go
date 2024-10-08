// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AAAFrontendLoginResponseCause - EXCEPTION: Some exception has occurred on the server
//
// swagger:model AAA_Frontend_LoginResponseCause
type AAAFrontendLoginResponseCause string

func NewAAAFrontendLoginResponseCause(value AAAFrontendLoginResponseCause) *AAAFrontendLoginResponseCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFrontendLoginResponseCause.
func (m AAAFrontendLoginResponseCause) Pointer() *AAAFrontendLoginResponseCause {
	return &m
}

const (

	// AAAFrontendLoginResponseCauseUNSPECIFIED captures enum value "UNSPECIFIED"
	AAAFrontendLoginResponseCauseUNSPECIFIED AAAFrontendLoginResponseCause = "UNSPECIFIED"

	// AAAFrontendLoginResponseCauseOK captures enum value "OK"
	AAAFrontendLoginResponseCauseOK AAAFrontendLoginResponseCause = "OK"

	// AAAFrontendLoginResponseCauseUNKNOWN captures enum value "UNKNOWN"
	AAAFrontendLoginResponseCauseUNKNOWN AAAFrontendLoginResponseCause = "UNKNOWN"

	// AAAFrontendLoginResponseCauseCREDENTIALS captures enum value "CREDENTIALS"
	AAAFrontendLoginResponseCauseCREDENTIALS AAAFrontendLoginResponseCause = "CREDENTIALS"

	// AAAFrontendLoginResponseCauseSUSPENDED captures enum value "SUSPENDED"
	AAAFrontendLoginResponseCauseSUSPENDED AAAFrontendLoginResponseCause = "SUSPENDED"

	// AAAFrontendLoginResponseCauseEXCEPTION captures enum value "EXCEPTION"
	AAAFrontendLoginResponseCauseEXCEPTION AAAFrontendLoginResponseCause = "EXCEPTION"

	// AAAFrontendLoginResponseCauseINACTIVE captures enum value "INACTIVE"
	AAAFrontendLoginResponseCauseINACTIVE AAAFrontendLoginResponseCause = "INACTIVE"

	// AAAFrontendLoginResponseCauseSIGNEDUPSTATE captures enum value "SIGNEDUPSTATE"
	AAAFrontendLoginResponseCauseSIGNEDUPSTATE AAAFrontendLoginResponseCause = "SIGNEDUPSTATE"

	// AAAFrontendLoginResponseCauseUNKNOWNSTATE captures enum value "UNKNOWNSTATE"
	AAAFrontendLoginResponseCauseUNKNOWNSTATE AAAFrontendLoginResponseCause = "UNKNOWNSTATE"

	// AAAFrontendLoginResponseCauseCREATEDSTATE captures enum value "CREATEDSTATE"
	AAAFrontendLoginResponseCauseCREATEDSTATE AAAFrontendLoginResponseCause = "CREATEDSTATE"

	// AAAFrontendLoginResponseCauseUSERUNKNOWN captures enum value "USER_UNKNOWN"
	AAAFrontendLoginResponseCauseUSERUNKNOWN AAAFrontendLoginResponseCause = "USER_UNKNOWN"

	// AAAFrontendLoginResponseCauseENTERPRISEUNKNOWN captures enum value "ENTERPRISE_UNKNOWN"
	AAAFrontendLoginResponseCauseENTERPRISEUNKNOWN AAAFrontendLoginResponseCause = "ENTERPRISE_UNKNOWN"

	// AAAFrontendLoginResponseCauseROLEUNKNOWN captures enum value "ROLE_UNKNOWN"
	AAAFrontendLoginResponseCauseROLEUNKNOWN AAAFrontendLoginResponseCause = "ROLE_UNKNOWN"

	// AAAFrontendLoginResponseCauseUSERUNKNOWNSTATE captures enum value "USER_UNKNOWNSTATE"
	AAAFrontendLoginResponseCauseUSERUNKNOWNSTATE AAAFrontendLoginResponseCause = "USER_UNKNOWNSTATE"

	// AAAFrontendLoginResponseCauseUSERINACTIVE captures enum value "USER_INACTIVE"
	AAAFrontendLoginResponseCauseUSERINACTIVE AAAFrontendLoginResponseCause = "USER_INACTIVE"

	// AAAFrontendLoginResponseCauseUSERSIGNEDUPSTATE captures enum value "USER_SIGNEDUPSTATE"
	AAAFrontendLoginResponseCauseUSERSIGNEDUPSTATE AAAFrontendLoginResponseCause = "USER_SIGNEDUPSTATE"

	// AAAFrontendLoginResponseCauseUSERCREATEDSTATE captures enum value "USER_CREATEDSTATE"
	AAAFrontendLoginResponseCauseUSERCREATEDSTATE AAAFrontendLoginResponseCause = "USER_CREATEDSTATE"

	// AAAFrontendLoginResponseCauseENTERPRISEUNKNOWNSTATE captures enum value "ENTERPRISE_UNKNOWNSTATE"
	AAAFrontendLoginResponseCauseENTERPRISEUNKNOWNSTATE AAAFrontendLoginResponseCause = "ENTERPRISE_UNKNOWNSTATE"

	// AAAFrontendLoginResponseCauseENTERPRISEINACTIVE captures enum value "ENTERPRISE_INACTIVE"
	AAAFrontendLoginResponseCauseENTERPRISEINACTIVE AAAFrontendLoginResponseCause = "ENTERPRISE_INACTIVE"

	// AAAFrontendLoginResponseCauseENTERPRISESIGNEDUPSTATE captures enum value "ENTERPRISE_SIGNEDUPSTATE"
	AAAFrontendLoginResponseCauseENTERPRISESIGNEDUPSTATE AAAFrontendLoginResponseCause = "ENTERPRISE_SIGNEDUPSTATE"

	// AAAFrontendLoginResponseCauseENTERPRISECREATEDSTATE captures enum value "ENTERPRISE_CREATEDSTATE"
	AAAFrontendLoginResponseCauseENTERPRISECREATEDSTATE AAAFrontendLoginResponseCause = "ENTERPRISE_CREATEDSTATE"

	// AAAFrontendLoginResponseCauseCREDENTIALNOTFOUND captures enum value "CREDENTIAL_NOTFOUND"
	AAAFrontendLoginResponseCauseCREDENTIALNOTFOUND AAAFrontendLoginResponseCause = "CREDENTIAL_NOTFOUND"

	// AAAFrontendLoginResponseCauseCREDENTIALMISMATCH captures enum value "CREDENTIAL_MISMATCH"
	AAAFrontendLoginResponseCauseCREDENTIALMISMATCH AAAFrontendLoginResponseCause = "CREDENTIAL_MISMATCH"

	// AAAFrontendLoginResponseCauseSCHEMEUNKNOWN captures enum value "SCHEME_UNKNOWN"
	AAAFrontendLoginResponseCauseSCHEMEUNKNOWN AAAFrontendLoginResponseCause = "SCHEME_UNKNOWN"

	// AAAFrontendLoginResponseCauseUPDATELOGINTIMEFAILED captures enum value "UPDATE_LOGINTIME_FAILED"
	AAAFrontendLoginResponseCauseUPDATELOGINTIMEFAILED AAAFrontendLoginResponseCause = "UPDATE_LOGINTIME_FAILED"

	// AAAFrontendLoginResponseCauseCREDENTIALMISMATCHMAXFAILEDPWDATTEMPT captures enum value "CREDENTIAL_MISMATCH_MAX_FAILED_PWD_ATTEMPT"
	AAAFrontendLoginResponseCauseCREDENTIALMISMATCHMAXFAILEDPWDATTEMPT AAAFrontendLoginResponseCause = "CREDENTIAL_MISMATCH_MAX_FAILED_PWD_ATTEMPT"

	// AAAFrontendLoginResponseCauseREDIRECT captures enum value "REDIRECT"
	AAAFrontendLoginResponseCauseREDIRECT AAAFrontendLoginResponseCause = "REDIRECT"

	// AAAFrontendLoginResponseCauseDECRYPTIONFAILED captures enum value "DECRYPTION_FAILED"
	AAAFrontendLoginResponseCauseDECRYPTIONFAILED AAAFrontendLoginResponseCause = "DECRYPTION_FAILED"

	// AAAFrontendLoginResponseCausePASSWORDEXPIRED captures enum value "PASSWORD_EXPIRED"
	AAAFrontendLoginResponseCausePASSWORDEXPIRED AAAFrontendLoginResponseCause = "PASSWORD_EXPIRED"

	// AAAFrontendLoginResponseCauseTOTPOK captures enum value "TOTP_OK"
	AAAFrontendLoginResponseCauseTOTPOK AAAFrontendLoginResponseCause = "TOTP_OK"

	// AAAFrontendLoginResponseCauseTOTPINVALIDCODE captures enum value "TOTP_INVALID_CODE"
	AAAFrontendLoginResponseCauseTOTPINVALIDCODE AAAFrontendLoginResponseCause = "TOTP_INVALID_CODE"

	// AAAFrontendLoginResponseCauseTOTPERR captures enum value "TOTP_ERR"
	AAAFrontendLoginResponseCauseTOTPERR AAAFrontendLoginResponseCause = "TOTP_ERR"
)

// for schema
var aAAFrontendLoginResponseCauseEnum []interface{}

func init() {
	var res []AAAFrontendLoginResponseCause
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","OK","UNKNOWN","CREDENTIALS","SUSPENDED","EXCEPTION","INACTIVE","SIGNEDUPSTATE","UNKNOWNSTATE","CREATEDSTATE","USER_UNKNOWN","ENTERPRISE_UNKNOWN","ROLE_UNKNOWN","USER_UNKNOWNSTATE","USER_INACTIVE","USER_SIGNEDUPSTATE","USER_CREATEDSTATE","ENTERPRISE_UNKNOWNSTATE","ENTERPRISE_INACTIVE","ENTERPRISE_SIGNEDUPSTATE","ENTERPRISE_CREATEDSTATE","CREDENTIAL_NOTFOUND","CREDENTIAL_MISMATCH","SCHEME_UNKNOWN","UPDATE_LOGINTIME_FAILED","CREDENTIAL_MISMATCH_MAX_FAILED_PWD_ATTEMPT","REDIRECT","DECRYPTION_FAILED","PASSWORD_EXPIRED","TOTP_OK","TOTP_INVALID_CODE","TOTP_ERR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFrontendLoginResponseCauseEnum = append(aAAFrontendLoginResponseCauseEnum, v)
	}
}

func (m AAAFrontendLoginResponseCause) validateAAAFrontendLoginResponseCauseEnum(path, location string, value AAAFrontendLoginResponseCause) error {
	if err := validate.EnumCase(path, location, value, aAAFrontendLoginResponseCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a frontend login response cause
func (m AAAFrontendLoginResponseCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFrontendLoginResponseCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a frontend login response cause based on context it is used
func (m AAAFrontendLoginResponseCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
