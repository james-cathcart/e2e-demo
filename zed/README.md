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