module View exposing (..)

import Html exposing (Html, a, button, div, h1, header, img, p, span, text)
import Html.Attributes exposing (class, href, src)
import Html.Events exposing (onClick)
import Model exposing (Model)
import Msg exposing (Msg(..))


view : Model -> Html Msg
view model =
    div [ class "container p-2" ]
        [ header [ class "grid-cols-3" ]
            [ span [] [ img [ src "/images/logo.png" ] [] ]
            , div [] [ span [ class "icon" ] [] ]
            , h1 [ class "text-2xl font-bold ml-2" ] [ text "Elm 0.19.1 Webpack Starter, with hot-reloading" ]
            ]
        , p [] [ text "Click on the button below to increment the state." ]
        , div [ class "flex flex-row justify-between" ]
            [ div [ class "flex flex-row items-center" ]
                [ button
                    [ class "border border-green-500 bg-green-500 text-white rounded-md px-4 py-2 m-2 transition duration-500 ease select-none hover:bg-green-600 focus:outline-none focus:shadow-outline"
                    ]
                    [ text "+ 3" ]
                ]
            , div [ class "flex flex-row items-center" ]
                [ button
                    [ class "border border-green-500 bg-green-500 text-white rounded-md px-4 py-2 m-2 transition duration-500 ease select-none hover:bg-green-600 focus:outline-none focus:shadow-outline"
                    , onClick TestServer
                    ]
                    [ text "ping dev server" ]
                , text model.serverMessage
                ]
            ]
        , p [] [ text "Then make a change to the source code and see how the state is retained after recompilation." ]
        , p []
            [ text "And now don't forget to add a star to the Github repo "
            , a [ href "https://github.com/simonh1000/elm-webpack-starter" ] [ text "elm-webpack-starter" ]
            ]
        ]
