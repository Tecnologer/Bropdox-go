package proto

import (
	"fmt"
	"path/filepath"
)

func NewFile(path string) *File {
	return &File{
		Path: path,
	}
}

func ParseErrorToResponse(err error) *Response {
	return CreateErrorResponse(fmt.Sprint(err))
}

func CreateErrorResponse(message string) *Response {
	return &Response{
		Content: &Response_ErrorResponse{
			ErrorResponse: &ErrorResponse{
				Message: message,
			},
		},
	}
}

func CreateFileResponse(file *File, _type TypeResponse) *Response {
	return &Response{
		Content: &Response_FileResponse{
			FileResponse: &FileResponse{
				File: file,
				Type: _type,
			},
		},
	}
}

func CreateFilesResponse(files []*File) *Response {
	return &Response{
		Content: &Response_FilesResponse{
			FilesResponse: &FilesResponse{
				Files: files,
			},
		},
	}
}

//GetFolderPath returns the path for the folder where the file is
func (f *File) GetFolderPath() string {
	return filepath.Dir(f.Path)
}
