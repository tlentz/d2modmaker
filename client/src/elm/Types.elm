module Types exposing (AdvancedIntMsg(..), AdvancedCheckboxOption, AdvancedCheckboxOptions, AdvancedNumberOption, BasicOption(..), initMinProps, CheckboxMsg(..), CheckboxName, Model, Msg(..), Mode(..), initAdvancedCheckboxOptions, ItemGenerationMode(..), Route(..), Screen, View(..), emptyModel)

import Browser.Dom as Dom
import Dict
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
    | SetCheckedState CheckboxMsg
    | SetAdvancedInt AdvancedIntMsg
    | SetSelectedMode Mode
    | SetSelectedBasicOption BasicOption
    | SetItemGenerationMode ItemGenerationMode
    | GenerateBasic
    | SaveConfig

type CheckboxMsg
    = ToggleCheckbox AdvancedCheckboxOptions CheckboxName
    | SetRandomize AdvancedCheckboxOptions
    | SetUseSeed AdvancedCheckboxOptions
    | SetSeed AdvancedCheckboxOptions Int
    | SetUseOSkills AdvancedCheckboxOptions
    | SetPerfectProps AdvancedCheckboxOptions
    | SetAllowDupProps AdvancedCheckboxOptions
    | SetIsBalanced AdvancedCheckboxOptions
    | SetBalancedPropCount AdvancedCheckboxOptions
    | SetMeleeSplash AdvancedCheckboxOptions
    | SetEnableTownSkills AdvancedCheckboxOptions
    | SetStartWithCube AdvancedCheckboxOptions
    | SetCowzzz AdvancedCheckboxOptions
    | SetIncreaseStackSizes AdvancedCheckboxOptions
    | SetRemoveLevelRequirements AdvancedCheckboxOptions
    | SetRemoveAttRequirements AdvancedCheckboxOptions
    | SetRemoveUniqueCharmLimit AdvancedCheckboxOptions
    | SetNoDropZero AdvancedCheckboxOptions
    | SetQuestDrops AdvancedCheckboxOptions

type AdvancedIntMsg
    = SetMinProps AdvancedCheckboxOptions Int
    | SetMaxProps AdvancedCheckboxOptions Int
    | SetMonsterDensity AdvancedCheckboxOptions Int
    | SetUniqueItemDropRate AdvancedCheckboxOptions Int
    | SetRuneDropRate AdvancedCheckboxOptions Int

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
    | Advanced AdvancedCheckboxOptions


type BasicOption
    = MinorQolEnhancement
    | QolOnly
    | Vanilla
    | Better
    | Good
    | Great
    | Fantastic
    | Zomg


type alias CheckboxName =
    String


type alias AdvancedCheckboxOptions =
    { test : Dict.Dict CheckboxName AdvancedCheckboxOption
    , randomize : AdvancedCheckboxOption
    , useSeed : AdvancedCheckboxOption
    , seed : Int
    , useOSkills : AdvancedCheckboxOption
    , perfectProps : AdvancedCheckboxOption
    , allowDupProps : AdvancedCheckboxOption
    , isBalanced : AdvancedCheckboxOption
    , balancedPropCount : AdvancedCheckboxOption
    , minProps : AdvancedNumberOption
    , maxProps : AdvancedNumberOption
    , meleeSplash : AdvancedCheckboxOption
    , monsterDensity : AdvancedNumberOption
    , enableTownSkills : AdvancedCheckboxOption
    , startWithCube : AdvancedCheckboxOption
    , cowzzz : AdvancedCheckboxOption
    , increaseStackSizes : AdvancedCheckboxOption
    , removeLevelRequirements : AdvancedCheckboxOption
    , removeAttRequirements : AdvancedCheckboxOption
    , removeUniqueCharmLimit : AdvancedCheckboxOption
    , noDropZero : AdvancedCheckboxOption
    , questDrops : AdvancedCheckboxOption
    , uniqueItemDropRate : AdvancedNumberOption
    , runeDropRate : AdvancedNumberOption
    , itemGenerationMode : ItemGenerationMode
    }
    

type ItemGenerationMode
    = None
    | Randomize
    | Generate


type alias AdvancedCheckboxOption =
    { isChecked : Bool
    , tooltip : String
    }


type alias AdvancedNumberOption =
    { value : Float
    , tooltip : String
    }

