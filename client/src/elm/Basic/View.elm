module Basic.View exposing (..)

import Basic.Model exposing (Model)
import Html exposing (Html, text)
import Msg exposing (Msg)


view : Model -> Html Msg
view model =
    text "Basic.View.view"
