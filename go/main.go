// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/genomics/v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// The following utility functions exist for the purpose of
// handling the OAuth2 authentication flow:
//
//  1- Check if credentials have been cached locally (in credentials.dat)
//  2- If no credential
//     a- Emit a URL for the user to paste into their browser
//        (to then authorize access to their genomic data)
//     b- Emit a prompt for the resulting authorization code
//     c- Save credentials to a local file (credentials.dat)
//  3- Create an HTTP client with the authorization attached

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile := tokenCacheFile()

	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser to authorize:\n%s\n\n"+
		"Enter the authorization code: ", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() string {
	return "credentials.dat"
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// End OAuth2 utility functions

func main() {
	// Set up command line parsing
	//
	// Usage:
	//   go run main.go [--client_secrets_filename <client_secrets.json>]
	// where the default client secrets file is ./client_secrets.json
	//
	var clientSecrets string
	flag.StringVar(&clientSecrets,
		"client_secrets_filename", "client_secrets.json",
		"The filename of a client_secrets.json file from a "+
			"Google `Client ID for native application` that "+
			"has the Genomics API enabled.")

	flag.Parse()

	// Read up the client secrets contents
	secrets, err := ioutil.ReadFile(clientSecrets)
	if err != nil {
		log.Fatal(err)
	}

	// Create an OAuth2 config object from the client secrets
	config, err := google.ConfigFromJSON(
		secrets, "https://www.googleapis.com/auth/genomics")
	if err != nil {
		log.Fatal(err)
	}

	// Run the OAuth2 authorization flow to get an HTTP client for requests
	client := getClient(context.Background(), config)

	// Create a genomics API service
	svc, err := genomics.New(client)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Authorized to query genomics data...
	// This example gets the read bases for NA12878 at specific a position
	//
	datasetId := "10473108253681171589" // This is the 1000 Genomes dataset ID
	sample := "NA12872"
	referenceName := "22"
	referencePosition := int64(51003835)

	// 1. First find the read group set ID for the sample
	rsRes, err := svc.Readgroupsets.Search(&genomics.SearchReadGroupSetsRequest{
		DatasetIds: []string{datasetId},
		Name:       sample,
	}).Fields("readGroupSets(id)").Do()
	if err != nil {
		log.Fatal(err)
	}
	if len(rsRes.ReadGroupSets) != 1 {
		fmt.Fprintln(os.Stderr, "Searching for "+sample+" didn't return the right number of read group sets")
		return
	}
	readGroupSetId := rsRes.ReadGroupSets[0].Id

	// 2. Once we have the read group set ID,
	// lookup the reads at the position we are interested in
	rRes, err := svc.Reads.Search(&genomics.SearchReadsRequest{
		ReadGroupSetIds: []string{readGroupSetId},
		ReferenceName:   referenceName,
		Start:           referencePosition,
		End:             referencePosition + 1,
	}).Fields("alignments(alignment,alignedSequence)").Do()
	if err != nil {
		log.Fatal(err)
	}

	bases := make(map[uint8]int)
	for _, read := range rRes.Alignments {
		// Note: This is simplistic - the cigar should be considered for real code
		base := read.AlignedSequence[referencePosition-int64(read.Alignment.Position.Position)]
		bases[base]++
	}

	fmt.Printf("%s bases on %s at %d are\n", sample, referenceName, referencePosition)
	for base, count := range bases {
		fmt.Printf("%c: %d\n", base, count)
	}

	//
	// This example gets the variants for a sample at a specific position
	//

	// 1. First find the call set ID for the sample
	csRes, err := svc.Callsets.Search(&genomics.SearchCallSetsRequest{
		VariantSetIds: []string{datasetId},
		Name:          sample,
	}).Fields("callSets(id)").Do()
	if err != nil {
		log.Fatal(err)
	}
	if len(csRes.CallSets) != 1 {
		fmt.Fprintln(os.Stderr, "Searching for "+sample+" didn't return the right number of call sets")
		return
	}
	callSetId := csRes.CallSets[0].Id

	// 2. Once we have the call set ID,
	// lookup the variants that overlap the position we are interested in
	vRes, err := svc.Variants.Search(&genomics.SearchVariantsRequest{
		CallSetIds:    []string{callSetId},
		ReferenceName: referenceName,
		Start:         referencePosition,
		End:           referencePosition + 1,
	}).Fields("variants(names,referenceBases,alternateBases,calls(genotype))").Do()
	if err != nil {
		log.Fatal(err)
	}

	variant := vRes.Variants[0]
	variantName := variant.Names[0]

	genotype := make([]string, 2)
	for i, g := range variant.Calls[0].Genotype {
		if g == 0 {
			genotype[i] = variant.ReferenceBases
		} else {
			genotype[i] = variant.AlternateBases[g-1]
		}
	}

	fmt.Printf("the called genotype is %s for %s\n", genotype, variantName)
}
