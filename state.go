package wireless

import (
	"strconv"
	"strings"
)

// State represents the current status of WPA
type State struct {
	BSSID          string `json:"bssid"`
	SSID           string `json:"ssid"`
	ID             string `json:"id"`
	Mode           string `json:"mode"`
	KeyManagement  string `json:"key_management"`
	WpaState       string `json:"wpa_state"`
	IPAddress      string `json:"ip_address"`
	Address        string `json:"address"`
	UUID           string `json:"uuid"`
	GroupCipher    string `json:"group_cipher"`
	PairwiseCipher string `json:"pairwise_cipher"`
}

// NewState will return the state of the WPA when given the raw output
func NewState(data string) State {
	s := State{}
	for _, l := range strings.Split(data, "\n") {
		bits := strings.Split(strings.TrimSpace(l), "=")
		if len(bits) < 2 {
			continue
		}

		switch bits[0] {
		case "bssid":
			s.BSSID = bits[1]
		case "ssid":
			s.SSID = bits[1]
		case "id":
			s.ID = bits[1]
		case "mode":
			s.Mode = bits[1]
		case "key_management":
			s.KeyManagement = bits[1]
		case "wpa_state":
			s.WpaState = bits[1]
		case "ip_address":
			s.IPAddress = bits[1]
		case "address":
			s.Address = bits[1]
		case "uuid":
			s.UUID = bits[1]
		case "group_cipher":
			s.GroupCipher = bits[1]
		case "pairwise_cipher":
			s.PairwiseCipher = bits[1]
		}
	}

	return s
}

type SignalPoll struct {
	Rssi      int    `json:"rssi"`
	LinkSpeed int    `json:"linkSpeed"`
	Noise     int    `json:"noise"`
	Frequency int    `json:"frequency"`
	Width     string `json:"width"`
}

func NewSignalPoll(data string) SignalPoll {
	s := SignalPoll{}
	for _, l := range strings.Split(data, "\n") {
		bits := strings.Split(strings.TrimSpace(l), "=")
		if len(bits) < 2 {
			continue
		}

		switch strings.ToUpper(bits[0]) {
		case "RSSI":
			s.Rssi, _ = strconv.Atoi(bits[1])
		case "LINKSPEED":
			s.LinkSpeed, _ = strconv.Atoi(bits[1])
		case "NOISE":
			s.Noise, _ = strconv.Atoi(bits[1])
		case "FREQUENCY":
			s.Frequency, _ = strconv.Atoi(bits[1])
		case "WIDTH":
			s.Width = bits[1]
		}
	}

	return s
}
