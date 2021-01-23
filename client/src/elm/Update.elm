module Update exposing (..)

import Http
import Json.Decode as Decode
import Model exposing (Model)
import Msg exposing (Msg(..))
import Ports
import Util exposing (httpErrorToString)


update : Msg -> Model -> ( Model, Cmd Msg )
update message model =
    case message of
        TestServer ->
            let
                expect =
                    Http.expectJson OnServerResponse (Decode.field "result" Decode.string)
            in
            ( model
            , Http.get { url = "/test", expect = expect }
            )

        OnServerResponse res ->
            case res of
                Ok r ->
                    ( { model | serverMessage = r }, Cmd.none )

                Err err ->
                    ( { model | serverMessage = "Error: " ++ httpErrorToString err }, Cmd.none )

        ChangePage page ->
            ( { model | page = page }, Cmd.none )
