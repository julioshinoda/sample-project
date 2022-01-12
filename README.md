# sample-project

## How to use

Use commnad `make run` than server running on PORT 8080

Following bellow request example

```curl
curl --request POST \
  --url http://localhost:8080/orders \
  --header 'Content-Type: application/json' \
  --data '
{
	"email": "gehrmann.j@nekom.com"
}'

```