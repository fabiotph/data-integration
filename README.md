# Data-integration
- - - -


## :arrow_forward: How to Run 
* [Docker](https://www.docker.com/) must be installed;
* If you are using Windows, you will need to download the GNU Make, you can do this with the package manager for Windows [Chocolatey](https://chocolatey.org/packages/make);
* Execute this command in the shell inside the project folder:  
`make start`


## :computer: Technologies used
* Go
* Postgres
* Gorm
* Docker


## :link: Endpoints
Route                                | Method | Request Structured | Description
------------------------------------ | ------ | -----------------    | -----
/company/all                         |  GET   | -----------------    | Returns all companies registered in the database.
/company?name={value1}&zip={value2}  |  GET   | -----------------    | Returns a company based on the "name" and "zip" query parameters sent. The "name" may be incomplete and both query parameters are mandatory.
/company/update                      | PATCH  |  multipart/form-data | Submit a CSV file to update the websites of companies already registered. The file must contain the name, zip code and website of each company to be updated.


## :book: Examples
### GET /company/all
#### Response success:
```json
[{
    "id": "9974e21e-a7e9-480e-bcd5-ba5e6ca7255f",
    "name": "PIZZA HUT",
    "zip": "44667",
    "website": "https://www.pizzahut.com"
  }, 
  ...
]  
```

### GET /company
#### Request example:  
`/company?name=saint&zip=55109`
#### Response success:
```json
{
  "id": "3a3c0ba3-09d3-42bc-bafd-7adceda4a914",
  "name": "SAINT PAUL RADIOLOGY",
  "zip": "55109",
  "website": "http://stpaulradiology.com"
} 
```

### PATCH /company/update
#### Resquest structure: 
`multipart/form-data`
#### File structure:
File with .csv format and must have the structure: 
 Name | ZipCode | WebSite
---|---|---
#### Response success:
```json
{
  "success": true,
  "message": "Updated companies."
}
```


## :white_check_mark: Tests
Execute this command in the shell inside the project folder:  
`make test`
