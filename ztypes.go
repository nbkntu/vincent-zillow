package zs

/**
 * This file contains type definitions corresponding to XML schema
 * of a search result from Zillow service.
 * It also contains XML tag names to unmarshal XML text to data object.
 * Currently the search result is only partially mapped. We take only the fields
 * we're interested in and send them to client.
 * If more fields are needed, we can add more types and include them.
 */

type ZSearchResults struct {
    Request ZRequest `xml:"request"`
    Message ZMessage `xml:"message"`
    Response ZResponse `xml:"response"`
}

type ZRequest struct {
    Address string `xml:"address"`
    CityStateZip string `xml:"citystatezip"`
}

type ZMessage struct {
    Text string `xml:"text"`
    Code int32 `xml:"code"`
}

type ZResponse struct {
    Result []ZResult `xml:"results>result"`
}

type ZResult struct {
    Zpid string `xml:"zpid"`
    Links ZLinks `xml:"links"`
    Address ZAddress `xml:"address"`
    Zestimate ZZestimate `xml:"zestimate"`
}

type ZLinks struct {
    HomeDetails string `xml:"homedetails"`
    GraphsAndData string `xml:"graphsanddata"`
    MapThisHome string `xml:"mapthishome"`
    Comparables string `xml:"comparables"`
}

type ZAddress struct {
    Street string `xml:"street"`
    ZipCode string `xml:"zipcode"`
    City string `xml:"city"`
    State string `xml:"state"`
    Latitude string `xml:"latitude"`
    Longitude string `xml:"longitude"`
}

type ZZestimate struct {
    Amount struct {
        Value string `xml:",chardata"`
        Currency string `xml:"currency,attr"`
    } `xml:"amount"`
}
