[![Github All Releases](https://img.shields.io/github/downloads/tlentz/d2modmaker/total.svg)]()
![release](https://img.shields.io/github/v/release/tlentz/d2modmaker)
[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Ftlentz%2Fd2modmaker&count_bg=%23E7AA5D&title_bg=%23555555&icon=&icon_color=%23BA4141&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

<p align="center">
  <img width="400" height="400" src="/images/D2_Mod_Logo_V3.png">
</p>

# D2 Mod Maker

The idea behind this project is to play Diablo II the way that you want to play.  Every option in the `ModConfig` is optional.

![v0.5.0_ux](/images/v0.5.0_ux.png)

# Links
* [Repository](https://github.com/tlentz/d2modmaker)
* [Releases](https://github.com/tlentz/d2modmaker/releases)
* [Installation Guide](https://docs.google.com/document/d/1M5uY67giX4DGnXHxmApb-Uf5AUZdN5yquidAr2BUR_c/edit?usp=sharing)
* [Discord Community](https://discord.gg/eQt2Z9b)
* [Support this Project](https://github.com/tlentz/d2modmaker#support)

# Support

<a href="https://www.buymeacoffee.com/tlentz" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)

# PlugY and Things
This should be compatible with PlugY and other things such as:
* [PlugY](http://plugy.free.fr/en/index.html) by Yohann Nicolas.
* [MultiRes / BH](https://www.reddit.com/r/slashdiablo/comments/7z5uy1/hd_mod_and_maphack_new_release/) by SlashDiablo.

# Options

The mod config is located in `cfg.json`.  You can change this config to your liking and run to produce a new `data` folder.

## General Options
#### SourceDir `string`
* Specifies the directory the source text files are read from
* If this is omitted, or set to "", the built-in 113c data files will be used. 
#### OutputDir `string`
* Specifies the data directory to write files to.  
* If omitted it will default to creating the data file tree directly underneath 
* the current directory, i.e. 
#### MeleeSplash `bool`
* Enables spawning of jewels that have the added property "Melee Splash"
* If the Generator is enabled it can generate items with this property.
* The Randomizer does not produce items with Melee Splash, so you'll have to use jewels.
#### IncreaseStackSizes `bool`
* Increases book of tp to 100
* Increases book of id to 100
* Increases arrows maxstack to 511
* Increases bolts maxstack to 511
* Increases key stack sizes to 100
#### IncreaseMonsterDensity `float`
* Will increase the density of all areas by the given multiplier
* `MIN: 0.0`
* `MAX: 30.0`
* Set to `-1` to omit
#### EnableTownSkills `bool`
* Enables all skills in town
#### BiggerGoldPiles `bool`
* 10x bigger, fewer gold piles
#### NoFlawGems `bool`
* (Mostly) Disables Flawed & Flawless gems from dropping on higher difficulties.
#### NoDropZero `bool`
* Sets "NoDrop" = 0 (Monsters will always drop items)
#### QuestDrops `bool`
* Enables quest drops for boss kills always
#### UniqueItemDropRate `float`
* Will increase the rate in which uniques/sets drop
* When using this setting, high values will prevent some monsters from dropping set items.
* Act bosses at approximately 10
* Mini bosses at approximately 85
* Other special monsters at approximately 200
* All other monsters at approximately 450
* Set to `-1` to omit
#### RuneDropRate `float`
* Valid values are from 1 (vanilla drop rate) - 100 (even chance for all runes)
* Does not change the maximum rune any enemy can drop.
* Scales exponetially:
* Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively
	* E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100.
#### StartWithCube `bool`
* Characters will start with cube when created
#### Cowzzz `bool`
* Enables ability to kill cow king and still make cow portal
* Adds ability to cube 1 town portal scroll to make the cow portal
#### RemoveLevelRequirements `bool`
* Removes level requirements from items. (Oskill level requirements still apply!)
#### RemoveAttRequirements `bool`
* Removes attribute requirements from items.
#### RemoveUniqCharmLimit `bool`
* Allows to carry more than 1 unique charm of the same type.
#### UseOSkills `bool`
* Will change class only skills to oskills
#### PerfectProps `bool`
* All props will have the max value for min/max values
#### SafeUnsocket `bool`
* Adds recipe (item + quiver) to unsocket an item, returning both the item and everything from its sockets.
#### PropDebug `bool`
* Adds recipe health potion + socketable weapon => debugging weapon.  General idea is to hand-edit the cubemain.txt file to add
* the property you are trying to debug, create and test it.

#### EnterToExit `bool`
* If this is true, this will require the user to press enter to close the program
* If false, it will not prompt user input
---
## RandomOptions `RandomOptions`
#### Randomize `bool`
* Will randomize if set to true
#### UseSeed `bool`
* Will use provided seed if set, generate random seed every run if not set
#### Seed `int`
* Will use this seed for randomization
* Set to `-1` to generate a random seed
#### EnhancedSets `bool`
* Removes all full set bonuses because they change on existing items every time d2mm is run
#### IsBalanced `bool`
* Allows props only from items within 10 levels of the base item so that you don't get crazy hell stats on normal items, but still get a wide range of randomization
#### AllowDupeProps `bool`
* If this value is false, the same prop type will not be placed on an item twice
* E.g. two instances of all resist will not get stacked on the same randomized item
#### BalancedPropCount `bool`
* Pick prop count on items based on counts from vanilla items
* Picks from items up to 10 levels higher when randomizing
* Enabling this setting will make MinProps and MaxProps unused
#### MinProps `int`
* Minimum number of non blank props that spawn on an item
* Set to `-1` to omit
#### MaxProps `int`
* Maximum number of non blank props that spawn on an item
* Set to `-1` to omit
#### ElementalSkills `bool`
* Add the ability to spawn + to cold skills, poison skills etc, not just + fire skill.
---
## GeneratorOptions `GeneratorOptions`
#### Generate `bool`
* Set to turn on the Prop Generator
#### UseSeed `bool`
* Will use provided seed if set, generate random seed every run if not set
#### Seed `int`
* Will use this seed for randomization
* Set to `-1` to generate a random seed
#### EnhancedSets `bool`
* Removes all full set bonuses because they change every time d2mm is run
* Configures all sets to have more partial set bonuses
#### BalancedPropCount `bool`
* Pick prop count on items based on counts from vanilla items
* Generates up to 4 props more than vanilla if needed to match the vanilla item's score.
* Enabling this setting will make MinProps and MaxProps unused
#### MinProps `int`
* Minimum number of non blank props that spawn on an item
* Set to `-1` to omit
#### MaxProps `int`
* Maximum number of non blank props that spawn on an item
* Set to `-1` to omit
#### NumClones `int`
* Number of clone unique items to create.  Clones will have
* same name but different generated properties.
#### PropScoreMultipler `int`
* The I Win lever.  1 = vanilla.  2 = 2x the score of the vanilla item.
#### ElementalSkills `bool`
* Add the ability to spawn + to cold skills, poison skills etc, not just + fire skill.

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
### Melee Splash
![Melee Splash](https://i.imgur.com/7qqDycZ.png)

# How to use UI
1. Launch d2modmaker binary
2. Go to http://localhost:8148
3. Press `Load Config`
4. Change all the things
5. Save Config
6. Run

**NOTE** The d2modmaker binary must be in the same directory as the `cfg.json`, and both the `templates` and `react-ui` folders.

**NOTE** `Load Config` will read the `cfg.json` in the same directory as the d2modmaker binary into the UI.

**NOTE** `Save Config` will write `cfg.json` to the same directory as the d2modmaker binary

**NOTE** `Run` will run the program with the current `cfg.json` loaded into the UI.  If you want to save the cfg that you just ran, you need to press `Save Config`

# Support
<a href="https://www.buymeacoffee.com/tlentz" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)

If you'd like to support the project, you can do so by [buying me a coffee](https://www.buymeacoffee.com/tlentz) or donating via [paypal](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=3FMDQNDZN5N8S&currency_code=USD&source=url)!

Donations will help support development in the project whether that is new features or bug fixes.  

Anyone who donates, will get recognition in the form of a role in the Discord.

Thanks!

# Change Log
## v0.6.0-alpha-23
* Adjustments to Aura availability. Ty iksargodzilla.
* Preliminary support for different mods.  Manual edit of cfg.json only for now.
* Removed UseSetsSeed/SetsSeed options in favor of EnhancedSets.
## v0.5.4
* Adding BiggerGoldPiles, NoFlawGems and SafeUnsocket.  
* Patch to elementalskills always being turned on.  This is causing problems with +fireskills coming up negative, and not showing +fireskills text.
* Added elementalskills option in UI in Randomizer section
* Changed AllowDupProps & AllowDuplicateProps to AllowDupeProps
* Fix version # being loaded from cfg.json
* Added ElementalSkills option (+Cold,Lightning,Magic,Poison Skills)

## v0.5.3
* Fixed version #

## v0.5.2
* [bugfix] - fixed density overlap in old code, which was squaring density for nightmare, no increase for hell.
* Upped density max to 45 and split between MonStats.txt & Levels.txt so that the density caps are not hit.

## v0.5.2-alpha-22
* [bugfix] PropScores lines weren't allowing Faith or Exiles Path

## v0.5.2-alpha-21
* Expanded level and item availability of auras for the Generator
* Merged in cleanup from obc/mobdensity branch
* [bugfix] Path fixes to get the srcdir option working again.

## v0.5.2-alpha-20
* Set Generator default to false to work around the ui bug.
* Fix for PerfectProps doing "r=" proptypes
## v0.5.2-alpha-19
* Set PropScoreMultiplier to 1 if user had set it to 0  (thx Speculator)
* Fix Seeds.txt code to write correct seed, and in addition write SetsSeed.
* Balance adjustment in PropScores.txt, some due to the new level limit based code from a17
* Fix for broken Monster Density code that was squaring uniques monster density in nm, and no increase in hell
* Increased Monster Density to Max and split density between Monstats.txt and Levels.txt.  Ty Necrolis (Issue #81)
## v0.5.2-alpha-18
* More prop balance by iksargodzilla
* Fixed issue score cap from a17 was not working.
## v0.5.2-alpha-17
* Added debug cube recipes: axe + 1 health potion = axe with 1 each light,cold,fire,poison, magic skills.  Not added to UI, this allows
* testing of new props to verify that they are working correctly.
* Changed elementalkills allow curses (-lightning skills, -cold skills, etc)  Beware this may shift existing items + elemental skills down by 2, possibly going negative.
* Added PropScores.txt ScoreMax column and capped max score to itemlevel * (1.2 + 0.1 * (PropScoreMultiplier-1)) but only when PropScoreMult < 4
* This prevents low level items from spawning with very high values for props that have low ScoreMax.
## v0.5.2-alpha-16
* Bugfix: Randomizer:IsBalanced was broken
## v0.5.2-alpha-15
* Fixed ElementalSkills
* Increased ranges of some damage props to allow higher values/scores
## v0.5.2-alpha-13
* Adding 2hander buff to Generator
* Fix Randomizer crash
## v0.5.2-alpha-12
* Added ElementalSkills handling in Randomizer
* Balance adjustment to SynergyGroup and ias (properties swing1-3)
* Fixed some broken curses (-AC etc.)
## v0.5.2-alpha-11
* Added "SafeUnsocket" runeword recipe
* Added "ElementalSkills": ability to Generate items with + cold, poison, lightning, or magic skills.
## v0.5.2-alpha-10
* Separated out Generator from Randomizer options
* Made OSkills and PerfectProps their own separate entities, capable of running against vanilla, randomized or generated items.
* Separated out Sets Seed from the other seed.
## v0.5.2-alpha-9
* Fix for ScoreLimit not applying correctly to proppartype s (skill)
* Made oskill conversion its own module, run after Generator or Randomizer
## v0.5.2-alpha-8
* Added PropScores.txt:ScoreLimit -- Limits the max score rolled for a new affix to % of vanilla item score.
## v0.5.2-alpha-7
* Many bugfixes.
* Widened out the allowable range of props rolled in the beginning.
* Added per-slot-ish probability buckets based on item code and type
## v0.5.2-alpha-6
* Fixes to +skills to prevent +5 all & +3 class showing up on same item
* Made +class skill cheaper on class specific
## v0.5.2-alpha-5
* Fix to stamdrain prop
## v0.5.2-alpha-4
* Fixes to Scorer for MinProps support
* Added r= PropParType to allow generation of props in multiples of 5 or 10
* Balancing, and more balancing in PropScores.txt mainly to do with damage, but also restricted 1-100% mana steal to staves, 15% max for everything else.
## v0.5.2-alpha-3
* Added Scorer & Generator modules.
   * Scorer reads in Unique, Sets, Setitems & Runes and calculates scores for each item.
   * Generator uses the scores from the Scorer to generate new weighted random properties with are then filtered according to score.  It will keep generating props until it reaches # props on vanilla item + 4 (if BalancedPropCount) or MaxProps.  It also scales the min & max values based on the target score.
## v0.5.1
* Fixed an issue that caused 1.14 game to crash with Cowzzz option enabled.
* Adds new feature `RemoveLevelRequirements`
    * Removes level requirements from items (Oskill level requirements still apply!).
* Adds new feature `RemoveAttRequirements` 
    * Removes attribute requirements from items.
* Adds new feature `RemoveUniqCharmLimit`
    * Allows to carry more than 1 unique charm of the same type.
* Fix for `BalancedPropCount` not applying

## v0.5.0
* Randomization has been reworked. Old seeds are invalidated.
* Refactors IsBalanced mode
   * Now item properties are selected from items up to 10 levels above
   * This replaces the 0-30, 31-60, and 61+ buckets
* Fixes bugs with Runeword property parsing and randomizing
   * Runeword props were previously always added to the 0-30 bucket, allowing them on all items regardles of runeword levels
   * Runewords were previously assigned props only from the 0-30 bucket.
   * Runeword properties and randomization are now assigned by the level requirement of their highest rune.
* Adds a new randomization option: BalancedPropCount
   * This option enables picking the prop count for items from the counts on vanilla items
   * The count is pulled from items up to 10 levels above the item being randomized
   * The MinProps and MaxProps settings will be ignored if this is enabled
* Adds a new randomization option: AllowDupeProps
   * If this property is false (default), the same property type will not be added to an item twice (e.g. two instances of resist all)
* Prevents two auras from being placed on the same item. This is bugged in the game, and one aura would not work.
* Adds an option to specify the directory to read source Diablo 2 text files from instead of using the built-in 1.13c data.
* Adds an option to specify the output directory.
* NEW USER INTERFACE!
![v0.5.0_ux](/images/v0.5.0_ux.png)

## v0.4.0
* Adds new feature `MeleeSplash`
   * This enables `Splash Damage` as an affix on jewels.
   * Can spawn on any magic or rare jewel.
   * ![Melee Splash](https://i.imgur.com/7qqDycZ.png)

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
    
# Credits
* [Dead Weight Design](https://www.instagram.com/deadweightdesign/) - Thanks for creating the logo!
* tlentz, Deadlock39, OldBeardedCoder/EMPY -- Teh Devs
* Amek for being the true moderating god he is and for his awesome tutorials and cat herding.
* iksargodzilla - Thank-you so much for doing 90% of the grunt work for the scoring engine
* macohan, Negative Inspiration, for helping with design and being a huge help with the newbies.
* The many others that reported bugs, proposed enhancements and gave moral support & encouragement.
