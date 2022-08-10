# moral-ambiguity

Moral Ambiguity is a SaaS solution for developers hoping to provide a means of protecting their digital content.

- [moral-ambiguity](#moral-ambiguity)
  - [Dependencies](#dependencies)
  - [Installation](#installation)

## Dependencies
  - [pre-commit](https://pre-commit.com/)
    - check `.pre-commit-config.yaml` for full list of hooks
  - Backend
    - [docker](https://www.docker.com/)
    - [go v1.18](https://go.dev/)
      - [go fiber](https://github.com/gofiber/fiber)
    - [PostgreSQL](https://www.postgresql.org/)
  - Frontend
    - [vite](https://vitejs.dev/)
    - [reactjs](https://reactjs.org/)

## Installation
- Clone the repository by using `git clone git@github.com:jquerin/moral-ambiguity.git`.
- Open the repo `cd moral-ambiguity`
- Modify the `.env` file with desired ports and database configuration
- Install docker
- Run the below commands to start the server

```bash
make build
make run
```
- Open browser to `localhost:<.env port>`
  - NOTE: to see swagger documentation, run `make docbuild` before `make run`
    - Then navigate to `localhost:<.env port>/swagger.index.html`
