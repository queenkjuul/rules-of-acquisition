# Rules of Acquisition

A digital compendium of all known Ferengi Rules of Acquisition. 

***[WARNING]: This project is not sactioned by the Ferengi Commerce Authority, and does not operate with any form of license. Use at your own risk!***

_(but remember #62: "*the riskier the road, the greater the profit*")_

[see it in action](https://api.queenkjuul.xyz/ferengi)

## Usage

### Command Line

The Rules of Acquisition database ([ferengi.json](./ferengi.json)) is embedded into the binary at build time. Therefore, you only need to download a single binary from the [releases](https://github.com/queenkjuul/rules-of-acquisition/releases) page, and then run it from your shell:

```sh
./rules-of-acquisition # get a random rule
./rules-of-acquisition -id NUMBER # get a specific rule
```

### Web Server (REST API)

You can also run the binary as a web server: a single REST GET endpoint is exposed, which provides a random rule or, optionally, a specific rule by ID.

```sh
./rules-of-acquisition -serve # default server at http://localhost:8080
# in your browser, go to http://localhost:8080
# or in another shell, run
curl http://localhost:8080 # get a random rule
curl http://localhost:8080/NUMBER # get a specific rule
```

You can customize the server:

```
Usage of rules-of-acquisition:
  -address string
        address to listen on (if in serve mode) (default ":8080")
  -cert string
        certificate file to use (if TLS enabled) (default "cert.pem")
  -id int
        number of the specific rule you'd like to retrieve (default -1)
  -key string
        key file to use (if TLS enabled) (default "key.pem")
  -route string
        route, including prefix, that server will listen on (such as when behind a reverse proxy, for example '/ferengi') (default "/")
  -serve
        set to true to run a REST API server
  -tls
        enable TLS (https), must have cert.pem and key.pem
```

## Limitations

Despite the diligent efforts of both the Federation Interspecies Research Agency and the authors of this package, the complete contents of the Rules of Acquisition are not yet known to exist in a human-readable (or even machine-translatable) format. As a result, there are significant gaps in the completeness of this record.

Additionally, some rules are known, but their number is *not* known. Such rules are assigned negative indices in our dataset, starting with `-2`. As long as unknown entries are only appended to this database, their ID number within the database will be stable.

If you have access to a more complete compendium of the Rules, or even just a single rule we have failed to include (or just, like, want to say hello), please open an Issue, or better yet, a Pull Request.

## Development / Modification / Compiling for Other Platforms

Go is super easy to use. Install Go (v1.22+) on your system, clone the repo, and run

```
go build
```

That's it. It'll spit out a single static binary with the data embedded, that you can run from wherever.