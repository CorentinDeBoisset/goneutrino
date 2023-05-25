Warning: This project is not ready. A lot of development is still required until an alpha could even be released.

<p align="center">
    <img src="https://github.com/corentindeboisset/neutrino/raw/main/frontend-src/src/assets/logo_large.svg" alt="Logo of the neutrino application" style="width: 8rem" />
</p>

# Neutrino

📮 Send your secrets securely to third-parties, using client-to-client encryption.

This project has many of its features inspired from [Datash](https://github.com/datash/datash).

## ❯ Install

Simply run:

    make
    ./bin/neutrino serve -c config.yml

Then go to [localhost:8080]

## ❯ Contributing

### ❯ Development setup

You can first setup a configuration file for your backend with the following structure:

```yaml
# example.yml
host: 127.0.0.1
port: 8081
trusted_proxies: [0.0.0.0/0]
```

Then you can build the backend:

    make dev_build
    ./bin/neutrino serve -c example.yml

Finally, you can install and start the frontend web server:

    cd frontend-src
    npm install
    npm run serve

The application is then available at [http://localhost:8080]

### ❯ Linting and Tests

If you want to open an MR, be sure to run the tests with:

    golangci-lint run
    go test ./...

If you want to run all these tests automatically before every commit, add the custom git-hooks with:

    git config core.hooksPath .githooks
