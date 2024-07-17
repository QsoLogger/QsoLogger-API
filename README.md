# QsoLogger-API

Here is the backend of QsoLogger, which will work with QsoLogger-UI

# Dependency
* [Golang](https://go.dev/dl/), v1.16+ is required, and the dev team will always use a head stable version
* [Gnu make](https://www.gnu.org/software/make/), is need for non-expert user
* Static libc is prefer if it is build for Linux

# Build
* For normal build, we prefer `make`
* For dynamic build that you will run this API with a same Env of your build box only, you can `make C_STATIC_LINK_FLAGS=`


# Run
You need the config file QsoLogger.ini, which will be search in the certain PATHs in order, that may copy from the template file `QsoLogger.ini.template`
* `./`
* `./etc`
* `/etc`

# For Developers and Experts
Go ahead, we need you


