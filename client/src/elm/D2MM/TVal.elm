module D2MM.TVal exposing (..)

import Json.Encode as Encode



--------------------------------
-- TVal
--------------------------------


type TVal
    = TString String
    | TInt Int
    | TFloat Float
    | TBool Bool
    | TList (List TVal)


encodeTVal : TVal -> Encode.Value
encodeTVal t =
    case t of
        TString str ->
            Encode.string str

        TInt int ->
            Encode.int int

        TFloat float ->
            Encode.float float

        TBool bool ->
            Encode.bool bool

        TList lst ->
            Encode.list encodeTVal lst
