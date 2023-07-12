# Foo Application
This application represents the deprecated advertisement system the engineering team is migrating from.
## The "bad" Logic
This application has some "bad" logic in it that will can be exposed by working through the example in the `etl/README.md` file. 

The issue is in `<repo>/foo/internal/ads/view/ad-view.go` starting on line 44 in the `getAdsByCustomer(w http.ResponseWriter, r *http.Request)` method:
```
// Note: the following if statement represents some old "bad" logic unknown to the dev
	// team and the client. This is what we will be trying to catch with the ETL
	// applications E2E tests
	customer := r.URL.Query().Get(`valeu`)
	if customer == `` {
		api.getAllAds(w, r)
		return
	}
	// end "bad" logic example
```
The retrieval of the customer value from the URL produces an empty result because word "value" in the `r.URL.Query().Get("valeu")` line is misspelled. Then when the code finds that the `customer` value is empty, the call is handed off to `getAllAds(w, r)` leading to the confusion. We were expecting only results for the customer we specified, but instead we got results for every customer. See the `etl` application README for details on how this could be prevented.