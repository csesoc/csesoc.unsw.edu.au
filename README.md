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

which kills your containers but keeps your images. Be sure to use `docker-compose --help` for any additional help or other options.

If the code has been changed, you need to rebuild your image. Due to the fact that we use a docker-compose.yml, please look at the service name you are looking to rebuild and run:

```
docker-compose build your-services-name
```

And then proceed with previous steps to get start local deployment.

The server will start on `0.0.0.0:1323` (`[::]:1323`) which serves both the frontend and API simultaneously.

## Running the tests

The project uses Github Actions for continuous integration and automated testing.

### Input Validation
To validate test, we are utilising a feature of the echo web framework that allows us to couple a validator package to validate structs that contain user inputs from requests. The package is golang's [package validator](https://pkg.go.dev/gopkg.in/go-playground/validator.v9?tab=doc#pkg-index). Everytime validation needs to occur for inserting into a database please use `echo-context.Validator(&struct)` to validate and handle any errors accordingly.

### Postman
API testing is performed using the platform Postman. To run standalone tests, please check the [Postman website](https://www.postman.com/). Please make sure you have newman downloaded, using `npm install -g newman`.

1. Before coding of the feature occurs, API testing must be done to ensure that integration is as smooth as possible.
2. After you have written the tests and applied it to your code, export the Postman collection (a collection of related test) and place it in tests/postman in the collective postman file. 
3. Open a pull request. This will automatically run previous test in the `dev` branch so that your code integrates and does not (hopefully) break existing code.
4. Fix any errors that do occur and once done then merge. Your new api tests should now be run every time there is a PR into dev.

As of this moment, a fix is taking place to get the github actions to run newman.

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
