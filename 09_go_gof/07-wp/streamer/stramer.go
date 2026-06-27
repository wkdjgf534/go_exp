package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

// ProcessingMessage is the information sent back to the client.
type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

// VideoProcessingJob is the unit of work to be performed. We wrap this type
// around a Video, which has all the information we need about the input source
// and what we want the output to look like.
type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

// Video is the type for a video that we wish to process.
type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyCan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	fmt.Println("NewVideo: new video created:", id, input)

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encType,
		NotifyChan:   notifyCan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		// encoode the video
		fmt.Println("v.enncode(): About to encode to mp4", v.ID)
		name, err := v.encodeToMP4()
		if err != nil {
			// send information to the notifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}

		fileName = fmt.Sprintf("%s.mp4", name)
	default:
		fmt.Println("v.encode(): error trying to encode video", v.ID)
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
		return
	}

	fmt.Println("v.encode(): sending success message for vide id", v.ID, "to notifyChan")
	v.sendToNotifyChan(
		true,
		fileName,
		fmt.Sprintf("video id %d processed and saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)),
	)
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""
	fmt.Println("v.encodeToMP4: about to try to encode video id", v.ID)
	if !v.Options.RenameOutput {
		// Get the base file name.
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		// TODO: Generate random file name
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}

	fmt.Println("v.encodeToMP4: successfully encoded video id", v.ID)
	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	fmt.Println("v.sendTONotifyChan: sneding message to notifyChan for video id", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

// New creates and returns a new worker pool.
func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	fmt.Println("New creating worker pool")
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// Processor logic
	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
