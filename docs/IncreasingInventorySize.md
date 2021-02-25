# Increasing Inventory Size

![ekm5015's Inventory](https://i.imgur.com/e6xGcKu.png)

# Steps
1. There are two files that need to be modded. Inventory.txt and invchar6.dc6. Inventory.txt defines the size of the inventory grid and invchar6.dc6 is the inventory graphic.
2. To use these files you will need to use the -direct - txt parameters in your plugy shortcut. Right click on the shortcut you use to run plugy.exe and modify the Target to add -direct and -txt to it. For example mine is "C:\Program Files (x86)\Diablo II\Mod PlugY\PlugY.exe" -w -direct -txt
3. In order for the -direct -txt arguments to do anything you need to create a folder structure in your plugy folder. For me, my diablo 2 installation is C:\Program Files (x86)\Diablo II. Inside this folder is a folder called Mod PlugY.
4. Create a new folder inside the Mod PlugY folder and call it "data"
    - Create a new folder inside the "data" folder and call it "global"
    - Create two new folders inside the "global" folder and call them "excel" and "ui"
    - Create a new folder inside "ui" and call it "Panel"
5. Once all your folders are created you need to copy the Inventory.txt file into the "excel" folder and invchar6.dc6 into the "Panel" folder. Here is a link to my inventory and invchar6.dc6 files: https://drive.google.com/drive/folders/1OCHT3w_6tTnrRoJSMg1DRMGGwwm1pRk9?usp=sharing
6. Run the game using the PlugY shortcut which has the -direct -txt parameters and your inventory should look like mine in the screen shot above. Good luck!



### Credits
- [/u/ekm5015](https://www.reddit.com/user/ekm5015)
- https://www.reddit.com/r/diablo2/comments/aw9a07/getting_plugy_to_run_with_expanded_inventory_mods/
