# Simple Microservice using GKIT

This project is a customer service with **CRUD** operation. **postgresql** used for data consistancy. for cloning this repository, use the following command
```shell
	git clone 
```
It runs on port 8000 

to create an account
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
		"customerid":"",
		"email":"some@email.com",
		"phone":xxxxxxxxxx
	}

To delete an account

DELETE: /account/{customerid}
