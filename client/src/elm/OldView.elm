module OldView exposing (..)

import Dict
import Html exposing (Html, div, input, label, text)
import Html.Attributes exposing (checked, class, classList, style, title, type_, value)
import Html.Events exposing (onClick, onInput)
import List
import Material.Icons exposing (help)
import Material.Icons.Types exposing (Coloring(..))
import Tailwind exposing (tailwind, withClasses)
import Tailwind.Classes exposing (border, border_black, flex, flex_col, flex_row, font_bold, items_center, justify_center, m_1, mb_10, mb_2, mt_2, p_1, p_2, p_3, pb_1, pl_1, pl_2, pr_1, pr_2, pt_1, rounded, text_2xl, text_left, w_16)


view : Model -> Html Msg
view model =
    let
        getCurrentMode newMode =
            case newMode of
                Just m ->
                    case m of
                        Basic o ->
                            [ segmentedItem "Basic" model.mode (Basic o) (SetSelectedMode <| Basic o)
                            , segmentedItem "Advanced" model.mode (Advanced initAdvancedCheckboxOptions) (SetSelectedMode <| Advanced initAdvancedCheckboxOptions)
                            ]

                        Advanced o ->
                            [ segmentedItem "Basic" model.mode (Basic Nothing) (SetSelectedMode <| Basic Nothing)
                            , segmentedItem "Advanced" model.mode (Advanced o) (SetSelectedMode <| Advanced o)
                            ]

                Nothing ->
                    [ segmentedItem "Basic" model.mode (Basic Nothing) (SetSelectedMode <| Basic Nothing)
                    , segmentedItem "Advanced" model.mode (Advanced initAdvancedCheckboxOptions) (SetSelectedMode <| Advanced initAdvancedCheckboxOptions)
                    ]

        header =
            [ getCurrentMode model.mode
                |> div [ tailwind <| withClasses [ "segmented-control" ] <| [ mb_10 ] ]
            ]
                |> div [ tailwind [ flex, justify_center ] ]

        body =
            case model.mode of
                Just m ->
                    case m of
                        Basic o ->
                            div [ tailwind [ flex, justify_center, flex_col ] ] <|
                                [ div [ tailwind <| withClasses [ "segmented-control" ] <| [ flex, justify_center, mb_10 ] ] <|
                                    [ segmentedItem "Minor QOL Enhancement" o MinorQolEnhancement (SetSelectedBasicOption MinorQolEnhancement)
                                    , segmentedItem "QOL Only" o QolOnly (SetSelectedBasicOption QolOnly)
                                    , segmentedItem "Vanilla" o Vanilla (SetSelectedBasicOption Vanilla)
                                    , segmentedItem "Better" o Better (SetSelectedBasicOption Better)
                                    , segmentedItem "Good" o Good (SetSelectedBasicOption Good)
                                    , segmentedItem "Great" o Great (SetSelectedBasicOption Great)
                                    , segmentedItem "Fantastic" o Fantastic (SetSelectedBasicOption Fantastic)
                                    , segmentedItem "Zomg" o Zomg (SetSelectedBasicOption Zomg)
                                    ]
                                , div [ tailwind [ flex, justify_center ] ] <|
                                    [ submitButton "Generate" (o == Nothing) GenerateBasic
                                    ]
                                ]

                        Advanced o ->
                            [ randomization o
                            , other o
                            , qualityOfLife o
                            , dropRates o
                            , [ div [ tailwind [ flex, justify_center, mt_2 ] ] <|
                                    [ submitButton "Save Config" False SaveConfig
                                    ]
                              ]
                            ]
                                |> List.concat
                                |> div [ tailwind [ flex, justify_center, flex_col ] ]

                Nothing ->
                    div [] []
    in
    div [ class "content-container" ]
        [ header
        , body
        ]


randomization : AdvancedCheckboxOptions -> List (Html Msg)
randomization advancedOptions =
    let
        value =
            Dict.get "UseSeed" advancedOptions.checkboxes

        seed =
            case value of
                Just v ->
                    if v.isChecked then
                        div [ tailwind [ flex, items_center, p_1 ] ] [ text <| "Seed: " ++ String.fromInt advancedOptions.seed ]

                    else
                        div [] []

                Nothing ->
                    div [] []
    in
    [ div [ tailwind [ flex, justify_center ] ]
        [ div [ tailwind [ flex_col ], style "min-width" "600px" ]
            [ div [ tailwind [ text_left, font_bold, text_2xl ] ] [ text "Randomization" ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ div [ tailwind [ flex_col ] ]
                    [ checkboxInput "Randomize" advancedOptions
                    , checkboxInput "AllowDupProps" advancedOptions
                    , checkboxInput "UseSeed" advancedOptions
                    ]
                , div [ tailwind [ flex_col ] ]
                    [ checkboxInput "UseOSkills" advancedOptions
                    , checkboxInput "PerfectProps" advancedOptions
                    , seed
                    ]
                , div [ tailwind [ flex_col ] ]
                    [ checkboxInput "BalancedPropCount" advancedOptions
                    , checkboxInput "IsBalanced" advancedOptions
                    ]
                ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ numberInput "MinProps" advancedOptions
                , numberInput "MaxProps" advancedOptions
                ]
            ]
        ]
    ]


