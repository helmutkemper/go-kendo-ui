package telerik

import (
	"bytes"
	"reflect"
)

//@see https://json-schema.org/
//@see https://tools.ietf.org/html/draft-handrews-json-schema-validation-01
//@see https://stackoverflow.com/questions/38717933/jsonschema-attribute-conditionally-required
//@see http://json-schema.org/draft-07/json-schema-release-notes.html
//@see https://cswr.github.io/JsonSchema/spec/generic_keywords/

// <input> elements of type "text" create basic single-line text fields.
//
// <input type="text">
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/text
type HtmlInputText struct {
	Apagar string `htmlAttr:"apagar" jsonSchema_complex:"{\"animation\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{\"close\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"object\"}],\"properties\":{\"duration\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"effects\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}}}},\"open\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"object\"}],\"properties\":{\"duration\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"effects\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}}}}}},\"autoWidth\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"clearButton\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"dataSource\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"object\"},{\"type\":\"array\"},{\"type\":\"kendo.data.datasource\"}],\"properties\":{}},\"dataTextField\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}},\"delay\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"enable\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"enforceMinLength\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"filter\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}},\"fixedGroupTemplate\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"footerTemplate\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"groupTemplate\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"headerTemplate\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"height\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"highlightFirst\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"ignoreCase\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"minLength\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"noDataTemplate\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"placeholder\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}},\"popup\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"object\"}],\"properties\":{}},\"separator\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"array\"}],\"properties\":{}},\"suggest\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"template\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"},{\"type\":\"function\"}],\"properties\":{}},\"value\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}},\"valuePrimitive\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{}},\"virtual\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"boolean\"}],\"properties\":{\"itemHeight\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"number\"}],\"properties\":{}},\"mapValueTo\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"string\"}],\"properties\":{}},\"valueMapper\":{\"description\":\"olá mundo!\",\"oneOf\":[{\"type\":\"function\"}],\"properties\":{}}}}}"`
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
	*/
	Name string `htmlAttr:"name" jsonSchema_complex:"{\"type\":\"string\",\"description\":\"The name of the control, which is submitted with the form data.\\n\\nIf you need to use the default values for the form's auto-completion feature, press ctrl+space to see the full list of values.\",\"oneOf\":[{\"pattern\":\"^(?!^name$|^honorific-prefix$|^given-name$|^additional-name$|^family-name$|^honorific-suffix$|^nickname$|^email$|^username$|^new-password$|^current-password$|^organization-title$|^organization$|^street-address$|^address-line1$|^address-line2$|^address-line3$|^address-level4$|^address-level3$|^address-level2$|^country$|^country-name$|^postal-code$|^cc-name$|^cc-given-name$|^cc-additional-name$|^cc-family-name$|^cc-number$|^cc-exp$|^cc-exp-month$|^cc-exp-year$|^cc-csc$|^cc-type$|^transaction-currency$|^transaction-amount$|^language$|^bday$|^bday-day$|^bday-month$|^bday-year$|^sex$|^tel$|^tel-country-code$|^tel-national$|^tel-area-code$|^tel-local$|^tel-local-prefix$|^tel-local-suffix$|^tel-extension$|^url$|^photo$)[a-zA-Z0-9_]+([[0-9]*])*$\",\"description\":\"html tag name\"},{\"const\":\"name\",\"description\":\"Default value for form autocomplete:\\n\\nName\"},{\"const\":\"honorific-prefix\",\"description\":\"Default value for form autocomplete:\\n\\nPrefix or title (e.g. \\\"Mr.\\\", \\\"Ms.\\\", \\\"Dr.\\\", \\\"Mlle\\\").\"},{\"const\":\"given-name\",\"description\":\"Default value for form autocomplete:\\n\\nFirst name.\"},{\"const\":\"additional-name\",\"description\":\"Default value for form autocomplete:\\n\\nMiddle name.\"},{\"const\":\"family-name\",\"description\":\"Default value for form autocomplete:\\n\\nLast name.\"},{\"const\":\"honorific-suffix\",\"description\":\"Default value for form autocomplete:\\n\\nSuffix (e.g. \\\"Jr.\\\", \\\"B.Sc.\\\", \\\"MBASW\\\", \\\"II\\\").\"},{\"const\":\"nickname\",\"description\":\"Default value for form autocomplete:\\n\\nNickname\"},{\"const\":\"email\",\"description\":\"Default value for form autocomplete:\\n\\nE-mail\"},{\"const\":\"username\",\"description\":\"Default value for form autocomplete:\\n\\nUser name\"},{\"const\":\"new-password\",\"description\":\"Default value for form autocomplete:\\n\\nA new password (e.g. when creating an account or changing a password).\"},{\"const\":\"current-password\",\"description\":\"Default value for form autocomplete:\\n\\nCurrent password\"},{\"const\":\"organization-title\",\"description\":\"Default value for form autocomplete:\\n\\nOrganization title\"},{\"const\":\"organization\",\"description\":\"Default value for form autocomplete:\\n\\nOrganization\"},{\"const\":\"street-address\",\"description\":\"Default value for form autocomplete:\\n\\nStreet address\"},{\"const\":\"address-line1\",\"description\":\"Default value for form autocomplete:\\n\\nAddress line1\"},{\"const\":\"address-line2\",\"description\":\"Default value for form autocomplete:\\n\\nAddress line2\"},{\"const\":\"address-line3\",\"description\":\"Default value for form autocomplete:\\n\\nAddress line3\"},{\"const\":\"address-level4\",\"description\":\"Default value for form autocomplete:\\n\\nAddress level4\"},{\"const\":\"address-level3\",\"description\":\"Default value for form autocomplete:\\n\\nAddress level3\"},{\"const\":\"address-level2\",\"description\":\"Default value for form autocomplete:\\n\\nAddress level2\"},{\"const\":\"address-level1\",\"description\":\"Default value for form autocomplete:\\n\\nAddress level1\"},{\"const\":\"country\",\"description\":\"Default value for form autocomplete:\\n\\nCountry\"},{\"const\":\"country-name\",\"description\":\"Default value for form autocomplete:\\n\\nCountry name\"},{\"const\":\"postal-code\",\"description\":\"Default value for form autocomplete:\\n\\nPostal code\"},{\"const\":\"cc-name\",\"description\":\"Default value for form autocomplete:\\n\\nFull name as given on the payment instrument.\"},{\"const\":\"cc-given-name\",\"description\":\"Default value for form autocomplete:\\n\\nFirst name as given on the payment instrument.\"},{\"const\":\"cc-additional-name\",\"description\":\"Default value for form autocomplete:\\n\\nMiddle name as given on the payment instrument.\"},{\"const\":\"cc-family-name\",\"description\":\"Default value for form autocomplete:\\n\\nLast name as given on the payment instrument.\"},{\"const\":\"cc-number\",\"description\":\"Default value for form autocomplete:\\n\\nCode identifying the payment instrument (e.g. the credit card number).\"},{\"const\":\"cc-exp\",\"description\":\"Default value for form autocomplete:\\n\\nExpiration date of the payment instrument.\"},{\"const\":\"cc-exp-month\",\"description\":\"Default value for form autocomplete:\\n\\nExpiration month of the payment instrument.\"},{\"const\":\"cc-exp-year\",\"description\":\"Default value for form autocomplete:\\n\\nExpiration year of the payment instrument.\"},{\"const\":\"cc-csc\",\"description\":\"Default value for form autocomplete:\\n\\nSecurity code for the payment instrument.\"},{\"const\":\"cc-type\",\"description\":\"Default value for form autocomplete:\\n\\nType of payment instrument (e.g. Visa).\"},{\"const\":\"transaction-currency\",\"description\":\"Default value for form autocomplete:\\n\\nTransaction currency for the payment instrument.\"},{\"const\":\"transaction-amount\",\"description\":\"Default value for form autocomplete:\\n\\nTransaction amount for the payment instrument.\"},{\"const\":\"language\",\"description\":\"Default value for form autocomplete:\\n\\nPreferred language; a valid BCP 47 language tag.\"},{\"const\":\"bday\",\"description\":\"Default value for form autocomplete:\\n\\nBirthday\"},{\"const\":\"bday-day\",\"description\":\"Default value for form autocomplete:\\n\\nDay for birthday\"},{\"const\":\"bday-month\",\"description\":\"Default value for form autocomplete:\\n\\nMonth for birthday\"},{\"const\":\"bday-year\",\"description\":\"Default value for form autocomplete:\\n\\nYear for birthday\"},{\"const\":\"sex\",\"description\":\"Default value for form autocomplete:\\n\\nGender identity (e.g. Female, Fa'afafine), free-form text, no newlines.\"},{\"const\":\"tel\",\"description\":\"Default value for form autocomplete:\\n\\nFull telephone number, including country code\"},{\"const\":\"tel-country-code\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone country code\"},{\"const\":\"tel-national\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone national code\"},{\"const\":\"tel-area-code\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone area code\"},{\"const\":\"tel-local\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone local code\"},{\"const\":\"tel-local-prefix\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone prefix\"},{\"const\":\"tel-local-suffix\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone suffix\"},{\"const\":\"tel-extension\",\"description\":\"Default value for form autocomplete:\\n\\nTelephone extension\"},{\"const\":\"url\",\"description\":\"Default value for form autocomplete:\\n\\nHome page or other Web page corresponding to the company, person, address, or contact information in the other fields associated with this field.\"},{\"const\":\"photo\",\"description\":\"Default value for form autocomplete:\\n\\nPhotograph, icon, or other image corresponding to the company, person, address, or contact information in the other fields associated with this field.\"}]}"`

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string `htmlAttr:"value" jsonSchema_description:"The initial value of the control. This attribute is optional except when the value of the type attribute is radio or checkbox.\nNote that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was changed before the reload."`

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `htmlAttr:"form" jsonSchema_description:"The form element that the input element is associated with (its form owner). The value of the attribute must be an id of a <form> element in the same document. If this attribute is not specified, this <input> element must be a descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not just as descendants of their form elements. An input can only be associated with one form."`

	/*
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean `htmlAttrSet:"disabled" jsonSchema_description:"This Boolean attribute indicates that the form control is not available for interaction. In particular, the click event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.\nUnlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use the autocomplete attribute to control this feature."`

	/*
	  This attribute indicates whether the value of the control can be automatically completed by the browser.
	  Possible values are:
	  off: The user must explicitly enter a value into this field for every use, or the document provides its own
	  auto-completion method. The browser does not automatically complete the entry.
	  on: The browser is allowed to automatically complete the value based on values that the user has entered during
	  previous uses, however on does not provide any further information about what kind of data the user might be expected
	  to enter.
	  @see typeNamesForAutocomplete.go
	*/
	AutoComplete Boolean `htmlAttrOnOff:"autocomplete" jsonSchema_description:"This attribute indicates whether the value of the control can be automatically completed by the browser.\nPossible values are:\n off: The user must explicitly enter a value into this field for every use, or the document provides its own auto-completion method. The browser does not automatically complete the entry.\non: The browser is allowed to automatically complete the value based on values that the user has entered during previous uses, however on does not provide any further information about what kind of data the user might be expected to enter."`

	/*
	  Identifies a list of pre-defined options to suggest to the user. The value must be the id of a <datalist> element in
	  the same document. The browser displays only options that are valid values for this input element. This attribute is
	  ignored when the type attribute's value is hidden, checkbox, radio, file, or a button type.
	*/
	List string `htmlAttr:"list" jsonSchema_description:"Identifies a list of pre-defined options to suggest to the user. The value must be the id of a <datalist> element in the same document. The browser displays only options that are valid values for this input element. This attribute is ignored when the type attribute's value is hidden, checkbox, radio, file, or a button type."`

	/*
	  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the maximum
	  number of characters (in UTF-16 code units) that the user can enter. For other control types, it is ignored. It can
	  exceed the value of the size attribute. If it is not specified, the user can enter an unlimited number of characters.
	  Specifying a negative number results in the default behavior (i.e. the user can enter an unlimited number of
	  characters). The constraint is evaluated only when the value of the attribute has been changed.
	*/
	MaxLength int `htmlAttr:"maxlength" jsonSchema_description:"If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the maximum number of characters (in UTF-16 code units) that the user can enter. For other control types, it is ignored. It can exceed the value of the size attribute. If it is not specified, the user can enter an unlimited number of characters.\nSpecifying a negative number results in the default behavior (i.e. the user can enter an unlimited number of characters). The constraint is evaluated only when the value of the attribute has been changed."`

	/*
	  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the minimum
	  number of characters (in Unicode code points) that the user can enter. For other control types, it is ignored.
	*/
	MinLength int `htmlAttr:"minlength" jsonSchema_description:"If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the minimum number of characters (in Unicode code points) that the user can enter. For other control types, it is ignored."`

	/*
	  A regular expression that the control's value is checked against. The pattern must match the entire value, not just
	  some subset. Use the title attribute to describe the pattern to help the user. This attribute applies when the value
	  of the type attribute is text, search, tel, url, email, or password, otherwise it is ignored. The regular expression
	  language is the same as JavaScript RegExp algorithm, with the 'u' parameter that makes it treat the pattern as a
	  sequence of unicode code points. The pattern is not surrounded by forward slashes.
	*/
	Pattern string `htmlAttr:"pattern" jsonSchema_description:"A regular expression that the control's value is checked against. The pattern must match the entire value, not just some subset. Use the title attribute to describe the pattern to help the user. This attribute applies when the value of the type attribute is text, search, tel, url, email, or password, otherwise it is ignored. The regular expression language is the same as JavaScript RegExp algorithm, with the 'u' parameter that makes it treat the pattern as a sequence of unicode code points. The pattern is not surrounded by forward slashes."`

	/*
	  A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or
	  line-feeds.
	*/
	PlaceHolder string `htmlAttr:"placeholder" jsonSchema_description:"A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or line-feeds."`

	/*
	  This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type
	  attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS
	  pseudo-classes will be applied to the field as appropriate.
	*/
	Required Boolean `htmlAttrSet:"required" jsonSchema_description:"This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS pseudo-classes will be applied to the field as appropriate."`

	/*
	  This attribute indicates that the user cannot modify the value of the control. The value of the attribute is
	  irrelevant. If you need read-write access to the input value, do not add the "readonly" attribute. It is ignored if
	  the value of the type attribute is hidden, range, color, checkbox, radio, file, or a button type (such as button or
	  submit).
	*/
	Readonly Boolean `htmlAttrSet:"readonly" jsonSchema_description:"This attribute indicates that the user cannot modify the value of the control. The value of the attribute is irrelevant. If you need read-write access to the input value, do not add the 'readonly' attribute. It is ignored if the value of the type attribute is hidden, range, color, checkbox, radio, file, or a button type (such as button or submit)."`

	/*
	  The initial size of the control. This value is in pixels unless the value of the type attribute is text or password,
	  in which case it is an integer number of characters. Starting in HTML5, this attribute applies only when the type
	  attribute is set to text, search, tel, url, email, or password, otherwise it is ignored. In addition, the size must be
	  greater than zero. If you do not specify a size, a default value of 20 is used. HTML5 simply states "the user agent
	  should ensure that at least that many characters are visible", but different characters can have different widths in
	  certain fonts. In some browsers, a certain string with x characters will not be entirely visible even if size is
	  defined to at least x.
	*/
	Size int `htmlAttr:"size" jsonSchema_description:"The initial size of the control. This value is in pixels unless the value of the type attribute is text or password, in which case it is an integer number of characters. Starting in HTML5, this attribute applies only when the type attribute is set to text, search, tel, url, email, or password, otherwise it is ignored. In addition, the size must be greater than zero. If you do not specify a size, a default value of 20 is used. HTML5 simply states 'the user agent should ensure that at least that many characters are visible', but different characters can have different widths in certain fonts. In some browsers, a certain string with x characters will not be entirely visible even if size is defined to at least x."`

	CheckCode string `htmlAttr:"checkcode" jsonSchema_description:""`

	Global HtmlGlobalAttributes `htmlAttr:"-" jsonSchema_description:"" jsonSchema_keyNewName:"global"`

	*ToJavaScriptConverter `htmlAttr:"-" jsonSchema_description:"-"`
}

func (el *HtmlInputText) ToJSonSchema() []byte {
	var buffer bytes.Buffer

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToSJsonSchema("HtmlInputText", "object", element)

	buffer.Write(data)

	return buffer.Bytes()
}
func (el *HtmlInputText) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<input type="text"`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))

	return buffer.Bytes()
}
func (el *HtmlInputText) GetId() []byte {
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}
	return []byte(el.Global.Id)
}
func (el *HtmlInputText) GetName() []byte {
	if el.Name == "" {
		el.Name = GetAutoId()
	}
	return []byte(el.Name)
}
func (el *HtmlInputText) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}

	ret.Write([]byte(`$("#` + el.Global.Id + `").addClass('k-textbox');`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
