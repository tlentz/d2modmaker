![release](https://img.shields.io/github/v/release/tlentz/d2modmaker?include_prereleases&sort=semver)

<a href="https://www.buymeacoffee.com/tlentz" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)

# D2 Mod Maker

The idea behind this project is to play Diablo II the way that you want to play.  Every option in the `ModConfig` is optional.

# Links
* [Repository](https://github.com/tlentz/d2modmaker)
* [Releases](https://github.com/tlentz/d2modmaker/releases)
* [Installation Guide](https://docs.google.com/document/d/1M5uY67giX4DGnXHxmApb-Uf5AUZdN5yquidAr2BUR_c/edit?usp=sharing)
* [Discord Community](https://discord.gg/eQt2Z9b)
* [Support this Project](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)

# PlugY and Things
This should be compatible with PlugY and other things such as:
* [PlugY](http://plugy.free.fr/en/index.html) by Yohann Nicolas.
* [MultiRes / BH](https://www.reddit.com/r/slashdiablo/comments/7z5uy1/hd_mod_and_maphack_new_release/) by SlashDiablo.

# Support
<a href="https://www.buymeacoffee.com/tlentz" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)

If you'd like to support the project, you can do so by [buying me a coffee](https://www.buymeacoffee.com/tlentz) or donating via [paypal](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)!

Donations will help support development in the project whether that is new features or bug fixes.  

Anyone who donates, will get recognition in the form of a role in the Discord.

Thanks!

# ModConfig

The mod config is located in `cfg.json`.  You can change this config to your liking to produce a new `data` folder.

## ModConfig Options
* **IncreaseStackSizes** `bool`
    * Increases book of tp to 100
    * Increases book of id to 100
    * Increases arrows maxstack to 511
    * Increases bolts maxstack to 511
    * Increases key stack sizes to 100
* **IncreaseMonsterDensity** `float`
    * Will increase the density of all areas by the given multiplier
    * `MAX: 30.0`
    * `MIN: 0.0`
    * Set to `-1` to omit
* **EnableTownSkills** `bool`
    * Enables all skills in town
* **NoDropZero** `bool`
    * Sets "NoDrop" = 0 (Monsters will always drop items)
* **QuestDrops** `bool`
    * Enables quest drops for boss kills always
* **UniqueItemDropRate** `float`
    * Will increase the rate in which uniques/sets drop
	* When using this setting, high values will prevent some monsters from dropping set items.
		* Act bosses at approximately 10
		* Mini bosses at approximately 85
		* Other special monsters at approximately 200
		* All other monsters at approximately 450
    * Set to `-1` to omit
* **RuneDropRate** `float`
	* Valid values are from 1 (vanilla drop rate) - 100 (even chance for all runes)
	* Does not change the maximum rune any enemy can drop.
	* Scales exponetially:
		* Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively
		* E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100.
* **StartWithCube** `bool`
    * Characters will start with cube when created
* **Cowzzz** `bool`
    * Enables ability to kill cow king and still make cow portal
    * Adds ability to cube 1 town portal scroll to make the cow portal
* **EnterToExit** `bool`
    * If this is true, this will require the user to press enter to close the program
    * If false, it will not prompt user input
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
{
    "IncreaseStackSizes": true,
    "IncreaseMonsterDensity": 1,
    "EnableTownSkills": true,
    "NoDropZero": true,
    "QuestDrops": true,
    "UniqueItemDropRate": 100,
    "RuneDropRate": -1,
    "StartWithCube": true,
    "Cowzzz": true,
    "EnterToExit": true,
    "RandomOptions": {
        "Randomize": true,
        "Seed": -1,
        "IsBalanced": false,
        "MinProps": -1,
        "MaxProps": -1,
        "UseOSkills": true,
        "PerfectProps": false
    }
}
```

# Screenshots
### Nagel
![Nagel](https://i.imgur.com/1zOKK3q.png)
### Raven Claw
![Raven Claw](https://i.imgur.com/tmxZpjc.png)
### Venom Ward
![Venom Ward](https://i.imgur.com/7cLQDBN.png)
### Angelic Halo
![Angelic Halo](https://i.imgur.com/N3Om8II.png)
### Wall of Eyeless
![Wall of Eyeless](https://i.imgur.com/QL07TKL.png)
### MonsterDensity: 30
![MonsterDensity: 30](https://i.imgur.com/d6iCBZA.png)

# Change Log

## v0.3.4
* Fixed an issue where `MaxProps` was effectively one less than the configured value
   * This will invalidate most seeds
   * If MaxProps was 7 or less, the seed can be preserved by setting it to one less than the previous value

## v0.3.3
* Added new option to cfg
	* **RuneDropRate** 'float'
		* Scales Rune drop rates from vanilla to even chance per rune.

## v0.3.2
* Fixed an issue where `MinProps` and `MaxProps` weren't working correctly
   * This may invalidate seeds that were using these options

## v0.3.1
* Fixed an issue where unique drop rate multipliers less than 10 would actually reduce drop rates

## v0.3.0
* Fixed an issue where seeding wasn't working correctly
* Fixed an issue where `IncreaseMonsterDensity` couldn't be a value between 0 and 1

## v0.2.0
* Fixed issue where buckets were incorrect for the `IsBalanced` option.
* Added new option to cfg
    * **EnterToExit** `bool`
        * If this is true, this will require the user to press enter to close the program
        * If false, it will not prompt user input

## v0.1.6
* **Cowzzz** `bool`
    * Enables ability to kill cow king and still make cow portal
    * Adds ability to cube 1 town portal scroll to make the cow portal
