***Konnect APP***

This api service will load the user's API services pages which they have added
to their Konnect portal to check the service status, their UP time, p99 latency,
API flows.

As part of this exercise developed following flow:
1. **Load API service homepage** with all the available services paginated and sort by default by updated_at desc.
   In response, we will get all the services as per limit and pager number with 200 response
   ````   
   Request: http://127.0.0.1:8000/api/api-services?page=2&per_page=3
   Response:
   {
      "data": {
      "service_detail": [
         {
            "id": "5",
            "name": "API Service 5",
            "description": "API Svc 1 desc test data",
            "published": false,
            "version_count": "0"
         },
         {
            "id": "6",
            "name": "API Service 6",
            "description": "API Svc 1 desc test data",
            "published": false,
            "version_count": "0"
         },
         {
            "id": "1",
            "name": "API Service 1",
            "description": "API Svc 1 desc test data",
            "published": false,
            "version_count": "0"
         }
         ]
      },
      "success": true,
      "errors": null
   }
2. **Search over the Service Name**. This api response is also paginated and sorted by updated_at desc.
   In response, we will get all the services matching the filter as per limit and pager number with 200 response
    ````   
   Request: http://127.0.0.1:8000/api/api-services?filter[name][contains]=Service%201&per_page=5
   Response:
   {
      "data": {
      "service_detail": [
            {
               "id": "10",
               "name": "API Service 10",
               "description": "API Svc 1 desc test data",
               "published": false,
               "version_count": "0"
            },
            {
               "id": "1",
               "name": "API Service 1",
               "description": "API Svc 1 desc test data",
               "published": false,
               "version_count": "0"
            }
         ]
      },
      "success": true,
      "errors": null
   }
3. **Get the service detail** per service where user will be able to view the service detail using serviceID in request param.
    ````   
   Request: http://127.0.0.1:8000/api/api-services/6
   Response: 
    {
        "data": {
        "id": "6",
        "name": "API Service 6",
        "description": "API Svc 1 desc test data",
        "published": false,
        "version_count": "0"
        },
        "success": true
    }

**PreRequisite:**

Golang Version: 1.22

Docker Up and Running for postgres

**To start the service**: `go build api`

**DATABASE**

Used **Postgresql:14** as database as the schema is predefined and in future case it will joins in adjoining child tables

**DB SCHEMA**
Link: https://dbdiagram.io/d/Konnect-DB-Schema-667c0a5e9939893dae52f94e
Here we have Services and Versions table which we have implemented. Documents and ServiceData can be added for extending
the functionality to provided more details for the api services and versions.

**TEST CASES**

1. UnitTests: Have mentioned what all unit test needed to be tested at the function level in the respective `*_test.go` file eg: `api_service_test.go`
2. Integration Tests: As part of this project we will have Two integration level test for each api's
   1. `/api/api-products` (To get all the api services matching criteria). It will test
      1. if invalid params are passed then 4xx error should come
      2. if all the valid and some error at service level then error should be handled
      3. if all the correct and services are found matching criteria then return []services with 200 success response
   4. `/api/api-products/` (To get all the service detail matching criteria). It will test
      1. if invalid params are passed then 4xx error should come
      2. if all the valid and some error at service level then error should be handled
      3. if all the correct and services are found matching criteria then return service details with 200 success response

**ENHANCEMENTS**

1. When version for a service will be added, we will have a transaction which will insert the version data to the versions
   table and will update the `version_count` in services which we will be handy to get the versions count for a service instead of computing it on the way
2. DB: as mentioned above  `Documents and ServiceData can be added for extending
   the functionality to provided more details for the api services and versions.`
3. **Authentication and authorising user** Every user who will be using this api should be authenticated by the login api's 
   and they should have `Authorization Header with Bearer token` which should be validate be validated at api gateway level for eg at kong proxy layer.
   After this call should come to upstream service where scopes present in this token should be checked for the request api 
   for the user which we can do as part of this func `isUserAuthorised` at middleware level. If authorised serve the request
   else return 401 error stating unauthorised call
4. **Caching** as we will be serving the home page for the api services we can cache the first page to increase our performance.
