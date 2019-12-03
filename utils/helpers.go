package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"time"
)

// GetAudioFileReadDiscriptor returns read file descriptor.
func GetAudioFileReadDiscriptor(fn string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fd, err := os.Open(dir + "/files/" + fn + ".mp4")
	if err != nil {
		return nil, err
	}

	return fd, nil
}

// response with an hash.
func getHashValue() string {
	ct := time.Now()
	ha := sha1.New()
	ha.Write([]byte(ct.String()))
	return hex.EncodeToString(ha.Sum(nil))
}

// GetAudioFileWriteDiscriptor returns write file descriptor.
func GetAudioFileWriteDiscriptor() (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fn := getHashValue()

	fd, err := os.Create(dir + "/files/" + fn + ".mp4")
	if err != nil {
		return nil, err
	}

	return fd, nil
}

// DelFile return true and err.
func DelFile(fn string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Remove(dir + "/files/" + fn + ".mp4")
	if err != nil {
		return err
	}

	return nil
}

// GetAccessFileDiscriptor returns write file descriptor.
// func GetAccessFileDiscriptor() (*os.File, error) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fd, err := os.Create(dir + "/config/ACCESS.txt")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return fd, nil
// }
