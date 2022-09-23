# Photon

This is a dispatcher component that is responsible for receiving webhooks from git providers and queueing actions to be taken as a response to events received. This component needs to be highly available to receive webhooks and other requests in large volumes.

## Building

- Makefile

 ```bash
 # install dependencies
 make deps
 # generate wire injectors  
 make wire
 # check linting
 make lint
 # build binary
 make build

 ```

- Docker usage

  ```bash
  docker build --rm --tag=photon .
  ```

- A sample shell file is also included to automate the build steps. `build.sh` in the root directory can be used for special adding build steps.

## Hot reloading

This project is configured to support [fresh](https://github.com/gravityblast/fresh) runner which reloads the application actively whenever any golang file (or any other file configured for hot reloading) changes. This is very useful while actively developing as it removes the need to recompile and run the application again and again. `runner.conf` in the root directory is used to configure the fresh runner. More information can be viewed on [their github project](https://github.com/gravityblast/fresh)

## Libraries used

- [sqlx](https://github.com/jmoiron/sqlx/)
- [wire](https://github.com/google/wire)
- [kafka](github.com/segmentio/kafka-go)
- [Viper](github.com/spf13/viper)
- [Cobra](github.com/spf13/cobra)

## TODO

- [ ] Add unit test cases
- [ ] Add swagger docs
- [ ] Add benchmarking test cases
