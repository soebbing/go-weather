# Go Weather

Just a little first Golang project, showing the weather of a location you define on your CLI.

## Preparations

You need to create an account with [OpenWeather](https://home.openweathermap.org/users/sign_up) and get yourself
an API key. This API key then can be provided as an `OPENWEATHER_APP_ID` environment variable to the program:
`export OPENWEATHER_APP_ID=(your app id)`.

## Usage

Just copy the binary somewhere to your path. After you've defined the `OPENWEATHER_APP_ID` environment variable,
you can do something like this:

```
$ weather munich
München: Bedeckt 🌥  bei 7.9°
```

## Todos

- Add tests
- Store last city, use as default