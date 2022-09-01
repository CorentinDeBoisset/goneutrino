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

If you want to open an MR, be sure to run the tests with:

    golangci-lint run
    go test ./...

If you want to run all these tests automatically before every commit, add the custom git-hooks with:

    git config core.hooksPath .githooks