initAdvancedCheckboxOptions : AdvancedCheckboxOptions
initAdvancedCheckboxOptions =
    { test = Dict.fromList
        [ ( "Test", initRandomize ) ]
    , randomize = initRandomize
    , useSeed = initUseSeed
    , seed = 1
    , useOSkills = initUseOSkills
    , perfectProps = initPerfectProps
    , allowDupProps = initAllowDupProps
    , isBalanced = initISBalanced
    , balancedPropCount = initBalancedPropCount
    , minProps = initMinProps
    , maxProps = initMaxProps
    , meleeSplash = initMeleeSplash
    , monsterDensity = initMonsterDensity
    , enableTownSkills = initEnableTownSkills
    , startWithCube = initStartWithCube
    , cowzzz = initCowzzz
    , increaseStackSizes = initIncreaseStackSizes
    , removeLevelRequirements = initRemoveLevelRequirements
    , removeAttRequirements = initRemoveAttRequirements
    , removeUniqueCharmLimit = initRemoveUniqueCharmLimit
    , noDropZero = initNoDropZero
    , questDrops = initQuestDrops
    , uniqueItemDropRate = initUniqueItemDropRate
    , runeDropRate = initRuneDropRate
    , itemGenerationMode = None
    }

initRandomize: AdvancedCheckboxOption
initRandomize =
    { isChecked = False
    , tooltip = "Randomize all all uniques, sets, and runewords."
    }

initUseSeed: AdvancedCheckboxOption
initUseSeed =
    { isChecked = False
    , tooltip = "Provide a specific seed to use.  Toggling on/off will generate a new seed."
    }

initUseOSkills: AdvancedCheckboxOption
initUseOSkills =
    { isChecked = False
    , tooltip = "Change class only skill props to spawn as oskills."
    }

initPerfectProps: AdvancedCheckboxOption
initPerfectProps =
    { isChecked = False
    , tooltip = "All props will have a perfect max value when spawning on an item."
    }

initAllowDupProps: AdvancedCheckboxOption
initAllowDupProps =
    { isChecked = False
    , tooltip = "If turned off, prevents the same prop from being placed on an item more than once. e.g. two instances of all resist will not get stacked on the same randomized item."
    }

initISBalanced: AdvancedCheckboxOption
initISBalanced =
    { isChecked = False
    , tooltip = "Allows props only from items within 10 levels of the base item so that you don't get crazy hell stats on normal items, but still get a wide range of randomization."
    }

initBalancedPropCount: AdvancedCheckboxOption
initBalancedPropCount =
    { isChecked = False
    , tooltip = "Pick prop count on items based on counts from vanilla items. Picks from items up to 10 levels higher when randomizing."
    }

initMeleeSplash: AdvancedCheckboxOption
initMeleeSplash =
    { isChecked = False
    , tooltip = "Enables Splash Damage.  Can spawn as an affix on magic and rare jewels."
    }

initEnableTownSkills: AdvancedCheckboxOption
initEnableTownSkills =
    { isChecked = False
    , tooltip = "Enable the ability to use all skills in town."
    }

initStartWithCube: AdvancedCheckboxOption
initStartWithCube =
    { isChecked = False
    , tooltip = "Newly created characters will start with a cube."
    }

initCowzzz: AdvancedCheckboxOption
initCowzzz =
    { isChecked = False
    , tooltip = "Enables the ability to recreate a cow portal after killing the cow king.  Adds cube recipe to cube a single tp scroll to create the cow portal4."
    }

initIncreaseStackSizes: AdvancedCheckboxOption
initIncreaseStackSizes =
    { isChecked = False
    , tooltip = "Increases tome sizes to 100.  Increases arrows/bolts stack sizes to 511.  Increases key stack sizes to 100."
    }

initRemoveLevelRequirements: AdvancedCheckboxOption
initRemoveLevelRequirements =
    { isChecked = False
    , tooltip = "Removes level requirements from items."
    }

initRemoveAttRequirements: AdvancedCheckboxOption
initRemoveAttRequirements =
    { isChecked = False
    , tooltip = "Removes stat requirements from items."
    }

initRemoveUniqueCharmLimit: AdvancedCheckboxOption
initRemoveUniqueCharmLimit =
    { isChecked = False
    , tooltip = "Removes unique charm limit in inventory."
    }

initNoDropZero: AdvancedCheckboxOption
initNoDropZero =
    { isChecked = False
    , tooltip = "Guarantees that a monster drops something upon death."
    }

initQuestDrops: AdvancedCheckboxOption
initQuestDrops =
    { isChecked = False
    , tooltip = "Act bosses will always drop quest drops."
    }


initMinProps: AdvancedNumberOption
initMinProps =
    { value = 0
    , tooltip = "Minimum number of props an item can have."
    }


initMaxProps: AdvancedNumberOption
initMaxProps =
    { value = 20
    , tooltip = "Maximum number of props an item can have."
    }

initMonsterDensity: AdvancedNumberOption
initMonsterDensity =
    { value = 1
    , tooltip = "Increases monster density throughout the map by the given factor."
    }

initUniqueItemDropRate: AdvancedNumberOption
initUniqueItemDropRate =
    { value = 1
    , tooltip = "Increases the drop rate of unique and set items.  When using this setting, high values prevent some monsters from dropping set items."
    }

initRuneDropRate: AdvancedNumberOption
initRuneDropRate =
    { value = 1
    , tooltip = "Increases rune drop rates. Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively. E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100."
    }
