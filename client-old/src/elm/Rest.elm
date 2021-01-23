module Rest exposing (postApiUpdateCanRank)

import Http
import Types exposing (Msg(..))


postApiUpdateCanRank : Cmd Msg
postApiUpdateCanRank =
    Http.request
        { method =
            "POST"
        , headers =
            []
        , url =
            String.join "/"
                [ "localhost:3001"
                , "hello"
                ]
        , body =
            Http.emptyBody
        , expect =
            Http.expectString GetResponse
        , timeout =
            Nothing
        , tracker = Nothing
        }
