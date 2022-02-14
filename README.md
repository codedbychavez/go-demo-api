# go-demo-api
A simple and effective api made with go

Prerequisites
Heroku CLI - This can be downloaded from https://devcenter.heroku.com/articles/heroku-cli

Go app - If you don’t have one, go ahead and clone this one to start with.

Heroku account - If you don’t have a Heroku account, head over to https://signup.heroku.com/ to create one.

# Steps
1. The first thing we need to do is create a folder called `compose` in the foot of our project.

We may need a different docker files for production and development, so create two sub folders called `prod` and `dev` in the `compose` folder.

Lets just focus on a single Dockerfile for dev:

Then create a `Dockerfile` in the dev folder:

    FROM golang:1.17.6

    # Add a work directory
    WORKDIR /app

    # Cache and install dependencies
    COPY go.mod go.sum ./
    RUN go mod download

    # Copy app files
    COPY . .

    # Expose port
    EXPOSE 4000

    # Start app
    CMD go run main.go

The above docker file downloads the go image, sets up a directory, installs our go app dependencies, copies the files, exposes a port (4000) for our go app to use then executes the go run command to start the app.


2. Secondly, create a `docker-compose.dev.yml` file in the dev folder as well

        version: "3.8"

        services:
        demo-go-api:
            container_name: demo-go-api-dev
            image: demo-go-api-dev
            build:
            context: .
            dockerfile: compose/dev/Dockerfile
            volumes:
            - .:/app
            environment:
            - PORT
            ports:
            - ${PORT}:${PORT}

In our docker-compose file, we define a single service called app, define our build, volumes, ports, and env variable.

It is important to know when it comes to Heroku, Heroku assigns a random dynamic port to our go app. 
To ensure that go uses that port when running our application, we’ll need to reference the default Heroku port environment variable.

A simple way to do this is to create a .env file in the root of our go project with the following content:

PORT=$PORT
ENVIRONMENT="development"

Notice `$PORT`, the value will be pulled from Heroku automatically.

Now it’s time to build and run the docker container locally.

### WAIT !
if you made any references to $POST as we talked about in the last step, you’ll need to change this temporarily to a real value to run the go app locally, I’ll be using 4000

Run the following command to start up your go app in docker:

`docker-compose -f docker-compose.dev.yml up`

After the container builds try to access your go app to ensure that the app runs locally via docker.
The next step we’ll move onto would be deploying the docker container that is running our app to Heroku. Note that we are now deploying the go app to Heroku in the standard way but instead we’re deploying the docker container.

## Deploying to Heroku

1. Before you can deploy your container to heroku, you need to push your container to heroku’s registry

2. Login to heroku from the command line:
heroku login

3. Login to the container registry
heroku container:login

4. Optional: If you ran into any authentication error while trying to login, try login into heroku’s container registry manually by running the following command:

`docker login - -username=YOUR-HEROKU-USERNAME - -password=$(heroku auth:token) registry.heroku.com`

5. Next, Change heroku’s default stack to container
    Note: you can see a list of heroku’s stacks by running the following command:

    `heroku stack -a NAME-OF-APP`

    Set the stack to container by running:

    `heroku stack:set container`

6. Then, push the container by running:

`heroku container:push web -a NAME-OF-APP-HEROKU`

Note: The word web in the above command. This is telling heroku that this container is to be handled by heroku’s `web` process type. 

7. Release the container so that your app on heroku can use it by running:

`heroku container:release web -a NAME-OF-APP-HEROKU`

8. Optional, if this does not happen automatically after releasing container, lets add a web dyno (Heroku’s name for virtual container) that we need our container to be served from.
heroku ps:scale web=1
