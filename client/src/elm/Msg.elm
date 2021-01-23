module Msg exposing (..)

import Http


type Msg
    = Inc
    | TestServer
    | OnServerResponse (Result Http.Error String)
