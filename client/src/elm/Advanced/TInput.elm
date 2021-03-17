module Advanced.TInput exposing (..)


type TInput
    = TNumberInput NumberInput
    | TCheckbox Bool -- isChecked
    | TTextInput String


type alias NumberInput =
    { value : Float
    , min : Maybe Float
    , max : Maybe Float
    }
