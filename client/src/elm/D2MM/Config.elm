module D2MM.Config exposing (..)

import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Extra exposing (andMap)
import Json.Encode as Encode


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


blankConfig : Config
blankConfig =
    { version = ""
    , sourceDir = ""
    , outputDir = ""
    , meleeSplash = False
    , increasedStackSizes = False
    , increaseMonsterDensity = 1
    , enableTownSkills = False
    , noDropZero = False
    , questDrops = False
    , uniqueItemDropRate = 1.0
    , runeDropRate = 1.0
    , startWithCube = False
    , cowzzz = False
    , removeLevelRequirements = False
    , removeAttRequirements = False
    , removeUniqCharmLimit = False
    , randomOptions = blankRandomOptions
    }


encodeConfig : Config -> Encode.Value
encodeConfig t =
    [ ( "Version", Encode.string t.version )
    , ( "SourceDir", Encode.string t.sourceDir )
    , ( "OutputDir", Encode.string t.outputDir )
    , ( "MeleeSplash", Encode.bool t.meleeSplash )
    , ( "IncreasedStackSizes", Encode.bool t.increasedStackSizes )
    , ( "IncreaseMonsterDensity", Encode.int t.increaseMonsterDensity )
    , ( "EnableTownSkills", Encode.bool t.enableTownSkills )
    , ( "NoDropZero", Encode.bool t.noDropZero )
    , ( "QuestDrops", Encode.bool t.questDrops )
    , ( "UniqueItemDropRate", Encode.float t.uniqueItemDropRate )
    , ( "RuneDropRate", Encode.float t.runeDropRate )
    , ( "StartWithCube", Encode.bool t.startWithCube )
    , ( "Cowzzz", Encode.bool t.cowzzz )
    , ( "RemoveLevelRequirements", Encode.bool t.removeLevelRequirements )
    , ( "RemoveAttRequirements", Encode.bool t.removeAttRequirements )
    , ( "RemoveUniqCharmLimit", Encode.bool t.removeUniqCharmLimit )
    , ( "RandomOptions", encodeRandomOptions t.randomOptions )
    ]
        |> Encode.object


decodeConfig : Decoder Config
decodeConfig =
    Decode.succeed Config
        |> andMap (Decode.field "Version" Decode.string)
        |> andMap (Decode.field "SourceDir" Decode.string)
        |> andMap (Decode.field "OutputDir" Decode.string)
        |> andMap (Decode.field "MeleeSplash" Decode.bool)
        |> andMap (Decode.field "IncreasedStackSizes" Decode.bool)
        |> andMap (Decode.field "IncreaseMonsterDensity" Decode.int)
        |> andMap (Decode.field "EnableTownSkills" Decode.bool)
        |> andMap (Decode.field "NoDropZero" Decode.bool)
        |> andMap (Decode.field "QuestDrops" Decode.bool)
        |> andMap (Decode.field "UniqueItemDropRate" Decode.float)
        |> andMap (Decode.field "RuneDropRate" Decode.float)
        |> andMap (Decode.field "StartWithCube" Decode.bool)
        |> andMap (Decode.field "Cowzzz" Decode.bool)
        |> andMap (Decode.field "RemoveLevelRequirements" Decode.bool)
        |> andMap (Decode.field "RemoveAttRequirements" Decode.bool)
        |> andMap (Decode.field "RemoveUniqCharmLimit" Decode.bool)
        |> andMap (Decode.field "RandomOptions" decodeRandomOptions)


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


blankRandomOptions : RandomOptions
blankRandomOptions =
    { randomize = False
    , seed = 0
    , isBalanced = False
    , balancedPropCount = False
    , allowDuplicateProps = False
    , minProps = -1
    , maxProps = -1
    , useOSkills = False
    , perfectProps = False
    }


encodeRandomOptions : RandomOptions -> Encode.Value
encodeRandomOptions t =
    [ ( "Randomize", Encode.bool t.randomize )
    , ( "Seed", Encode.int t.seed )
    , ( "IsBalanced", Encode.bool t.isBalanced )
    , ( "BalancedPropCount", Encode.bool t.balancedPropCount )
    , ( "AllowDuplicateProps", Encode.bool t.allowDuplicateProps )
    , ( "MinProps", Encode.int t.minProps )
    , ( "MaxProps", Encode.int t.maxProps )
    , ( "UseOSkills", Encode.bool t.useOSkills )
    , ( "PerfectProps", Encode.bool t.perfectProps )
    ]
        |> Encode.object


decodeRandomOptions : Decoder RandomOptions
decodeRandomOptions =
    Decode.succeed RandomOptions
        |> andMap (Decode.field "Randomize" Decode.bool)
        |> andMap (Decode.field "Seed" Decode.int)
        |> andMap (Decode.field "IsBalanced" Decode.bool)
        |> andMap (Decode.field "BalancedPropCount" Decode.bool)
        |> andMap (Decode.field "AllowDuplicateProps" Decode.bool)
        |> andMap (Decode.field "MinProps" Decode.int)
        |> andMap (Decode.field "MaxProps" Decode.int)
        |> andMap (Decode.field "UseOSkills" Decode.bool)
        |> andMap (Decode.field "PerfectProps" Decode.bool)
