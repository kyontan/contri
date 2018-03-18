# contri

Executes docker container simply like as CGI.

## About it / Motivation

I want to execute container from remote, not a web server, without complex settings, port mapping, or something else.

Like webhook, trigger container, and container stops (not waiting another request) after each request.

Settings like below makes server allows endpoint `POST /` end returns stdout of `date` to response body.

```yaml
endpoints:
 - path: /
   image: alpine
   cmd: date
```

A little bit practical settings will be supported.

```yaml
endpoints:
 - path: /
   image: alpine
   cmd: date
 - path: /yeah
   options: -v /srv:/src
   image: some_image
   cmd: some_command
   args: "-foo {.foo} -bar {.bar}"
```

If server is requested with `POST /yeah` with data `foo=hello&bar=contri`, contri executes like `docker run -v /srv:/srv some_image some_command -foo hello -bar contri`.

Also, if contri support `cgi_mode: true`, it trigger container with CGI environment, and container must print response header and body.
