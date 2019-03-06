package coco

//Information is basic image and Coco information and is shared between all the data formats
type Information struct {
	Year        int    `json:"year,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Contributor string `json:"contributor,omitempty"`
	URL         string `json:"url,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}

//Image is the image information and is shared between all the dataformats
type Image struct {
	ID           int    `json:"id,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	License      int    `json:"license,omitempty"`
	FlickrURL    string `json:"flickr_url,omitempty"`
	CocoURL      string `json:"coco_url,omitempty"`
	DateCaptured string `json:"date_captured,omitempty"`
}

//License is the license information and is shared between all the formats
type License struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

//Segment interface a placeholder for Segmentation data structures
//It can either be an RLE or a Polygon
type Segment interface {
}

/*

These are helpers to work with the shared json catagory.


*/

//RLEgo is a struct that will take the info from json to an RLE format
type RLEgo struct {
	Counts []uint32
	Size   []uint32
}

//Polygon from what ive seen looks to be in the form of [1][x0, x1,x2,x3 . . .] but it could be more on the first part.
type Polygon [][]float32

//SegmentationHelper is used for segmentation
type SegmentationHelper struct {
	Poly Polygon
	Rle  RLEgo
}
