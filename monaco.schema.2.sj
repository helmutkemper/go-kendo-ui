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





















        "HtmlInputText": {
          "properties": {
            "apagar": {
              "properties": {
                "animation": {
                  "oneOf": [
                    {
                      "type": "boolean"
                    },
                    {
                      "type":"object",
                      "close": {
                        "properties": {
                          "description": "https://docs.telerik.com/kendo-ui/api/javascript/ui/animation.close\\n\\nType: Object\\n\\n\\n\\nExample - configure the close animation\\n\\n\u003cinput id=\\\"autocomplete\\\" /\u003e\\n\u003cscript\u003e\\n$(\\\"#autocomplete\\\").kendoAutoComplete({\\n  animation: {\\n   close: {\\n     effects: \\\"zoom:out\\\",\\n     duration: 300\\n   }\\n  }\\n});\\n\u003c/script\u003e\\n",
                          "oneOf": [
                            {
                              "duration": {
                                "properties": {
                                  "description": "https://docs.telerik.com/kendo-ui/api/javascript/ui/animation.close.duration\\n\\nType: Number (default: 100)\\n\\n\\n\\n",
                                  "oneOf": [
                                    {
                                      "type": "number"
                                    }
                                  ]
                                }
                              },
                              "effects": {
                                "properties": {
                                  "description": "https://docs.telerik.com/kendo-ui/api/javascript/ui/animation.close.effects\\n\\nType: String\\n\\n\\n\\n",
                                  "oneOf": [
                                    {
                                      "type": "string"
                                    }
                                  ]
                                }
                              },
                              "type": "object"
                            }
                          ]
                        }
                      },

                    }
                  ]
                },

              },
              "global": {
                "properties": {
                  "accesskey": {
                    "description": "Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.",
                    "type": "string"
                  },
                  "class": {
                    "description": "Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access specific elements via the class selectors or functions like the method Document.getElementsByClassName().",
                    "type": "string"
                  },
                  "contenteditable": {
                    "description": "Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its widget to allow editing. The attribute must take one of the following values:\n true or the empty string, which indicates that the element must be editable;\nfalse, which indicates that the element must not be editable.",
                    "type": "boolean"
                  },
                  "contextmenu": {
                    "description": "Is the id of an \u003cmenu\u003e to use as the contextual menu for this element.",
                    "type": "string"
                  },
                  "data": {
                    "type": "object"
                  },
                  "dir": {
                    "description": "Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:\nltr, which means left to right and is to be used for languages that are written from the left to the right (like English);\nrtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);\nauto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until it finds a character with a strong directionality, then apply that directionality to the whole element.",
                    "enum": [
                      "ltr",
                      "rtl",
                      "auto"
                    ],
                    "type": "string"
                  },
                  "draggable": {
                    "description": "Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used to hide elements of the page that can't be used until the login process has been completed. The browser won't render such elements. This attribute must not be used to hide content that could legitimately be shown.",
                    "type": "boolean"
                  },
                  "dropzone": {
                    "enum": [
                      "copy",
                      "move",
                      "link"
                    ],
                    "type": "string"
                  },
                  "extra": {
                    "type": "object"
                  },
                  "id": {
                    "description": "Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element when linking (using a fragment identifier), scripting, or styling (with CSS).",
                    "type": "string"
                  },
                  "itemdrop": {
                    "description": "Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop consists of a name and value pair.",
                    "type": "string"
                  },
                  "itemid": {
                    "description": "The unique, global identifier of an item.",
                    "type": "string"
                  },
                  "itemref": {
                    "description": "Properties that are not descendants of an element with the itemscope attribute can be associated with the item using an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.",
                    "type": "string"
                  },
                  "itemscope": {
                    "description": "itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of a vocabulary (such as schema.org) that describes the item and its properties context.",
                    "type": "string"
                  },
                  "itemtype": {
                    "description": "Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure. itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.",
                    "type": "string"
                  },
                  "lang": {
                    "description": "Participates in defining the language of the element, the language that non-editable elements are written in or the language that editable elements should be written in. The tag contains one single entry value in the format defined in the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.",
                    "type": "string"
                  },
                  "onabort": {
                    "description": "An event handler for abort events sent to the window. (Not available with Firefox 2 or Safari.)\nWhile the standard for aborting a document load is defined, HTML issue #3525 suggests that browsers do not currently fire the 'abort' event on window that would trigger onabort to be called.",
                    "type": "string"
                  },
                  "onautocomplete": {
                    "description": "The HTML autocomplete attribute is available on several kinds of \u003cinput\u003e elements—those that take a text or numeric value as input. autocomplete lets web developers specify what if any permission the user agent has to provide automated assistance in filling out form field values, as well as guidance to the browser as to the type of information expected in the field.\nThe source of the suggested values is generally up to the browser; typically values come from past values entered by the user, but they may also come from pre-configured values. For instance, a browser might let the user save their name, address, phone number, and email addresses for autocomplete purposes. Perhaps the browser offers the ability to save encrypted credit card information, for autocompletion following an authentication procedure.\nIf an \u003cinput\u003e element has no autocomplete attribute, then browsers use the autocomplete attribute of the element's form owner, which is either the \u003cform\u003e element that the \u003cinput\u003e element is a descendant of, or the \u003cform\u003e whose id is specified by the form attribute of the \u003cinput\u003e element.",
                    "type": "string"
                  },
                  "onautocompleteerror": {
                    "type": "string"
                  },
                  "onblur": {
                    "type": "string"
                  },
                  "oncancel": {
                    "type": "string"
                  },
                  "oncanplay": {
                    "type": "string"
                  },
                  "oncanplaythrough": {
                    "type": "string"
                  },
                  "onchange": {
                    "type": "string"
                  },
                  "onclick": {
                    "type": "string"
                  },
                  "onclose": {
                    "type": "string"
                  },
                  "oncontextmenu": {
                    "type": "string"
                  },
                  "oncuechange": {
                    "type": "string"
                  },
                  "ondblclick": {
                    "type": "string"
                  },
                  "ondrag": {
                    "type": "string"
                  },
                  "ondragend": {
                    "type": "string"
                  },
                  "ondragenter": {
                    "type": "string"
                  },
                  "ondragexit": {
                    "type": "string"
                  },
                  "ondragleave": {
                    "type": "string"
                  },
                  "ondragover": {
                    "type": "string"
                  },
                  "ondragstart": {
                    "type": "string"
                  },
                  "ondrop": {
                    "type": "string"
                  },
                  "ondurationchange": {
                    "type": "string"
                  },
                  "onemptied": {
                    "type": "string"
                  },
                  "onended": {
                    "type": "string"
                  },
                  "onerror": {
                    "type": "string"
                  },
                  "onfocus": {
                    "type": "string"
                  },
                  "oninput": {
                    "type": "string"
                  },
                  "oninvalid": {
                    "type": "string"
                  },
                  "onkeydown": {
                    "type": "string"
                  },
                  "onkeypress": {
                    "type": "string"
                  },
                  "onkeyup": {
                    "type": "string"
                  },
                  "onload": {
                    "type": "string"
                  },
                  "onloadeddata": {
                    "type": "string"
                  },
                  "onloadedmetadata": {
                    "type": "string"
                  },
                  "onloadstart": {
                    "type": "string"
                  },
                  "onmousedown": {
                    "type": "string"
                  },
                  "onmouseenter": {
                    "type": "string"
                  },
                  "onmouseleave": {
                    "type": "string"
                  },
                  "onmousemove": {
                    "type": "string"
                  },
                  "onmouseout": {
                    "type": "string"
                  },
                  "onmouseover": {
                    "type": "string"
                  },
                  "onmouseup": {
                    "type": "string"
                  },
                  "onmousewheel": {
                    "type": "string"
                  },
                  "onpause": {
                    "type": "string"
                  },
                  "onplay": {
                    "type": "string"
                  },
                  "onplaying": {
                    "type": "string"
                  },
                  "onprogress": {
                    "type": "string"
                  },
                  "onratechange": {
                    "type": "string"
                  },
                  "onreset": {
                    "type": "string"
                  },
                  "onresize": {
                    "type": "string"
                  },
                  "onscroll": {
                    "type": "string"
                  },
                  "onseeked": {
                    "type": "string"
                  },
                  "onseeking": {
                    "type": "string"
                  },
                  "onselect": {
                    "type": "string"
                  },
                  "onshow": {
                    "type": "string"
                  },
                  "onsort": {
                    "type": "string"
                  },
                  "onstalled": {
                    "type": "string"
                  },
                  "onsubmit": {
                    "type": "string"
                  },
                  "onsuspend": {
                    "type": "string"
                  },
                  "ontimeupdate": {
                    "type": "string"
                  },
                  "ontoggle": {
                    "type": "string"
                  },
                  "onvolumechange": {
                    "type": "string"
                  },
                  "onwaiting": {
                    "type": "string"
                  },
                  "sort": {
                    "description": "Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot created by the \u003cslot\u003e element whose name attribute's value matches that slot attribute's value.",
                    "type": "string"
                  },
                  "spellcheck": {
                    "description": "Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following values:\ntrue, which indicates that the element should be, if possible, checked for spelling errors;\nfalse, which indicates that the element should not be checked for spelling errors.",
                    "type": "boolean"
                  },
                  "style": {
                    "description": "Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined in a separate file or files. This attribute and the \u003cstyle\u003e element have mainly the purpose of allowing for quick styling, for example for testing purposes.",
                    "type": "string"
                  },
                  "tabindex": {
                    "description": "Is an integer attribute indicating if the element can take input focus (is focusable), if it should participate to sequential keyboard navigation, and if so, at what position. It can takes several values:\na negative value means that the element should be focusable, but should not be reachable via sequential keyboard navigation;\n0 means that the element should be focusable and reachable via sequential keyboard navigation, but its relative order is defined by the platform convention;\na positive value means that the element should be focusable and reachable via sequential keyboard navigation; the order in which the elements are focused is the increasing value of the tabindex. If several elements share the same tabindex, their relative order follows their relative positions in the document.",
                    "type": "number"
                  },
                  "title": {
                    "description": "Contains a text representing advisory information related to the element it belongs to. Such information can typically, but not necessarily, be presented to the user as a tooltip.",
                    "type": "string"
                  },
                  "translate": {
                    "description": "Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the following values:\nempty string and 'yes', which indicates that the element will be translated.\n'no', which indicates that the element will not be translated.",
                    "type": "boolean"
                  }
                },
                "required": [

                ],
                "type": "object"
              },
              "name": {
                "description": "The name of the control, which is submitted with the form data.\n\nIf you need to use the default values for the form's auto-completion feature, press ctrl+space to see the full list of values.",
                "oneOf": [
                  {
                    "description": "html tag name",
                    "pattern": "^(?!^name$|^honorific-prefix$|^given-name$|^additional-name$|^family-name$|^honorific-suffix$|^nickname$|^email$|^username$|^new-password$|^current-password$|^organization-title$|^organization$|^street-address$|^address-line1$|^address-line2$|^address-line3$|^address-level4$|^address-level3$|^address-level2$|^country$|^country-name$|^postal-code$|^cc-name$|^cc-given-name$|^cc-additional-name$|^cc-family-name$|^cc-number$|^cc-exp$|^cc-exp-month$|^cc-exp-year$|^cc-csc$|^cc-type$|^transaction-currency$|^transaction-amount$|^language$|^bday$|^bday-day$|^bday-month$|^bday-year$|^sex$|^tel$|^tel-country-code$|^tel-national$|^tel-area-code$|^tel-local$|^tel-local-prefix$|^tel-local-suffix$|^tel-extension$|^url$|^photo$)[a-zA-Z0-9_]+([[0-9]*])*$"
                  },
                  {
                    "const": "name",
                    "description": "Default value for form autocomplete:\n\nName"
                  },
                  {
                    "const": "honorific-prefix",
                    "description": "Default value for form autocomplete:\n\nPrefix or title (e.g. \"Mr.\", \"Ms.\", \"Dr.\", \"Mlle\")."
                  },
                  {
                    "const": "given-name",
                    "description": "Default value for form autocomplete:\n\nFirst name."
                  },
                  {
                    "const": "additional-name",
                    "description": "Default value for form autocomplete:\n\nMiddle name."
                  },
                  {
                    "const": "family-name",
                    "description": "Default value for form autocomplete:\n\nLast name."
                  },
                  {
                    "const": "honorific-suffix",
                    "description": "Default value for form autocomplete:\n\nSuffix (e.g. \"Jr.\", \"B.Sc.\", \"MBASW\", \"II\")."
                  },
                  {
                    "const": "nickname",
                    "description": "Default value for form autocomplete:\n\nNickname"
                  },
                  {
                    "const": "email",
                    "description": "Default value for form autocomplete:\n\nE-mail"
                  },
                  {
                    "const": "username",
                    "description": "Default value for form autocomplete:\n\nUser name"
                  },
                  {
                    "const": "new-password",
                    "description": "Default value for form autocomplete:\n\nA new password (e.g. when creating an account or changing a password)."
                  },
                  {
                    "const": "current-password",
                    "description": "Default value for form autocomplete:\n\nCurrent password"
                  },
                  {
                    "const": "organization-title",
                    "description": "Default value for form autocomplete:\n\nOrganization title"
                  },
                  {
                    "const": "organization",
                    "description": "Default value for form autocomplete:\n\nOrganization"
                  },
                  {
                    "const": "street-address",
                    "description": "Default value for form autocomplete:\n\nStreet address"
                  },
                  {
                    "const": "address-line1",
                    "description": "Default value for form autocomplete:\n\nAddress line1"
                  },
                  {
                    "const": "address-line2",
                    "description": "Default value for form autocomplete:\n\nAddress line2"
                  },
                  {
                    "const": "address-line3",
                    "description": "Default value for form autocomplete:\n\nAddress line3"
                  },
                  {
                    "const": "address-level4",
                    "description": "Default value for form autocomplete:\n\nAddress level4"
                  },
                  {
                    "const": "address-level3",
                    "description": "Default value for form autocomplete:\n\nAddress level3"
                  },
                  {
                    "const": "address-level2",
                    "description": "Default value for form autocomplete:\n\nAddress level2"
                  },
                  {
                    "const": "address-level1",
                    "description": "Default value for form autocomplete:\n\nAddress level1"
                  },
                  {
                    "const": "country",
                    "description": "Default value for form autocomplete:\n\nCountry"
                  },
                  {
                    "const": "country-name",
                    "description": "Default value for form autocomplete:\n\nCountry name"
                  },
                  {
                    "const": "postal-code",
                    "description": "Default value for form autocomplete:\n\nPostal code"
                  },
                  {
                    "const": "cc-name",
                    "description": "Default value for form autocomplete:\n\nFull name as given on the payment instrument."
                  },
                  {
                    "const": "cc-given-name",
                    "description": "Default value for form autocomplete:\n\nFirst name as given on the payment instrument."
                  },
                  {
                    "const": "cc-additional-name",
                    "description": "Default value for form autocomplete:\n\nMiddle name as given on the payment instrument."
                  },
                  {
                    "const": "cc-family-name",
                    "description": "Default value for form autocomplete:\n\nLast name as given on the payment instrument."
                  },
                  {
                    "const": "cc-number",
                    "description": "Default value for form autocomplete:\n\nCode identifying the payment instrument (e.g. the credit card number)."
                  },
                  {
                    "const": "cc-exp",
                    "description": "Default value for form autocomplete:\n\nExpiration date of the payment instrument."
                  },
                  {
                    "const": "cc-exp-month",
                    "description": "Default value for form autocomplete:\n\nExpiration month of the payment instrument."
                  },
                  {
                    "const": "cc-exp-year",
                    "description": "Default value for form autocomplete:\n\nExpiration year of the payment instrument."
                  },
                  {
                    "const": "cc-csc",
                    "description": "Default value for form autocomplete:\n\nSecurity code for the payment instrument."
                  },
                  {
                    "const": "cc-type",
                    "description": "Default value for form autocomplete:\n\nType of payment instrument (e.g. Visa)."
                  },
                  {
                    "const": "transaction-currency",
                    "description": "Default value for form autocomplete:\n\nTransaction currency for the payment instrument."
                  },
                  {
                    "const": "transaction-amount",
                    "description": "Default value for form autocomplete:\n\nTransaction amount for the payment instrument."
                  },
                  {
                    "const": "language",
                    "description": "Default value for form autocomplete:\n\nPreferred language; a valid BCP 47 language tag."
                  },
                  {
                    "const": "bday",
                    "description": "Default value for form autocomplete:\n\nBirthday"
                  },
                  {
                    "const": "bday-day",
                    "description": "Default value for form autocomplete:\n\nDay for birthday"
                  },
                  {
                    "const": "bday-month",
                    "description": "Default value for form autocomplete:\n\nMonth for birthday"
                  },
                  {
                    "const": "bday-year",
                    "description": "Default value for form autocomplete:\n\nYear for birthday"
                  },
                  {
                    "const": "sex",
                    "description": "Default value for form autocomplete:\n\nGender identity (e.g. Female, Fa'afafine), free-form text, no newlines."
                  },
                  {
                    "const": "tel",
                    "description": "Default value for form autocomplete:\n\nFull telephone number, including country code"
                  },
                  {
                    "const": "tel-country-code",
                    "description": "Default value for form autocomplete:\n\nTelephone country code"
                  },
                  {
                    "const": "tel-national",
                    "description": "Default value for form autocomplete:\n\nTelephone national code"
                  },
                  {
                    "const": "tel-area-code",
                    "description": "Default value for form autocomplete:\n\nTelephone area code"
                  },
                  {
                    "const": "tel-local",
                    "description": "Default value for form autocomplete:\n\nTelephone local code"
                  },
                  {
                    "const": "tel-local-prefix",
                    "description": "Default value for form autocomplete:\n\nTelephone prefix"
                  },
                  {
                    "const": "tel-local-suffix",
                    "description": "Default value for form autocomplete:\n\nTelephone suffix"
                  },
                  {
                    "const": "tel-extension",
                    "description": "Default value for form autocomplete:\n\nTelephone extension"
                  },
                  {
                    "const": "url",
                    "description": "Default value for form autocomplete:\n\nHome page or other Web page corresponding to the company, person, address, or contact information in the other fields associated with this field."
                  },
                  {
                    "const": "photo",
                    "description": "Default value for form autocomplete:\n\nPhotograph, icon, or other image corresponding to the company, person, address, or contact information in the other fields associated with this field."
                  }
                ],
                "type": "string"
              }
            },
            "required": [

            ],
            "type": "object"
          }
        }













































































        ,
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
              oneOf:[
                {
                  properties: {
                    not:{
                      outra: {
                        type: "object"
                      }
                    },
                    global: {
                      $ref: "#/definitions/HtmlInputText"
                    },
                    label: {
                      type: "string",
                      description: "text of html element label"
                    },
                    type: {
                      type: "string",
                      enum: ["text", "date"]
                    }
                  }
                },
                {
                  properties: {
                    outra: {
                      type: "object"
                    },
                    new: {
                      type: "string",
                      enum: ["text", "date"]
                    }
                  }
                }
              ]
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