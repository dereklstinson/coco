package coco

/*
Result
*/

//ObjectBB is the form for results for Bounding Boxes Object detection.
/*From cocodataset.org documentation:
Note: box coordinates are floats measured from the top left image corner
(and are 0-indexed). We recommend rounding coordinates to the nearest tenth
of a pixel to reduce resulting JSON file size.
*/
type ObjectBB struct {
	ImageID    int        `json:"image_id"`
	CategoryID int        `json:"category_id"`
	BBox       [4]float32 `json:"bbox"` // x,y widith, height,
	Score      float32    `json:"score"`
}

//ObjectSeg is the form for results for Segmentation Object detection.
/*From cocodataset.org documentation:
Note: a binary mask containing an object segment should be encoded to RLE
using the MaskApi function encode().
*/
type ObjectSeg struct {
	ImageID      int     `json:"image_id"`
	CategoryID   int     `json:"category_id"`
	Segmentation RLE     `json:"segmentation"`
	Score        float32 `json:"score"`
}

//KeyPoint is the form for results of Keypoint Detection
/*From cocodataset.org documentation:
Note: keypoint coordinates are floats measured from the top left image corner
(and are 0-indexed). We recommend rounding coordinates to the nearest pixel
to reduce file size. Note also that the visibility flags vi are not
currently used (except for controlling visualization), we recommend simply setting vi=1.
*/
type KeyPoint struct {
	ImageID    int       `json:"image_id"`
	CategoryID int       `json:"category_id"`
	Keypoints  []float32 `json:"keypoints"`
	Score      float32   `json:"score"`
}

//StuffSeg is the form for results of Stuff Segmentation
/*From cocodataset.org documentation:
The stuff segmentation format is identical to the object segmentation format
except the score field is not necessary. Note: We recommend encoding each
label that occurs in an image with a single binary mask. Binary masks should
be encoded via RLE using the MaskApi function encode().
*/
type StuffSeg struct {
	ImageID      int `json:"image_id"`
	CategoryID   int `json:"category_id"`
	Segmentation RLE `json:"segmentation"`
}

//PanopticSeg is the form for results for Panoptic Segmentation
/*From cocodataset.org documentation:
For the panoptic task, each per-image annotation should have two parts:
(1) a PNG that stores the class-agnostic image segmentation
(2) a JSON struct that stores the semantic information for each image segment.
The PNGs should be located in the folder annotations/name/*,
where annotations/name.json is the JSON file.
For details see the ground truth format for panoptic segmentation.
Results for evaluation should contain both the JSON and the PNGs.
*/
type PanopticSeg struct {
	ImageID      int                 `json:"image_id"`
	FileName     string              `json:"file_name"`
	SegmentsInfo []SegmentInfoResult `json:"segments_info"`
}

//SegmentInfoResult is used in PanopticSeg
type SegmentInfoResult struct {
	ID         int `json:"id"`
	CategoryID int `json:"category_id"`
}

//ImageCaptioning is the form for results for Image Captioning
type ImageCaptioning struct {
	ImageID int    `json:"image_id"`
	Caption string `json:"caption"`
}
