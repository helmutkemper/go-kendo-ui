package telerik

type KendoOfflineStorage struct{
  GetItem                       String                    `jsObject:"getItem"`
  SetItem                       String                    `jsObject:"setItem"`
}

func(el KendoOfflineStorage) String() string {
  return `{ getItem: function() { return JSON.parse(sessionStorage.getItem("` + el.GetItem.String() + `")); }, setItem: function(item) { sessionStorage.setItem("` + el.SetItem.String() + `", JSON.stringify(item)); } }`
}