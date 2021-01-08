module View exposing (view)

import Element exposing (Element, column, el, fill, height, row, padding, spacing, width, text, centerX, centerY)
import Element.Background as Background exposing (color)
import Element.Border as Border exposing (solid, rounded)
import Element.Input as Input
import Framework.Button exposing (button)
import Html exposing (Html, div)
import Html.Attributes exposing (class, classList)
import Html.Events exposing (onClick)
import Types
    exposing
        ( BasicOption(..)
        , ColorTheme(..)
        , initAdvancedOptions
        , Model
        , Mode(..)
        , Msg(..)
        , View(..)
        )
import Style as Style


view : Model -> Html Msg
view model =
    let
        errorTxt =
            case model.errorMessage of
                Just txt ->
                    text txt
                        |> el []

                Nothing ->
                    Element.none
        
        mode =
            model.mode
        
        getCurrentMode newMode =
            case newMode of
                Just m ->
                    case m of
                        Basic o ->
                            [ segmentedItem "Basic" model.mode (Basic o) (SetSelectedMode <| Basic o)
                            , segmentedItem "Advanced" model.mode (Advanced initAdvancedOptions) (SetSelectedMode <| Advanced initAdvancedOptions)
                            ]
                        
                        Advanced o ->
                            [ segmentedItem "Basic" model.mode (Basic Nothing) (SetSelectedMode <| Basic Nothing)
                            , segmentedItem "Advanced" model.mode (Advanced o) (SetSelectedMode <| Advanced o)
                            ]
                
                Nothing ->
                    [ segmentedItem "Basic" model.mode (Basic Nothing) (SetSelectedMode <| Basic Nothing)
                    , segmentedItem "Advanced" model.mode (Advanced initAdvancedOptions) (SetSelectedMode <| Advanced initAdvancedOptions)
                    ]

        header =
            [ getCurrentMode mode
                |> div [ class "segmented-control" ]
                |> Element.html
            ]
                |> column [ centerX ]

        body =
            case mode of
                Just m ->
                    case m of
                        Basic o ->
                            column [ centerX, spacing 25 ] <|
                                [ column [] <|
                                    [ Element.html <|
                                        div [ class "segmented-control" ] <|
                                            [ segmentedItem "Minor QOL Enhancement" o MinorQolEnhancement (SetSelectedBasicOption MinorQolEnhancement)
                                            , segmentedItem "QOL Only" o QolOnly (SetSelectedBasicOption QolOnly)
                                            , segmentedItem "Vanilla" o Vanilla (SetSelectedBasicOption Vanilla)
                                            , segmentedItem "Better" o Better (SetSelectedBasicOption Better)
                                            , segmentedItem "Good" o Good (SetSelectedBasicOption Good)
                                            , segmentedItem "Great" o Great (SetSelectedBasicOption Great)
                                            , segmentedItem "Fantastic" o Fantastic (SetSelectedBasicOption Fantastic)
                                            , segmentedItem "Zomg" o Zomg (SetSelectedBasicOption Zomg)
                                            ]
                                    ]
                                , column [ centerX ] <|
                                    [ Input.button
                                        [ padding 5, Border.width 2, Border.solid, Border.color Style.borderTheme, rounded 3 ]
                                        { onPress = Just GenerateBasic
                                        , label = text "Generate"
                                        }
                                    ]
                                ]
                        
                        Advanced o ->
                            row [ centerX ] <|
                                [ Input.checkbox []
                                    { onChange = CheckRandomize
                                    , icon = Input.defaultCheckbox
                                    , checked = o.randomize
                                    , label =
                                        Input.labelRight []
                                            (text "Randomize")
                                    }
                                ]
                
                Nothing ->
                    Element.none
    in
    render <|
        column
            [ height <| Element.minimum model.screen.height <| fill
            , width fill
            , Background.color Style.backgroundTheme
            ]
            [ el
                [ padding 50
                , width fill
                ]
                <| header
                , body
            ]


segmentedItem : String -> Maybe a -> a -> Msg -> Html Msg
segmentedItem text original new msg =
    div [ onClick <| msg, classList [ ("segmented-control-item", True), ( "selected", original == Just new) ] ] [ Html.text text ]


render : Element Msg -> Html Msg
render =
    Element.layoutWith
        { options =
            [ Element.focusStyle
                { borderColor = Nothing
                , backgroundColor = Nothing
                , shadow = Nothing
                }
            ]
        }
        [ height fill
        , width fill
        ]