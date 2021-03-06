// Copyright 2015 YP LLC.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

var (
	//testDriver cephRBDVolumeDriver
	testPool  = "test"
	testImage = "image"
)

func TestShUnlockImage(t *testing.T) {
	// lock it first ... ?
	locker, err := testDriver.lockImage(testPool, testImage)
	if err != nil {
		log.Printf("WARN: Unable to lock image in preparation for test: %s", err)
	}
	// print the locker
    log.Printf("Locker: %s", locker)
	// now unlock it
	err = testDriver.unlockImage(testPool, testImage, locker)
	assert.Nil(t, err, fmt.Sprintf("Unable to unlock image using sh rbd: %s", err))
}

func TestLockerOverride(t *testing.T) {
	_, err := testDriver.lockImage(testPool, testImage)
	if err != nil {
		log.Printf("WARN: Unable to lock image in preparation for test: %s", err)
	}

	// Try to lock it again
	locker, err := testDriver.lockImage(testPool, testImage)
	if err != nil {
		log.Printf("WARN: Unable to relock image: %s", err)
	}
	// print the locker
	log.Printf("Locker: %s", locker)
	// now unlock it
	err = testDriver.unlockImage(testPool, testImage, locker)
	assert.Nil(t, err, fmt.Sprintf("Unable to unlock image using sh rbd: %s", err))
}

func test_sh(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	log.Printf("DEBUG: sh CMD: %q", cmd)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Printf("ERROR: %q: %s", err, stderr)
	}
	return strings.Trim(string(out), " \n"), err
}