other : AdvancedCheckboxOptions -> List (Html Msg)
other advancedOptions =
    [ div [ tailwind [ flex, justify_center ] ]
        [ div [ tailwind [ flex_col ], style "min-width" "600px" ]
            [ div [ tailwind [ text_left, font_bold, text_2xl ] ] [ text "Other Awesome Options" ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ checkboxInput "MeleeSplash" advancedOptions
                ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ numberInput "MonsterDensity" advancedOptions
                ]
            , div [ tailwind <| withClasses [ "segmented-control" ] <| [ p_3, flex, mb_2 ] ] <|
                [ segmentedItem "None" (Just None) advancedOptions.itemGenerationMode (SetItemGenerationMode None)
                , segmentedItem "Randomize" (Just Randomize) advancedOptions.itemGenerationMode (SetItemGenerationMode Randomize)
                , segmentedItem "Generate" (Just Generate) advancedOptions.itemGenerationMode (SetItemGenerationMode Generate)
                ]
            ]
        ]
    ]


qualityOfLife : AdvancedCheckboxOptions -> List (Html Msg)
qualityOfLife advancedOptions =
    [ div [ tailwind [ flex, justify_center ] ]
        [ div [ tailwind [ flex_col ], style "min-width" "600px" ]
            [ div [ tailwind [ text_left, font_bold, text_2xl ] ] [ text "Quality of Life" ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ div [ tailwind [ flex_col ] ]
                    [ checkboxInput "EnableTownSkills" advancedOptions
                    , checkboxInput "IncreaseStackSizes" advancedOptions
                    , checkboxInput "RemoveUniqueCharmLimit" advancedOptions
                    , checkboxInput "RemoveLevelRequirements" advancedOptions
                    ]
                , div [ tailwind [ flex_col ] ]
                    [ checkboxInput "StartWithCube" advancedOptions
                    , checkboxInput "RemoveAttRequirements" advancedOptions
                    , checkboxInput "Cowzzz" advancedOptions
                    ]
                ]
            ]
        ]
    ]


dropRates : AdvancedCheckboxOptions -> List (Html Msg)
dropRates advancedOptions =
    [ div [ tailwind [ flex, justify_center ] ]
        [ div [ tailwind [ flex_col ], style "min-width" "600px" ]
            [ div [ tailwind [ text_left, font_bold, text_2xl ] ] [ text "Drop Rates" ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ div [ tailwind [ flex_col ] ]
                    [ checkboxInput "NoDropZero" advancedOptions
                    ]
                , div [ tailwind [ flex_col ] ]
                    [ checkboxInput "QuestDrops" advancedOptions
                    ]
                ]
            , div [ tailwind [ pl_2, pr_2, pt_1, pb_1, flex, flex_row ] ]
                [ numberInput "UniqueItemDropRate" advancedOptions
                , numberInput "RuneDropRate" advancedOptions
                ]
            ]
        ]
    ]


checkboxInput : String -> AdvancedCheckboxOptions -> Html Msg
checkboxInput key advancedOptions =
    let
        value =
            Dict.get key advancedOptions.checkboxes

        ( checkboxValue, checkboxTooltip ) =
            case value of
                Just v ->
                    ( v.isChecked, v.tooltip )

                Nothing ->
                    ( False, "" )
    in
    div [ tailwind [ flex, items_center, p_1 ] ]
        [ label [ tailwind [ flex ], style "cursor" "pointer" ]
            [ input
                [ type_ "checkbox"
                , checked checkboxValue
                , onClick <| (SetCheckedState <| ToggleCheckbox advancedOptions key)
                , tailwind [ m_1 ]
                , style "width" "1em"
                , style "height" "1em"
                ]
                []
            , text key
            ]
        , label [ title checkboxTooltip ] [ help 20 Inherit ]
        ]


numberInput : String -> AdvancedCheckboxOptions -> Html Msg
numberInput key advancedOptions =
    let
        dictValue =
            Dict.get key advancedOptions.numberInputs

        ( numValue, numberTooltip ) =
            case dictValue of
                Just v ->
                    ( v.value, v.tooltip )

                Nothing ->
                    ( 0, "" )
    in
    div [ tailwind [ flex, items_center, p_1 ] ]
        [ label []
            [ input
                [ type_ "number"
                , value <| String.fromFloat numValue
                , onInput <| (SetAdvancedInt << SetInputValue advancedOptions key << Maybe.withDefault 0 << String.toFloat)
                , tailwind [ m_1, w_16, border, border_black, rounded, pl_1, pr_1 ]
                ]
                []
            , text key
            ]
        , label [ title numberTooltip ] [ help 20 Inherit ]
        ]


segmentedItem : String -> Maybe a -> a -> Msg -> Html Msg
segmentedItem textLabel original new msg =
    div [ onClick <| msg, classList [ ( "segmented-control-item", True ), ( "selected", original == Just new ) ] ] [ text textLabel ]


submitButton : String -> Bool -> Msg -> Html Msg
submitButton buttonText isDisabled msg =
    let
        onClickMsg =
            if isDisabled then
                class ""

            else
                onClick msg
    in
    div [ onClickMsg, classList [ ( "submit-button", True ), ( "disabled", isDisabled ) ] ] [ text buttonText ]
