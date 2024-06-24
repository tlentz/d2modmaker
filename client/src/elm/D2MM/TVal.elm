module D2MM.TVal exposing (..)

import GenericDict as Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Extra exposing (fromResult)
import Json.Encode as Encode
import List as L



--------------------------------
-- TVal
--------------------------------


type TVal
    = TString String
    | TInt Int
    | TFloat Float
    | TBool Bool
    | TList (List TVal)
    | TObj (Dict TKey TVal)


encodeTVal : TVal -> Encode.Value
encodeTVal t =
    case t of
        TString str ->
            Encode.string str

        TInt int ->
            Encode.int int

        TFloat float ->
            Encode.float float

        TBool bool ->
            Encode.bool bool

        TList lst ->
            Encode.list encodeTVal lst

        TObj dict ->
            dict
                |> Dict.toList
                |> L.map (Tuple.mapBoth tKeyToString encodeTVal)
                |> Encode.object



--------------------------------
-- TKey
--------------------------------


type TKey
    = Version
    | SourceDir
    | OutputDir
    | MeleeSplash
    | IncreasedStackSizes
    | IncreaseMonsterDensity
    | EnableTownSkills
    | NoDropZero
    | QuestDrops
    | UniqueItemDropRate
    | RuneDropRate
    | StartWithCube
    | Cowzzz
    | RemoveLevelRequirements
    | RemoveAttRequirements
    | RemoveUniqCharmLimit
    | RandomOptions
    | Randomize
    | UseSeed
    | Seed
    | IsBalanced
    | BalancedPropCount
    | AllowDupProps
    | MinProps
    | MaxProps
    | UseOSkills
    | PerfectProps


encodeTKey : TKey -> Encode.Value
encodeTKey =
    tKeyToString >> Encode.string


decodeTKey : String -> Decoder TKey
decodeTKey =
    tKeyFromString >> fromResult


tKeyToString : TKey -> String
tKeyToString t =
    case t of
        Version ->
            "Version"

        SourceDir ->
            "SourceDir"

        OutputDir ->
            "OutputDir"

        MeleeSplash ->
            "MeleeSplash"

        IncreasedStackSizes ->
            "IncreasedStackSizes"

        IncreaseMonsterDensity ->
            "IncreaseMonsterDensity"

        EnableTownSkills ->
            "EnableTownSkills"

        NoDropZero ->
            "NoDropZero"

        QuestDrops ->
            "QuestDrops"

        UniqueItemDropRate ->
            "UniqueItemDropRate"

        RuneDropRate ->
            "RuneDropRate"

        StartWithCube ->
            "StartWithCube"

        Cowzzz ->
            "Cowzzz"

        RemoveLevelRequirements ->
            "RemoveLevelRequirements"

        RemoveAttRequirements ->
            "RemoveAttRequirements"

        RemoveUniqCharmLimit ->
            "RemoveUniqCharmLimit"

        RandomOptions ->
            "RandomOptions"

        Randomize ->
            "Randomize"

        Seed ->
            "Seed"

        IsBalanced ->
            "IsBalanced"

        BalancedPropCount ->
            "BalancedPropCount"

        AllowDupProps ->
            "AllowDupProps"

        MinProps ->
            "MinProps"

        MaxProps ->
            "MaxProps"

        UseOSkills ->
            "UseOSkills"

        PerfectProps ->
            "PerfectProps"

        UseSeed ->
            "UseSeed"


tKeyFromString : String -> Result String TKey
tKeyFromString s =
    case s of
        "Version" ->
            Ok Version

        "SourceDir" ->
            Ok SourceDir

        "OutputDir" ->
            Ok OutputDir

        "MeleeSplash" ->
            Ok MeleeSplash

        "IncreasedStackSizes" ->
            Ok IncreasedStackSizes

        "IncreaseMonsterDensity" ->
            Ok IncreaseMonsterDensity

        "EnableTownSkills" ->
            Ok EnableTownSkills

        "NoDropZero" ->
            Ok NoDropZero

        "QuestDrops" ->
            Ok QuestDrops

        "UniqueItemDropRate" ->
            Ok UniqueItemDropRate

        "RuneDropRate" ->
            Ok RuneDropRate

        "StartWithCube" ->
            Ok StartWithCube

        "Cowzzz" ->
            Ok Cowzzz

        "RemoveLevelRequirements" ->
            Ok RemoveLevelRequirements

        "RemoveAttRequirements" ->
            Ok RemoveAttRequirements

        "RemoveUniqCharmLimit" ->
            Ok RemoveUniqCharmLimit

        "RandomOptions" ->
            Ok RandomOptions

        "Randomize" ->
            Ok Randomize

        "Seed" ->
            Ok Seed

        "IsBalanced" ->
            Ok IsBalanced

        "BalancedPropCount" ->
            Ok BalancedPropCount

        "AllowDupProps" ->
            Ok AllowDupProps

        "MinProps" ->
            Ok MinProps

        "MaxProps" ->
            Ok MaxProps

        "UseOSkills" ->
            Ok UseOSkills

        "PerfectProps" ->
            Ok PerfectProps

        "UseSeed" ->
            Ok UseSeed

        _ ->
            Err <| "[TKey.fromStr] unknown key : " ++ s
