module View exposing (..)

import Advanced.View
import Basic.View
import Html exposing (Html)
import Model exposing (Model, Page(..))
import Msg exposing (Msg(..))


view : Model -> Html Msg
view model =
    case model.page of
        PageBasic m ->
            Basic.View.view m

        PageAdvanced m ->
            Advanced.View.view m
