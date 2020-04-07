package main

import "encoding/xml"

type EventMessage struct {
	XMLName     xml.Name `xml:"event_message"`
	Text        string   `xml:",chardata"`
	AlgVers     string   `xml:"alg_vers,attr"`
	Category    string   `xml:"category,attr"`
	Instance    string   `xml:"instance,attr"`
	MessageType string   `xml:"message_type,attr"`
	OrigSys     string   `xml:"orig_sys,attr"`
	RefID       string   `xml:"ref_id,attr"`
	RefSrc      string   `xml:"ref_src,attr"`
	Timestamp   string   `xml:"timestamp,attr"`
	Version     string   `xml:"version,attr"`
	CoreInfo    struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Mag  struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"mag"`
		MagUncer struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"mag_uncer"`
		Lat struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"lat"`
		LatUncer struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"lat_uncer"`
		Lon struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"lon"`
		LonUncer struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"lon_uncer"`
		Depth struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"depth"`
		DepthUncer struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"depth_uncer"`
		OrigTime struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"orig_time"`
		OrigTimeUncer struct {
			Text  string `xml:",chardata"`
			Units string `xml:"units,attr"`
		} `xml:"orig_time_uncer"`
		Likelihood  string `xml:"likelihood"`
		NumStations string `xml:"num_stations"`
	} `xml:"core_info"`
	Contributors struct {
		Text        string `xml:",chardata"`
		Contributor struct {
			Text        string `xml:",chardata"`
			AlgInstance string `xml:"alg_instance,attr"`
			AlgName     string `xml:"alg_name,attr"`
			AlgVersion  string `xml:"alg_version,attr"`
			Category    string `xml:"category,attr"`
			EventID     string `xml:"event_id,attr"`
			Version     string `xml:"version,attr"`
		} `xml:"contributor"`
	} `xml:"contributors"`
	GmInfo struct {
		Text          string `xml:",chardata"`
		GmcontourPred struct {
			Text    string `xml:",chardata"`
			Number  string `xml:"number,attr"`
			Contour []struct {
				Text string `xml:",chardata"`
				MMI  struct {
					Text  string `xml:",chardata"`
					Units string `xml:"units,attr"`
				} `xml:"MMI"`
				PGA struct {
					Text  string `xml:",chardata"`
					Units string `xml:"units,attr"`
				} `xml:"PGA"`
				PGV struct {
					Text  string `xml:",chardata"`
					Units string `xml:"units,attr"`
				} `xml:"PGV"`
				Polygon struct {
					Text   string `xml:",chardata"`
					Number string `xml:"number,attr"`
				} `xml:"polygon"`
			} `xml:"contour"`
		} `xml:"gmcontour_pred"`
	} `xml:"gm_info"`
}


