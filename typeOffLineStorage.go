package telerik

type OfflineStorage struct{
  GetItem                       String                    `jsObject:"getItem"`
  SetItem                       String                    `jsObject:"setItem"`
}

func(el OfflineStorage) String() string {
  return `{ getItem: function() { return JSON.parse(sessionStorage.getItem("` + el.GetItem.String() + `")); }, setItem: function(item) { sessionStorage.setItem("` + el.SetItem.String() + `", JSON.stringify(item)); } }`
}