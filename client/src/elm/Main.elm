module Main exposing (init, main)

import Browser
import Model exposing (Model)
import Msg exposing (Msg(..))
import Update exposing (update)
import View exposing (view)


main : Program Int Model Msg
main =
    Browser.document
        { init = init
        , update = update
        , view =
            \m ->
                { title = "Elm 0.19 starter"
                , body = [ view m ]
                }
        , subscriptions = \_ -> Sub.none
        }


init : Int -> ( Model, Cmd Msg )
init flags =
    ( { serverMessage = "" }, Cmd.none )
