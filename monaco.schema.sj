// Configures two JSON schemas, with references.

var jsonCode = [
  '{',
  '    ',
  "}"
].join('\n');
var modelUri = monaco.Uri.parse("a://b/foo.json"); // a made up unique URI for our model
var model = monaco.editor.createModel(jsonCode, "json", modelUri);

// configure the JSON language support with schemas and schema associations
monaco.languages.json.jsonDefaults.setDiagnosticsOptions({
  validate: true,
  schemas: [{
    uri: "http://myserver/foo-schema.json", // id of the first schema
    fileMatch: [modelUri.toString()], // associate with our model
    schema: {
      $id: "https://example.com/arrays.schema.json",
      $schema: "http://json-schema.org/draft-07/schema#",
      description: "A representation of a person, company, organization, or place",
      type: "object",
      properties: {
        fruits: {
          type: "array",
          items: {
            type: "string"
          }
        },
        vegetables: {
          type: "array",
          description: "Acompanhamento de reps",
          items: {
            $ref: "#/definitions/veggie"
          }
        },
        form: {
          $ref: "#/definitions/formConfig",
          description: "Acompanhamento de reps"
        }
      },
      definitions: {

        HtmlInputText: {
          type: "object",
          properties: {
            name: {
              type: "string",
              pattern: "^$|^[a-zA-Z0-9_]+$",
              description: "The name of the control, which is submitted with the form data.",
            },
            value: {
              type: "string",
              pattern: "^$|^[a-zA-Z0-9_]+$",
              description: "The initial value of the control. This attribute is optional except when the value of the type attribute is radio or checkbox.\nNote that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was changed before the reload.",
            }
          }
        },
        HtmlInputText: {
          type: "object",
          properties: {
            name: {
              type: "string",
              pattern: "^$|^[a-zA-Z0-9_]+$",
              description: "The name of the control, which is submitted with the form data.",
            },
            value: {
              type: "string",
              pattern: "^$|^[a-zA-Z0-9_]+$",
              description: "The initial value of the control. This attribute is optional except when the value of the type attribute is radio or checkbox.\nNote that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was changed before the reload.",
            }
          }
        },
        HtmlInputTime: {
          type: "object",
          required: [
            "id"
          ],
          properties: {
            accesskey: {
              type: "string",
              pattern: "^$|^[a-zA-Z0-9_]+$",
              description: "Provides a hint for generating a keyboard shortcut for the current element.\nThis attribute consists of a space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.",
            },
          }
        },

        "test":{
          "type": "object",
          "properties": {
            "name": {
              "type": "string",
              description: "The name of the control, which is submitted with the form data.\nIf you need to use the default values for the form's auto-completion feature, press ctrl + space to see the full list of values.",
              "oneOf": [
                {
                  "pattern": "^[a-zA-Z0-9_]+(\\[[0-9]*\\])*$",
                  description:"html tag name"
                },
                {
                  const: "name",
                  description: "Default value for form autocomplete:\n\nName"
                },
                {
                  const: "honorific-prefix",
                  description: "Default value for form autocomplete:\n\nPrefix or title (e.g. \"Mr.\", \"Ms.\", \"Dr.\", \"Mlle\")."
                },
                {
                  const: "given-name",
                  description: "Default value for form autocomplete:\n\nFirst name."
                },
                {
                  const: "additional-name",
                  description: "Default value for form autocomplete:\n\nMiddle name."
                },
                {
                  const: "family-name",
                  description: "Default value for form autocomplete:\n\nLast name."
                },
                {
                  const: "honorific-suffix",
                  description: "Default value for form autocomplete:\n\nSuffix (e.g. \"Jr.\", \"B.Sc.\", \"MBASW\", \"II\")."
                },
                {
                  const: "nickname",
                  description: "Default value for form autocomplete:\n\nNickname"
                },
                {
                  const: "email",
                  description: "Default value for form autocomplete:\n\nE-mail"
                },
                {
                  const: "username",
                  description: "Default value for form autocomplete:\n\nUser name"
                },
                {
                  const: "new-password",
                  description: "Default value for form autocomplete:\n\nA new password (e.g. when creating an account or changing a password)."
                },
                {
                  const: "current-password",
                  description: "Default value for form autocomplete:\n\nCurrent password"
                },
                {
                  const: "organization-title",
                  description: "Default value for form autocomplete:\n\nOrganization title"
                },
                {
                  const: "organization",
                  description: "Default value for form autocomplete:\n\nOrganization"
                },
                {
                  const: "street-address",
                  description: "Default value for form autocomplete:\n\nStreet address"
                },
                {
                  const: "address-line1",
                  description: "Default value for form autocomplete:\n\nAddress line1"
                },
                {
                  const: "address-line2",
                  description: "Default value for form autocomplete:\n\nAddress line2"
                },
                {
                  const: "address-line3",
                  description: "Default value for form autocomplete:\n\nAddress line3"
                },
                {
                  const: "address-level4",
                  description: "Default value for form autocomplete:\n\nAddress level4"
                },
                {
                  const: "address-level3",
                  description: "Default value for form autocomplete:\n\nAddress level3"
                },
                {
                  const: "address-level2",
                  description: "Default value for form autocomplete:\n\nAddress level2"
                },
                {
                  const: "address-level1",
                  description: "Default value for form autocomplete:\n\nAddress level1"
                },
                {
                  const: "country",
                  description: "Default value for form autocomplete:\n\nCountry"
                },
                {
                  const: "country-name",
                  description: "Default value for form autocomplete:\n\nCountry name"
                },
                {
                  const: "postal-code",
                  description: "Default value for form autocomplete:\n\nPostal code"
                },
                {
                  const: "cc-name",
                  description: "Default value for form autocomplete:\n\nFull name as given on the payment instrument."
                },
                {
                  const: "cc-given-name",
                  description: "Default value for form autocomplete:\n\nFirst name as given on the payment instrument."
                },
                {
                  const: "cc-additional-name",
                  description: "Default value for form autocomplete:\n\nMiddle name as given on the payment instrument."
                },
                {
                  const: "cc-family-name",
                  description: "Default value for form autocomplete:\n\nLast name as given on the payment instrument."
                },
                {
                  const: "cc-number",
                  description: "Default value for form autocomplete:\n\nCode identifying the payment instrument (e.g. the credit card number)."
                },
                {
                  const: "cc-exp",
                  description: "Default value for form autocomplete:\n\nExpiration date of the payment instrument."
                },
                {
                  const: "cc-exp-month",
                  description: "Default value for form autocomplete:\n\nExpiration month of the payment instrument."
                },
                {
                  const: "cc-exp-year",
                  description: "Default value for form autocomplete:\n\nExpiration year of the payment instrument."
                },
                {
                  const: "cc-csc",
                  description: "Default value for form autocomplete:\n\nSecurity code for the payment instrument."
                },
                {
                  const: "cc-type",
                  description: "Default value for form autocomplete:\n\nType of payment instrument (e.g. Visa)."
                },
                {
                  const: "transaction-currency",
                  description: "Default value for form autocomplete:\n\nTransaction currency for the payment instrument."
                },
                {
                  const: "transaction-amount",
                  description: "Default value for form autocomplete:\n\nTransaction amount for the payment instrument."
                },
                {
                  const: "language",
                  description: "Default value for form autocomplete:\n\nPreferred language; a valid BCP 47 language tag."
                },
                {
                  const: "bday",
                  description: "Default value for form autocomplete:\n\nBirthday"
                },
                {
                  const: "bday-day",
                  description: "Default value for form autocomplete:\n\nDay for birthday"
                },
                {
                  const: "bday-month",
                  description: "Default value for form autocomplete:\n\nMonth for birthday"
                },
                {
                  const: "bday-year",
                  description: "Default value for form autocomplete:\n\nYear for birthday"
                },
                {
                  const: "sex",
                  description: "Default value for form autocomplete:\n\nGender identity (e.g. Female, Fa'afafine), free-form text, no newlines."
                },
                {
                  const: "tel",
                  description: "Default value for form autocomplete:\n\nFull telephone number, including country code"
                },
                {
                  const: "tel-country-code",
                  description: "Default value for form autocomplete:\n\nTelephone country code"
                },
                {
                  const: "tel-national",
                  description: "Default value for form autocomplete:\n\nTelephone national code"
                },
                {
                  const: "tel-area-code",
                  description: "Default value for form autocomplete:\n\nTelephone area code"
                },
                {
                  const: "tel-local",
                  description: "Default value for form autocomplete:\n\nTelephone local code"
                },
                {
                  const: "tel-local-prefix",
                  description: "Default value for form autocomplete:\n\nTelephone prefix"
                },
                {
                  const: "tel-local-suffix",
                  description: "Default value for form autocomplete:\n\nTelephone suffix"
                },
                {
                  const: "tel-extension",
                  description: "Default value for form autocomplete:\n\nTelephone extension"
                },
                {
                  const: "url",
                  description: "Default value for form autocomplete:\n\nHome page or other Web page corresponding to the company, person, address, or contact information in the other fields associated with this field."
                },
                {
                  const: "photo",
                  description: "Default value for form autocomplete:\n\nPhotograph, icon, or other image corresponding to the company, person, address, or contact information in the other fields associated with this field."
                }
              ]
            },
            "bar": { "type": "string" }
          }
        },
        "HtmlInputText":{"properties":{"autocomplete":{"description":"This attribute indicates whether the value of the control can be automatically completed by the browser.\nPossible values are:\n off: The user must explicitly enter a value into this field for every use, or the document provides its own auto-completion method. The browser does not automatically complete the entry.\non: The browser is allowed to automatically complete the value based on values that the user has entered during previous uses, however on does not provide any further information about what kind of data the user might be expected to enter.","type":"boolean"},"checkcode":{"type":"string"},"disabled":{"description":"This Boolean attribute indicates that the form control is not available for interaction. In particular, the click event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.\nUnlike other browsers, Firefox will by default persist the dynamic disabled state of an \u003cinput\u003e across page loads. Use the autocomplete attribute to control this feature.","type":"boolean"},"form":{"description":"The form element that the input element is associated with (its form owner). The value of the attribute must be an id of a \u003cform\u003e element in the same document. If this attribute is not specified, this \u003cinput\u003e element must be a descendant of a \u003cform\u003e element. This attribute enables you to place \u003cinput\u003e elements anywhere within a document, not just as descendants of their form elements. An input can only be associated with one form.","type":"string"},"list":{"description":"Identifies a list of pre-defined options to suggest to the user. The value must be the id of a \u003cdatalist\u003e element in the same document. The browser displays only options that are valid values for this input element. This attribute is ignored when the type attribute's value is hidden, checkbox, radio, file, or a button type.","type":"string"},"maxlength":{"description":"If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the maximum number of characters (in UTF-16 code units) that the user can enter. For other control types, it is ignored. It can exceed the value of the size attribute. If it is not specified, the user can enter an unlimited number of characters.\nSpecifying a negative number results in the default behavior (i.e. the user can enter an unlimited number of characters). The constraint is evaluated only when the value of the attribute has been changed.","type":"number"},"minlength":{"description":"If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the minimum number of characters (in Unicode code points) that the user can enter. For other control types, it is ignored.","type":"number"},"name":{"description":"The name of the control, which is submitted with the form data.\n\nIf you need to use the default values for the form's auto-completion feature, press ctrl+space to see the full list of values.","oneOf":[{"description":"html tag name","pattern":"^(?!^name$|^honorific-prefix$|^given-name$|^additional-name$|^family-name$|^honorific-suffix$|^nickname$|^email$|^username$|^new-password$|^current-password$|^organization-title$|^organization$|^street-address$|^address-line1$|^address-line2$|^address-line3$|^address-level4$|^address-level3$|^address-level2$|^country$|^country-name$|^postal-code$|^cc-name$|^cc-given-name$|^cc-additional-name$|^cc-family-name$|^cc-number$|^cc-exp$|^cc-exp-month$|^cc-exp-year$|^cc-csc$|^cc-type$|^transaction-currency$|^transaction-amount$|^language$|^bday$|^bday-day$|^bday-month$|^bday-year$|^sex$|^tel$|^tel-country-code$|^tel-national$|^tel-area-code$|^tel-local$|^tel-local-prefix$|^tel-local-suffix$|^tel-extension$|^url$|^photo$)[a-zA-Z0-9_]+([[0-9]*])*$"},{"const":"name","description":"Default value for form autocomplete:\n\nName"},{"const":"honorific-prefix","description":"Default value for form autocomplete:\n\nPrefix or title (e.g. \"Mr.\", \"Ms.\", \"Dr.\", \"Mlle\")."},{"const":"given-name","description":"Default value for form autocomplete:\n\nFirst name."},{"const":"additional-name","description":"Default value for form autocomplete:\n\nMiddle name."},{"const":"family-name","description":"Default value for form autocomplete:\n\nLast name."},{"const":"honorific-suffix","description":"Default value for form autocomplete:\n\nSuffix (e.g. \"Jr.\", \"B.Sc.\", \"MBASW\", \"II\")."},{"const":"nickname","description":"Default value for form autocomplete:\n\nNickname"},{"const":"email","description":"Default value for form autocomplete:\n\nE-mail"},{"const":"username","description":"Default value for form autocomplete:\n\nUser name"},{"const":"new-password","description":"Default value for form autocomplete:\n\nA new password (e.g. when creating an account or changing a password)."},{"const":"current-password","description":"Default value for form autocomplete:\n\nCurrent password"},{"const":"organization-title","description":"Default value for form autocomplete:\n\nOrganization title"},{"const":"organization","description":"Default value for form autocomplete:\n\nOrganization"},{"const":"street-address","description":"Default value for form autocomplete:\n\nStreet address"},{"const":"address-line1","description":"Default value for form autocomplete:\n\nAddress line1"},{"const":"address-line2","description":"Default value for form autocomplete:\n\nAddress line2"},{"const":"address-line3","description":"Default value for form autocomplete:\n\nAddress line3"},{"const":"address-level4","description":"Default value for form autocomplete:\n\nAddress level4"},{"const":"address-level3","description":"Default value for form autocomplete:\n\nAddress level3"},{"const":"address-level2","description":"Default value for form autocomplete:\n\nAddress level2"},{"const":"address-level1","description":"Default value for form autocomplete:\n\nAddress level1"},{"const":"country","description":"Default value for form autocomplete:\n\nCountry"},{"const":"country-name","description":"Default value for form autocomplete:\n\nCountry name"},{"const":"postal-code","description":"Default value for form autocomplete:\n\nPostal code"},{"const":"cc-name","description":"Default value for form autocomplete:\n\nFull name as given on the payment instrument."},{"const":"cc-given-name","description":"Default value for form autocomplete:\n\nFirst name as given on the payment instrument."},{"const":"cc-additional-name","description":"Default value for form autocomplete:\n\nMiddle name as given on the payment instrument."},{"const":"cc-family-name","description":"Default value for form autocomplete:\n\nLast name as given on the payment instrument."},{"const":"cc-number","description":"Default value for form autocomplete:\n\nCode identifying the payment instrument (e.g. the credit card number)."},{"const":"cc-exp","description":"Default value for form autocomplete:\n\nExpiration date of the payment instrument."},{"const":"cc-exp-month","description":"Default value for form autocomplete:\n\nExpiration month of the payment instrument."},{"const":"cc-exp-year","description":"Default value for form autocomplete:\n\nExpiration year of the payment instrument."},{"const":"cc-csc","description":"Default value for form autocomplete:\n\nSecurity code for the payment instrument."},{"const":"cc-type","description":"Default value for form autocomplete:\n\nType of payment instrument (e.g. Visa)."},{"const":"transaction-currency","description":"Default value for form autocomplete:\n\nTransaction currency for the payment instrument."},{"const":"transaction-amount","description":"Default value for form autocomplete:\n\nTransaction amount for the payment instrument."},{"const":"language","description":"Default value for form autocomplete:\n\nPreferred language; a valid BCP 47 language tag."},{"const":"bday","description":"Default value for form autocomplete:\n\nBirthday"},{"const":"bday-day","description":"Default value for form autocomplete:\n\nDay for birthday"},{"const":"bday-month","description":"Default value for form autocomplete:\n\nMonth for birthday"},{"const":"bday-year","description":"Default value for form autocomplete:\n\nYear for birthday"},{"const":"sex","description":"Default value for form autocomplete:\n\nGender identity (e.g. Female, Fa'afafine), free-form text, no newlines."},{"const":"tel","description":"Default value for form autocomplete:\n\nFull telephone number, including country code"},{"const":"tel-country-code","description":"Default value for form autocomplete:\n\nTelephone country code"},{"const":"tel-national","description":"Default value for form autocomplete:\n\nTelephone national code"},{"const":"tel-area-code","description":"Default value for form autocomplete:\n\nTelephone area code"},{"const":"tel-local","description":"Default value for form autocomplete:\n\nTelephone local code"},{"const":"tel-local-prefix","description":"Default value for form autocomplete:\n\nTelephone prefix"},{"const":"tel-local-suffix","description":"Default value for form autocomplete:\n\nTelephone suffix"},{"const":"tel-extension","description":"Default value for form autocomplete:\n\nTelephone extension"},{"const":"url","description":"Default value for form autocomplete:\n\nHome page or other Web page corresponding to the company, person, address, or contact information in the other fields associated with this field."},{"const":"photo","description":"Default value for form autocomplete:\n\nPhotograph, icon, or other image corresponding to the company, person, address, or contact information in the other fields associated with this field."}],"type":"string"},"pattern":{"description":"A regular expression that the control's value is checked against. The pattern must match the entire value, not just some subset. Use the title attribute to describe the pattern to help the user. This attribute applies when the value of the type attribute is text, search, tel, url, email, or password, otherwise it is ignored. The regular expression language is the same as JavaScript RegExp algorithm, with the 'u' parameter that makes it treat the pattern as a sequence of unicode code points. The pattern is not surrounded by forward slashes.","type":"string"},"placeholder":{"description":"A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or line-feeds.","type":"string"},"readonly":{"description":"This attribute indicates that the user cannot modify the value of the control. The value of the attribute is irrelevant. If you need read-write access to the input value, do not add the 'readonly' attribute. It is ignored if the value of the type attribute is hidden, range, color, checkbox, radio, file, or a button type (such as button or submit).","type":"boolean"},"required":{"description":"This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS pseudo-classes will be applied to the field as appropriate.","type":"boolean"},"size":{"description":"The initial size of the control. This value is in pixels unless the value of the type attribute is text or password, in which case it is an integer number of characters. Starting in HTML5, this attribute applies only when the type attribute is set to text, search, tel, url, email, or password, otherwise it is ignored. In addition, the size must be greater than zero. If you do not specify a size, a default value of 20 is used. HTML5 simply states 'the user agent should ensure that at least that many characters are visible', but different characters can have different widths in certain fonts. In some browsers, a certain string with x characters will not be entirely visible even if size is defined to at least x.","type":"number"},"value":{"description":"The initial value of the control. This attribute is optional except when the value of the type attribute is radio or checkbox.\nNote that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was changed before the reload.","type":"string"}},"required":[],"type":"object"},
        formConfig: {
          type: "object",
          required: [
            "id",
            "content"
          ],
          properties: {
            id: {
              type: "string",
              pattern: "[a-zA-Z0-9_]+",
              description: "html tag id. must be unique",
            },
            test: {
              $ref: "#/definitions/test"

            },
            HtmlInputText: {
              $ref: "#/definitions/HtmlInputText"
            },
            class: {
              type: "string",
              description: "tag class"
            },
            content: {
              type: "array",
              items: {
                $ref: "#/definitions/formContent"


              }
            }
          }
        },
        formContent: {
          type: "object",
          properties: {
            label: {
              type: "string",
              description: "text of html element label"
            },
            input: {
              type: "object",
              properties: {
                global: {
                  $ref: "#/definitions/HtmlInputText"
                },
                label: {
                  type: "string",
                  description: "text of html element label"
                },
                type: {
                  type: "object",
                  enum: ["text", "date"]
                }
              }
            }
          }
        },
        "veggie": {
          "type": "object",
          "required": [
            "veggieName",
            "veggieLike"
          ],
          "properties": {
            "label": {
              "type": "string",
              "description": "Name of html element label"
            },
            "input": {
              "type": "object",
              $ref: "#/definitions/HtmlInputText"


            },
            "veggieName": {
              "type": "string",
              "description": "The name of the vegetable."
            },
            "veggieLike": {
              "type": "boolean",
              "description": "Do I like this vegetable?"
            }
          }
        }
      }
    }
  }]
});

monaco.editor.create(document.getElementById("container"), {
  model: model
});