# csesoc.unsw.edu.au

This repo houses the CSESoc website.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Docker needs to be installed on your system because local deployment uses containerisation to standardise deployments across machines. For further information please read Docker docs. https://docs.docker.com/.

For Mac users, proceed to https://docs.docker.com/docker-for-mac/install/ and follow the instructions there. From here on, all terminal commands are written for MacOS and a zsh terminal. Make changes as necessary.

For Windows users, proceed to https://docs.docker.com/docker-for-windows/install/ and follow the instructions there. Be aware that if you have a Windows Home account with a Windows Subsystem for Linux, then additional steps need to be taken before downloading Docker, of which you can find here https://medium.com/@sebagomez/installing-the-docker-client-on-ubuntus-windows-subsystem-for-linux-612b392a44c4.

Once Docker has been downloaded, check by running the command in your terminal:
``` script
docker --version
```

Once that's done, clone the repo. `git clone https://github.com/csesoc/csesoc.unsw.edu.au`

### Setting up environment variables

All tokens are being treated as environment variables and are purposely left out of the GitHub repository. The easiest way is to inject these secrets into the container as environment variables.

Contact the website project lead to get the tokens sent directly to you via secure mediums.

Docker expects a file named `.env` in the root directory of the project containing environment variables in the following format:
```
{variable_name}={variable_value}
```
Here is a valid `.env` file:
```
ENV_VAR1=0123456789
ENV_VAR2=super_secure_secret
```

These environment variables are only reacheable during the building process of the Docker containers, not in the container themselves (by default). To pass them into their relevant container the `docker-compose.yml` file specifies the required env variables under the *environment* field for each service.

### Installing

*Please check back regularly as deployment steps may change at any point within these early stages of development. Thank you for your patience.*

With the repo cloned, proceed to checkout to the dev branch, `git checkout dev`.

From the root folder of the dev branch, run the following command in your terminal.

``` script
docker-compose up -d --build
```

This will automatically build the images required for the containers, as well as the containers for the first time. After this images will not need to be built again until changes have been made to the repo files. The '-d' is to start the container in the background and leave them running. There will be three containers that start up `frontend`, `backend` and `mongo`. Once you are finished with local deployment, run:

``` script
docker-compose down
```

which kills your containers but keeps your iamges. Be sure to use `docker-compose --help` for any additional help or other options.

If the frontend code has been changed, rebuild using `docker-compose up -d --build`. For backend or mongo containers they are reloaded live and any changes made on your system will be automatically refelected in the container and recompiled accordingly.

To access the website, the static files will be served on `0.0.0.0:8080` (`[::]:8080`) while the backend APIs are served on `0.0.0.0:1323` (`[::]:1323`). Make sure when you are making calls from the frontend to the backend in development stage, you use the suffix of the api call and not call with the domain e.g

``` javascript
/api/v1/sponsors
```

as opposed to

``` javascript
https://localhost:1323/api/v1/sponsors
```

## API Documentation

The API documentation is handled by [Swagger](https://swagger.io/) and can be found by navigating to `0.0.0.0:1323/swagger/index.html` (`[::]:1323/swagger/index.html`). Notice that it's also in the port that serves the APIs themselves. Swagger was adopted as we looked at employing a 'docs-as-code' approach to allow developers to quickly and efficiently write documentation ad-hoc, as well as having a permanent space for future teams to read up on API while working with it. Lastly, Swagger is intuitive and provides an interactive way to contact the APIs.

## Running the tests

The project uses Github Actions for continuous integration and automated testing. Testing will always be written at the beginning of each sprint and run every time a push is detected on your feat/fix/hotfix branch or a merge/push to dev.

### Input Validation
To validate structs, we are utilising a feature of the echo web framework that allows us to couple a validator package to validate structs that contain user inputs from requests. The package is golang's [package validator](https://pkg.go.dev/gopkg.in/go-playground/validator.v9?tab=doc#pkg-index). Everytime validation needs to occur for inserting into a database please use `echo-context.Validator(&struct)` to validate and handle any errors accordingly.

### Unit and Integration Testing
These tests are written in golang's standard testing package and are written in the same package as the file that they are testing. The tests are named `*_test.go` and the testing package has to be imported. For more information please read the Golang documentation for the [package testing](https://golang.org/pkg/testing/). To have these test run, please ensure docker is running and your containers for development are running. Then to run the actual test, go to the directory that contains the go files and run `go test`. This will run all test in child directories. 

It is imperative that you test frequently to spot bugs and errors early on. Do not rely on Github Actions when you conduct a PR to check because it is used as an integration tool so that it is a final check before changes are merged onto dev.

### Github Actions
Github Actions is the CI tool that we are using because of the relative ease of use and the ability to make changes as a developer to the CI workflow as needed. The script to run Github Actions is in the `.github` directory and is named `ci.yml`.

At the current moment it builds the docker images, runs the containers and performs the `go test` directive on those containers. Github actions will have increased functionality as we move away from working with Go backend files and look at testing other aspects of our website.

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

## System Architecture

For more on the system's architecture, please head to the [confluence page](https://compclub.atlassian.net/wiki/spaces/Projects/pages/845414415/Architectural+Guide)

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Tommy Truong** - *Initial work for README.md file* - [glebme](https://github.com/glebme)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Thanks to [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2) for the README template.
