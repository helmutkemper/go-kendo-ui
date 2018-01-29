package telerik

import (
  "bytes"
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
  AccessKey                   String

  /*
  Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access
  specific elements via the class selectors or functions like the method Document.getElementsByClassName().
  */
  Class                       String

  /*
  Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its
  widget to allow editing. The attribute must take one of the following values:
  true or the empty String, which indicates that the element must be editable;
  false, which indicates that the element must not be editable.
  */
  ContentEditable             Boolean

  /*
  Is the id of an <menu> to use as the contextual menu for this element.
  */
  ContextMenu                 String

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
  Dir                         String

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
  Id                          String

  /*
  The unique, global identifier of an item.
  */
  ItemId                      String

  /*
  Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop
  consists of a name and value pair.
  */
  ItemDrop                    String

  /*
  Properties that are not descendants of an element with the itemscope attribute can be associated with the item using
  an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
  */
  ItemRef                     String

  /*
  itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
  item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of
  a vocabulary (such as schema.org) that describes the item and its properties context.
  */
  ItemScope                   String

  /*
  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure.
  itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
  */
  ItemType                    String

  /*
  Participates in defining the language of the element, the language that non-editable elements are written in or the
  language that editable elements should be written in. The tag contains one single entry value in the format defined in
  the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.
  */
  Lang                        String

  /*
  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot
  created by the <slot> element whose name attribute's value matches that slot attribute's value.
  */
  Sort                        String

  /*
  Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following
  values:
  true, which indicates that the element should be, if possible, checked for spelling errors;
  false, which indicates that the element should not be checked for spelling errors.
  */
  SpellCheck                  String

  /*
  Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined
  in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick
  styling, for example for testing purposes.
  */
  Style                       String

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
  TabIndex                    Int

  /*
  Contains a text representing advisory information related to the element it belongs to. Such information can
  typically, but not necessarily, be presented to the user as a tooltip.
  */
  Title                       String

  /*
  Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text
  node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the
  following values:
  empty String and "yes", which indicates that the element will be translated.
  "no", which indicates that the element will not be translated.
  */
  Translate                   Boolean

  Extra                       map[string]string

  OnAbort					            String
  OnAutoComplete		    			String
  OnAutoCompleteError					String
  OnBlur            					String
  OnCancel					          String
  OnCanPlay		          			String
  OnCanPlayThrough			  		String
  OnChange	          				String
  OnClick					            String
  OnClose					            String
  OnContextMenu					      String
  OnCueChange					        String
  OnDblClick					        String
  OnDrag					            String
  OnDragEnd					          String
  OnDragEnter					        String
  OnDragExit					        String
  OnDragLeave					        String
  OnDragOver					        String
  OnDragStart					        String
  OnDrop					            String
  OnDurationChange					  String
  OnEmptied					          String
  OnEnded					            String
  OnError					            String
  OnFocus					            String
  OnInput					            String
  OnInvalid					          String
  OnKeyDown					          String
  OnKeyPress					        String
  OnKeyUp					            String
  OnLoad					            String
  OnLoadedData					      String
  OnLoadedMetadata					  String
  OnLoadStart					        String
  OnMouseDown					        String
  OnMouseEnter					      String
  OnMouseLeave					      String
  OnMouseMove					        String
  OnMouseOut					        String
  OnMouseOver					        String
  OnMouseUp					          String
  OnMouseWheel					      String
  OnPause					            String
  OnPlay					            String
  OnPlaying					          String
  OnProgress					        String
  OnRateChange					      String
  OnReset					            String
  OnResize					          String
  OnScroll					          String
  OnSeeked					          String
  OnSeeking					          String
  OnSelect					          String
  OnShow					            String
  OnSort					            String
  OnStalled					          String
  OnSubmit					          String
  OnSuspend					          String
  OnTimeUpdate					      String
  OnToggle					          String
  OnVolumeChange					    String
  OnWaiting                   String
}
func(el *HtmlGlobalAttributes)String() string {

  var out bytes.Buffer

  if el.AccessKey != "" {
    out.WriteString( el.AccessKey.ToAttr("accesskey") )
  }

  if el.Class != "" {
    out.WriteString( el.Class.ToAttr("class") )
  }

  if el.ContextMenu != "" {
    out.WriteString( el.ContextMenu.ToAttr("contextmenu") )
  }

  if el.Dir != "" {
    out.WriteString( el.Dir.ToAttr("dir") )
  }

  if el.Id != "" {
    out.WriteString( el.Id.ToAttr("id") )
  }

  if el.ItemId != "" {
    out.WriteString( el.ItemId.ToAttr("itemid") )
  }

  if el.ItemDrop != "" {
    out.WriteString( el.ItemDrop.ToAttr("itemdrop") )
  }

  if el.ItemRef != "" {
    out.WriteString( el.ItemRef.ToAttr("itemref") )
  }

  if el.ItemScope != "" {
    out.WriteString( el.ItemScope.ToAttr("itemscope") )
  }

  if el.ItemType != "" {
    out.WriteString( el.ItemType.ToAttr("itemtype") )
  }

  if el.Lang != "" {
    out.WriteString( el.Lang.ToAttr("lang") )
  }

  if el.Sort != "" {
    out.WriteString( el.Sort.ToAttr("Sort") )
  }

  if el.SpellCheck != "" {
    out.WriteString( el.SpellCheck.ToAttr("spellcheck") )
  }

  if el.Style != "" {
    out.WriteString( el.Style.ToAttr("style") )
  }

  if el.Title != "" {
    out.WriteString( el.Title.ToAttr("title") )
  }

  if el.OnAbort != "" {
    out.WriteString( el.OnAbort.ToAttr("onabort") )
  }

  if el.OnAutoComplete != "" {
    out.WriteString( el.OnAutoComplete.ToAttr("onautocomplete") )
  }

  if el.OnAutoCompleteError != "" {
    out.WriteString( el.OnAutoCompleteError.ToAttr("onautocompleteerror") )
  }

  if el.OnBlur != "" {
    out.WriteString( el.OnBlur.ToAttr("onblur") )
  }

  if el.OnCancel != "" {
    out.WriteString( el.OnCancel.ToAttr("oncancel") )
  }

  if el.OnCanPlay != "" {
    out.WriteString( el.OnCanPlay.ToAttr("oncanplay") )
  }

  if el.OnCanPlayThrough != "" {
    out.WriteString( el.OnCanPlayThrough.ToAttr("oncanplaythrough") )
  }

  if el.OnChange != "" {
    out.WriteString( el.OnChange.ToAttr("onchange") )
  }

  if el.OnClick != "" {
    out.WriteString( el.OnClick.ToAttr("onclick") )
  }

  if el.OnClose != "" {
    out.WriteString( el.OnClose.ToAttr("onclose") )
  }

  if el.OnContextMenu != "" {
    out.WriteString( el.OnContextMenu.ToAttr("oncontextmenu") )
  }

  if el.OnCueChange != "" {
    out.WriteString( el.OnCueChange.ToAttr("oncuechange") )
  }

  if el.OnDblClick != "" {
    out.WriteString( el.OnDblClick.ToAttr("ondblclick") )
  }

  if el.OnDrag != "" {
    out.WriteString( el.OnDrag.ToAttr("ondrag") )
  }

  if el.OnDragEnd != "" {
    out.WriteString( el.OnDragEnd.ToAttr("ondragend") )
  }

  if el.OnDragEnter != "" {
    out.WriteString( el.OnDragEnter.ToAttr("ondragenter") )
  }

  if el.OnDragExit != "" {
    out.WriteString( el.OnDragExit.ToAttr("ondragexit") )
  }

  if el.OnDragLeave != "" {
    out.WriteString( el.OnDragLeave.ToAttr("ondragleave") )
  }

  if el.OnDragOver != "" {
    out.WriteString( el.OnDragOver.ToAttr("ondragover") )
  }

  if el.OnDragStart != "" {
    out.WriteString( el.OnDragStart.ToAttr("ondragstart") )
  }

  if el.OnDrop != "" {
    out.WriteString( el.OnDrop.ToAttr("ondrop") )
  }

  if el.OnDurationChange != "" {
    out.WriteString( el.OnDurationChange.ToAttr("ondurationchange") )
  }

  if el.OnEmptied != "" {
    out.WriteString( el.OnEmptied.ToAttr("onemptied") )
  }

  if el.OnEnded != "" {
    out.WriteString( el.OnEnded.ToAttr("onended") )
  }

  if el.OnError != "" {
    out.WriteString( el.OnError.ToAttr("onerror") )
  }

  if el.OnFocus != "" {
    out.WriteString( el.OnFocus.ToAttr("onfocus") )
  }

  if el.OnInput != "" {
    out.WriteString( el.OnInput.ToAttr("oninput") )
  }

  if el.OnInvalid != "" {
    out.WriteString( el.OnInvalid.ToAttr("oninvalid") )
  }

  if el.OnKeyDown != "" {
    out.WriteString( el.OnKeyDown.ToAttr("onkeydown") )
  }

  if el.OnKeyPress != "" {
    out.WriteString( el.OnKeyPress.ToAttr("onkeypress") )
  }

  if el.OnKeyUp != "" {
    out.WriteString( el.OnKeyUp.ToAttr("onkeyup") )
  }

  if el.OnLoad != "" {
    out.WriteString( el.OnLoad.ToAttr("onload") )
  }

  if el.OnLoadedData != "" {
    out.WriteString( el.OnLoadedData.ToAttr("onloadeddata") )
  }

  if el.OnLoadedMetadata != "" {
    out.WriteString( el.OnLoadedMetadata.ToAttr("onloadedmetadata") )
  }

  if el.OnLoadStart != "" {
    out.WriteString( el.OnLoadStart.ToAttr("onloadstart") )
  }

  if el.OnMouseDown != "" {
    out.WriteString( el.OnMouseDown.ToAttr("onmousedown") )
  }

  if el.OnMouseEnter != "" {
    out.WriteString( el.OnMouseEnter.ToAttr("onmouseenter") )
  }

  if el.OnMouseLeave != "" {
    out.WriteString( el.OnMouseLeave.ToAttr("onmouseleave") )
  }

  if el.OnMouseMove != "" {
    out.WriteString( el.OnMouseMove.ToAttr("onmousemove") )
  }

  if el.OnMouseOut != "" {
    out.WriteString( el.OnMouseOut.ToAttr("onmouseout") )
  }

  if el.OnMouseOver != "" {
    out.WriteString( el.OnMouseOver.ToAttr("onmouseover") )
  }

  if el.OnMouseUp != "" {
    out.WriteString( el.OnMouseUp.ToAttr("onmouseup") )
  }

  if el.OnMouseWheel != "" {
    out.WriteString( el.OnMouseWheel.ToAttr("onmousewheel") )
  }

  if el.OnPause != "" {
    out.WriteString( el.OnPause.ToAttr("onpause") )
  }

  if el.OnPlay != "" {
    out.WriteString( el.OnPlay.ToAttr("onplay") )
  }

  if el.OnPlaying != "" {
    out.WriteString( el.OnPlaying.ToAttr("onplaying") )
  }

  if el.OnProgress != "" {
    out.WriteString( el.OnProgress.ToAttr("onprogress") )
  }

  if el.OnRateChange != "" {
    out.WriteString( el.OnRateChange.ToAttr("onratechange") )
  }

  if el.OnReset != "" {
    out.WriteString( el.OnReset.ToAttr("onreset") )
  }

  if el.OnResize != "" {
    out.WriteString( el.OnResize.ToAttr("OnResize") )
  }

  if el.OnScroll != "" {
    out.WriteString( el.OnScroll.ToAttr("onscroll") )
  }

  if el.OnSeeked != "" {
    out.WriteString( el.OnSeeked.ToAttr("onseeked") )
  }

  if el.OnSeeking != "" {
    out.WriteString( el.OnSeeking.ToAttr("onseeking") )
  }

  if el.OnSelect != "" {
    out.WriteString( el.OnSelect.ToAttr("onselect") )
  }

  if el.OnShow != "" {
    out.WriteString( el.OnShow.ToAttr("onshow") )
  }

  if el.OnSort != "" {
    out.WriteString( el.OnSort.ToAttr("onsort") )
  }

  if el.OnStalled != "" {
    out.WriteString( el.OnStalled.ToAttr("onstalled") )
  }

  if el.OnSubmit != "" {
    out.WriteString( el.OnSubmit.ToAttr("onsubmit") )
  }

  if el.OnSuspend != "" {
    out.WriteString( el.OnSuspend.ToAttr("onsuspend") )
  }

  if el.OnTimeUpdate != "" {
    out.WriteString( el.OnTimeUpdate.ToAttr("ontimeupdate") )
  }

  if el.OnToggle != "" {
    out.WriteString( el.OnToggle.ToAttr("ontoggle") )
  }

  if el.OnVolumeChange != "" {
    out.WriteString( el.OnVolumeChange.ToAttr("onvolumechange") )
  }

  if el.OnWaiting != "" {
    out.WriteString( el.OnWaiting.ToAttr("onwaiting") )
  }

  if el.ContentEditable.IsSet() {
    out.WriteString( el.ContentEditable.ToAttr("contenteditable") )
  }

  if el.Draggable.IsSet() {
    out.WriteString(el.Draggable.ToAttr("draggable") )
  }

  if el.Hidden.IsSet() {
    out.WriteString( el.Hidden.ToAttrSet("hidden") )
  }

  if el.Translate.IsSet() {
    out.WriteString( el.Translate.ToAttr("translate") )
  }

  for k, v := range el.Data{
    out.WriteString(`data-` + k + `="` + v + `" `)
  }

  for k, v := range el.Extra{
    out.WriteString(k + `="` + v + `" `)
  }

  if el.TabIndex != 0 {
    out.WriteString( el.TabIndex.ToAttr("tabindex") )
  }

  if el.DropZone != 0 {
    out.WriteString(`DropZone="` + el.DropZone.String() + `" `)
  }

  return out.String()
}
