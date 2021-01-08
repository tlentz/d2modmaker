module Helper exposing (colorThemeToString, ifThenElse, isMobile, mkCmd, mouseDownPreventDefault, onRightClick, return, style, when, whenAttr, whenJust)

import Element exposing (Attribute, Element)
import Html
import Html.Attributes as Attrs exposing (style)
import Html.Events
import Json.Decode as Json
import Maybe.Extra
import Task
import Types
    exposing
        ( ColorTheme(..)
        , Msg(..)
        , Screen
        )


return : a -> List (Cmd b) -> ( a, Cmd b )
return model cmds =
    ( model, Cmd.batch cmds )


mkCmd : msg -> Cmd msg
mkCmd msg =
    Task.perform (always msg) (Task.succeed msg)


onRightClick : a -> Html.Attribute a
onRightClick message =
    Json.succeed
        { message = message
        , stopPropagation = False
        , preventDefault = True
        } 
        |> Html.Events.custom "contextmenu"


whenJust : (a -> Element msg) -> Maybe a -> Element msg
whenJust =
    Maybe.Extra.unwrap Element.none


whenAttr : Bool -> Attribute msg -> Attribute msg
whenAttr bool =
    if bool then
        identity

    else
        Attrs.classList []
            |> Element.htmlAttribute
            |> always


when : Bool -> Element msg -> Element msg
when b =
    if b then
        identity

    else
        always Element.none


ifThenElse : Bool -> a -> a -> a
ifThenElse bool a b =
    if bool then
        a

    else
        b


isMobile : Screen -> Bool
isMobile { width } =
    width < 450


colorThemeToString : ColorTheme -> String
colorThemeToString colorTheme =
    case colorTheme of
        Light ->
            "Light"

        Dark ->
            "Dark"


mouseDownPreventDefault : Msg -> Html.Attribute Msg
mouseDownPreventDefault msg =
    Html.Events.custom "click"
        (Json.succeed
            { message = msg
            , stopPropagation = True
            , preventDefault = True
            }
        )


style : String -> String -> Attribute msg
style k v =
    Attrs.style k v
        |> Element.htmlAttribute
