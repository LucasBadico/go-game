package asset

import (
	"example.raylib.gamedev/game/asset/enum"
)

type Asset struct {
	Key  string
	Path string
	Type enum.AssetType
}

type Loaded struct {
	Asset  Asset
	Loaded interface{}
}

type AssetsGroup interface {
	GetAssets() []*Asset

	// currentGroup(A,B,C) <> nextGroup(A,C) = (B)
	Diffs(*AssetsGroup) AssetsGroup

	// currentGroup(A,B,C) >< nextGroup(A,D) = (A)
	Intersects(*AssetsGroup) AssetsGroup

	// currentGroup(A,B,C) <= nextGroup(B) x= true
	Contains(*AssetsGroup) bool

	// currentGroup(A,B) => nextGroup(A,B,C) = true
	IsContained(*AssetsGroup) bool
}

type AssetGroup struct {
	Assets []*Asset
}

func (g *AssetGroup) GetAssets() []*Asset {
	return g.Assets
}

func (g *AssetGroup) Diffs(o *AssetsGroup) AssetsGroup {
	return &AssetGroup{}
}

func (g *AssetGroup) Intersects(o *AssetsGroup) AssetsGroup {
	return &AssetGroup{}
}

func (g *AssetGroup) Contains(o *AssetsGroup) bool {
	return false
}

func (g *AssetGroup) IsContained(o *AssetsGroup) bool {
	return false
}

func NewAssetGroup(toGroup []*Asset) AssetsGroup {
	return &AssetGroup{
		Assets: toGroup,
	}
}
