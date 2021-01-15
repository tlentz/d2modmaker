module Update exposing (update)

import Browser.Dom as Dom
import Dict
import Helper exposing (return, mkCmd)
import Http exposing (Error(..))
import Random
import Task
import Types
    exposing
        ( AdvancedIntMsg(..)
        , AdvancedCheckboxOption
        , AdvancedCheckboxOptions
        , CheckboxMsg(..)
        , CheckboxName
        , Model
        , Mode(..)
        , Msg(..)
        , Route(..)
        , View(..)
        )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        DoNothing ->
            return model []

        Resize screen ->
            ( { model | screen = screen }, Cmd.none )

        SetViewportCb ->
            ( model, Cmd.none )

        FocusOn id ->
            ( model, Dom.focus id |> Task.attempt FocusResult )

        FocusResult _ ->
            ( model, Cmd.none )
        
        GetResponse _ ->
            ( model, Cmd.none )
        
        SetCheckedState checkboxMsg ->
            updateCheckboxState checkboxMsg
                |> setUpdatedOptionsOnModel model
        
        SetAdvancedInt advancedIntMsg ->
            updateAdvancedInt advancedIntMsg
                |> setUpdatedIntOptionsOnModel model

        SetSelectedMode mode ->
            ( { model | mode = Just mode}, Cmd.none )
        
        SetSelectedBasicOption basicOption ->
            let
                mode = model.mode

                updatedMode =
                    case mode of
                        Just m ->
                            let
                                newMode =
                                    case m of
                                        Basic _ ->
                                            Just basicOption
                                                |> Basic
                                        Advanced options ->
                                            m
                            in
                            Just newMode

                        Nothing ->
                            mode
                
                updatedModel =
                    { model | mode = updatedMode }

            in
            ( updatedModel, Cmd.none )
        
        SetItemGenerationMode itemGenerationMode ->
            let
                mode = model.mode

                updatedMode =
                    case mode of
                        Just m ->
                            let
                                newMode =
                                    case m of
                                        Basic _ ->
                                            m

                                        Advanced options ->
                                            let
                                                updatedOptions =
                                                    { options | itemGenerationMode = itemGenerationMode }
                                            in
                                            Advanced updatedOptions
                            in
                            Just newMode

                        Nothing ->
                            mode
                
                updatedModel =
                    { model | mode = updatedMode }

            in
            ( updatedModel, Cmd.none )

        GenerateBasic ->
            ( model, Cmd.none )
        
        SaveConfig ->
            ( model, Cmd.none )


toggle : comparable -> Dict.Dict comparable AdvancedCheckboxOption -> Dict.Dict comparable AdvancedCheckboxOption
toggle key dict =
    Dict.update key
        (\oldValue ->
            case oldValue of
                Just value ->
                    Just <| { value | isChecked = not value.isChecked }
                Nothing ->
                    Nothing
        )
        dict


