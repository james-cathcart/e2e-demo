# Zed Application
This application provides the data layer for the distributed system

## Datastore Details
This application uses a `Sqlite3` database with the following schema and values
```
CREATE TABLE ads (
	customer TEXT NOT NULL,
	info TEXT NOT NULL
);

INSERT INTO ads (customer, info) VALUES
	('BigCo', 'Best cookware on earth!'),
	('SmallCo', 'We clean carpets cheap!'),
	('AlphaCo', 'Northern Colorado sports gear'),
	('BravoCo', 'Southern Colorado sports gear');
```
# E2E Testing
These e2e tests demonstrate how to use the Postman environment to pass information from one test to the next. This can be used to keep track of the ID for an automatically created record in a previous test. This can be used by subsequent requests to perform other tests and eventually the record can be deleted as a part of the test cleanup. See the `e2e/e2e-zed.postman_collection.json` collection for more details.

The tests in this example will:
- create a new record
- add the ID of the new record to the postman environment
- confirm that a record with the new ID exists
- update the values for said record
- confirm that the record has been updated
- delete the record using the stored ID
- confirm that no record with that ID exists

## Run E2E Tests
```
newman run e2e-zed.postman_collection.json -e zed-local.postman_environment.json
```
Expected Output
```
$ newman run e2e-zed.postman_collection.json -e zed-local.postman_environment.json 
newman

E2E Zed Tests

→ Get All Ads
  GET http://localhost:8082/ads [200 OK, 399B, 25ms]
  ✓  Status code is 200

→ Create Ad
  POST http://localhost:8082/ads [201 Created, 129B, 5ms]
  ✓  Status code is 201
  ✓  Your test name

→ Get Ad By ID
  GET http://localhost:8082/ads?filter=id&value=5 [200 OK, 197B, 3ms]
  ✓  Status code is 200
  ✓  JSON field values should be as expected

→ Get Ad By Customer
  GET http://localhost:8082/ads?filter=customer&value=e2eTestCustomer [200 OK, 188B, 3ms]
  ✓  Status code is 200
  ✓  Should have at least one result

→ Update Ad
  PUT http://localhost:8082/ads [200 OK, 75B, 3ms]
  ✓  Status code is 200

→ Confirm Update
  GET http://localhost:8082/ads?filter=id&value=5 [200 OK, 203B, 2ms]
  ✓  Status code is 200
  ✓  Your test name

→ Delete Ad (clean-up)
  DELETE http://localhost:8082/ads?id=5 [200 OK, 75B, 4ms]
  ✓  Status code is 200

→ Confirm Ad Deletion
  GET http://localhost:8082/ads?filter=id&value=5 [204 No Content, 64B, 2ms]
  ✓  Status code is 204

┌─────────────────────────┬─────────────────┬─────────────────┐
│                         │        executed │          failed │
├─────────────────────────┼─────────────────┼─────────────────┤
│              iterations │               1 │               0 │
├─────────────────────────┼─────────────────┼─────────────────┤
│                requests │               8 │               0 │
├─────────────────────────┼─────────────────┼─────────────────┤
│            test-scripts │               8 │               0 │
├─────────────────────────┼─────────────────┼─────────────────┤
│      prerequest-scripts │               0 │               0 │
├─────────────────────────┼─────────────────┼─────────────────┤
│              assertions │              12 │               0 │
├─────────────────────────┴─────────────────┴─────────────────┤
│ total run duration: 146ms                                   │
├─────────────────────────────────────────────────────────────┤
│ total data received: 535B (approx)                          │
├─────────────────────────────────────────────────────────────┤
│ average response time: 5ms [min: 2ms, max: 25ms, s.d.: 7ms] │
└─────────────────────────────────────────────────────────────┘
```