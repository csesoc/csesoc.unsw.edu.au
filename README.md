# csesoc.unsw.edu.au

This repo houses the CSESoc website.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Docker needs to be installed on your system because local deployment uses containerisation to standardise deployments across machines. For further information please read Docker docs. https://docs.docker.com/.

For Mac users, proceed to https://docs.docker.com/docker-for-mac/install/ and follow the instructions there. From here on, all terminal commands are written for MacOS and a zsh terminal. Make changes as necessary.

For Windows users, proceed to https://docs.docker.com/docker-for-windows/install/ and follow the instructions there. Be aware that if you have a Windows Home account with a Windows Subsystem for Linux, then additional steps need to be taken before downloading Docker, of which you can find here https://medium.com/@sebagomez/installing-the-docker-client-on-ubuntus-windows-subsystem-for-linux-612b392a44c4.

Once Docker has been downloaded, check by running the command in your terminal:
```
docker --version
```

Once that's done, clone the repo. `git clone https://github.com/csesoc/csesoc.unsw.edu.au`

### Installing

*Please check back regularly as deployment steps may change at any point within these early stages of development. Thank you for your patience.*

With the repo cloned, proceed to checkout to the dev branch, `git checkout dev`.

From the root folder of the dev branch, run the following command in your terminal.

```
docker-compose up -d  
```

This will automatically build the images required for the containers, as well as the containers for the first time. After this images will not need to be built again until changes have been made to the repo files. The '-d' is to start the container in the background and leave them running. Once you are finished with local deployment, run:

```
docker-compose down
```

which kills your containers but keeps your iamges. Be sure to use `docker-compose --help` for any additional help or other options.

If the code has been changed, you need to rebuild your image. Due to the fact that we use a docker-compose.yml, please look at the service name you are looking to rebuild and run:

```
docker-compose build your-services-name
```

And then proceed with previous steps to get start local deployment.

The server will start on `0.0.0.0:1323` (`[::]:1323`) which serves both the frontend and API simultaneously.

## Running the tests

The project uses Github Actions for continuous integration and automated testing. Testing will always be written at the beginning of each sprint and run every time a push is detected on your feat/fix/hotfix branch or a merge/push to dev.

### Input Validation
To validate structs, we are utilising a feature of the echo web framework that allows us to couple a validator package to validate structs that contain user inputs from requests. The package is golang's [package validator](https://pkg.go.dev/gopkg.in/go-playground/validator.v9?tab=doc#pkg-index). Everytime validation needs to occur for inserting into a database please use `echo-context.Validator(&struct)` to validate and handle any errors accordingly.

### Unit and Integration Testing
These tests are written in golang's standard testing package and are written in the same package as the file that they are testing. To have these test run, please ensure docker is running and your containers for development are running. Then to run the actual test, go to the directory that contains the go files and run `go test`. This will run all test in child directories. 

It is imperative that you test frequently to spot bugs and errors early on. Do not rely on Github Actions when you conduct a PR to check because it is used as an integration tool so that it is a final check before changes are merged onto dev.

### Github Actions
Github Actions is the CI tool that we are using because of the relative ease of use and the ability to make changes as a developer to the CI workflow as needed. The script to run Github Actions is in the `.github` directory and is named `ci.yml`.

At the current moment it builds the docker images, runs the containers and performs the `go test` directive on those containers. Github actions will have increased functionality as we move away from working with Go backend files and look at testing other aspects of our website.

## Deployment

[Gordon Zhong](https://github.com/gawdn) has written up deployment steps on CSESOC servers which can be viewed here: [How to deploy a project on Wheatley](https://compclub.atlassian.net/wiki/spaces/Projects/pages/733118519/How+to+deploy+a+project+on+Wheatley)

## Built With

* [Vue + Vuetify](https://vuejs.org/) - The web framework used
* [MongoDB](https://www.mongodb.com/) - Database
* [Golang](https://golang.org/) - Used to write the API and Server backend. To view "[api docs](https://gawdn.com/api-docs)" (link may eventually break in the future).

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
