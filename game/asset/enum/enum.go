package enum

// Step 1: Define a custom type for the `type` field
type AssetType string

// Step 2: Define constants for the allowed values
const (
	AssetType_2D_TEXTURE  AssetType = "AssetType.2D_TEXTURE"
	AssetType_AUDIO_MUSIC AssetType = "AssetType.AUDIO.MUSIC"
	AssetType_AUDIO_SOUND AssetType = "AssetType.AUDIO.SOUND"
)