updateCheckboxState : CheckboxMsg -> (AdvancedCheckboxOptions, Cmd CheckboxMsg)
updateCheckboxState checkboxMsg =
    case checkboxMsg of
        ToggleCheckbox advancedOptions checkboxName ->
            ( { advancedOptions | test = toggle checkboxName advancedOptions.test }, Cmd.none )

        SetRandomize advancedOptions ->
            let
                advancedOption =
                    advancedOptions.randomize

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked}

            in
            ({ advancedOptions | randomize = updatedOption }, Cmd.none)
        
        SetUseSeed advancedOptions ->
            let
                advancedOption =
                    advancedOptions.useSeed

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }
                
                newAdvancedOptions =
                    { advancedOptions | useSeed = updatedOption }
            in
            (newAdvancedOptions, Random.generate (SetSeed newAdvancedOptions) (Random.int 1 Random.maxInt))

        SetSeed advancedOptions newSeed ->
            ({ advancedOptions | seed = newSeed }, Cmd.none)

        SetUseOSkills advancedOptions ->
            let
                advancedOption =
                    advancedOptions.useOSkills

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | useOSkills = updatedOption }, Cmd.none)

        SetPerfectProps advancedOptions ->
            let
                advancedOption =
                    advancedOptions.perfectProps

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | perfectProps = updatedOption }, Cmd.none)

        SetAllowDupProps advancedOptions ->
            let
                advancedOption =
                    advancedOptions.allowDupProps

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | allowDupProps = updatedOption }, Cmd.none)

        SetIsBalanced advancedOptions ->
            let
                advancedOption =
                    advancedOptions.isBalanced

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | isBalanced = updatedOption }, Cmd.none)

        SetBalancedPropCount advancedOptions ->
            let
                advancedOption =
                    advancedOptions.balancedPropCount

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | balancedPropCount = updatedOption }, Cmd.none)

        SetMeleeSplash advancedOptions ->
            let
                advancedOption =
                    advancedOptions.meleeSplash

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | meleeSplash = updatedOption }, Cmd.none)

        SetEnableTownSkills advancedOptions ->
            let
                advancedOption =
                    advancedOptions.enableTownSkills

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | enableTownSkills = updatedOption }, Cmd.none)

        SetStartWithCube advancedOptions ->
            let
                advancedOption =
                    advancedOptions.startWithCube

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | startWithCube = updatedOption }, Cmd.none)

        SetCowzzz advancedOptions ->
            let
                advancedOption =
                    advancedOptions.cowzzz

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | cowzzz = updatedOption }, Cmd.none)

        SetIncreaseStackSizes advancedOptions ->
            let
                advancedOption =
                    advancedOptions.increaseStackSizes

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | increaseStackSizes = updatedOption }, Cmd.none)

        SetRemoveLevelRequirements advancedOptions ->
            let
                advancedOption =
                    advancedOptions.removeLevelRequirements

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | removeLevelRequirements = updatedOption }, Cmd.none)

        SetRemoveAttRequirements advancedOptions ->
            let
                advancedOption =
                    advancedOptions.removeAttRequirements

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | removeAttRequirements = updatedOption }, Cmd.none)

        SetRemoveUniqueCharmLimit advancedOptions ->
            let
                advancedOption =
                    advancedOptions.removeUniqueCharmLimit

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | removeUniqueCharmLimit = updatedOption }, Cmd.none)

        SetNoDropZero advancedOptions ->
            let
                advancedOption =
                    advancedOptions.noDropZero

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | noDropZero = updatedOption }, Cmd.none)

        SetQuestDrops advancedOptions ->
            let
                advancedOption =
                    advancedOptions.questDrops

                updatedOption = 
                    { advancedOption | isChecked = not advancedOption.isChecked }

            in
            ({ advancedOptions | questDrops = updatedOption }, Cmd.none)


updateAdvancedInt : AdvancedIntMsg -> AdvancedCheckboxOptions
updateAdvancedInt advancedIntMsg =
    case advancedIntMsg of
        SetMinProps advancedOptions value ->
            let
                advancedOption =
                   advancedOptions.minProps

                updatedOption = 
                    { advancedOption | value = toFloat(max 0 (min 20 value)) }
            in
            { advancedOptions | minProps = updatedOption }
        
        SetMaxProps advancedOptions value ->
            let
                advancedOption =
                    advancedOptions.maxProps

                updatedOption = 
                    { advancedOption | value = toFloat(max 0 (min 20 value)) }
            in
            { advancedOptions | maxProps = updatedOption }
        
        SetMonsterDensity advancedOptions value ->
            let
                advancedOption =
                    advancedOptions.monsterDensity

                updatedOption = 
                    { advancedOption | value = toFloat(max 1 (min 30 value)) }
            in
            { advancedOptions | monsterDensity = updatedOption }
        
        SetUniqueItemDropRate advancedOptions value ->
            let
                advancedOption =
                    advancedOptions.uniqueItemDropRate

                updatedOption = 
                    { advancedOption | value = toFloat(max 1 (min 100 value)) }
            in
            { advancedOptions | uniqueItemDropRate = updatedOption }
        
        SetRuneDropRate advancedOptions value ->
            let
                advancedOption =
                    advancedOptions.runeDropRate

                updatedOption = 
                    { advancedOption | value = toFloat(max 1 (min 100 value)) }
            in
            { advancedOptions | runeDropRate = updatedOption }


setUpdatedOptionsOnModel : Model -> (AdvancedCheckboxOptions, Cmd CheckboxMsg) -> (Model, Cmd Msg)
setUpdatedOptionsOnModel model advancedOptions =
    let
        mode = model.mode

        updatedMode =
            case mode of
                Just m ->
                    let
                        newMode =
                            case m of
                                Basic _ ->
                                   m
                                Advanced options ->
                                    Advanced <| Tuple.first advancedOptions
                    in
                    Just newMode

                Nothing ->
                    mode
        
        updatedModel =
            { model | mode = updatedMode }

    in
    (updatedModel, Cmd.map SetCheckedState (Tuple.second advancedOptions))


setUpdatedIntOptionsOnModel : Model -> AdvancedCheckboxOptions -> (Model, Cmd Msg)
setUpdatedIntOptionsOnModel model advancedOptions =
    let
        mode = model.mode

        updatedMode =
            case mode of
                Just m ->
                    let
                        newMode =
                            case m of
                                Basic _ ->
                                   m
                                Advanced options ->
                                    Advanced <| advancedOptions
                    in
                    Just newMode

                Nothing ->
                    mode
        
        updatedModel =
            { model | mode = updatedMode }

    in
    (updatedModel, Cmd.none)