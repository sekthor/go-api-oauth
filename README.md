# go-api-oauth

This is a simple mock api.
It is secured with `jwt`s to be provided as bearer tokens.
These are validated with `jkws` fetched from a given endpoint.

## Configuration

Provide the following environment variables:

| Variable | Description                                                               |
|:---------|:--------------------------------------------------------------------------| 
| JWKSURL  | the url from which the jwks may be fetched from                           |
| HOST     | the bind interface for our api server (e.g. `0.0.0.0:8081`, `:8081`, ...) |

## Usage

- `/public` is accessible regardless if a valid token is provided or not. If it was, then the response will include the `sub` Subject ID from the token.
- `/authenticated` requires a valid token. If none was provided, it responds with a `401 Unauthorized`
- `/authorized/:id` requires a valid token and the `:id` to match the subject ID (`sub`) from the token. This is simulating accessing a resource for which we need to be the resource owner