package flac

import (
	"io"

	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/metadata"
	"github.com/audioid/audioid/utils"
	"github.com/valyala/bytebufferpool"
)

func (block *MetadataBlock) ApplyTo(t *metadata.Track) {
	if block.Type == BlockTypeVorbisComment {
		block.Data.(*VorbisComment).Apply(t)
	}
	if block.Type == BlockTypeStreamInfo {
		block.Data.(*StreamInfo).Apply(t)
	}
	if block.Type == BlockTypePicture {
		block.Data.(*Picture).Apply(t)
	}
}

func parseBlock(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) (*MetadataBlock, error) {
	bb.Reset()

	if err := utils.ReadBytes(bb, f, 1); err != nil {
		return nil, errors.Wrap("coud not read flac block type", err)
	}

	block := &MetadataBlock{}

	blockType := BlockType(bb.B[0])

	if (bb.B[0]>>7)&0x1 == 1 {
		blockType ^= (1 << 7)
		block.IsLast = true
	}

	bb.Reset()
	blockLen, err := utils.ReadInt(bb, f, 3)
	if err != nil {
		return nil, err
	}

	bb.Reset()
	switch blockType {
	case BlockTypeVorbisComment:
		err = block.LoadVorbisComment(f, bb)

	case BlockTypeStreamInfo:
		err = block.LoadStreamInfo(f, bb)
	// case BlockTypePicture:
	// 	err = block.LoadPictureBlock(f, bb)

	default:
		block.Type = BlockTypeInvalid
		_, err = f.Seek(int64(blockLen), io.SeekCurrent)
	}

	if err != nil {
		return nil, err
	}

	return block, nil
}
