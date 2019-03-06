package coco

//ImageCaption is used for the json annotation for image captioning
/*from cocodataset.org documentation
These annotations are used to store image captions. Each caption describes
the specified image and each image has at least 5 captions (some images have more).
See also the captioning task.
*/
type ImageCaption struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []ICAnnotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`
}

//ICAnnotation is used for image capture annotation
type ICAnnotation struct {
	ID      int    `json:"id,omitempty"`
	ImageID int    `json:"image_id,omitempty"`
	Caption string `json:"caption,omitempty"`
}
