## Use an official Python runtime as a parent image
#FROM golang
#
## Set the working directory to /app
##WORKDIR /app
#
## Copy the current directory contents into the container at /app
#ADD . /go/src/github.com/AndreyShamis/golang
#
## Install any needed packages specified in requirements.txt
##RUN pip install --trusted-host pypi.python.org -r requirements.txt
#
#RUN go install github.com/AndreyShamis/golang
#
## Make port 80 available to the world outside this container
#EXPOSE 80
#
## Define environment variable
#ENV NAME World
#
## Run app.py when the container launches
##CMD ["python", "app.py"]
#CMD ["go", "build"]
#CMD ["pwd", ""]
#CMD ["./golang", ""]


FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u -t github.com/gorilla/mux
RUN go get -u -t github.com/gorilla/context
RUN go build -o main .
CMD ["/app/main"]