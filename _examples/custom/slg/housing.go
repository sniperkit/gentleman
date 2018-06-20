package slg

import (
	"encoding/xml"
	"time"
)

type Photo struct {
	URL string `xml:"bigUrl"`
}

type Time struct {
	time time.Time
}

func (c *Time) Time() time.Time {
	return c.time
}

func (c *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02T15:04:05"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = Time{parse}
	return nil
}

type Housing struct {
	ID           int     `xml:"idAnnonce"`
	CreatedAt    Time    `xml:"dtCreation"`
	UpdatedAt    Time    `xml:"dtFraicheur"`
	Photos       []Photo `xml:"photos>photo"`
	Price        float64 `xml:"prix,omitempty"`
	Unit         string  `xml:"prixUnite"`
	URL          string  `xml:"permaLien"`
	RoomCount    int     `xml:"nbPiece"`
	Surface      float64 `xml:"surface,omitempty"`
	ZipCode      string  `xml:"cp"`
	FeesIncluded bool    `xml:"si_cc"`
	Fees         float64 `xml:"charges,omitempty"`
	Latitude     float64 `xml:"latitude,omitempty"`
	Longitude    float64 `xml:"longitude,omitempty"`
}
