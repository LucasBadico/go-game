package hellogame

import (
	"example.raylib.gamedev/game/asset"
	assetenum "example.raylib.gamedev/game/asset/enum"
)

var grassAsset = asset.Asset{
	Key:  "simple.grass",
	Type: assetenum.AssetType_2D_TEXTURE,
	Path: "resource/grass.png",
}

var characterAsset = asset.Asset{
	Key:  "character.basic",
	Type: assetenum.AssetType_2D_TEXTURE,
	Path: "resource/basic_charakter_spritesheet.png",
}

var backgroundMusicAsset = asset.Asset{
	Key: "background.music",
	Type: assetenum.AssetType_AUDIO_MUSIC,
	Path: "resource/bg-music.mp3",
}
