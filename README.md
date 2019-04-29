[Currently under development] This is my first project in Go. 

![Alt text](https://travis-ci.org/danfujita/solarium-data-collector.svg?branch=master)

Solarium Data Collector
======
![Alt text](https://github.com/danfujita/solarium-golang/blob/master/web/public/logo.png)


Solarium Data Collector is a Golang version of data collection backend microservice, which is currently under development for Apsis Science & Technology LLC. 

The goal of this application is to collect data from Solarium devices, and to upload the telemetry data to InfluxDB and the payload data to S3. 
The application works with a separate customer facing web portal service, which is currently under development. 

The Python hardware code for Solarium can be found at https://github.com/apsistech/Solarium

### Layout
This project follows Standard Go Project Layout from https://github.com/golang-standards/project-layout

### Building from source

You can also clone the repository and build using make:

    $ git clone https://github.com/danfujita/solarium-data-collector.git
    $ make
    $ ./main

### Docker images

Docker images are available on Docker Hub.

    $ docker pull danfujita55/solarium-data-collector
    $ docker run -d -p 8080:8080 danfujita55/solarium-data-collector
