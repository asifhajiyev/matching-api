Matching API\
It is used to match rider with the driver. There two endpoints
* GET: http://127.0.0.1:8090/api/auth/get-token
This endpoint is for getting token. Response example is given below. After getting token other endpoints could be used. To use them put the token in Authorization header with Bearer key.
```json
{
    "code": 200,
    "message": "OK",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlLCJpc3MiOiJtYXRjaGluZy1hcGkiLCJleHAiOjE2NDkwMjc0MDF9.VZH3P9akipu1LRFXvRRtkvKzfmCt881ulw_HHoBHZu4"
    }
}
```

* GET: http://127.0.0.1:8090/api/match \
It has two query parameters: longitude and latitude. Radius is constant number (9750000) in the code.
  * Request example: \
    http://127.0.0.1:8090/api/match?longitude=-73.9667&latitude=40.78
  * Response example: \
  ```json
  {
    "code": 200,
    "message": "OK",
    "data": {
        "rideInfo": {
            "driverInfo": {
                "location": {
                    "coordinates": [
                        40.62189228,
                        30.04352028
                    ]
                }
            },
            "distance": 9661.68
        }
    }
  }
  ```


