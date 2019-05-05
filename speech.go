// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the audio file to transcribe.
	// For pathen til filen kan man ikke bruke en \.
	// Som man ser under gjenkjenner syntaxen dobbel \ (\\) som en, slik at pathen blir gjenkjent
	// som en path og ikke noe annet.
	filename := "C:\\Users\\Stein Ove\\Downloads\\audio-file.FLAC"

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Detects speech in the audio file.
	// I "Encoding" feltet må man definere hvilken type format man ønsker å kjøre. I dette tilfellet
	// blir det FLAC fordi formatet til denne filen er FLAC.
	// I "SampleRate" definerer man hertzen som de har brukt for formatet.
	// Filen som ble brukt ble ikke kjørt fordi "SampleRateHertz" var for lav.
	// Man kan spesifisere fra 8000 til 48000hz. FLAC-filen bruker 44100hz.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: 44100, // Endret fra 16000hz til 44100hz.
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}
