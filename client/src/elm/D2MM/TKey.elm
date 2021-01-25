module D2MM.TKey exposing (..)

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
    | Seed
    | IsBalanced
    | BalancedPropCount
    | AllowDuplicateProps
    | MinProps
    | MaxProps
    | UseOSkills
    | PerfectProps


toStr : TKey -> String
toStr t =
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

        AllowDuplicateProps ->
            "AllowDuplicateProps"

        MinProps ->
            "MinProps"

        MaxProps ->
            "MaxProps"

        UseOSkills ->
            "UseOSkills"

        PerfectProps ->
            "PerfectProps"


fromStr : String -> Result String TKey
fromStr s =
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

        "AllowDuplicateProps" ->
            Ok AllowDuplicateProps

        "MinProps" ->
            Ok MinProps

        "MaxProps" ->
            Ok MaxProps

        "UseOSkills" ->
            Ok UseOSkills

        "PerfectProps" ->
            Ok PerfectProps

        _ ->
            Err <| "[TKey.fromStr] unknown key : " ++ s
