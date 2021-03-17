module Basic.Model exposing (..)


type alias Model =
    { preset : Preset }


initModel : Model
initModel =
    { preset = MinorQolEnhancement }


type Preset
    = MinorQolEnhancement
    | QolOnly
    | Vanilla
    | Better
    | Good
    | Great
    | Fantastic
    | Zomg
