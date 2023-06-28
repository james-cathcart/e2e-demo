# ETL Application
This application reads records from the `Foo` application and loads those records into the `Bar` application.

# Running E2E Tests
Navigate to the e2e directory and run the collection
```
cd <repo>/etl/e2e && newman run e2e-etl.postman_collection.json
```
Expected Output:
```
$ newman run e2e-etl.postman_collection.json 
newman

E2E ETL Tests

❏ ETL
↳ Migrate Records - Dry Run
  POST http://localhost:8084/etl [201 Created, 406B, 25ms]
  ✓  Status code is 201
  ✓  Ads array should not be empty
  ┌
  │ 'checking index: ', 0
  │ 'checking index: ', 1
  │ 'checking index: ', 2
  │ 'checking index: ', 3
  └
  ✓  All fields in all elements of Ads array should be populated
  ✓  Migrated value should be greater than 0

┌─────────────────────────┬──────────────────┬──────────────────┐
│                         │         executed │           failed │
├─────────────────────────┼──────────────────┼──────────────────┤
│              iterations │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│                requests │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│            test-scripts │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│      prerequest-scripts │                0 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│              assertions │                4 │                0 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 51ms                                      │
├───────────────────────────────────────────────────────────────┤
│ total data received: 292B (approx)                            │
├───────────────────────────────────────────────────────────────┤
│ average response time: 25ms [min: 25ms, max: 25ms, s.d.: 0µs] │
└───────────────────────────────────────────────────────────────┘
```