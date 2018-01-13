# Dockerfile will define what goes on in the environment inside your container.
# Access to resources like networking interfaces and disk drives is virtualized inside this environment,
# which is isolated from the rest of your system, so you have to map ports to the outside world, and be specific
# about what files you want to “copy in” to that environment. However, after doing that, you can expect that
# the build of your app defined in this Dockerfile will behave exactly the same wherever it runs.


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
RUN go test -v ./...
CMD ["/app/main"]