# Simple CRUD Service Using GOKIT

Clone this project using below command:
```shell
git clone https://github.com/foadHoseyni/test-gokit.git
```
Run the project using docker-compose:
```shell
docker-compose up
```
It runs on port 8000:
```shell
http://loaclhost:8000
```

To create an account 
POST:  /account

	{
		"email":"some@email.com",
		"phone":"xxxxxxxxxx"
	}

To get an account by id

GET:	/account/{customerid}

To get all the customers details

GET: /account/getAll

To update an account

PATCH:	/account/update

	{
		"customerid":1,
		"email":"some@email.com",
		"phone":"xxxxxxxxxx"
	}

To delete an account

DELETE: /account/{customerid}

**test-gokit** project also contains **prometheus** and **grafana** for monitoring purposes which are running on ports **:9090** and **:3000** respectivley.

The default username and password for grafana is **admin** and **admin**. 