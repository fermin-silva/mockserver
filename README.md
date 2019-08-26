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
mkdir sources
mkdir sources/users
```

2. Create two mocked server responses
```sh
echo '{ "message" : "hello world" }' > sources/index.json
echo '{ "user" : "Martin Scorsese" }' > sources/users/1.json
```

3. Start the `mockserver` poiting to the `sources` directory
```sh
./mockserver sources
```

4. Hit the HTTP server from your browser or from curl and see the response
```sh
curl "localhost:8080"
{ "message" : "hello world" }

curl "localhost:8080/users/1"
{ "user" : "Martin Scorsese" }
```

## How does it work?

Basically `mockserver` is a web server that listens on a given port for all
incoming URLs. When its hitted, it looks for a file that matches the URL within
the serving directory, parses the file, renders it as a template (if applicable)
and returns it as a response, along with any custom header you specify.

So the two main concepts to understand are how the URL matching works and how
the file is rendered as a template. Both of them depend on the `Front Matter` of
each of your files. All of these concepts are explained below:

### Front Matter of files

The Front Matter of your files allows you to customize the response rendering,
the file matching and much more. It's inspired in [Jekyll's YAML Front Matter](https://jekyllrb.com/docs/front-matter/), but the format and variables are different.

The front matter must be the first thing in the file and must take the form of
a valid TOML (and not YAML like Jekyll) set between triple-dashed lines. Here is
a basic example:

```
---
template = true

match = [
	"/index.html",
    "/index.php"
]

[Headers]
Content-Type = "Application/json"
Whatever = "You Like"
---
... rest of your file goes here ...
```

Between these triple-dashed lines, you can set predefined variables or even
create custom ones of your own. These variables will then be available for you
to access in the template rendering, and customize the URL matching. See the
full reference of front matter variables [HERE]().

If in doubt of how to write a valid TOML document you can check the [official
language specs](https://github.com/toml-lang/toml) or the specific [TOML
library](https://github.com/BurntSushi/toml) we use.


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

Template rendering helps you to return a dynamic responses, based on query
parameters, configuration files, etc. You might also want to use templates to
simplify or modularize files, so they are easier to mantain.

We support the [Django template language](https://docs.djangoproject.com/en/2.2/ref/templates/language/) through the go library [pongo2](https://github.com/flosch/pongo2).
You shouldn't need to see much of pongo2, unless you find a Django feature that
is not working as expected.

In order to enable template processing, you need to add `template = true` to the
Front Matter of a file and then use normal Django tags. For example with an
`index.json` file like this:

```json
---
template = true

[MyCustomData]
names = [ "Juan", "Pedro", "Miguel" ]
---
{
    "hello" : "{{ Request.Query("name") }}",
    "other_names" : [
        {% for name in File.Get("MyCustomData").names %}
            "{{name}}"{% if not forloop.Last %},{% endif %}
        {% endfor %}
    ]
}
```

If we do `curl localhost:8080/?name=john` we get:

```json
{
    "hello" : "john",
    "other_names" : [
            "Juan",
            "Pedro",
            "Miguel"
    ]
}
```

This is just a simple example of templating a file with custom data, but you can
use much more complex logic, imports, ifs, etc. See the full reference of template variables and examples [HERE]().

## Configuration

COMING SOON

## Examples

COMING SOON

## For Developers

COMING SOON

### How to compile from source

1. Install go: https://golang.org/doc/install#install (version 1.11.2 is being
used to develop the project, although older versions might work)

2. Clone the repository into the correct folder of your $GOPATH, normally:
```sh
git clone https://github.com/fermin-silva/mockserver $GOPATH/src/github.com/fermin-silva/mockserver
```

3. Install the Glide dependency management tool: https://glide.sh

4. Install the dependencies of the project into the local `vendor` project:
```sh
cd $GOPATH/src/github.com/fermin-silva/mockserver
glide install
```

5. Compile the project:
```sh
go build
```

This will leave an executable file `mockserver` (or `mockserver.exe` if you are
on Windows). From there you can run it just like the downloaded version.

If you need to cross compile the code (for ex. compiling on Mac but want to run
`mockserver` in your linux server), please read this article on how
[How To Build Go Executables for Multiple Platforms](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

### How to contribute

COMING SOON