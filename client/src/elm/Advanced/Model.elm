module Advanced.Model exposing (..)

import Common.Checkbox exposing (Checkbox)
import Common.NumberInput exposing (NumberInput)
import Dict exposing (Dict)


type alias Model =
    { checkboxes : Dict String Checkbox
    , numberInputs : Dict String NumberInput
    , seed : Maybe Int
    }


initModel : Model
initModel =
    { checkboxes = initCheckboxes
    , numberInputs = initNumberInputs
    , seed = Nothing
    }


initCheckboxes : Dict String Checkbox
initCheckboxes =
    Dict.fromList
        [ ( "Randomize", initRandomize )
        , ( "UseSeed", initUseSeed )
        , ( "UseOSkills", initUseOSkills )
        , ( "PerfectProps", initPerfectProps )
        , ( "AllowDupProps", initAllowDupProps )
        , ( "IsBalanced", initIsBalanced )
        , ( "BalancedPropCount", initBalancedPropCount )
        , ( "MeleeSplash", initMeleeSplash )
        , ( "EnableTownSkills", initEnableTownSkills )
        , ( "StartWithCube", initStartWithCube )
        , ( "Cowzzz", initCowzzz )
        , ( "IncreaseStackSizes", initIncreaseStackSizes )
        , ( "RemoveLevelRequirements", initRemoveLevelRequirements )
        , ( "RemoveAttRequirements", initRemoveAttRequirements )
        , ( "RemoveUniqueCharmLimit", initRemoveUniqueCharmLimit )
        , ( "NoDropZero", initNoDropZero )
        , ( "QuestDrops", initQuestDrops )
        ]


initNumberInputs : Dict.Dict String NumberInput
initNumberInputs =
    Dict.fromList
        [ ( "MinProps", initMinProps )
        , ( "MaxProps", initMaxProps )
        , ( "MonsterDensity", initMonsterDensity )
        , ( "UniqueItemDropRate", initUniqueItemDropRate )
        , ( "RuneDropRate", initRuneDropRate )
        ]


initRandomize : Checkbox
initRandomize =
    { isChecked = True
    , tooltip = "Randomize all all uniques, sets, and runewords."
    }


initUseSeed : Checkbox
initUseSeed =
    { isChecked = False
    , tooltip = "Provide a specific seed to use.  Toggling on/off will generate a new seed."
    }


initUseOSkills : Checkbox
initUseOSkills =
    { isChecked = True
    , tooltip = "Change class only skill props to spawn as oskills."
    }


initPerfectProps : Checkbox
initPerfectProps =
    { isChecked = False
    , tooltip = "All props will have a perfect max value when spawning on an item."
    }


initAllowDupProps : Checkbox
initAllowDupProps =
    { isChecked = False
    , tooltip = "If turned off, prevents the same prop from being placed on an item more than once. e.g. two instances of all resist will not get stacked on the same randomized item."
    }


initIsBalanced : Checkbox
initIsBalanced =
    { isChecked = True
    , tooltip = "Allows props only from items within 10 levels of the base item so that you don't get crazy hell stats on normal items, but still get a wide range of randomization."
    }


initBalancedPropCount : Checkbox
initBalancedPropCount =
    { isChecked = True
    , tooltip = "Pick prop count on items based on counts from vanilla items. Picks from items up to 10 levels higher when randomizing."
    }


initMeleeSplash : Checkbox
initMeleeSplash =
    { isChecked = True
    , tooltip = "Enables Splash Damage.  Can spawn as an affix on magic and rare jewels."
    }


initEnableTownSkills : Checkbox
initEnableTownSkills =
    { isChecked = True
    , tooltip = "Enable the ability to use all skills in town."
    }


initStartWithCube : Checkbox
initStartWithCube =
    { isChecked = True
    , tooltip = "Newly created characters will start with a cube."
    }


initCowzzz : Checkbox
initCowzzz =
    { isChecked = True
    , tooltip = "Enables the ability to recreate a cow portal after killing the cow king.  Adds cube recipe to cube a single tp scroll to create the cow portal4."
    }


initIncreaseStackSizes : Checkbox
initIncreaseStackSizes =
    { isChecked = True
    , tooltip = "Increases tome sizes to 100.  Increases arrows/bolts stack sizes to 511.  Increases key stack sizes to 100."
    }


initRemoveLevelRequirements : Checkbox
initRemoveLevelRequirements =
    { isChecked = False
    , tooltip = "Removes level requirements from items."
    }


initRemoveAttRequirements : Checkbox
initRemoveAttRequirements =
    { isChecked = False
    , tooltip = "Removes stat requirements from items."
    }


initRemoveUniqueCharmLimit : Checkbox
initRemoveUniqueCharmLimit =
    { isChecked = False
    , tooltip = "Removes unique charm limit in inventory."
    }


initNoDropZero : Checkbox
initNoDropZero =
    { isChecked = True
    , tooltip = "Guarantees that a monster drops something upon death."
    }


initQuestDrops : Checkbox
initQuestDrops =
    { isChecked = True
    , tooltip = "Act bosses will always drop quest drops."
    }


initMinProps : NumberInput
initMinProps =
    { value = 0
    , min = 0
    , max = 20
    , tooltip = "Minimum number of props an item can have."
    }


initMaxProps : NumberInput
initMaxProps =
    { value = 20
    , min = 0
    , max = 20
    , tooltip = "Maximum number of props an item can have."
    }


initMonsterDensity : NumberInput
initMonsterDensity =
    { value = 1
    , min = 1
    , max = 30
    , tooltip = "Increases monster density throughout the map by the given factor."
    }


initUniqueItemDropRate : NumberInput
initUniqueItemDropRate =
    { value = 1
    , min = 1
    , max = 100
    , tooltip = "Increases the drop rate of unique and set items.  When using this setting, high values prevent some monsters from dropping set items."
    }


initRuneDropRate : NumberInput
initRuneDropRate =
    { value = 1
    , min = 1
    , max = 100
    , tooltip = "Increases rune drop rates. Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively. E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100."
    }
