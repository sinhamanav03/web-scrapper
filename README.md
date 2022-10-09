# WEB Scrapper Application

Applicaion scrapes an Amazon web page given its URL


## To run Scrapper Service
run
```
docker compose up -d
go run cmd/scrapper/main.go

```
you can hit ther service at postman at localhost:8081/scrape with url in body 
```
{
    "url":"https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/"
}
```

## To run scrapper-store service
run
```
docker compose up -d
go run cmd/scrapper-store/main.go

```
you can hit ther service at postman at localhost:8088/add with url in body 

```
{
    "url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/",
    "product": {
        "name": "PlayStation 4 Pro 1TB Console",
        "imageURL": "https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg",
        "description": "Heighten our experiences.\n Enrich your adventures.",
        "price": "$348.00",
        "totalReviews": 4136
    }
}

```
other endpoints available are
```
\get{id}
```
```
\update{id} expects in body
     {
        "name": "PlayStation 4 Pro 1TB Console",
        "imageURL": "https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg",
        "description": "Heighten our experiences.\n Enrich your adventures.",
        "price": "$348.00",
        "totalReviews": 4136
    }

```
```
\delete{id}
```


## Note
all env variables can be set in config.json