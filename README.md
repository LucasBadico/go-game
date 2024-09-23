# Go Game
A simple exercice to create an golang game using Raylib.

The main IDEA is to have all game related stuffs inside GAME package, child packages being:
- scenes and scene manager
- character
- asset

The whole custom logic is suposed to run outside this layer, in any other package, I'm using hello-game as for now. But it could be on main.

# Engine Interface and Raylib
Engine is the interface layer to the game engine. For now I'm using raylib, but in future interations, want to make ready for other game lib/engines that have bindings in golang.

The raylib package is the packages that implements the Engine interface to Raylib. 

So all "ugly code" from the raylib o any other game engine, stays out of sigth.

## Assets
Assets is anything that leaves and go to the memmory of the game. Needs the Enginge.Loader to put on memory and also the Engine.Unload to remove it.

Assets came with the AssetGroup helper. THat gives us collection operations, that the game uses to improve performance on loads and unloads.

