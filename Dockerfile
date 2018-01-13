# Use an official Python runtime as a parent image
FROM golang:latest
RUN mkdir /app

#Install any needed packages specified in requirements.txt
#RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Copy the current directory contents into the container at /app
ADD . /app/
# Set the working directory to /app
WORKDIR /app

## Make port 80 available to the world outside this container
#EXPOSE 80
#
## Define environment variable
#ENV NAME World

RUN go get -u -t github.com/gorilla/mux
RUN go get -u -t github.com/gorilla/context
RUN go build -o main .
CMD ["/app/main"]