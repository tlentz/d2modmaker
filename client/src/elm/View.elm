module View exposing (view)

import Color
import Html exposing (Html, div, text, input)
import Html.Attributes as Attrs exposing (checked, class, classList, title, type_, style)
import Html.Events exposing (onClick, onCheck)
import List
import Material.Icons exposing (help)
import Material.Icons.Types exposing (Coloring(..))
import Tailwind exposing (tailwind, withClasses)
import Tailwind.Classes exposing (content_center, justify_center, flex, flex_col, m_5, m_10, mb_10, m_24, p_5, text_left)
import Types
    exposing
        ( AdvancedCheckboxOptions
        , AdvancedIntMsg(..)
        , BasicOption(..)
        , CheckboxMsg(..)
        , initAdvancedCheckboxOptions
        , ItemGenerationMode(..)
        , initMinProps
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
                    [ text txt ]
                        |> div []

                Nothing ->
                    div [] []
        
        mode =
            model.mode
        
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
            [ getCurrentMode mode
                |> div [ tailwind <| withClasses [ "segmented-control" ] <| [ mb_10 ] ]
            ]
                |> div [ tailwind [ flex, justify_center ] ]

        body =
            case mode of
                Just m ->
                    case m of
                        Basic o ->
                            div [ tailwind [ flex, justify_center, flex_col ] ] <|
                                [ div [ tailwind <| withClasses [ "segmented-control"] <| [ flex, justify_center, mb_10 ] ] <|
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
                            [ 
                            --     randomization o
                            other o
                            -- , qualityOfLife o
                            -- , dropRates o
                            , [ div [ tailwind [ flex, justify_center ] ] <|
                                    [ submitButton "Save Config" False SaveConfig
                                    ]
                                ]
                            ]
                                |> List.concat
                                |> div [ tailwind [ flex, justify_center, flex_col ] ]

                
                Nothing ->
                    div [] []
    in
    div
        [ style "height" "100%"
        , style "width" "100%"
        , style "padding" "25px"
        , style "background-color" Style.backgroundTheme
        , class "content-container"
        ]
        [ header
        , body
        ]


-- randomization : AdvancedCheckboxOptions -> List (Element Msg)
-- randomization advancedOptions =
--     let
--         seed =
--             if advancedOptions.useSeed.isChecked then
--                 column [] <| 
--                     [ text <| String.fromInt advancedOptions.seed ]
--             else
--                 Element.none
--     in
--     [ row [ heading 1, alignLeft ] <| [ h2 [] <| text "Randomization" ]
--     , row [ spacing 75, paddingXY 15 0 ] 
--         [ column [ spacing 15 ] <|
--             [ 
--         --         Input.checkbox []
--         --         { onChange = SetCheckedState << SetRandomize advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.randomize.isChecked
--         --         , label = inputLabel "Randomize" advancedOptions.randomize.tooltip
--         --         }
--         --     , Input.checkbox []
--         --         { onChange = SetCheckedState << SetAllowDupProps advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.allowDupProps.isChecked
--         --         , label = inputLabel "AllowDupProps" advancedOptions.allowDupProps.tooltip
--         --         }
--         --     , Input.checkbox []
--         --         { onChange = SetCheckedState << SetBalancedPropCount advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.balancedPropCount.isChecked
--         --         , label = inputLabel "BalancedPropCount" advancedOptions.balancedPropCount.tooltip
--         --         }
--         --     ]
--         -- , column [ spacing 15, alignTop ] <|
--         --     [ Input.checkbox []
--         --         { onChange = SetCheckedState << SetUseOSkills advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.useOSkills.isChecked
--         --         , label = inputLabel "UseOSkills" advancedOptions.useOSkills.tooltip
--         --         }
--         --     , Input.checkbox []
--         --         { onChange = SetCheckedState << SetPerfectProps advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.perfectProps.isChecked
--         --         , label = inputLabel "PerfectProps" advancedOptions.perfectProps.tooltip
--         --         }                    
--         --     ]
--         -- , column [ spacing 15, alignTop ] <|
--         --     [ row [ spacing 25 ] 
--         --         [ column [] 
--         --             [ Input.checkbox []
--         --                 { onChange = SetCheckedState << SetUseSeed advancedOptions
--         --                 , icon = Input.defaultCheckbox
--         --                 , checked = advancedOptions.useSeed.isChecked
--         --                 , label = inputLabel "UseSeed" advancedOptions.useSeed.tooltip
--         --                 }
--         --             ]
--         --         , seed
--         --         ]
--         --     , Input.checkbox []
--         --         { onChange = SetCheckedState << SetIsBalanced advancedOptions
--         --         , icon = Input.defaultCheckbox
--         --         , checked = advancedOptions.isBalanced.isChecked
--         --         , label = inputLabel "IsBalanced" advancedOptions.useOSkills.tooltip
--         --         }                       
--         --     ]
--         ]
--     , row [ spacing 75, paddingXY 15 0 ]
--         [ column [ spacing 15 ] <|
--             [ Input.text [ padding 3, width <| minimum 60 fill, htmlAttribute <| type_ "number", htmlAttribute <| Attrs.min "0", htmlAttribute <| Attrs.max "20" ]
--                 { onChange = SetAdvancedInt << SetMinProps advancedOptions << Maybe.withDefault 0 << String.toInt
--                 , text = String.fromFloat advancedOptions.minProps.value
--                 , placeholder = Just <| Input.placeholder [] <| text "0"
--                 , label = inputLabel "MinProps" advancedOptions.minProps.tooltip
--                 }
--             ]
--         , column [ spacing 15 ] <|
--             [ Input.text [ padding 3, width <| minimum 60 fill, htmlAttribute <| type_ "number", htmlAttribute <| Attrs.min "0", htmlAttribute <| Attrs.max "20" ]
--                 { onChange = SetAdvancedInt << SetMaxProps advancedOptions << Maybe.withDefault 0 << String.toInt
--                 , text = String.fromFloat advancedOptions.maxProps.value
--                 , placeholder = Just <| Input.placeholder [] <| text "0"
--                 , label = inputLabel "MaxProps" advancedOptions.maxProps.tooltip
--                 }
--             ]
--         ]
--     ]


other : AdvancedCheckboxOptions -> List (Html Msg)
other advancedOptions =
    [ div [ tailwind [ text_left ] ] [ text "Other Awesome Options" ]
    , div [ tailwind [ m_10, p_5, flex_col ] ] 
        [ div [ tailwind [ m_5 ] ]
            [ 
                -- input [ type_ "checkbox", onClick <| SetCheckedState << SetMeleeSplash advancedOptions, checked advancedOptions.meleeSplash.isChecked ] []
                -- Input.checkbox []
                -- { onChange = SetCheckedState << SetMeleeSplash advancedOptions
                -- , icon = Input.defaultCheckbox
                -- , checked = advancedOptions.meleeSplash.isChecked
                -- , label = inputLabel "MeleeSplash" advancedOptions.meleeSplash.tooltip
                -- }
            ]
    --     , column [ spacing 15 ] <|
    --         [ Input.text [ padding 3, width <| minimum 60 fill, htmlAttribute <| type_ "number", htmlAttribute <| Attrs.min "1", htmlAttribute <| Attrs.max "30" ]
    --             { onChange = SetAdvancedInt << SetMonsterDensity advancedOptions << Maybe.withDefault 0 << String.toInt
    --             , text = String.fromFloat advancedOptions.monsterDensity.value
    --             , placeholder = Just <| Input.placeholder [] <| text "0"
    --             , label = inputLabel "MonsterDensity" advancedOptions.monsterDensity.tooltip
    --             }
    --         ]
        ]
    -- , row [ spacing 75, paddingXY 15 0 ]
    --     [ column [ spacing 15 ] <|
    --         [ row [] 
    --             [ Element.html <|
    --                 div [ class "segmented-control", style "margin-right" "5px" ] <|
    --                     [ segmentedItem "None" (Just None) advancedOptions.itemGenerationMode (SetItemGenerationMode None)
    --                     , segmentedItem "Randomize" (Just Randomize) advancedOptions.itemGenerationMode (SetItemGenerationMode Randomize)
    --                     , segmentedItem "Generate" (Just Generate) advancedOptions.itemGenerationMode (SetItemGenerationMode Generate)
    --                     ]
    --             , el [] <| text "ItemGenerationMode"
    --             ]
    --         ]
    --     ]
    ]


-- qualityOfLife : AdvancedCheckboxOptions -> List (Element Msg)
-- qualityOfLife advancedOptions =
--     [ row [ heading 1, alignLeft ] <| [ h2 [] <| text "Quality of Life" ]
--     , row [ spacing 75, paddingXY 15 0 ] 
--         [ column [ spacing 15 ] <|
--             [ Input.checkbox []
--                 { onChange = SetCheckedState << SetEnableTownSkills advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.enableTownSkills.isChecked
--                 , label = inputLabel "EnableTownSkills" advancedOptions.enableTownSkills.tooltip
--                 }
--             , Input.checkbox []
--                 { onChange = SetCheckedState << SetIncreaseStackSizes advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.increaseStackSizes.isChecked
--                 , label = inputLabel "IncreaseStackSizes" advancedOptions.increaseStackSizes.tooltip
--                 }
--             , Input.checkbox []
--                 { onChange = SetCheckedState << SetRemoveUniqueCharmLimit advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.removeUniqueCharmLimit.isChecked
--                 , label = inputLabel "RemoveUniqueCharmLimit" advancedOptions.removeUniqueCharmLimit.tooltip
--                 }
--             ]
--         , column [ spacing 15, alignTop ] <|
--             [ Input.checkbox []
--                 { onChange = SetCheckedState << SetStartWithCube advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.startWithCube.isChecked
--                 , label = inputLabel "StartWithCube" advancedOptions.startWithCube.tooltip
--                 }
--             , Input.checkbox []
--                 { onChange = SetCheckedState << SetRemoveAttRequirements advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.removeAttRequirements.isChecked
--                 , label = inputLabel "RemoveAttRequirements" advancedOptions.removeAttRequirements.tooltip
--                 }                               
--             ]
--         , column [ spacing 15, alignTop ] <|
--             [ Input.checkbox []
--                 { onChange = SetCheckedState << SetCowzzz advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.cowzzz.isChecked
--                 , label = inputLabel "Cowzzz" advancedOptions.cowzzz.tooltip
--                 }
--             , Input.checkbox []
--                 { onChange = SetCheckedState << SetRemoveLevelRequirements advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.removeLevelRequirements.isChecked
--                 , label = inputLabel "RemoveLevelRequirements" advancedOptions.removeLevelRequirements.tooltip
--                 }                         
--             ]
--         ]
--     ]


-- dropRates : AdvancedCheckboxOptions -> List (Element Msg)
-- dropRates advancedOptions =
--     [ row [ heading 1, alignLeft ] <| [ h2 [] <| text "Drop Rates" ]
--     , row [ spacing 75, paddingXY 15 0 ] 
--         [ column [ spacing 15 ] <|
--             [ Input.checkbox []
--                 { onChange = SetCheckedState << SetNoDropZero advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.noDropZero.isChecked
--                 , label = inputLabel "NoDropZero" advancedOptions.noDropZero.tooltip
--                 }
--             ]
--         , column [ spacing 15, alignTop ] <|
--             [ Input.checkbox []
--                 { onChange = SetCheckedState << SetQuestDrops advancedOptions
--                 , icon = Input.defaultCheckbox
--                 , checked = advancedOptions.questDrops.isChecked
--                 , label = inputLabel "QuestDrops" advancedOptions.questDrops.tooltip
--                 }                            
--             ]
--         ]
--     , row [ spacing 75, paddingXY 15 0 ]
--         [ column [ spacing 15 ] <|
--             [ Input.text [ padding 3, width <| minimum 60 fill, htmlAttribute <| type_ "number", htmlAttribute <| Attrs.min "1", htmlAttribute <| Attrs.max "100" ]
--                 { onChange = SetAdvancedInt << SetUniqueItemDropRate advancedOptions << Maybe.withDefault 0 << String.toInt
--                 , text = String.fromFloat advancedOptions.uniqueItemDropRate.value
--                 , placeholder = Just <| Input.placeholder [] <| text "0"
--                 , label = inputLabel "UniqueItemDropRate" advancedOptions.uniqueItemDropRate.tooltip
--                 }
--             ]
--         , column [ spacing 15 ] <|
--             [ Input.text [ padding 3, width <| minimum 60 fill, htmlAttribute <| type_ "number", htmlAttribute <| Attrs.min "1", htmlAttribute <| Attrs.max "100" ]
--                 { onChange = SetAdvancedInt << SetRuneDropRate advancedOptions << Maybe.withDefault 0 << String.toInt
--                 , text = String.fromFloat advancedOptions.runeDropRate.value
--                 , placeholder = Just <| Input.placeholder [] <| text "0"
--                 , label = inputLabel "RuneDropRate" advancedOptions.runeDropRate.tooltip
--                 }
--             ]
--         ]
--     ]


-- inputLabel : String -> String -> Input.Label Msg
-- inputLabel label tooltip =
--     Input.labelRight [ centerY ] <|
--         column [ onRight
--             (Element.html <| div [ title tooltip, style "margin-left" "5px", style "cursor" "pointer" ] [ help 20 Inherit ]
--             ) ] <| [ text label ]


segmentedItem : String -> Maybe a -> a -> Msg -> Html Msg
segmentedItem textLabel original new msg =
    div [ onClick <| msg, classList [ ("segmented-control-item", True), ( "selected", original == Just new) ] ] [ text textLabel ]


submitButton : String -> Bool -> Msg -> Html Msg
submitButton buttonText isDisabled msg =
    let
        onClickMsg =
            if isDisabled then
                DoNothing
            
            else
                msg
    in
    div [ onClick onClickMsg, classList [ ("submit-button", True), ( "disabled", isDisabled) ] ] [ text buttonText ]