package coco

//PanopticSegmentation is used for Panoptic Segmentation task
/*From cocodataset.org Documentation:

	For the panoptic task, each annotation struct is a per-image annotation rather than
a per-object annotation. Each per-image annotation has two parts:
(1) a PNG that stores the class-agnostic image segmentation
(2) a JSON struct that stores the semantic information for each image segment.
In more detail:

1.	To match an annotation with an image, use the image_id field (that is annotation.image_id==image.id).

2.	For each annotation, per-pixel segment ids are stored as a single PNG at annotation.file_name.
	The PNGs are in a folder with the same name as the JSON, i.e., annotations/name/ for annotations/name.json.
	Each segment (whether it's a stuff or thing segment) is assigned a unique id.
 	Unlabeled pixels (void) are assigned a value of 0. Note that when you load
	the PNG as an RGB image, you will need to compute the ids via ids=R+G*256+B*256^2.

3.	For each annotation, per-segment info is stored in annotation.segments_info.
	segment_info.id stores the unique id of the segment and is used to retrieve the corresponding
	mask from the PNG (ids==segment_info.id). category_id gives the semantic category and
	iscrowd indicates the segment encompasses a group of objects (relevant for thing categories only).
	The bbox and area fields provide additional info about the segment.

4.	The COCO panoptic task has the same thing categories as the detection task, whereas the
	stuff categories differ from those in the stuff task (for details see the panoptic evaluation page).
	Finally, each category struct has two additional fields: isthing that distinguishes
	stuff and thing categories and color that is useful for consistent visualization.

*/
type PanopticSegmentation struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []PSAnnotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`
	Categories  []PSCategories `json:"categories,omitempty"`
}

//PSAnnotation is for the Panoptic segmentation
type PSAnnotation struct {
	ImageID      int             `json:"image_id,omitempty"`
	FileName     string          `json:"file_name,omitempty"`
	SegmentsInfo []PSSegmentInfo `json:"segments_info,omitempty"`
}

//PSSegmentInfo contains segment info for the annotation
type PSSegmentInfo struct {
	ID         int        `json:"id,omitempty"`
	CategoryID int        `json:"category_id,omitempty"`
	Area       int        `json:"area,omitempty"`
	Bbox       [4]float32 `json:"bbox,omitempty"`
	Iscrowd    byte       `json:"iscrowd,omitempty"`
}

//PSCategories contains category information for the PanopticSegmentation json file
type PSCategories struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Supercategory string    `json:"supercategory,omitempty"`
	Isthing       byte      `json:"isthing,omitempty"`
	Color         [3]uint32 `json:"color,omitempty"`
}
