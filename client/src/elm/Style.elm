module Style exposing (style, backgroundTheme, borderTheme)

import Element exposing (Attribute, Color)
import Html.Attributes


style : String -> String -> Attribute msg
style k v =
    Html.Attributes.style k v
        |> Element.htmlAttribute


backgroundTheme : Color
backgroundTheme =
    Element.rgb255 247 247 248


borderTheme : Color
borderTheme =
    Element.rgb255 74 74 74