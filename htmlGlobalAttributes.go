package telerik

import (
	"bytes"
	"reflect"
)

// Global attributes are attributes common to all HTML elements; they can be used on all elements, though they may have
// no effect on some elements.
//
// Global attributes may be specified on all HTML elements, even those not specified in the standard. That means that
// any non-standard elements must still permit these attributes, even though using those elements means that the
// document is no longer HTML5-compliant. For example, HTML5-compliant browsers hide content marked as <foo hidden>...
// <foo>, even though <foo> is not a valid HTML element.
type HtmlGlobalAttributes struct {
	/*
	  This field was created for internal use and should not be accessed directly
	*/
	DoNotUseThisFieldOmitHtml Boolean `htmlAttr:"-" jsonSchema_description:"-"`
	/*
	  Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a
	  space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.
	*/
	AccessKey string `htmlAttr:"accesskey" jsonSchema_description:"Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout."`

	/*
	  Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access
	  specific elements via the class selectors or functions like the method Document.getElementsByClassName().
	*/
	Class string `htmlAttr:"class" jsonSchema_description:"Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access specific elements via the class selectors or functions like the method Document.getElementsByClassName()."`

	/*
	  Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its
	  widget to allow editing. The attribute must take one of the following values:
	  true or the empty string, which indicates that the element must be editable;
	  false, which indicates that the element must not be editable.
	*/
	ContentEditable Boolean `htmlAttr:"contenteditable" jsonSchema_description:"Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its widget to allow editing. The attribute must take one of the following values:\n true or the empty string, which indicates that the element must be editable;\nfalse, which indicates that the element must not be editable."`

	/*
	  Is the id of an <menu> to use as the contextual menu for this element.
	*/
	ContextMenu string `htmlAttr:"contextmenu" jsonSchema_description:"Is the id of an <menu> to use as the contextual menu for this element."`

	/*
	  Forms a class of attributes, called custom data attributes, that allow proprietary information to be exchanged between
	  the HTML and its DOM representation that may be used by scripts. All such custom data are available via the
	  HTMLElement interface of the element the attribute is set on. The HTMLElement.dataset property gives access to them.
	*/
	Data map[string]string `htmlAttr:"data" jsonSchema_description:"Forms a class of attributes, called custom data attributes, that allow proprietary information to be exchanged between the HTML and its DOM representation that may be used by scripts. All such custom data are available via the HTMLElement interface of the element the attribute is set on. The HTMLElement.dataset property gives access to them."`

	/*
	  Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:
	  ltr, which means left to right and is to be used for languages that are written from the left to the right (like English);
	  rtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);
	  auto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until
	  it finds a character with a strong directionality, then apply that directionality to the whole element.
	*/
	Dir string `htmlAttr:"dir" jsonSchema_description:"Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:\nltr, which means left to right and is to be used for languages that are written from the left to the right (like English);\nrtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);\nauto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until it finds a character with a strong directionality, then apply that directionality to the whole element." jsonSchema_enum:"['ltr', 'rtl', 'auto']"`

	/*
	  Is an enumerated attribute indicating whether the element can be dragged, using the Drag and Drop API. It can have the
	  following values:
	  true, which indicates that the element may be dragged
	  false, which indicates that the element may not be dragged.
	*/
	Draggable Boolean `htmlAttr:"draggable" jsonSchema_description:"Is an enumerated attribute indicating whether the element can be dragged, using the Drag and Drop API. It can have the following values:\ntrue, which indicates that the element may be dragged\nfalse, which indicates that the element may not be dragged."`

	/*
	  Is an enumerated attribute indicating what types of content can be dropped on an element, using the Drag and Drop API.
	  It can have the following values:
	  copy, which indicates that dropping will create a copy of the element that was dragged
	  move, which indicates that the element that was dragged will be moved to this new location.
	  link, will create a link to the dragged data.
	*/
	DropZone TypeHtmlDropZone `htmlAttr:"dropzone" jsonSchema_description:"Is an enumerated attribute indicating what types of content can be dropped on an element, using the Drag and Drop API.\nIt can have the following values:\ncopy, which indicates that dropping will create a copy of the element that was dragged\nmove, which indicates that the element that was dragged will be moved to this new location.\nlink, will create a link to the dragged data." jsonSchema_enum:"['copy', 'move', 'link']"`

	/*
	  Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used
	  to hide elements of the page that can't be used until the login process has been completed. The browser won't render
	  such elements. This attribute must not be used to hide content that could legitimately be shown.
	*/
	Hidden Boolean `htmlAttrSet:"draggable" jsonSchema_description:"Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used to hide elements of the page that can't be used until the login process has been completed. The browser won't render such elements. This attribute must not be used to hide content that could legitimately be shown."`

	/*
	  Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element
	  when linking (using a fragment identifier), scripting, or styling (with CSS).
	*/
	Id string `htmlAttr:"id" jsonSchema_description:"Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element when linking (using a fragment identifier), scripting, or styling (with CSS)."`

	/*
	  The unique, global identifier of an item.
	*/
	ItemId string `htmlAttr:"itemid" jsonSchema_description:"The unique, global identifier of an item."`

	/*
	  Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop
	  consists of a name and value pair.
	*/
	ItemDrop string `htmlAttr:"itemdrop" jsonSchema_description:"Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop consists of a name and value pair."`

	/*
	  Properties that are not descendants of an element with the itemscope attribute can be associated with the item using
	  an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
	*/
	ItemRef string `htmlAttr:"itemref" jsonSchema_description:"Properties that are not descendants of an element with the itemscope attribute can be associated with the item using an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document."`

	/*
	  itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
	  item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of
	  a vocabulary (such as schema.org) that describes the item and its properties context.
	*/
	ItemScope string `htmlAttr:"itemscope" jsonSchema_description:"itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of a vocabulary (such as schema.org) that describes the item and its properties context."`

	/*
	  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure.
	  itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
	*/
	ItemType string `htmlAttr:"itemtype" jsonSchema_description:"Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure. itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active."`

	/*
	  Participates in defining the language of the element, the language that non-editable elements are written in or the
	  language that editable elements should be written in. The tag contains one single entry value in the format defined in
	  the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.
	*/
	Lang string `htmlAttr:"lang" jsonSchema_description:"Participates in defining the language of the element, the language that non-editable elements are written in or the language that editable elements should be written in. The tag contains one single entry value in the format defined in the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it."`

	/*
	  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot
	  created by the <slot> element whose name attribute's value matches that slot attribute's value.
	*/
	Sort string `htmlAttr:"sort" jsonSchema_description:"Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot created by the <slot> element whose name attribute's value matches that slot attribute's value."`

	/*
	  Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following
	  values:
	  true, which indicates that the element should be, if possible, checked for spelling errors;
	  false, which indicates that the element should not be checked for spelling errors.
	*/
	SpellCheck Boolean `htmlAttrSet:"spellcheck" jsonSchema_description:"Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following values:\ntrue, which indicates that the element should be, if possible, checked for spelling errors;\nfalse, which indicates that the element should not be checked for spelling errors."`

	/*
	  Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined
	  in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick
	  styling, for example for testing purposes.
	*/
	Style string `htmlAttr:"style" jsonSchema_description:"Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick styling, for example for testing purposes."`

	/*
	  Is an integer attribute indicating if the element can take input focus (is focusable), if it should participate to
	  sequential keyboard navigation, and if so, at what position. It can takes several values:
	  a negative value means that the element should be focusable, but should not be reachable via sequential keyboard
	  navigation;
	  0 means that the element should be focusable and reachable via sequential keyboard navigation, but its relative order
	  is defined by the platform convention;
	  a positive value means that the element should be focusable and reachable via sequential keyboard navigation; the
	  order in which the elements are focused is the increasing value of the tabindex. If several elements share the same
	  tabindex, their relative order follows their relative positions in the document.
	*/
	TabIndex int `htmlAttr:"tabindex" jsonSchema_description:"Is an integer attribute indicating if the element can take input focus (is focusable), if it should participate to sequential keyboard navigation, and if so, at what position. It can takes several values:\na negative value means that the element should be focusable, but should not be reachable via sequential keyboard navigation;\n0 means that the element should be focusable and reachable via sequential keyboard navigation, but its relative order is defined by the platform convention;\na positive value means that the element should be focusable and reachable via sequential keyboard navigation; the order in which the elements are focused is the increasing value of the tabindex. If several elements share the same tabindex, their relative order follows their relative positions in the document."`

	/*
	  Contains a text representing advisory information related to the element it belongs to. Such information can
	  typically, but not necessarily, be presented to the user as a tooltip.
	*/
	Title string `htmlAttr:"title" jsonSchema_description:"Contains a text representing advisory information related to the element it belongs to. Such information can typically, but not necessarily, be presented to the user as a tooltip."`

	/*
	  Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text
	  node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the
	  following values:
	  empty string and "yes", which indicates that the element will be translated.
	  "no", which indicates that the element will not be translated.
	*/
	Translate Boolean `htmlAttr:"translate" jsonSchema_description:"Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the following values:\nempty string and 'yes', which indicates that the element will be translated.\n'no', which indicates that the element will not be translated."`

	Extra map[string]interface{} `htmlAttr:"extra" jsonSchema_description:""`

	/*
			  An event handler for abort events sent to the window. (Not available with Firefox 2 or Safari.)
		    While the standard for aborting a document load is defined, HTML issue #3525 suggests that browsers do not currently fire the "abort" event on window that would trigger onabort to be called.
	*/
	OnAbort string `htmlAttr:"onabort" jsonSchema_description:"An event handler for abort events sent to the window. (Not available with Firefox 2 or Safari.)\nWhile the standard for aborting a document load is defined, HTML issue #3525 suggests that browsers do not currently fire the 'abort' event on window that would trigger onabort to be called."`

	/*
			  The HTML autocomplete attribute is available on several kinds of <input> elements—those that take a text or numeric value as input. autocomplete lets web developers specify what if any permission the user agent has to provide automated assistance in filling out form field values, as well as guidance to the browser as to the type of information expected in the field.
		    The source of the suggested values is generally up to the browser; typically values come from past values entered by the user, but they may also come from pre-configured values. For instance, a browser might let the user save their name, address, phone number, and email addresses for autocomplete purposes. Perhaps the browser offers the ability to save encrypted credit card information, for autocompletion following an authentication procedure.
		    If an <input> element has no autocomplete attribute, then browsers use the autocomplete attribute of the element's form owner, which is either the <form> element that the <input> element is a descendant of, or the <form> whose id is specified by the form attribute of the <input> element.
	*/
	OnAutoComplete string `htmlAttr:"onautocomplete" jsonSchema_description:"The HTML autocomplete attribute is available on several kinds of <input> elements—those that take a text or numeric value as input. autocomplete lets web developers specify what if any permission the user agent has to provide automated assistance in filling out form field values, as well as guidance to the browser as to the type of information expected in the field.\nThe source of the suggested values is generally up to the browser; typically values come from past values entered by the user, but they may also come from pre-configured values. For instance, a browser might let the user save their name, address, phone number, and email addresses for autocomplete purposes. Perhaps the browser offers the ability to save encrypted credit card information, for autocompletion following an authentication procedure.\nIf an <input> element has no autocomplete attribute, then browsers use the autocomplete attribute of the element's form owner, which is either the <form> element that the <input> element is a descendant of, or the <form> whose id is specified by the form attribute of the <input> element."`

	OnAutoCompleteError string `htmlAttr:"onautocompleteerror" jsonSchema_description:""`
	OnBlur              string `htmlAttr:"onblur" jsonSchema_description:""`
	OnCancel            string `htmlAttr:"oncancel" jsonSchema_description:""`
	OnCanPlay           string `htmlAttr:"oncanplay" jsonSchema_description:""`
	OnCanPlayThrough    string `htmlAttr:"oncanplaythrough" jsonSchema_description:""`
	OnChange            string `htmlAttr:"onchange" jsonSchema_description:""`
	OnClick             string `htmlAttr:"onclick" jsonSchema_description:""`
	OnClose             string `htmlAttr:"onclose" jsonSchema_description:""`
	OnContextMenu       string `htmlAttr:"oncontextmenu" jsonSchema_description:""`
	OnCueChange         string `htmlAttr:"oncuechange" jsonSchema_description:""`
	OnDblClick          string `htmlAttr:"ondblclick" jsonSchema_description:""`
	OnDrag              string `htmlAttr:"ondrag" jsonSchema_description:""`
	OnDragEnd           string `htmlAttr:"ondragend" jsonSchema_description:""`
	OnDragEnter         string `htmlAttr:"ondragenter" jsonSchema_description:""`
	OnDragExit          string `htmlAttr:"ondragexit" jsonSchema_description:""`
	OnDragLeave         string `htmlAttr:"ondragleave" jsonSchema_description:""`
	OnDragOver          string `htmlAttr:"ondragover" jsonSchema_description:""`
	OnDragStart         string `htmlAttr:"ondragstart" jsonSchema_description:""`
	OnDrop              string `htmlAttr:"ondrop" jsonSchema_description:""`
	OnDurationChange    string `htmlAttr:"ondurationchange" jsonSchema_description:""`
	OnEmptied           string `htmlAttr:"onemptied" jsonSchema_description:""`
	OnEnded             string `htmlAttr:"onended" jsonSchema_description:""`
	OnError             string `htmlAttr:"onerror" jsonSchema_description:""`
	OnFocus             string `htmlAttr:"onfocus" jsonSchema_description:""`
	OnInput             string `htmlAttr:"oninput" jsonSchema_description:""`
	OnInvalid           string `htmlAttr:"oninvalid" jsonSchema_description:""`
	OnKeyDown           string `htmlAttr:"onkeydown" jsonSchema_description:""`
	OnKeyPress          string `htmlAttr:"onkeypress" jsonSchema_description:""`
	OnKeyUp             string `htmlAttr:"onkeyup" jsonSchema_description:""`
	OnLoad              string `htmlAttr:"onload" jsonSchema_description:""`
	OnLoadedData        string `htmlAttr:"onloadeddata" jsonSchema_description:""`
	OnLoadedMetadata    string `htmlAttr:"onloadedmetadata" jsonSchema_description:""`
	OnLoadStart         string `htmlAttr:"onloadstart" jsonSchema_description:""`
	OnMouseDown         string `htmlAttr:"onmousedown" jsonSchema_description:""`
	OnMouseEnter        string `htmlAttr:"onmouseenter" jsonSchema_description:""`
	OnMouseLeave        string `htmlAttr:"onmouseleave" jsonSchema_description:""`
	OnMouseMove         string `htmlAttr:"onmousemove" jsonSchema_description:""`
	OnMouseOut          string `htmlAttr:"onmouseout" jsonSchema_description:""`
	OnMouseOver         string `htmlAttr:"onmouseover" jsonSchema_description:""`
	OnMouseUp           string `htmlAttr:"onmouseup" jsonSchema_description:""`
	OnMouseWheel        string `htmlAttr:"onmousewheel" jsonSchema_description:""`
	OnPause             string `htmlAttr:"onpause" jsonSchema_description:""`
	OnPlay              string `htmlAttr:"onplay" jsonSchema_description:""`
	OnPlaying           string `htmlAttr:"onplaying" jsonSchema_description:""`
	OnProgress          string `htmlAttr:"onprogress" jsonSchema_description:""`
	OnRateChange        string `htmlAttr:"onratechange" jsonSchema_description:""`
	OnReset             string `htmlAttr:"onreset" jsonSchema_description:""`
	OnResize            string `htmlAttr:"onresize" jsonSchema_description:""`
	OnScroll            string `htmlAttr:"onscroll" jsonSchema_description:""`
	OnSeeked            string `htmlAttr:"onseeked" jsonSchema_description:""`
	OnSeeking           string `htmlAttr:"onseeking" jsonSchema_description:""`
	OnSelect            string `htmlAttr:"onselect" jsonSchema_description:""`
	OnShow              string `htmlAttr:"onshow" jsonSchema_description:""`
	OnSort              string `htmlAttr:"onsort" jsonSchema_description:""`
	OnStalled           string `htmlAttr:"onstalled" jsonSchema_description:""`
	OnSubmit            string `htmlAttr:"onsubmit" jsonSchema_description:""`
	OnSuspend           string `htmlAttr:"onsuspend" jsonSchema_description:""`
	OnTimeUpdate        string `htmlAttr:"ontimeupdate" jsonSchema_description:""`
	OnToggle            string `htmlAttr:"ontoggle" jsonSchema_description:""`
	OnVolumeChange      string `htmlAttr:"onvolumechange" jsonSchema_description:""`
	OnWaiting           string `htmlAttr:"onwaiting" jsonSchema_description:""`

	*ToJavaScriptConverter `htmlAttr:"-" jsonSchema_description:""`
}

func (el *HtmlGlobalAttributes) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Id == "" {
		el.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)
	buffer.Write(data)

	return buffer.Bytes()
}
func (el *HtmlGlobalAttributes) GetId() []byte {
	var buffer bytes.Buffer

	if el.Id == "" {
		el.Id = GetAutoId()
	}

	buffer.WriteString(el.Id)

	return buffer.Bytes()
}
