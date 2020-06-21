# D2 Mod Maker

This is a program is used to create your own Diablo II mod, the way you want it to be given a simple config.

## How It Works
1. Clone the repository
2. `cd d2modmaker && ./bin/d2modmaker`

* This will generate a `data.zip` which will hold the `data` folder that you can put into your Diablo II folder.
* The generated files are based on the original 1.13c files
* Add `-direct -txt` to your Diablo II shortcut
* If using Plugy add `-direct -txt` to your plugy shortcut

## ModConfig

The mod conifg is located in `cfg.json`.  You can change this config to your liking to produce a new `data.zip` folder.

### ModConfig Options
* **IncreaseStackSizes** `bool`
    * Increases book of tp to 100
    * Increases book of id to 100
* **IncreaseMonsterDensity** `int`
    * Will increase the density of all areas by the given multiplier
    * `MAX: 30`