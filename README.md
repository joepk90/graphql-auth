
### GraphQL Auth (Go)

### Configuration
The following configuration options are available;

|ENV_VAR|argument|Default Value|Description|
|-------|--------|-------------|-----------|
|HTTP_BIND|--http-bind|:80|The adapter:port to listen for plaintext connections|
|OPS_PORT|--ops-port|8081|Port to bind to for metrics ect.ect|
|LOG_LEVEL|--log-level|info|Logging verbosity|
|LOG_FORMAT|--log-format|nil|Logging format|
|AUTH_SERVICE_URL|--auth-service-url|http://localhost:8080|Authorization Service to connect too (Cerbos)|

### Available metrics
The following metrics are currently available for this service

|Name|Type|Vectors|Description|
|----|----|-------|-----------|
|graphql_query_total|Counter Vec|query_type=either query or mutation, method_name=The method that was called| count of successful request |
|graphql_query_error_total|Counter Vec|query_type=either query or mutation, method_name=The method that was called| a counter of graphql query errors |
|request_size|Counter| | Size of requests |
|response_time_(bucket/count/total)|Histogram| | Request response times |
|response_size|Counter| | Size of request responses |
|request_total|Counter| | Count of requests received |

### Building
The easiest way of building this service is to use `make docker-image` as this container will support the necessary build tools;
* go

If you want to build locally using `make build` ensure that you meet the above requirements as well as the correct version of openssl and libssl setup and available. Openssl version mis matches will cause either the build to fail of completely unexpected behaviour. 

### Running (local)
Run the build using the following command `make dev`. The Graphql service can then be access at [localhost:8080](http://localhost:8080)


### GraphQLi Authorization Headers
```
{
  "authorization": "Bearer XXX"
}
```
