module Update exposing (update)

import Browser.Dom as Dom
import Helper exposing (return)
import Http exposing (Error(..))
import Task
import Types
    exposing
        ( AdvancedOptions
        , CheckboxMsg(..)
        , ColorTheme(..)
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

        GenerateBasic ->
            ( model, Cmd.none )


updateCheckboxState : CheckboxMsg -> AdvancedOptions
updateCheckboxState checkboxMsg =
    case checkboxMsg of
        SetRandomize advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | randomize = isChecked }

            in
            updatedMode
        
        SetUseSeed advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | useSeed = isChecked }

            in
            updatedMode

        SetUseOSkills advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | useOSkills = isChecked }

            in
            updatedMode

        SetPerfectProps advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | perfectProps = isChecked }

            in
            updatedMode

        SetAllowDupProps advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | allowDupProps = isChecked }

            in
            updatedMode

        SetIsBalanced advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | isBalanced = isChecked }

            in
            updatedMode

        SetBalancedPropCount advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | balancedPropCount = isChecked }

            in
            updatedMode

        SetMeleeSplash advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | meleeSplash = isChecked }

            in
            updatedMode

        SetEnableTownSkills advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | enableTownSkills = isChecked }

            in
            updatedMode

        SetStartWithCube advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | startWithCube = isChecked }

            in
            updatedMode

        SetCowzzz advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | cowzzz = isChecked }

            in
            updatedMode

        SetIncreaseStackSizes advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | increaseStackSizes = isChecked }

            in
            updatedMode

        SetRemoveLevelRequirements advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | removeLevelRequirements = isChecked }

            in
            updatedMode

        SetRemoveAttRequirements advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | removeAttRequirements = isChecked }

            in
            updatedMode

        SetRemoveUniqueCharmLimit advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | removeUniqueCharmLimit = isChecked }

            in
            updatedMode

        SetNoDropZero advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | noDropZero = isChecked }

            in
            updatedMode

        SetQuestDrops advancedOptions isChecked ->
            let
                updatedMode = 
                    { advancedOptions | questDrops = isChecked }

            in
            updatedMode

setUpdatedOptionsOnModel : Model -> AdvancedOptions -> (Model, Cmd Msg)
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
                                    Advanced advancedOptions
                    in
                    Just newMode

                Nothing ->
                    mode
        
        updatedModel =
            { model | mode = updatedMode }

    in
    (updatedModel, Cmd.none)