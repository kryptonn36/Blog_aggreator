package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request,err:= http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err!= nil{
		return nil,fmt.Errorf("Error in creating New request: %v",err)
	}
	request.Header.Set("User-Agent", "gator")

	resp, err:= client.Do(request)
	if err!= nil{
		return nil,fmt.Errorf("Error in getting response: %v",err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("unexpected status code: %v",resp.StatusCode)
	}

	body, err:= io.ReadAll(resp.Body)

	rssfeed := RSSFeed{}


	err = xml.Unmarshal(body,&rssfeed)
	if err!=nil{
		return &rssfeed,err
	}

	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)

	for i := range rssfeed.Channel.Item{
		rssfeed.Channel.Item[i].Title = html.UnescapeString(rssfeed.Channel.Item[i].Title)
		rssfeed.Channel.Item[i].Description = html.UnescapeString(rssfeed.Channel.Item[i].Description)
	}
	return &rssfeed,nil
	
}