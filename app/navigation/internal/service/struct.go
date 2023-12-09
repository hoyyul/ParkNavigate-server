package service

type SearchRequest struct {
	Resource  []string `json:"resource"`
	Keyword   string   `json:"keyword"`
	Latitude  float32  `json:"latitude"`
	Longitude float32  `json:"longitude"`
	Radius    int      `json:"radius"`
	SortType  int      `json:"sortType"`
}
type Extensions struct {
	ParkingAvailabilityFlag bool `json:"parkingAvailabilityFlag"`
}

// Response 根据提供的 JSON 数据定义的结构体
type SearchResponse struct {
	ResultNum int     `json:"resultNum"`
	PlaceList []Place `json:"placeList"`
}

// Place 代表 placeList 中的对象
type Place struct {
	ProviderType    string            `json:"providerType"`
	PoiId           string            `json:"poiId"`
	DataVersion     string            `json:"dataVersion"`
	Name            string            `json:"name"`
	Yomi            string            `json:"yomi"`
	Reading         string            `json:"reading"`
	Description     string            `json:"description"`
	Latitude        float32           `json:"latitude"`
	Longitude       float32           `json:"longitude"`
	Tel             string            `json:"tel"`
	AlcFlg          bool              `json:"alcFlg"`
	CigarFlg        bool              `json:"cigarFlg"`
	DrvThrFlg       bool              `json:"drvThrFlg"`
	AtmFlg          bool              `json:"atmFlg"`
	ParkFlg         bool              `json:"parkFlg"`
	GenreM          []GenreInfo       `json:"genreM"`
	GenreS          []GenreInfo       `json:"genreS"`
	GenreSS         []GenreInfo       `json:"genreSS"`
	Majorgroupcode  string            `json:"majorgroupcode"`
	GuidePoint      GuidePoint        `json:"guidePoint"`
	PrefCode        string            `json:"prefCode"`
	CityCode        string            `json:"cityCode"`
	Address         string            `json:"address"`
	AddressYomi     string            `json:"addressYomi"`
	ParkingStatus   int               `json:"parkingStatus"`
	ParentChildInfo []ParentChildInfo `json:"parentChildInfo"`
}

// GenreInfo 代表 genreM、genreS 和 genreSS 中的对象
type GenreInfo struct {
	GenreKey  int    `json:"genreKey"`
	GenreName string `json:"genreName"`
	GenreYomi string `json:"genreYomi"`
}

// GuidePoint 代表 guidePoint 对象
type GuidePoint struct {
	GuidePointNum  int               `json:"guidePointNum"`
	GuidePointList []GuidePointEntry `json:"guidePointList"`
}

// GuidePointEntry 代表 guidePointList 中的对象
type GuidePointEntry struct {
	Kind      string  `json:"kind"`
	Priority  int     `json:"priority"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ParentChildInfo 代表 parentChildInfo 中的对象
type ParentChildInfo struct {
	ParentPoiId  string `json:"parentPoiId"`
	ChildPoiId   string `json:"childPoiId"`
	RelationFlag string `json:"relationFlag"`
}
