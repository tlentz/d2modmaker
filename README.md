# D2 Mod Maker

This is a program is used to create your own Diablo II mod, the way you want it to be given a simple config.

# How It Works
1. Clone the repository
2. `cd d2modmaker && ./bin/d2modmaker`

* This will generate a `data.zip` which will hold the `data` folder that you can put into your Diablo II folder.
* The generated files are based on the original 1.13c files
* Add `-direct -txt` to your Diablo II shortcut
* If using Plugy add `-direct -txt` to your plugy shortcut

# ModConfig

The mod conifg is located in `cfg.json`.  You can change this config to your liking to produce a new `data.zip` folder.


## ModConfig Options
* **IncreaseStackSizes** `bool`
    * Increases book of tp to 100
    * Increases book of id to 100
* **IncreaseMonsterDensity** `float`
    * Will increase the density of all areas by the given multiplier
    * `MAX: 30.0`
    * `MIN 0.0`
    * Set to `-1` to omit
* **EnableTownSkills** `bool`
    * Enables teleport in town
* **NoDropZero** `bool`
    * Sets "NoDrop" = 0 (Monsters will always drop items)
* **QuestDrops** `bool`
    * Enables quest drops for boss kills always
* **UniqueItemDropRate** `float`
    * Will increase the rate in which uniques/sets drop
    * Set to `-1` to omit
* **StartWithCube** `bool`
    * Characters will start with cube when created
* **RandomOptions** `RandomOptions`
    * **Randomize** `bool`
        * Will randomize if set to true
    * **Seed** `int`
        * Will use this seed for randomization
        * Set to `-1` to generate a random seed
    * **IsBalanced** `bool`
        * bucketizes props by levels `[0-30] [31-60] [61+]` so that you don't get crazy hell stats on normal items, but still get a wide range of randomization
    * **MinProps** `int`
        * Minimum number of non blank props that spawn on an item
        * Set to `-1` to omit
    * **MaxProps** `int`
        * Maximum number of non blank props that spawn on an item
        * Set to `-1` to omit
    * **UseOSkills** `bool`
        * Will change class only skills to oskills
    * **PerfectProps** `bool`
        * All props will have the max value for min/max values

## Example ModConfig
```json
    "IncreaseStackSizes": true,
    "IncreaseMonsterDensity": 3.5,
    "EnableTownSkills": true,
    "NoDropZero": true,
    "QuestDrops": true,
    "UniqueItemDropRate": 100,
    "StartWithCube": true,
    "RandomOptions": {
        "Randomize": true,
        "Seed": -1,
        "IsBalanced": false,
        "MinProps": 20,
        "MaxProps": -1,
        "UseOSkills": true,
        "PerfectProps": true
    }
```