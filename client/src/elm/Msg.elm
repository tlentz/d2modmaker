module Msg exposing (..)

import Http


type Msg
    = TestServer
    | OnServerResponse (Result Http.Error String)
