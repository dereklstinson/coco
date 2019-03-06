package coco

//ObjectDetection is used for object detection jason format
/*from cocodataset.org documentation

	Each object instance annotation contains a series of fields, including the category
id and segmentation mask of the object. The segmentation format depends on whether
the instance represents a single object (iscrowd=0 in which case polygons are used)
or a collection of objects (iscrowd=1 in which case RLE is used). Note that a single
object (iscrowd=0) may require multiple polygons, for example if occluded. Crowd annotations
(iscrowd=1) are used to label large groups of objects (e.g. a crowd of people).
In addition, an enclosing bounding box is provided for each object
(box coordinates are measured from the top left image corner and are 0-indexed).
Finally, the categories field of the annotation structure stores the mapping of
category id to category and supercategory names. See also the detection task.

*/
type ObjectDetection struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []ODAnnotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`
	Categories  []ODCategories `json:"categories,omitempty"`
}

//ODAnnotation is the object detection annotation
type ODAnnotation struct {
	ID           int        `json:"id,omitempty"`
	ImageID      int        `json:"image_id,omitempty"`
	CategoryID   int        `json:"category_id,omitempty"`
	Segmentation Segment    `json:"segmentation,omitempty"`
	Area         float32    `json:"area,omitempty"`
	Bbox         [4]float32 `json:"bbox,omitempty"`
	Iscrowd      byte       `json:"iscrowd,omitempty"`
}

//ODCategories is the object detection categories.
type ODCategories struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Supercategory string `json:"supercategory,omitempty"`
}
