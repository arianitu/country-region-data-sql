# country-region-data-sql
This repo contains a .sql file that can be used on MySQL and Postgres for country and region data used in forms. It also contains a go script to generate the data from a maintained region and country repo https://github.com/country-regions/country-region-data


### Using

1. Insert the schema for your database:

`schema_mysql.sql`

2. Insert the data (it should work for both Postgres and MySQL schema)

`data.sql`


### Re-generating the data

The data is generated from the JSON at https://github.com/country-regions/country-region-data , you can download the JSON there and then run `go run generate_data.sql > data.sql`

