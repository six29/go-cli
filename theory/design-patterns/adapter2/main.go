// https://www.youtube.com/watch?v=pNv0dXBqKsM

package main

import "fmt"

// target
type Mobile interface {
	chargeAppleMobile()
}

// concrete prototy implementation
type Apple struct{}

func (a *Apple) chargeAppleMobile() {
	fmt.Println("Charge APPLE mobile")
}

// Adaptee
type Android struct{}

func (a *Android) chargeAndroidMobile() {
	fmt.Println("Charge ANDROID mobile")
}

// Adapter
type AndroidAdapter struct {
	android *Android
}

func (ad *AndroidAdapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
}

// Linux
// adaptee
type Linux struct{}

func (l *Linux) chargeLinuxMobile() {
	fmt.Println("Charge LINUX mobile")
}

// adapter
type LinuxAdapter struct {
	linux *Linux
}

func (ld *LinuxAdapter) chargeAppleMobile() {
	ld.linux.chargeLinuxMobile()
}

// client
type Client struct{}

func (c *Client) chargeMobile(mob Mobile) {
	mob.chargeAppleMobile()
}

func main() {
	apple := &Apple{}
	client := &Client{}
	client.chargeMobile(apple)

	// using the adapter
	android := &Android{}
	androidAd := &AndroidAdapter{
		android: android,
	}
	client.chargeMobile(androidAd)

	// using linux adapter
	linux := &Linux{}
	linuxAd := &LinuxAdapter{
		linux: linux,
	}
	client.chargeMobile(linuxAd)
}
