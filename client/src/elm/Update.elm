module Update exposing (update)

import Browser.Dom as Dom
import Helper exposing (return)
import Http exposing (Error(..))
import Task
import Types
    exposing
        ( ColorTheme(..)
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

        CheckRandomize isChecked ->
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
                                                newOptions =
                                                    { options | randomize = isChecked }
                                            in
                                            Advanced newOptions
                            in
                            Just newMode

                        Nothing ->
                            mode
                
                updatedModel =
                    { model | mode = updatedMode }

            in
            ( updatedModel, Cmd.none )
        
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