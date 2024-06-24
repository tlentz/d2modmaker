module Advanced.Tooltip exposing (..)

import D2MM.TVal exposing (TKey(..))


tooltip : TKey -> String
tooltip t =
    case t of
        Version ->
            "Current Version of D2MM,"

        SourceDir ->
            "Specifies the directory the source text files are read from. If this is left blank, the built-in 1.13c data files will be used."

        OutputDir ->
            "Specifies the directory that the output data folder will be generated."

        MeleeSplash ->
            "Enables Splash Damage.  Can spawn as an affix on magic and rare jewels."

        IncreasedStackSizes ->
            "Increases tome sizes to 100.  Increases arrows/bolts stack sizes to 511.  Increases key stack sizes to 100."

        IncreaseMonsterDensity ->
            "Increases monster density throughout the map by the given factor."

        EnableTownSkills ->
            "Enable the ability to use all skills in town."

        NoDropZero ->
            "Guarantees that a monster drops something upon death."

        QuestDrops ->
            "Act bosses will always drop quest drops."

        UniqueItemDropRate ->
            "Increases the drop rate of unique and set items.  When using this setting, high values prevent some monsters from dropping set items."

        RuneDropRate ->
            "Increases rune drop rates. Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively. E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100."

        StartWithCube ->
            "Newly created characters will start with a cube."

        Cowzzz ->
            "Enables the ability to recreate a cow portal after killing the cow king.  Adds cube recipe to cube a single tp scroll to create the cow portal."

        RemoveLevelRequirements ->
            "Removes level requirements from items."

        RemoveAttRequirements ->
            "Removes stat requirements from items."

        RemoveUniqCharmLimit ->
            "Removes unique charm limit in inventory."

        RandomOptions ->
            "RandomOptions"

        Randomize ->
            "Randomize all all uniques, sets, and runewords."

        Seed ->
            "Seed"

        IsBalanced ->
            "Allows props only from items within 10 levels of the base item so that you don't get crazy hell stats on normal items, but still get a wide range of randomization."

        BalancedPropCount ->
            "Pick prop count on items based on counts from vanilla items. Picks from items up to 10 levels higher when randomizing."

        AllowDupProps ->
            "If turned off, prevents the same prop from being placed on an item more than once. e.g. two instances of all resist will not get stacked on the same randomized item."

        MinProps ->
            "Minimum number of props an item can have."

        MaxProps ->
            "Maximum number of props an item can have."

        UseOSkills ->
            "Change class only skill props to spawn as oskills."

        PerfectProps ->
            "All props will have a perfect max value when spawning on an item."

        UseSeed ->
            "Provide a specific seed to use.  Toggling on/off will generate a new seed."
