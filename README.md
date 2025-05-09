# Nerding Out Backend API

==========================

Welcome to the backend API of Nerding Out, a blogging app built with Go and Fiber!

## Overview
----------

This repository contains the source code for the Nerding Out backend API, which is powered by [Fiber](https://github.com/gofiber/fiber) and deployed using [Fly.io](https://fly.io). The API provides an interface for users to create, read, update, and delete (CRUD) blog posts.

## Getting Started
-------------------

### Prerequisites

* Go installed on your machine
* The frontend installed (more instructions [here](https://github.com/TS22082/nerdingout_fe))
* Mongodb installed on your computer (min version 7)

### Local Development
1. Clone this repository: `git clone https://github.com/TS22082/nerdingout_be`
2. Install dependencies: `go get ./...` or use a tool like Dep to manage dependencies.
3. Run the API locally using Docker: `docker run -p 8080:80 nerdingout-backend`

### Deployment
The Nerding Out backend API is deployed on Fly.io, which provides automatic scaling and load balancing.

Running Godocs:
1. navigate to root of project
2. run `godoc -http=:6060`
3. navigate to `http://localhost:6060
4. Under "Packages" click "Third Party" to be scrolled to the repos docs.


## License
---------

This project is licensed under the MIT License. See `LICENSE` for details.