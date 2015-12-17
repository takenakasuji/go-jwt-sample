package main

import (
  "log"
  "net/http"
  "time"

  "github.com/StephanDollberg/go-json-rest-middleware-jwt"
  "github.com/ant0ine/go-json-rest/rest"
)

func handle_auth(w rest.ResponceWriter, r *rest.Request) {
  w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}

func main() {
  jwt_middleware := &jwt.JWTMiddleware{
    Key:  []byte("secret key"),
    Realm:  "jwt auth",
    Timeout: time.Hour,
    MaxRefresh: time.Hour * 24

  }
}
