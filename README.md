Template file writer, with EnvKey.com support.

- definte a template file, use `${ENV_VAR_NAME}` syntax in it
- make sure an [envkey](https://envkey.com) file or env variable is present
- write a new file


`v1.0.0`


# Example


Your template: `some.conf.tpl`:

```

host "${HOSTNAME}"

target "${SECRET_URL}"

```

then run:

`envkey-template some.conf.tpl some.conf`


Assuming that HOSTNAME and SECRET_URL are set in either the environment or in your EnvKey app then your output will be:

```
host "bananas.dev"

target "123:s3cr1t@foobar.internal"

```
