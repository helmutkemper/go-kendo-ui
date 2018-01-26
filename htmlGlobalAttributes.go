package telerik

import (
  "bytes"
  "strconv"
)

// Global attributes are attributes common to all HTML elements; they can be used on all elements, though they may have
// no effect on some elements.
//
// Global attributes may be specified on all HTML elements, even those not specified in the standard. That means that
// any non-standard elements must still permit these attributes, even though using those elements means that the
// document is no longer HTML5-compliant. For example, HTML5-compliant browsers hide content marked as <foo hidden>...
// <foo>, even though <foo> is not a valid HTML element.
type HtmlGlobalAttributes struct{
  /*
  Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a
  space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.
  */
  AccessKey                   string

  /*
  Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access
  specific elements via the class selectors or functions like the method Document.getElementsByClassName().
  */
  Class                       string

  /*
  Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its
  widget to allow editing. The attribute must take one of the following values:
  true or the empty string, which indicates that the element must be editable;
  false, which indicates that the element must not be editable.
  */
  ContentEditable             Boolean

  /*
  Is the id of an <menu> to use as the contextual menu for this element.
  */
  ContextMenu                 string

  /*
  Forms a class of attributes, called custom data attributes, that allow proprietary information to be exchanged between
  the HTML and its DOM representation that may be used by scripts. All such custom data are available via the
  HTMLElement interface of the element the attribute is set on. The HTMLElement.dataset property gives access to them.
  */
  Data                        map[string]string

  /*
  Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:
  ltr, which means left to right and is to be used for languages that are written from the left to the right (like English);
  rtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);
  auto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until
  it finds a character with a strong directionality, then apply that directionality to the whole element.
  */
  Dir                         string

  /*
  Is an enumerated attribute indicating whether the element can be dragged, using the Drag and Drop API. It can have the
  following values:
  true, which indicates that the element may be dragged
  false, which indicates that the element may not be dragged.
  */
  Draggable                   Boolean

  /*
  Is an enumerated attribute indicating what types of content can be dropped on an element, using the Drag and Drop API.
  It can have the following values:
  copy, which indicates that dropping will create a copy of the element that was dragged
  move, which indicates that the element that was dragged will be moved to this new location.
  link, will create a link to the dragged data.
  */
  DropZone                    TypeHtmlDropZone

  /*
  Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used
  to hide elements of the page that can't be used until the login process has been completed. The browser won't render
  such elements. This attribute must not be used to hide content that could legitimately be shown.
  */
  Hidden                      Boolean

  /*
  Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element
  when linking (using a fragment identifier), scripting, or styling (with CSS).
  */
  Id                          string

  /*
  The unique, global identifier of an item.
  */
  ItemId                      string

  /*
  Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop
  consists of a name and value pair.
  */
  ItemDrop                    string

  /*
  Properties that are not descendants of an element with the itemscope attribute can be associated with the item using
  an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
  */
  ItemRef                     string

  /*
  itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
  item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of
  a vocabulary (such as schema.org) that describes the item and its properties context.
  */
  ItemScope                   string

  /*
  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure.
  itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
  */
  ItemType                    string

  /*
  Participates in defining the language of the element, the language that non-editable elements are written in or the
  language that editable elements should be written in. The tag contains one single entry value in the format defined in
  the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.
  */
  Lang                        string

  /*
  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot
  created by the <slot> element whose name attribute's value matches that slot attribute's value.
  */
  Sort                        string

  /*
  Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following
  values:
  true, which indicates that the element should be, if possible, checked for spelling errors;
  false, which indicates that the element should not be checked for spelling errors.
  */
  SpellCheck                  string

  /*
  Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined
  in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick
  styling, for example for testing purposes.
  */
  Style                       string

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
  TabIndex                    int

  /*
  Contains a text representing advisory information related to the element it belongs to. Such information can
  typically, but not necessarily, be presented to the user as a tooltip.
  */
  Title                       string

  /*
  Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text
  node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the
  following values:
  empty string and "yes", which indicates that the element will be translated.
  "no", which indicates that the element will not be translated.
  */
  Translate                   Boolean

  OnAbort					            string
  OnAutoComplete		    			string
  OnAutoCompleteError					string
  OnBlur            					string
  OnCancel					          string
  OnCanPlay		          			string
  OnCanPlayThrough			  		string
  OnChange	          				string
  OnClick					            string
  OnClose					            string
  OnContextMenu					      string
  OnCueChange					        string
  OnDblClick					        string
  OnDrag					            string
  OnDragEnd					          string
  OnDragEnter					        string
  OnDragExit					        string
  OnDragLeave					        string
  OnDragOver					        string
  OnDragStart					        string
  OnDrop					            string
  OnDurationChange					  string
  OnEmptied					          string
  OnEnded					            string
  OnError					            string
  OnFocus					            string
  OnInput					            string
  OnInvalid					          string
  OnKeyDown					          string
  OnKeyPress					        string
  OnKeyUp					            string
  OnLoad					            string
  OnLoadedData					      string
  OnLoadedMetadata					  string
  OnLoadStart					        string
  OnMouseDown					        string
  OnMouseEnter					      string
  OnMouseLeave					      string
  OnMouseMove					        string
  OnMouseOut					        string
  OnMouseOver					        string
  OnMouseUp					          string
  OnMouseWheel					      string
  OnPause					            string
  OnPlay					            string
  OnPlaying					          string
  OnProgress					        string
  OnRateChange					      string
  OnReset					            string
  OnResize					          string
  OnScroll					          string
  OnSeeked					          string
  OnSeeking					          string
  OnSelect					          string
  OnShow					            string
  OnSort					            string
  OnStalled					          string
  OnSubmit					          string
  OnSuspend					          string
  OnTimeUpdate					      string
  OnToggle					          string
  OnVolumeChange					    string
  OnWaiting                   string
}
func(el *HtmlGlobalAttributes)String() string {

  var out bytes.Buffer

  if el.AccessKey != "" {
    out.WriteString(`AccessKey="` + el.AccessKey + `" `)
  }

  if el.Class != "" {
    out.WriteString(`Class="` + el.Class + `" `)
  }

  if el.ContextMenu != "" {
    out.WriteString(`ContextMenu="` + el.ContextMenu + `" `)
  }

  if el.Dir != "" {
    out.WriteString(`Dir="` + el.Dir + `" `)
  }

  if el.Id != "" {
    out.WriteString(`Id="` + el.Id + `" `)
  }

  if el.ItemId != "" {
    out.WriteString(`ItemId="` + el.ItemId + `" `)
  }

  if el.ItemDrop != "" {
    out.WriteString(`ItemDrop="` + el.ItemDrop + `" `)
  }

  if el.ItemRef != "" {
    out.WriteString(`ItemRef="` + el.ItemRef + `" `)
  }

  if el.ItemScope != "" {
    out.WriteString(`ItemScope="` + el.ItemScope + `" `)
  }

  if el.ItemType != "" {
    out.WriteString(`ItemType="` + el.ItemType + `" `)
  }

  if el.Lang != "" {
    out.WriteString(`Lang="` + el.Lang + `" `)
  }

  if el.Sort != "" {
    out.WriteString(`Sort="` + el.Sort + `" `)
  }

  if el.SpellCheck != "" {
    out.WriteString(`SpellCheck="` + el.SpellCheck + `" `)
  }

  if el.Style != "" {
    out.WriteString(`Style="` + el.Style + `" `)
  }

  if el.Title != "" {
    out.WriteString(`Title="` + el.Title + `" `)
  }

  if el.OnAbort != "" {
    out.WriteString(`OnAbort="` + el.OnAbort + `" `)
  }

  if el.OnAutoComplete != "" {
    out.WriteString(`OnAutoComplete="` + el.OnAutoComplete + `" `)
  }

  if el.OnAutoCompleteError != "" {
    out.WriteString(`OnAutoCompleteError="` + el.OnAutoCompleteError + `" `)
  }

  if el.OnBlur != "" {
    out.WriteString(`OnBlur="` + el.OnBlur + `" `)
  }

  if el.OnCancel != "" {
    out.WriteString(`OnCancel="` + el.OnCancel + `" `)
  }

  if el.OnCanPlay != "" {
    out.WriteString(`OnCanPlay="` + el.OnCanPlay + `" `)
  }

  if el.OnCanPlayThrough != "" {
    out.WriteString(`OnCanPlayThrough="` + el.OnCanPlayThrough + `" `)
  }

  if el.OnChange != "" {
    out.WriteString(`OnChange="` + el.OnChange + `" `)
  }

  if el.OnClick != "" {
    out.WriteString(`OnClick="` + el.OnClick + `" `)
  }

  if el.OnClose != "" {
    out.WriteString(`OnClose="` + el.OnClose + `" `)
  }

  if el.OnContextMenu != "" {
    out.WriteString(`OnContextMenu="` + el.OnContextMenu + `" `)
  }

  if el.OnCueChange != "" {
    out.WriteString(`OnCueChange="` + el.OnCueChange + `" `)
  }

  if el.OnDblClick != "" {
    out.WriteString(`OnDblClick="` + el.OnDblClick + `" `)
  }

  if el.OnDrag != "" {
    out.WriteString(`OnDrag="` + el.OnDrag + `" `)
  }

  if el.OnDragEnd != "" {
    out.WriteString(`OnDragEnd="` + el.OnDragEnd + `" `)
  }

  if el.OnDragEnter != "" {
    out.WriteString(`OnDragEnter="` + el.OnDragEnter + `" `)
  }

  if el.OnDragExit != "" {
    out.WriteString(`OnDragExit="` + el.OnDragExit + `" `)
  }

  if el.OnDragLeave != "" {
    out.WriteString(`OnDragLeave="` + el.OnDragLeave + `" `)
  }

  if el.OnDragOver != "" {
    out.WriteString(`OnDragOver="` + el.OnDragOver + `" `)
  }

  if el.OnDragStart != "" {
    out.WriteString(`OnDragStart="` + el.OnDragStart + `" `)
  }

  if el.OnDrop != "" {
    out.WriteString(`OnDrop="` + el.OnDrop + `" `)
  }

  if el.OnDurationChange != "" {
    out.WriteString(`OnDurationChange="` + el.OnDurationChange + `" `)
  }

  if el.OnEmptied != "" {
    out.WriteString(`OnEmptied="` + el.OnEmptied + `" `)
  }

  if el.OnEnded != "" {
    out.WriteString(`OnEnded="` + el.OnEnded + `" `)
  }

  if el.OnError != "" {
    out.WriteString(`OnError="` + el.OnError + `" `)
  }

  if el.OnFocus != "" {
    out.WriteString(`OnFocus="` + el.OnFocus + `" `)
  }

  if el.OnInput != "" {
    out.WriteString(`OnInput="` + el.OnInput + `" `)
  }

  if el.OnInvalid != "" {
    out.WriteString(`OnInvalid="` + el.OnInvalid + `" `)
  }

  if el.OnKeyDown != "" {
    out.WriteString(`OnKeyDown="` + el.OnKeyDown + `" `)
  }

  if el.OnKeyPress != "" {
    out.WriteString(`OnKeyPress="` + el.OnKeyPress + `" `)
  }

  if el.OnKeyUp != "" {
    out.WriteString(`OnKeyUp="` + el.OnKeyUp + `" `)
  }

  if el.OnLoad != "" {
    out.WriteString(`OnLoad="` + el.OnLoad + `" `)
  }

  if el.OnLoadedData != "" {
    out.WriteString(`OnLoadedData="` + el.OnLoadedData + `" `)
  }

  if el.OnLoadedMetadata != "" {
    out.WriteString(`OnLoadedMetadata="` + el.OnLoadedMetadata + `" `)
  }

  if el.OnLoadStart != "" {
    out.WriteString(`OnLoadStart="` + el.OnLoadStart + `" `)
  }

  if el.OnMouseDown != "" {
    out.WriteString(`OnMouseDown="` + el.OnMouseDown + `" `)
  }

  if el.OnMouseEnter != "" {
    out.WriteString(`OnMouseEnter="` + el.OnMouseEnter + `" `)
  }

  if el.OnMouseLeave != "" {
    out.WriteString(`OnMouseLeave="` + el.OnMouseLeave + `" `)
  }

  if el.OnMouseMove != "" {
    out.WriteString(`OnMouseMove="` + el.OnMouseMove + `" `)
  }

  if el.OnMouseOut != "" {
    out.WriteString(`OnMouseOut="` + el.OnMouseOut + `" `)
  }

  if el.OnMouseOver != "" {
    out.WriteString(`OnMouseOver="` + el.OnMouseOver + `" `)
  }

  if el.OnMouseUp != "" {
    out.WriteString(`OnMouseUp="` + el.OnMouseUp + `" `)
  }

  if el.OnMouseWheel != "" {
    out.WriteString(`OnMouseWheel="` + el.OnMouseWheel + `" `)
  }

  if el.OnPause != "" {
    out.WriteString(`OnPause="` + el.OnPause + `" `)
  }

  if el.OnPlay != "" {
    out.WriteString(`OnPlay="` + el.OnPlay + `" `)
  }

  if el.OnPlaying != "" {
    out.WriteString(`OnPlaying="` + el.OnPlaying + `" `)
  }

  if el.OnProgress != "" {
    out.WriteString(`OnProgress="` + el.OnProgress + `" `)
  }

  if el.OnRateChange != "" {
    out.WriteString(`OnRateChange="` + el.OnRateChange + `" `)
  }

  if el.OnReset != "" {
    out.WriteString(`OnReset="` + el.OnReset + `" `)
  }

  if el.OnResize != "" {
    out.WriteString(`OnResize="` + el.OnResize + `" `)
  }

  if el.OnScroll != "" {
    out.WriteString(`OnScroll="` + el.OnScroll + `" `)
  }

  if el.OnSeeked != "" {
    out.WriteString(`OnSeeked="` + el.OnSeeked + `" `)
  }

  if el.OnSeeking != "" {
    out.WriteString(`OnSeeking="` + el.OnSeeking + `" `)
  }

  if el.OnSelect != "" {
    out.WriteString(`OnSelect="` + el.OnSelect + `" `)
  }

  if el.OnShow != "" {
    out.WriteString(`OnShow="` + el.OnShow + `" `)
  }

  if el.OnSort != "" {
    out.WriteString(`OnSort="` + el.OnSort + `" `)
  }

  if el.OnStalled != "" {
    out.WriteString(`OnStalled="` + el.OnStalled + `" `)
  }

  if el.OnSubmit != "" {
    out.WriteString(`OnSubmit="` + el.OnSubmit + `" `)
  }

  if el.OnSuspend != "" {
    out.WriteString(`OnSuspend="` + el.OnSuspend + `" `)
  }

  if el.OnTimeUpdate != "" {
    out.WriteString(`OnTimeUpdate="` + el.OnTimeUpdate + `" `)
  }

  if el.OnToggle != "" {
    out.WriteString(`OnToggle="` + el.OnToggle + `" `)
  }

  if el.OnVolumeChange != "" {
    out.WriteString(`OnVolumeChange="` + el.OnVolumeChange + `" `)
  }

  if el.OnWaiting != "" {
    out.WriteString(`OnWaiting="` + el.OnWaiting + `" `)
  }

  if el.ContentEditable.IsSet() {
    out.WriteString(`ContentEditable="` + el.ContentEditable.String() + `" `)
  }

  if el.Draggable.IsSet() {
    out.WriteString(`Draggable="` + el.Draggable.String() + `" `)
  }

  if el.Hidden.IsSet() {
    out.WriteString(`Hidden="` + el.Hidden.String() + `" `)
  }

  if el.Translate.IsSet() {
    out.WriteString(`Translate="` + el.Translate.String() + `" `)
  }

  for k, v := range el.Data{
    out.WriteString(`data-` + k + `="` + v + `" `)
  }

  if el.TabIndex != 0 {
    out.WriteString(`TabIndex="` + strconv.Itoa( el.TabIndex ) + `" `)
  }

  if el.DropZone != 0 {
    out.WriteString(`DropZone="` + el.DropZone.String() + `" `)
  }

  return out.String()
}
