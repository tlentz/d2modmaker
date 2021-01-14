module Update exposing (update)

import Browser.Dom as Dom
import Helper exposing (return, mkCmd)
import Http exposing (Error(..))
import Random
import Task
import Types
    exposing
        ( AdvancedIntMsg(..)
        , AdvancedCheckboxOptions
        , CheckboxMsg(..)
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


updateCheckboxState : CheckboxMsg -> (AdvancedCheckboxOptions, Cmd CheckboxMsg)
updateCheckboxState checkboxMsg =
    case checkboxMsg of
        SetRandomize advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.randomize

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | randomize = updatedOption }, Cmd.none)
        
        SetUseSeed advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.useSeed

                updatedOption = 
                    { advancedOption | isChecked = isChecked }
                
                newAdvancedOptions =
                    { advancedOptions | useSeed = updatedOption }
            in
            (newAdvancedOptions, Random.generate (SetSeed newAdvancedOptions) (Random.int 1 Random.maxInt))

        SetSeed advancedOptions newSeed ->
            ({ advancedOptions | seed = newSeed }, Cmd.none)

        SetUseOSkills advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.useOSkills

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | useOSkills = updatedOption }, Cmd.none)

        SetPerfectProps advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.perfectProps

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | perfectProps = updatedOption }, Cmd.none)

        SetAllowDupProps advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.allowDupProps

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | allowDupProps = updatedOption }, Cmd.none)

        SetIsBalanced advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.isBalanced

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | isBalanced = updatedOption }, Cmd.none)

        SetBalancedPropCount advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.balancedPropCount

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | balancedPropCount = updatedOption }, Cmd.none)

        SetMeleeSplash advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.meleeSplash

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | meleeSplash = updatedOption }, Cmd.none)

        SetEnableTownSkills advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.enableTownSkills

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | enableTownSkills = updatedOption }, Cmd.none)

        SetStartWithCube advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.startWithCube

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | startWithCube = updatedOption }, Cmd.none)

        SetCowzzz advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.cowzzz

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | cowzzz = updatedOption }, Cmd.none)

        SetIncreaseStackSizes advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.increaseStackSizes

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | increaseStackSizes = updatedOption }, Cmd.none)

        SetRemoveLevelRequirements advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.removeLevelRequirements

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | removeLevelRequirements = updatedOption }, Cmd.none)

        SetRemoveAttRequirements advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.removeAttRequirements

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | removeAttRequirements = updatedOption }, Cmd.none)

        SetRemoveUniqueCharmLimit advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.removeUniqueCharmLimit

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | removeUniqueCharmLimit = updatedOption }, Cmd.none)

        SetNoDropZero advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.noDropZero

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

            in
            ({ advancedOptions | noDropZero = updatedOption }, Cmd.none)

        SetQuestDrops advancedOptions isChecked ->
            let
                advancedOption =
                    advancedOptions.questDrops

                updatedOption = 
                    { advancedOption | isChecked = isChecked }

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