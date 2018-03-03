package telerik

type KendoTransport struct{
  /*
  The configuration used when the data source destroys data items. Those are items removed from the data source via the remove method.
  The data source uses jQuery.ajax to make an HTTP request to the remote service. The value configured via transport.destroy is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via transport.destroy except the success and error callback functions which are used by the transport.

  If the value of transport.destroy is a function, the data source invokes that function instead of jQuery.ajax.

  If the value of transport.destroy is a string, the data source uses this string as the URL of the remote service.
  All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as objects. Mixing the different configuration alternatives is not possible.

  Example - set the destroy remote service
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      read: {
        url: "https://demos.telerik.com/kendo-ui/service/products",
        dataType: "jsonp"
      },
      // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/destroy
      destroy: {
        url: "https://demos.telerik.com/kendo-ui/service/products/destroy",
        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
      },
      parameterMap: function(data, type) {
        if (type == "destroy") {
          // send the destroyed data items as the "models" service parameter encoded in JSON
          return { models: kendo.stringify(data.models) }
        }
      }
    },
    batch: true,
    schema: {
      model: { id: "ProductID" }
    }
  });
  dataSource.fetch(function() {
    var products = dataSource.data();
    // remove the first data item
    dataSource.remove(products[0]);
    // send the destroyed data item to the remote service
    dataSource.sync();
  });
  </script>

  Example - set destroy as a function
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      read: function(options) {
        $.ajax({
          url: "https://demos.telerik.com/kendo-ui/service/products",
          dataType: "jsonp",
          success: function(result) {
            options.success(result);
          }
        });
      },
      destroy: function (options) {
        // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/destroy
        $.ajax({
          url: "https://demos.telerik.com/kendo-ui/service/products/destroy",
          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          // send the destroyed data items as the "models" service parameter encoded in JSON
          data: {
            models: kendo.stringify(options.data.models)
          },
          success: function(result) {
            // notify the data source that the request succeeded
            options.success(result);
          },
          error: function(result) {
            // notify the data source that the request failed
            options.error(result);
          }
        });
      }
    },
    batch: true,
    schema: {
      model: { id: "ProductID" }
    }
  });
  dataSource.fetch(function() {
    var products = dataSource.data();
    dataSource.remove(products[0]);
    dataSource.sync();
  });
  </script>
  */
  Destroy                                         *KendoDestroy                                 `jsObject:"destroy" jsType:"*KendoDestroy,*JavaScript,string"`
  ParameterMap                                    *KendoDestroy                                 `jsObject:"parameterMap"`
}
