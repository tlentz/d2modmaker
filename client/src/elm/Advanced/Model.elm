module Advanced.Model exposing (..)

import Advanced.FormField exposing (FormFields, formFields)


type alias Model =
    { formFields : FormFields
    }


initModel : Model
initModel =
    { formFields = formFields
    }
