module Model exposing (..)

import Advanced.Model as AdvancedM
import Basic.Model as BasicM


type alias Model =
    { serverMessage : String
    , page : Page
    }


initModel : Model
initModel =
    { serverMessage = ""
    , page = PageBasic BasicM.initModel
    }


type Page
    = PageBasic BasicM.Model
    | PageAdvanced AdvancedM.Model
