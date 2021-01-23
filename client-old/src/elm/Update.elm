module Update exposing (update)

import Browser.Dom as Dom
import Dict
import Helper exposing (return)
import Http exposing (Error(..))
import Random
import Task
import Types
    exposing
        ( AdvancedCheckboxOption
        , AdvancedCheckboxOptions
        , AdvancedIntMsg(..)
        , AdvancedNumberOption
        , CheckboxMsg(..)
        , InputName
        , Mode(..)
        , Model
        , Msg(..)
        )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        SetCheckedState checkboxMsg ->
            updateCheckboxState checkboxMsg
                |> setUpdatedOptionsOnModel model

        SetAdvancedInt advancedIntMsg ->
            updateAdvancedInt advancedIntMsg
                |> setUpdatedIntOptionsOnModel model

        SetSelectedMode mode ->
            ( { model | mode = Just mode }, Cmd.none )

        SetSelectedBasicOption basicOption ->
            let
                mode =
                    model.mode

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
                mode =
                    model.mode

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


updateCheckboxState : CheckboxMsg -> ( AdvancedCheckboxOptions, Cmd CheckboxMsg )
updateCheckboxState checkboxMsg =
    case checkboxMsg of
        ToggleCheckbox advancedOptions checkboxName ->
            let
                newAdvancedOptions =
                    { advancedOptions | checkboxes = toggle checkboxName advancedOptions.checkboxes }

                cmd =
                    if checkboxName == "UseSeed" then
                        Random.generate (SetSeed newAdvancedOptions) (Random.int 1 Random.maxInt)

                    else
                        Cmd.none
            in
            ( newAdvancedOptions, cmd )

        SetSeed advancedOptions newSeed ->
            ( { advancedOptions | seed = newSeed }, Cmd.none )


updateAdvancedInt : AdvancedIntMsg -> AdvancedCheckboxOptions
updateAdvancedInt advancedIntMsg =
    case advancedIntMsg of
        SetInputValue advancedOptions inputName value ->
            { advancedOptions | numberInputs = setInputValue value inputName advancedOptions.numberInputs }


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


setInputValue : Float -> comparable -> Dict.Dict comparable AdvancedNumberOption -> Dict.Dict comparable AdvancedNumberOption
setInputValue newValue key dict =
    Dict.update key
        (\oldOption ->
            case oldOption of
                Just o ->
                    Just <| { o | value = max o.min (min o.max newValue) }

                Nothing ->
                    Nothing
        )
        dict


setUpdatedOptionsOnModel : Model -> ( AdvancedCheckboxOptions, Cmd CheckboxMsg ) -> ( Model, Cmd Msg )
setUpdatedOptionsOnModel model advancedOptions =
    let
        mode =
            model.mode

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
    ( updatedModel, Cmd.map SetCheckedState (Tuple.second advancedOptions) )


setUpdatedIntOptionsOnModel : Model -> AdvancedCheckboxOptions -> ( Model, Cmd Msg )
setUpdatedIntOptionsOnModel model advancedOptions =
    let
        mode =
            model.mode

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
    ( updatedModel, Cmd.none )
