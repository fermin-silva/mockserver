# mockserver
The hassle-free way of mocking a backend server API, match your json files with dynamic urls and do templating. Useful for integration testing, frontend development and even production websites or apps.

## Install

Download the single executable file from the releases section and run it in a
directory with your source files or Install from source

## What is mockserver ?

`mockserver` is a web server that you run from the command line. It scans a
directory of source files and serves them dynamically based on the received URL.

For example, if the endpoint you want to mock is `/users/1234`, simply create a
directory `users` and inside of it create the file `1234.json`. `mockserver` will
then take care of serving it correctly.

It supports dynamic URLs and templating inside any kind of file in case you need
dynamic json files, see the Features section below for more information.

## Why mockserver ?

Frontend development needs super fast iterations and more often than not the
backend development is not fast enough. Current mocking solutions are either
complicated to setup, are hosted online and-or are paid.

With `mockserver` you can start developing a frontend with no backend and build
it as you go, creating source files with the response you need from the API.

You can also use it for integration tests of your frontend, or to reproduce a
hard to catch bug by returning specific responses from the backend.

## Quick start

1. Create folder where we will store our source json files (or any type really),
called the serving directory
```sh
$ mkdir sources
$ mkdir sources/users
```

2. Create two mocked server responses
```sh
$ echo '{ "message" : "hello world" }' > sources/index.json
$ echo '{ "user" : "Martin Scorsese" }' > sources/users/1.json
```

3. Start the `mockserver` poiting to the `sources` directory
```sh
$ ./mockserver sources
```

4. Hit the HTTP server from your browser or from curl and see the response
```sh
$ curl "localhost:8080"
{ "message" : "hello world" }

$ curl "localhost:8080/users/1"
{ "user" : "Martin Scorsese" }
```

## How does it work?

Basically `mockserver` is a web server that listens on a given port for all
incoming URLs. When its hitted, it looks for a file that matches the URL within
the serving directory, parses the file, renders it as a template (if applicable)
and returns it as a response, along with any custom header you specify.

So the two main concepts to understand are how the URL matching works and how is
the file rendered as a template. Both of them depend on the `Front Matter` of each
of your files. All of these concepts are explained below:

### Front Matter of files

COMING SOON

### URL to File Matching

The server always tries to return a file that exactly matches the URL (including
file extension), but most commonly you might want to have specific URLs that are
hard (or impractical) to express with exact file names, so these are the fuzzy
matching rules tried by the server (in the order they are tried):

1. If the path corresponds to a directory: return the first index file inside
the directory. For example, if the URL is `/web/users/` the server would catch
the file `$serving-dir/web/users/index*` (notice the wildcard symbol `*`, the
file name ending is not important)

2. If the path corresponds to an exact match of a file name: return it without
trying anything else. For example, if the URL is `/web/get_users.php` the server
would return the file `$serving-dir/web/get_users.php`. That file would have the
contents of the JSON response you need (and not the php source code!).

3. Try to see if there is a file with the same prefix name (try wildcard at the
end) in the same directory. For example, if the URL is `/web/users/1` the server
would catch the file `$serving-dir/web/users/1.json`, but **not** a file in a deeper
directory like `$serving-dir/web/users/1/index.json`

4. Try to return a 404 file inside. Does this even work? Its not a file nor a dir,
so looking inside for a 404 makes no sense

5. Try to find a 404 file in the parent directory of the URL. For example, if
the URL is `/web/users/2` and there's no file `$serving-dir/web/users/2*` (notice
the wildcard), the server would catch the file `$serving-dir/web/users/404.json`
('parent' URL because it drops the trailing `2` and tries a `404` file)

6. Try to find a 404 file in the serving directory. For example, if the URL is
`/web/users/2` and all other checks failed, the server would return the file
`$serving-dir/404.json`.

Notice: for the sake of this exaplaination we used mostly json, but keep in mind
you could have any extension, and the contents could be anything (plain text, xml,
json, whatever), as nothing is really enforced.

### Template Rendering

COMING SOON

## Configuration

COMING SOON

## Examples

COMING SOON

## For Developers

COMING SOON

### How to compile from source

COMING SOON

### How to contribute

COMING SOON