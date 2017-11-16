# Twitter Api Search

Using Golang with the go-twitter package this project interacts with Twitters Search Api to display tweets based on search terms.

[Live Demo](https://twitter-aoiivrdcxu.now.sh )
## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
You can download either. 

* [Docker](https://www.docker.com/community-edition#/download) - Download Docker CE for your computer
* [Golang](https://golang.org/doc/install) - Download the Go 1.9.2



## Setting Up Environment
Be sure to download your own twitter 
There two ways to run this project either locally or with dockers

###Locally

```
git clone https://github.com/ImperiousEnterprise/TwitterSearchApi.git
cd /{directory where you saved this project}
go run *.go
```
From here you can connect
http://localhost:8080

##Dockers

You will need to build your own image the directions below will explain how to:

```
git clone https://github.com/ImperiousEnterprise/TwitterSearchApi.git
cd /{directory where you saved this project}
docker build -t twitterapp:0.0.1 .
docker run --rm -p8080:8080 --name twitterapp twitterapp:0.0.1
```
From here you can connect to your through your docker ip. 
http://{docker-ip}:8080


## Running the tests

This section explains how to run tests.
All tests are written in main_test.go

```
go test
```

## Deployment

As mentioned earlier since dockers is used.
This project is deployed to a docker container.

## Things Left To Do

* Fix Media Not Showing in tweets
* Add Advanced Options in search (Currently only ten Tweets are returned at a time).
  Possible options include setting the amount returned and specifying date range.
* Add few more tests for API testing in main.go


## Author

* **Adefemi Adeyemi**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


