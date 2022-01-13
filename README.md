# sample-project



## How to use


Use command `make run` than server running on PORT 8001 or `make docker` to run the app inside a container.

Following bellow request example

```curl
curl --request POST \
  --url http://localhost:8001/orders \
  --header 'Content-Type: application/json' \
  --data '
{
	"email": "gehrmann.j@nekom.com"
}'

```

Following bellow a response example

```json
[
	{
		"uuid": "fdfd497828effdfda2b06c8b78efa26c",
		"aktiv": 1,
		"orderType": "order",
		"customer": {
			"uuid": "fdfd497828effdfda2b04949b015fd5a",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"email": "gehrmann.j@nekom.com",
			"customerid": 4711
		},
		"orderNr": 1048363,
		"billingAddress": {
			"uuid": "fdfd497828effdfda2b0495a93c16c93",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"street": "Bründlstr.",
			"houseNr": "16",
			"city": "Geisenfeld",
			"zip": "85290",
			"country": "DEU"
		},
		"shippingAddress": {
			"uuid": "fdfd497828effdfda2b0495a93c16c93",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"street": "Bründlstr.",
			"houseNr": "16",
			"city": "Geisenfeld",
			"zip": "85290",
			"country": "DEU"
		},
		"shippingType": "DPI",
		"paymentType": "Rechnung (manuell)",
		"orderStatus": 1,
		"orderValueGross": 5438500,
		"shippingCosts": 49500,
		"orderDate": 20220111,
		"orderTime": 104337,
		"currency": "EUR",
		"orderposition": [
			{
				"uuid": "fdfd497828effdfda2b0d4ef3e938b15",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b0749",
					"itemnumber": "6600000_01_50",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1790000,
				"positionvalue": 1790000,
				"taxrate": 19
			},
			{
				"uuid": "fdfd497828effdfda2b0d4ef3e938b28",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b0715",
					"itemnumber": "6600000_01_52",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1799500,
				"positionvalue": 1799500,
				"taxrate": 19
			},
			{
				"uuid": "fdfd497828effdfda2b0d4ef3e938b3e",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b3eef",
					"itemnumber": "6600000_01_54",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1799500,
				"positionvalue": 1799500,
				"taxrate": 19
			}
		],
		"channel": "Manuell"
	},
	{
		"uuid": "fdfd497828effdfda2b06c8b78c18b78",
		"aktiv": 1,
		"orderType": "order",
		"customer": {
			"uuid": "fdfd497828effdfda2b04949b015fd5a",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"email": "gehrmann.j@nekom.com",
			"customerid": 4711
		},
		"orderNr": 1048027,
		"billingAddress": {
			"uuid": "fdfd497828effdfda2b0495a93c16c93",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"street": "Bründlstr.",
			"houseNr": "16",
			"city": "Geisenfeld",
			"zip": "85290",
			"country": "DEU"
		},
		"shippingAddress": {
			"uuid": "fdfd497828effdfda2b0495a93c16c93",
			"firstName": "jürgen",
			"lastName": "Gährmönn",
			"street": "Bründlstr.",
			"houseNr": "16",
			"city": "Geisenfeld",
			"zip": "85290",
			"country": "DEU"
		},
		"shippingType": "DPI",
		"paymentType": "Rechnung (manuell)",
		"orderStatus": 1,
		"orderValueGross": 5438500,
		"shippingCosts": 49500,
		"orderDate": 20211208,
		"orderTime": 104653,
		"currency": "EUR",
		"orderposition": [
			{
				"uuid": "fdfd497828effdfda2b0d4ef3ea28b28",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b0749",
					"itemnumber": "6600000_01_50",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1790000,
				"positionvalue": 1790000,
				"taxrate": 19
			},
			{
				"uuid": "fdfd497828effdfda2b0d4ef3ea28b3e",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b0715",
					"itemnumber": "6600000_01_52",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1799500,
				"positionvalue": 1799500,
				"taxrate": 19
			},
			{
				"uuid": "fdfd497828effdfda2b0d4ef3ea28b07",
				"aktiv": 1,
				"item": {
					"uuid": "fdfd497828effdfda2b0d4496c8b3eef",
					"itemnumber": "6600000_01_54",
					"designation": "Herren Skihose Anton slim",
					"imagepath": "https://res.cloudinary.com/storer-handels-gmbh/image/upload/v1599026123/products-stage/6600000_01_a.jpg"
				},
				"itemdescription": "Herren Skihose Anton slim",
				"orderquantity": 1,
				"purchaseprice": 644300,
				"sellingprice": 1799500,
				"positionvalue": 1799500,
				"taxrate": 19
			}
		],
		"channel": "Manuell"
	}
] 
```