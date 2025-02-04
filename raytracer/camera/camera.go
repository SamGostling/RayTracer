package camera

type Camera struct {
	imageWidth, imageHeight int
}

func NewCamera(imageWidth int) Camera {
	aspectRatio := 16.0 / 9.0
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	return Camera{
		imageWidth:  imageWidth,
		imageHeight: imageHeight,
	}
}
