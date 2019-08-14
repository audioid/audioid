// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package flac

type Application uint32

const (
	// 41544348 - "ATCH"	FlacFile
	ApplicationATCH Application = 0x41544348
	// 42534F4C - "BSOL"	beSolo
	ApplicationBSOL = 0x42534F4C
	// 42554753 - "BUGS"	Bugs Player
	ApplicationBUGS = 0x42554753
	// 43756573 - "Cues"	GoldWave cue points
	ApplicationCUES = 0x43756573
	// 46696361 - "Fica"	CUE Splitter
	ApplicationFICA = 0x46696361
	// 46746F6C - "Ftol"	flac-tools
	ApplicationFTOL = 0x46746F6C
	// 4D4F5442 - "MOTB"	MOTB MetaCzar
	ApplicationMOTB = 0x4D4F5442
	// 4D505345 - "MPSE"	MP3 Stream Editor
	ApplicationMPSE = 0x4D505345
	// 4D754D4C - "MuML"	MusicML: Music Metadata Language
	ApplicationMUML = 0x4D754D4C
	// 52494646 - "RIFF"	Sound Devices RIFF chunk storage
	ApplicationSoundDevicesRIFF = 0x52494646
	// 5346464C - "SFFL"	Sound Font FLAC
	ApplicationSFFL = 0x5346464C
	// 534F4E59 - "SONY"	Sony Creative Software
	ApplicationSONY = 0x534F4E59
	// 5351455A - "SQEZ"	flacsqueeze
	ApplicationSQEZ = 0x5351455A
	// 54745776 - "TtWv"	TwistedWave
	ApplicationTTWV = 0x54745776
	// 55495453 - "UITS"	UITS Embedding tools
	ApplicationUITS = 0x55495453
	// 61696666 - "aiff"	FLAC AIFF chunk storage
	ApplicationAIFF = 0x61696666
	// 696D6167 - "imag"	flac-image application for storing arbitrary files in APPLICATION metadata blocks
	ApplicationIMAG = 0x696D6167
	// 7065656D - "peem"	Parseable Embedded Extensible Metadata
	ApplicationPEEM = 0x7065656D
	// 71667374 - "qfst"	QFLAC Studio
	ApplicationQFST = 0x71667374
	// 72696666 - "riff"	FLAC RIFF chunk storage
	ApplicationFlacRIFF = 0x72696666
	// 74756E65 - "tune"	TagTuner
	ApplicationTUNE = 0x74756E65
	// 78626174 - "xbat"	XBAT
	ApplicationXBAT = 0x78626174
	// 786D6364 - "xmcd"	xmcd
	ApplicationXMCD = 0x786D6364
)

// ToName() returns application name.
// Empty string implies unknown application
func (app Application) ToName() string {
	switch app {
	case ApplicationATCH:
		return "FlacFile"
	case ApplicationBSOL:
		return "beSolo"
	case ApplicationBUGS:
		return "Bugs Player"
	case ApplicationCUES:
		return "GoldWave cue points"
	case ApplicationFICA:
		return "CUE Splitter"
	case ApplicationFTOL:
		return "flac-tools"
	case ApplicationMOTB:
		return "MOTB MetaCzar"
	case ApplicationMPSE:
		return "MP3 Stream Editor"
	case ApplicationMUML:
		return "MusicML: Music Metadata Language"
	case ApplicationSoundDevicesRIFF:
		return "Sound Devices RIFF chunk storage"
	case ApplicationSFFL:
		return "Sound Font FLAC"
	case ApplicationSONY:
		return "Sony Creative Software"
	case ApplicationSQEZ:
		return "flacsqueeze"
	case ApplicationTTWV:
		return "TwistedWave"
	case ApplicationUITS:
		return "UITS Embedding tools"
	case ApplicationAIFF:
		return "FLAC AIFF chunk storage"
	case ApplicationIMAG:
		return "flac-image"
	case ApplicationPEEM:
		return "Parseable Embedded Extensible Metadata"
	case ApplicationQFST:
		return "QFLAC Studio"
	case ApplicationFlacRIFF:
		return "FLAC RIFF chunk storage"
	case ApplicationTUNE:
		return "TagTuner"
	case ApplicationXBAT:
		return "XBAT"
	case ApplicationXMCD:
		return "xmcd"
	default:
		return ""
	}
}
