module Msg exposing (..)

import Http
import Model exposing (Page)


type Msg
    = TestServer
    | OnServerResponse (Result Http.Error String)
    | ChangePage Page
