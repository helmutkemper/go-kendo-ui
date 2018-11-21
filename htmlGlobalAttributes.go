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
	DoNotUseThisFieldOmitHtml Boolean `htmlAttr:"-" description:""`
	/*
	  Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a
	  space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.
	*/
	AccessKey string `htmlAttr:"accesskey" description:""`

	/*
	  Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access
	  specific elements via the class selectors or functions like the method Document.getElementsByClassName().
	*/
	Class string `htmlAttr:"class" description:""`

	/*
	  Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its
	  widget to allow editing. The attribute must take one of the following values:
	  true or the empty string, which indicates that the element must be editable;
	  false, which indicates that the element must not be editable.
	*/
	ContentEditable Boolean `htmlAttr:"contenteditable" description:""`

	/*
	  Is the id of an <menu> to use as the contextual menu for this element.
	*/
	ContextMenu string `htmlAttr:"contextmenu" description:""`

	/*
	  Forms a class of attributes, called custom data attributes, that allow proprietary information to be exchanged between
	  the HTML and its DOM representation that may be used by scripts. All such custom data are available via the
	  HTMLElement interface of the element the attribute is set on. The HTMLElement.dataset property gives access to them.
	*/
	Data map[string]string `htmlAttr:"data" description:""`

	/*
	  Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:
	  ltr, which means left to right and is to be used for languages that are written from the left to the right (like English);
	  rtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);
	  auto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until
	  it finds a character with a strong directionality, then apply that directionality to the whole element.
	*/
	Dir string `htmlAttr:"dir" description:""`

	/*
	  Is an enumerated attribute indicating whether the element can be dragged, using the Drag and Drop API. It can have the
	  following values:
	  true, which indicates that the element may be dragged
	  false, which indicates that the element may not be dragged.
	*/
	Draggable Boolean `htmlAttr:"draggable" description:""`

	/*
	  Is an enumerated attribute indicating what types of content can be dropped on an element, using the Drag and Drop API.
	  It can have the following values:
	  copy, which indicates that dropping will create a copy of the element that was dragged
	  move, which indicates that the element that was dragged will be moved to this new location.
	  link, will create a link to the dragged data.
	*/
	DropZone TypeHtmlDropZone `htmlAttr:"dropzone" description:""`

	/*
	  Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used
	  to hide elements of the page that can't be used until the login process has been completed. The browser won't render
	  such elements. This attribute must not be used to hide content that could legitimately be shown.
	*/
	Hidden Boolean `htmlAttrSet:"draggable" description:""`

	/*
	  Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element
	  when linking (using a fragment identifier), scripting, or styling (with CSS).
	*/
	Id string `htmlAttr:"id" description:""`

	/*
	  The unique, global identifier of an item.
	*/
	ItemId string `htmlAttr:"itemid" description:""`

	/*
	  Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop
	  consists of a name and value pair.
	*/
	ItemDrop string `htmlAttr:"itemdrop" description:""`

	/*
	  Properties that are not descendants of an element with the itemscope attribute can be associated with the item using
	  an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
	*/
	ItemRef string `htmlAttr:"itemref" description:""`

	/*
	  itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
	  item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of
	  a vocabulary (such as schema.org) that describes the item and its properties context.
	*/
	ItemScope string `htmlAttr:"itemscope" description:""`

	/*
	  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure.
	  itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
	*/
	ItemType string `htmlAttr:"itemtype" description:""`

	/*
	  Participates in defining the language of the element, the language that non-editable elements are written in or the
	  language that editable elements should be written in. The tag contains one single entry value in the format defined in
	  the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.
	*/
	Lang string `htmlAttr:"lang" description:""`

	/*
	  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot
	  created by the <slot> element whose name attribute's value matches that slot attribute's value.
	*/
	Sort string `htmlAttr:"sort" description:""`

	/*
	  Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following
	  values:
	  true, which indicates that the element should be, if possible, checked for spelling errors;
	  false, which indicates that the element should not be checked for spelling errors.
	*/
	SpellCheck string `htmlAttr:"spellcheck" description:""`

	/*
	  Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined
	  in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick
	  styling, for example for testing purposes.
	*/
	Style string `htmlAttr:"style" description:""`

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
	TabIndex int `htmlAttr:"tabindex" description:""`

	/*
	  Contains a text representing advisory information related to the element it belongs to. Such information can
	  typically, but not necessarily, be presented to the user as a tooltip.
	*/
	Title string `htmlAttr:"title" description:""`

	/*
	  Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text
	  node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the
	  following values:
	  empty string and "yes", which indicates that the element will be translated.
	  "no", which indicates that the element will not be translated.
	*/
	Translate Boolean `htmlAttr:"translate" description:""`

	Extra map[string]interface{} `htmlAttr:"extra" description:""`

	OnAbort             string `htmlAttr:"onabort" description:""`
	OnAutoComplete      string `htmlAttr:"onautocomplete" description:""`
	OnAutoCompleteError string `htmlAttr:"onautocompleteerror" description:""`
	OnBlur              string `htmlAttr:"onblur" description:""`
	OnCancel            string `htmlAttr:"oncancel" description:""`
	OnCanPlay           string `htmlAttr:"oncanplay" description:""`
	OnCanPlayThrough    string `htmlAttr:"oncanplaythrough" description:""`
	OnChange            string `htmlAttr:"onchange" description:""`
	OnClick             string `htmlAttr:"onclick" description:""`
	OnClose             string `htmlAttr:"onclose" description:""`
	OnContextMenu       string `htmlAttr:"oncontextmenu" description:""`
	OnCueChange         string `htmlAttr:"oncuechange" description:""`
	OnDblClick          string `htmlAttr:"ondblclick" description:""`
	OnDrag              string `htmlAttr:"ondrag" description:""`
	OnDragEnd           string `htmlAttr:"ondragend" description:""`
	OnDragEnter         string `htmlAttr:"ondragenter" description:""`
	OnDragExit          string `htmlAttr:"ondragexit" description:""`
	OnDragLeave         string `htmlAttr:"ondragleave" description:""`
	OnDragOver          string `htmlAttr:"ondragover" description:""`
	OnDragStart         string `htmlAttr:"ondragstart" description:""`
	OnDrop              string `htmlAttr:"ondrop" description:""`
	OnDurationChange    string `htmlAttr:"ondurationchange" description:""`
	OnEmptied           string `htmlAttr:"onemptied" description:""`
	OnEnded             string `htmlAttr:"onended" description:""`
	OnError             string `htmlAttr:"onerror" description:""`
	OnFocus             string `htmlAttr:"onfocus" description:""`
	OnInput             string `htmlAttr:"oninput" description:""`
	OnInvalid           string `htmlAttr:"oninvalid" description:""`
	OnKeyDown           string `htmlAttr:"onkeydown" description:""`
	OnKeyPress          string `htmlAttr:"onkeypress" description:""`
	OnKeyUp             string `htmlAttr:"onkeyup" description:""`
	OnLoad              string `htmlAttr:"onload" description:""`
	OnLoadedData        string `htmlAttr:"onloadeddata" description:""`
	OnLoadedMetadata    string `htmlAttr:"onloadedmetadata" description:""`
	OnLoadStart         string `htmlAttr:"onloadstart" description:""`
	OnMouseDown         string `htmlAttr:"onmousedown" description:""`
	OnMouseEnter        string `htmlAttr:"onmouseenter" description:""`
	OnMouseLeave        string `htmlAttr:"onmouseleave" description:""`
	OnMouseMove         string `htmlAttr:"onmousemove" description:""`
	OnMouseOut          string `htmlAttr:"onmouseout" description:""`
	OnMouseOver         string `htmlAttr:"onmouseover" description:""`
	OnMouseUp           string `htmlAttr:"onmouseup" description:""`
	OnMouseWheel        string `htmlAttr:"onmousewheel" description:""`
	OnPause             string `htmlAttr:"onpause" description:""`
	OnPlay              string `htmlAttr:"onplay" description:""`
	OnPlaying           string `htmlAttr:"onplaying" description:""`
	OnProgress          string `htmlAttr:"onprogress" description:""`
	OnRateChange        string `htmlAttr:"onratechange" description:""`
	OnReset             string `htmlAttr:"onreset" description:""`
	OnResize            string `htmlAttr:"onresize" description:""`
	OnScroll            string `htmlAttr:"onscroll" description:""`
	OnSeeked            string `htmlAttr:"onseeked" description:""`
	OnSeeking           string `htmlAttr:"onseeking" description:""`
	OnSelect            string `htmlAttr:"onselect" description:""`
	OnShow              string `htmlAttr:"onshow" description:""`
	OnSort              string `htmlAttr:"onsort" description:""`
	OnStalled           string `htmlAttr:"onstalled" description:""`
	OnSubmit            string `htmlAttr:"onsubmit" description:""`
	OnSuspend           string `htmlAttr:"onsuspend" description:""`
	OnTimeUpdate        string `htmlAttr:"ontimeupdate" description:""`
	OnToggle            string `htmlAttr:"ontoggle" description:""`
	OnVolumeChange      string `htmlAttr:"onvolumechange" description:""`
	OnWaiting           string `htmlAttr:"onwaiting" description:""`

	*ToJavaScriptConverter `htmlAttr:"-" description:""`
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
