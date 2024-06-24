module Advanced.FormField exposing (..)

import Advanced.TInput exposing (NumberInput, TInput(..))
import Advanced.Tooltip as Tooltip
import D2MM.TVal exposing (TKey(..), TVal(..), tKeyToString)
import GenericDict as Dict exposing (Dict)
import List as L


type alias FormFields =
    Dict TKey FormField


type alias FormField =
    { tooltip : String
    , value : TVal
    , input : TInput
    }


formFields : FormFields
formFields =
    [ ( SourceDir, TString "", TTextInput "" )
    , ( OutputDir, TString "", TTextInput "" )
    , ( MeleeSplash, TBool False, TCheckbox False )
    , ( IncreasedStackSizes, TBool False, TCheckbox False )
    , ( IncreaseMonsterDensity, TBool False, TCheckbox False )
    , ( EnableTownSkills, TBool False, TCheckbox False )
    , ( NoDropZero, TBool False, TCheckbox False )
    , ( QuestDrops, TBool False, TCheckbox False )
    , ( UniqueItemDropRate, TFloat 1, TNumberInput <| NumberInput 1 (Just 1) (Just 100) )
    , ( RuneDropRate, TFloat 1, TNumberInput <| NumberInput 1 (Just 1) (Just 100) )
    , ( StartWithCube, TBool False, TCheckbox False )
    , ( Cowzzz, TBool False, TCheckbox False )
    , ( RemoveLevelRequirements, TBool False, TCheckbox False )
    , ( RemoveAttRequirements, TBool False, TCheckbox False )
    , ( RemoveUniqCharmLimit, TBool False, TCheckbox False )
    , -- RandomOptions
      ( Randomize, TBool False, TCheckbox False )
    , ( UseSeed, TBool False, TCheckbox False )
    , ( Seed, TInt -1, TNumberInput <| NumberInput -1 Nothing Nothing )
    , ( IsBalanced, TBool False, TCheckbox False )
    , ( BalancedPropCount, TBool False, TCheckbox False )
    , ( MinProps, TInt 0, TNumberInput <| NumberInput 0 (Just 0) (Just 20) )
    , ( MaxProps, TInt 20, TNumberInput <| NumberInput 20 (Just 0) (Just 20) )
    , ( UseOSkills, TBool False, TCheckbox False )
    , ( PerfectProps, TBool False, TCheckbox False )
    ]
        |> L.map toFormField
        |> Dict.fromList tKeyToString


toFormField : ( TKey, TVal, TInput ) -> ( TKey, FormField )
toFormField ( key, val, input ) =
    ( key
    , { tooltip = Tooltip.tooltip key
      , value = val
      , input = input
      }
    )
