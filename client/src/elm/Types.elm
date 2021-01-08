module Types exposing (BasicOption(..), ColorTheme(..), Model, Msg(..), Mode(..), initAdvancedOptions, ItemGenerationMode(..), Route(..), Screen, View(..), emptyModel)

import Browser.Dom as Dom
import Http


type alias Model =
    { screen : Screen
    , view : View
    , errorMessage : Maybe String
    , mode : Maybe Mode
    }


emptyModel : Model
emptyModel =
    { screen = { width = 0, height = 0 }
    , view = ViewHome
    , errorMessage = Nothing
    , mode = Nothing
    }

type Msg
    = DoNothing
    | Resize Screen
    | FocusOn String
    | FocusResult (Result Dom.Error ())
    | SetViewportCb
    | GetResponse (Result Http.Error String)
    | CheckRandomize Bool
    | SetSelectedMode Mode
    | SetSelectedBasicOption BasicOption
    | GenerateBasic
    

type alias Screen =
    { width : Int
    , height : Int
    }


type View
    = ViewHome
    | ViewAbout


type Route
    = RouteHome
    | RouteAbout
    | RouteNotFound


type Mode
    = Basic (Maybe BasicOption)
    | Advanced AdvancedOptions


type ItemGenerationMode
    = Randomize
    | Generate


type alias AdvancedOptions =
    { randomize : Bool
    , useSeed : Bool
    , useOSkills : Bool
    , perfectProps : Bool
    , allowDupProducts : Bool
    , isBalanced : Bool
    , balancedPropCount : Bool
    , minProps : Int
    , maxProps : Int
    , meleeSplash : Bool
    , monsterDensity : Int
    , enableTownSkills : Bool
    , startWithCube : Bool
    , cowzzz : Bool
    , increaseStackSizes : Bool
    , removeLevelRequirements : Bool
    , removeAttRequirements : Bool
    , removeUniqueCharmLimit : Bool
    , noDropZero : Bool
    , questDrops : Bool
    , uniqueItemDropRate : Int
    , runeDropRate : Int
    , itemGenerationMode : ItemGenerationMode
    }


initAdvancedOptions : AdvancedOptions
initAdvancedOptions =
    { randomize = False
    , useSeed = False
    , useOSkills = False
    , perfectProps = False
    , allowDupProducts = False
    , isBalanced = False
    , balancedPropCount = False
    , minProps = 0
    , maxProps = 0
    , meleeSplash = False
    , monsterDensity = 1
    , enableTownSkills = False
    , startWithCube = False
    , cowzzz = False
    , increaseStackSizes = False
    , removeLevelRequirements = False
    , removeAttRequirements = False
    , removeUniqueCharmLimit = False
    , noDropZero = False
    , questDrops = False
    , uniqueItemDropRate = 1
    , runeDropRate = 1
    , itemGenerationMode = Randomize
    }


type BasicOption
    = MinorQolEnhancement
    | QolOnly
    | Vanilla
    | Better
    | Good
    | Great
    | Fantastic
    | Zomg

type ColorTheme
    = Light
    | Dark
