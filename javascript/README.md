# Getting started in javascript

This example application demonstrates access to the Google Genomics API
using Javascript. With this application, one can perform either
anonymous or authenticated access to sample public genomic data.

This sample is derived from the 
[Getting Started](https://developers.google.com/api-client-library/javascript/start/start-js)
sample, demonstrating how to call Google APIs from JavaScript.
You are highly encouraged to read through the Getting Started instructions.

In this example you will run a local web server on port 8000, which serves
a single HTML + Javascript page. After loading the page, all subsequent
requests from the browser are to Google API servers.

1. Clone or fork this repository

2. Enable the Genomics API on a new or existing Google Cloud Project using
the [Cloud Console](https://console.cloud.google.com/flows/enableapi?apiid=genomics&redirect=https://console.cloud.google.com).

3. Follow the instructions to [Get access keys for your application](https://developers.google.com/api-client-library/javascript/start/start-js#get-access-keys-for-your-application). Get both an API key and an OAuth 2.0 client ID.

    When creating the API key, be sure to set the **Accept requests from these HTTP referrers** to `localhost:8000`.

    When creating the client ID, be sure to:
      * select **Web application** for the **Application type**
      * set the **Authorized JavaScript origins** to `http://localhost:8000`.

4. Update the Javascript in the index.html page with your API key and client ID:

        var API_KEY = 'YOUR-API-KEY';

    and 

        var CLIENT_ID = 'YOUR-CLIENT-ID';

4. Run a local http server (this requires [python](https://www.python.org/download/)):
    ```
    python -m SimpleHTTPServer 8000
    ```

    Note: this server will be available on all IP interfaces on your machine.
    If you would like to limit the server to serving on `localhost`, see this
    [Tech Tip](http://www.linuxjournal.com/content/tech-tip-really-simple-http-server-python)

4. Load the sample application web page:

   http://localhost:8000

5. Use the drop down list under "Call the API" to issue requests against
the Genomics API.

