## Foodji Tinder


Main packages of application:

The application has been structured keeping in view the prcinples of Ben johnson.

* **foodji** contain all the domain model

* **database** deals with databse layer

* **http** package is for all the rest endpoints and server information


The configuration required to connect to the postgreSQL database is present in the `internal/config/config.yaml`

### A sample interaction with appliction using rest api's


#### 1. Create a product:

 > POST Request on `http://localhost:8080/products/create`

body:
```

{
    "ID": "1",
    "NAME": "first product"
}
```

Output:
{"id":"1","name":"first product"}

with Status '200 OK'

#### 2. Create a new session 

> POST request on `http://localhost:8080/sessions/create`

Output:

{"id":"57bf2a7b-272e-4a3b-bcd3-446830e3ffe9","createdAt":"2023-04-09T16:33:36.575133+02:00"}

#### 3. Get Sessions

Output:
> GET request on `http://localhost:8080/sessions/all`

Output:

[{"id":"57bf2a7b-272e-4a3b-bcd3-446830e3ffe9","createdAt":"2023-04-09T16:33:36.575133Z"}]


#### 4. Store your vote for the product in a session

> POST request on `http://localhost:8080/votes/store`

body:
```
{
    "id": 1,
    "sessionId": "57bf2a7b-272e-4a3b-bcd3-446830e3ffe9",
    "productId": "1",
    "liked": true
}
```

Output:
 {"id":1,"sessionId":"57bf2a7b-272e-4a3b-bcd3-446830e3ffe9","productId":"1","liked":true}

#### 5. Get aggregated avergae scores

> GET request on `http://localhost:8080/votes/aggscores`

Output:
{"1":1}


### Possible Improvements

* Generating ids for Product and Vote instead of sending it
* Using the `crypto/rand package`package instead of `uuid`, this technique is not guaranteed to produce unique IDs, but the probability of a collision is very low.
* I tried using GORM initially but was having an issue, which I figured out now. I can maybe create another version of application with that
