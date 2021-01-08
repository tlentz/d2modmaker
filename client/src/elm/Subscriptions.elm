module Subscriptions exposing (subscriptions)

import Types exposing (Model, Msg(..))


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none
