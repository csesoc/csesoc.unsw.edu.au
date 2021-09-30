# csesoc.unsw.edu.au

This repo houses the static CSESoc website which will stand in for the new CMS-based website while it is in development.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Docker needs to be installed on your system because local deployment uses containerisation to standardise deployments across machines. For further information please read [Docker docs](https://docs.docker.com/).

For Mac users, proceed to [Docker for Mac](https://docs.docker.com/docker-for-mac/install/) and follow the instructions there. From here on, all terminal commands are written for MacOS and a zsh terminal. Make changes as necessary. For Windows users, proceed to [Docker for Windows](https://docs.docker.com/docker-for-windows/install/) and follow the instructions there.

Once Docker has been downloaded, check by running the command in your terminal:

``` script
docker --version
```

Once that's done, clone the repo.

``` script
git clone https://github.com/csesoc/csesoc.unsw.edu.au
```

For the sake of development, please also have Golang installed on your local machine and yarn as a your JS package manager.

### Installing

With the repo cloned, proceed to checkout to the dev branch.

From the root folder of the dev branch, run the following command in your terminal.

``` script
docker-compose -f docker-compose-dev.yml up -d --build
```

This will automatically build the images required for the containers, as well as the containers for the first time. After this images will not need to be built again until changes have been made to the dependencies. For subsequent runs, remove the `--build`.

The '-d' is to start the container in the background and leave them running. There will be three containers that start up `frontend`, `backend` and `mongo`.
Note: once you have built the containers from here, you can do most docker operations from the UI such as starting/stopping the containers. You will only *need* to do this for your first setup or if you delete the containers / purge your docker data.

If you would rather use the commandline, once you are finished developing, run:

``` script
docker-compose down -v
```

which kills your containers and removes any bind mounts and named volumes but keeps your images. Be sure to use `docker-compose --help` for any additional help or other options.

You can also run just
``` script
docker-compose down
```

if you want to erase your images from docker too (which can save some disk space).

You may want to be wary of docker as it is *extremely resource hungry*. It can occupy several gigs of RAM and will noticably cause increased battery drainage because it is constantly making use of the CPU. For your sanity you may want to make sure you *stop all containers and shut docker down* if you aren't working on the repo for your own sanity - if you have 32gb of RAM or greater <and aren't using battery power> this probably won't be an issue. 

To access the website, the static files will be served on `0.0.0.0:8080` (`[::]:8080`) while the backend APIs are served on `0.0.0.0:1323` (`[::]:1323`). Make sure when you are making calls from the frontend to the backend in development stage, you use the suffix of the api call and not call with the domain e.g

``` javascript
/api/v1/sponsors
```

as opposed to

``` http
https://localhost:1323/api/v1/sponsors
```

## API Documentation

The API documentation is handled by [Swagger](https://swagger.io/) and can be found by navigating to `0.0.0.0:1323/swagger/index.html` (`[::]:1323/swagger/index.html`). Notice that it's also in the port that serves the APIs themselves. Swagger was adopted to employ a 'docs-as-code' approach to allow developers to quickly and efficiently write documentation ad-hoc, as well as having a permanent space for future teams to read up on API while working with it. Lastly, Swagger is intuitive and provides an interactive way to contact the APIs.

## Living Style Guide

To help create a unique and consistent brand identity, we have looked at creating a living style guide for developers to utilise to be able to build components on our website more easily while still fitting in with the overall vision of wireframes, stylesheets and design guidelines. To run the living style guide on your local machine, run `yarn run kss` from the frontend folder. This will create css files for a local server you use to serve the files necessary to visualise the guide. Then once a local server is running proceed to the `frontend/src/styleguide/` and open the html file to begin browsing.

_There is currently no way to visualise the fonts used. Fix in progress._

## Running tests

The project uses Github Actions for continuous integration and automated testing. Testing will always be written at the beginning of each sprint and run every time a push is detected on your feat/fix/hotfix branch or a merge/push to dev.

### Input Validation

To validate structs, we are utilising a feature of the echo web framework that allows us to couple a validator package to validate structs that contain user inputs from requests. The package is golang's [package validator](https://pkg.go.dev/gopkg.in/go-playground/validator.v9?tab=doc#pkg-index). Everytime validation needs to occur for inserting into a database please use `echo-context.Validator(&struct)` to validate and handle any errors accordingly.

### Backend Unit Testing

These tests are written in Golang's standard testing package and are written in the same package as the file that they are testing. The tests are named `*_test.go` and the testing package has to be imported. For more information please read the Golang documentation for the [package testing](https://golang.org/pkg/testing/). To have these test run, please ensure docker is running and your containers for development are running. Then to run the actual test, go to the backend directory and run `go test ./...`. This will run all test in child directories.

### Frontend Unit Testing

While frontend unit testing of components is usually never thought of by developers still at university, it is something our teams (past and present) have given thought to. We have settled on using the Cypress.io framework to run E2E tests on our system. The framework was choosen because of its ease of use and low learning curve. It fits well with our current needs to provide quality and visual assurance that our frontend code appears and is used as how an end user would interact with the website. 

To get it running, go to the `frontend` folder. Make sure that you have `yarn install`ed previously when setting up to make sure that you have cypress ready to run. To launch the console please use `yarn run cypress open` which will load an application window to view tests that cypress has detected. For further information please go to the [Cypress installation documentation](https://docs.cypress.io/guides/getting-started/installing-cypress.html#Opening-Cypress).

Cypress is quite extensively documented but if there are any difficulties a good starting point is to learn Chai.js and look at the tutorials on the website.

### Github Actions

Github Actions is the CI tool that we are using because of the relative ease of use and the ability to make changes as a developer to the CI workflow as needed. The script to run Github Actions is in the `.github` directory and is named `ci.yml`.

At the current moment it builds the docker images, runs the containers and performs the `go test` directive on those containers. Github actions will have increased functionality as we move away from working with Go backend files and look at testing other aspects of our website.

It is imperative that you test frequently to spot bugs and errors early on. Do not rely on Github Actions when you conduct a PR to check because it is used as an integration tool so that it is a final check before changes are merged onto dev.

## Deployment

[Gordon Zhong](https://github.com/gawdn) has written up deployment steps on CSESOC servers which can be viewed here: [How to deploy a project on Wheatley](https://compclub.atlassian.net/wiki/spaces/Projects/pages/733118519/How+to+deploy+a+project+on+Wheatley)

## Security

Tokens and API keys should never be commited to a git repository. To overcome this we are using environment variables that are shared among developers through secure mediums. These environment variables are injected into the relevant containers when they are initialized.

To simplify things we are using environment variables instead of Docker Secrets. The reasoning behind this is because Rancher stores our production secrets internally as Kubernetes Secrets (and just like Docker Secrets are also encrypted at rest). The easiest way is to inject these secrets into the container is as environment variables which also simplifies (and by extension improves the security of) the CI/CD process. This is just as secure as using Docker Secrets since in both cases the secrets are plain-text in the container.

## Built With

* [Vue + Vuetify](https://vuejs.org/) - The web framework used
* [MongoDB](https://www.mongodb.com/) - Database
* [Golang](https://golang.org/) - Used to write the API and server backend.
* [Swagger](https://swagger.io/) - API documentation framework


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Thanks to [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2) for the README template.
