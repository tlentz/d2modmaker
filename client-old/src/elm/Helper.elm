module Helper exposing (..)


return : a -> List (Cmd b) -> ( a, Cmd b )
return model cmds =
    ( model, Cmd.batch cmds )
