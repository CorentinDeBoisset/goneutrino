# Secret-send

📮 Send your secrets securely to third-parties, using client-to-client encryption.

This project is heavily inspired from [Datash](https://github.com/datash/datash).

## ❯ Install

Simply run:

    make
    ./bin/secret-send serve -c config.yml

Then go to [localhost:8080]

## ❯ Contributing

If you want to open an MR, be sure to run the tests with:

    golangci-lint run
    go test ./...

If you want to run all these tests automatically before every commit, add the custom git-hooks with:

    git config core.hooksPath .githooks
