package coco

//StuffSegmentation is a lot like object detection but the annotations don't include is crowd.
/*From cocodataset.org Documentation:

	The stuff annotation format is identical and fully compatible to the object detection format above
(except iscrowd is unnecessary and set to 0 by default). We provide annotations in both JSON
and png format for easier access, as well as conversion scripts between the two formats.
In the JSON format, each category present in an image is encoded with a single RLE annotation
The category_id represents the id of the current stuff category.
For more details on stuff categories and supercategories see the stuff evaluation page. See also the stuff task.
*/
type StuffSegmentation struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []SGAnnotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`
	Categories  []ODCategories `json:"categories,omitempty"`
}

//SGAnnotation is the object detection annotation
type SGAnnotation struct {
	ID           int        `json:"id,omitempty"`
	ImageID      int        `json:"image_id,omitempty"`
	CategoryID   int        `json:"category_id,omitempty"`
	Segmentation Segment    `json:"segmentation,omitempty"`
	Area         float32    `json:"area,omitempty"`
	Bbox         [4]float32 `json:"bbox,omitempty"`
}
