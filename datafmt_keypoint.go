package coco

//KeypointDetection are Key Point Detection data types in the json format
/*From cocodataset.org Documentation:

	A keypoint annotation contains all the data of the object annotation (including id, bbox, etc.)
and two additional fields. First, "keypoints" is a length 3k array where k is the total number
of keypoints defined for the category. Each keypoint has a 0-indexed location x,y and a visibility flag v
defined as v=0: not labeled (in which case x=y=0), v=1: labeled but not visible, and v=2: labeled and visible.
A keypoint is considered visible if it falls inside the object segment.
"num_keypoints" indicates the number of labeled keypoints (v>0) for a given object
(many objects, e.g. crowds and small objects, will have num_keypoints=0). Finally, for each category,
the categories struct has two additional fields: "keypoints," which is a length k array of keypoint names,
and "skeleton", which defines connectivity via a list of keypoint edge pairs and is used for visualization.
Currently keypoints are only labeled for the person category (for most medium/large non-crowd person instances).
See also the keypoint task.
*/
type KeypointDetection struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []KPAnnotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`
	Categories  []KPAnnotation `json:"categories,omitempty"`
}

//KPAnnotation contains the keypoint annotation information
type KPAnnotation struct {
	Keypoints    []float32  `json:"keypoints,omitempty"`
	NumKeypoints int        `json:"num_keypoints,omitempty"`
	ID           int        `json:"id,omitempty"`
	ImageID      int        `json:"image_id,omitempty"`
	CategoryID   int        `json:"category_id,omitempty"`
	Segmentation Segment    `json:"segmentation,omitempty"`
	Area         float32    `json:"area,omitempty"`
	Bbox         [4]float32 `json:"bbox,omitempty"`
	Iscrowd      byte       `json:"iscrowd,omitempty"`
}

//KPCategories are the the catagories for key point
type KPCategories struct {
	Keypoints     []string `json:"keypoints,omitempty"`
	Skeleton      []Edge   `json:"skeleton,omitempty"`
	ID            int      `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Supercategory string   `json:"supercategory,omitempty"`
}

//Edge desribes a 2 point edge Probably [x,y] I haven't tested it yet
type Edge [2]uint32
