# Getting started in Go

1. If you have not already done so, follow the Google Genomics
   [sign up instructions](https://cloud.google.com/genomics/install-genomics-tools#authenticate)
   to generate and download a valid ``client_secrets.json`` file.

2. [Install go](http://golang.org/doc/install)

3. Set a [$GOPATH](https://github.com/golang/go/wiki/GOPATH) for
   installation of dependent packages. For example to set it to the
   current working directory:

    ```
    export GOPATH=$(pwd)
    ```

4. Get the client library and oauth dependencies.

    ```
    go get golang.org/x/oauth2
    go get golang.org/x/oauth2/google
    go get google.golang.org/api/genomics/v1
    ```

5. Run the code:

    ```
    go run main.go
    ```

# More information

* [Google Genomics client library](https://cloud.google.com/genomics/v1/libraries)
* [GoDoc reference for the Genomics API](https://godoc.org/google.golang.org/api/genomics/v1)

# Troubleshooting

[File an issue](https://github.com/googlegenomics/getting-started-with-the-api/issues/new)
if you run into problems and we'll do our best to help!
