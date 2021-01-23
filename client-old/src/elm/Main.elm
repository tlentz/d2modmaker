module Main exposing (main)

import Browser
import Types exposing (Model, Msg(..), emptyModel)
import Update exposing (update)
import View exposing (view)


main : Program () Model Msg
main =
    Browser.element
        { init = \_ -> init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        }


init : ( Model, Cmd Msg )
init =
    ( emptyModel, Cmd.none )
