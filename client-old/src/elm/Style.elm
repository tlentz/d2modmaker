module Style exposing (style, borderTheme)

import Element exposing (Attribute, Color)
import Html.Attributes


style : String -> String -> Attribute msg
style k v =
    Html.Attributes.style k v
        |> Element.htmlAttribute


borderTheme : Color
borderTheme =
    Element.rgb255 74 74 74