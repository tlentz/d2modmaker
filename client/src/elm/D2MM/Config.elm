module D2MM.Config exposing (..)


type alias Config =
    { version : String
    , sourceDir : String
    , outputDir : String
    , meleeSplash : Bool
    , increasedStackSizes : Bool
    , increaseMonsterDensity : Int
    , enableTownSkills : Bool
    , noDropZero : Bool
    , questDrops : Bool
    , uniqueItemDropRate : Float
    , runeDropRate : Float
    , startWithCube : Bool
    , cowzzz : Bool
    , removeLevelRequirements : Bool
    , removeAttRequirements : Bool
    , removeUniqCharmLimit : Bool
    , randomOptions : RandomOptions
    }


type alias RandomOptions =
    { randomize : Bool
    , seed : Int
    , isBalanced : Bool
    , balancedPropCount : Bool
    , allowDuplicateProps : Bool
    , minProps : Int
    , maxProps : Int
    , useOSkills : Bool
    , perfectProps : Bool
    }
